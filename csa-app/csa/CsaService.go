/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"

	"csa-app/db"
	"csa-app/model"
	"csa-app/report"
	"csa-app/util"

	"github.com/antchfx/xmlquery"
	"gopkg.in/yaml.v3"
)

// The Engine that does file parsing and rule matching
type CsaService struct {
	ruleRepository       db.RuleRepository
	runRepository        db.RunRepository
	findingRepository    db.FindingRepository
	reportDataRepository db.ReportDataRepository
	slocRepository       db.SlocRepository
	scoringRepository    db.ScoringRepository
	reportService        *report.ReportService
	fileUtil             *util.FileUtil
	saveChan             chan interface{} // = make(chan interface{}, *util.MaxBuffer)
	indexChan            chan interface{}
	saveDone             chan interface{} //  = make(chan interface{}, *util.MaxBuffer)
	indexDone            chan interface{} //  = make(chan interface{}, *util.MaxBuffer)
	analysisDone         bool             //= false
	savingDone           bool             //= false
	findingsSaved        int              //= 0
	findingsIndexed      int              //= 0
	yamlDocs             map[string](*yaml.Node)
	yamlMux              sync.Mutex
	xmlDocs              map[string](*xmlquery.Node)
	xmlMux               sync.Mutex
}

func NewCsaSvc(mgr *db.Repositories, reportService *report.ReportService) *CsaService {
	return NewCsaService(mgr.Rules, mgr.Run, mgr.Findings, mgr.Reports, mgr.Sloc, mgr.Scoring, reportService)
}

func NewCsaService(ruleRepository db.RuleRepository,
	runRepository db.RunRepository,
	findingRepository db.FindingRepository,
	reportDataRepository db.ReportDataRepository,
	slocRepository db.SlocRepository,
	scoringRepo db.ScoringRepository,
	reportService *report.ReportService) *CsaService {

	return &CsaService{
		ruleRepository:       ruleRepository,
		runRepository:        runRepository,
		findingRepository:    findingRepository,
		reportDataRepository: reportDataRepository,
		slocRepository:       slocRepository,
		scoringRepository:    scoringRepo,
		reportService:        reportService,
		fileUtil:             util.NewFileUtil(),
		saveChan:             make(chan interface{}, *util.MaxBuffer),
		indexChan:            make(chan interface{}, *util.MaxBuffer),
		saveDone:             make(chan interface{}),
		indexDone:            make(chan interface{}),
		analysisDone:         false,
		savingDone:           false,
		findingsSaved:        0,
		findingsIndexed:      0,
		xmlDocs:              make(map[string](*xmlquery.Node)),
		yamlDocs:             make(map[string](*yaml.Node)),
	}

}

func (csaService *CsaService) PerformAnalysis(run *model.Run) {

	csaService.startRun(run)
	csaService.gatherFiles(run)
	if !util.ProcessHadErrors("gathering") {
		if !*util.WriteConfigsOnly {
			if len(run.Applications) > 0 {
				saveWorkerCnt, indexWorkerCnt := csaService.startWorkers(run)
				if *util.SerialAppAnalysis {
					csaService.SerialAnalysis(run)
				} else {
					csaService.concurrentAnalysis(run)
				}
				csaService.waitForSavingAndIndexingToComplete(run, saveWorkerCnt, indexWorkerCnt)
				csaService.generateSloc(run)
				csaService.scoreApps(run)

				//---- Detect no rules applied ----

				for i := 0; i < len(run.Applications); i++ {
					if run.Applications[i].CIFindings == 0 && run.Applications[i].Score >= 9 {
						// --- the scoring model is reapplied in the UI, so we need this raw score to yield a zero
						//     in the formula: 10 - log10(rawScore)
						run.Applications[i].Score = 0
						run.Applications[i].RawScore = 10000000000
					}
				}

				//---- End language specific metrics to correct perfect scores----

				csaService.generateReports(run)
			}
		}
	}

	csaService.stopRun(run)

	if util.HasErrors() && *util.DisplayErrors {
		util.DumpErrors(run.ID)
	} else if !*util.WriteConfigsOnly {
		//Generate these last so they don't distort run duration measurement! (even though its only milliseconds...)
		csaService.genRuleMetrics(run)
	}

}

func (csaService *CsaService) PerformRecalculate(run *model.Run) {
	run.SetAlias(*util.RecalculateRunLabel)
	csaService.startRun(run)
	priorRun := model.Run{
		ID: *util.RecalculateRunId,
	}
	rules := &[]model.Rule{}
	errors := &[]error{}
	err := csaService.runRepository.CloneRun(priorRun.ID, run)
	if err == nil {
		*rules, err = csaService.ruleRepository.GetRulesForRun(&priorRun)
		if err == nil {
			*errors = csaService.findingRepository.UpdateFindingsForRecalculate(run, *rules)
			if len(*errors) == 0 {
				for a := range run.Applications {
					if err != nil {
						*errors = append(*errors, err)
						break
					}
					app := run.Applications[a]
					sm := &model.ScoringModel{}
					sm, err = csaService.scoringRepository.GetModelByName(app.ScoringModel)
					if err == nil {
						err = app.CalculateScore(sm)
					}
				}
			}
		} else {
			*errors = append(*errors, err)
		}
	} else {
		*errors = append(*errors, err)
	}
	if len(*errors) > 0 {
		log.Error("Error(s) during PerformRecalculate!")
		for e := range *errors {
			log.Error(fmt.Errorf("%w", (*errors)[e]))
		}
	} else {
		csaService.generateReports(run)
	}
	csaService.stopRun(run)
}

func (csaService *CsaService) concurrentAnalysis(run *model.Run) {

	var errors []error

	run.StartActivity("analysis")
	waitGroup := sync.WaitGroup{}

	apps := run.AppsOrdered()
	util.InitializeSpinners(len(apps))
	//app analysis
	for i := range apps {
		waitGroup.Add(1)
		go func(idx int) {
			defer waitGroup.Done()
			errors = append(errors, csaService.analyzeApp(run, apps[idx], csaService.saveChan)...)
		}(i)
	}

	waitGroup.Wait()

	msg := "done!"

	if util.ProcessHadErrors("Analysis") {
		msg = "errors!"
	}

	if !*util.Xtract {
		run.StopActivityLF("analysis", fmt.Sprintf("Analyzing...%s", msg), false, true)
	} else {
		run.StopActivityLF("analysis", "", false, false)
	}

	csaService.analysisDone = true
	close(csaService.saveChan)

}

func (csaService *CsaService) SerialAnalysis(run *model.Run) {
	var errors []error

	run.StartActivity("analysis")

	apps := run.AppsOrdered()
	util.InitializeSpinners(1)

	//app analysis
	for i := range apps {
		errors = append(errors, csaService.analyzeApp(run, apps[i], csaService.saveChan)...)
	}

	msg := "done!"

	if util.ProcessHadErrors("Analysis") {
		msg = "errors!"
	}
	if !*util.Xtract {
		run.StopActivityLF("analysis", fmt.Sprintf("Analyzing...%s", msg), false, true)
	} else {
		run.StopActivityLF("analysis", "", false, false)
	}

	csaService.analysisDone = true
	close(csaService.saveChan)
}

func (csaService *CsaService) startRun(run *model.Run) {
	err := csaService.runRepository.StartRun(run)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Starting Analysis Run! Details: %v", err)
		os.Exit(1)
	}
}

func (csaService *CsaService) stopRun(run *model.Run) {
	err := csaService.runRepository.StopRun(run)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Stoping Analysis Run! Details: %v", err)
	}
}
