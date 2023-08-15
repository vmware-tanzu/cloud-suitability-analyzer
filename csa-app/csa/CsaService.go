/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa

import (
	"fmt"
	"os"
	"strings"
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

				//---- Add language specific metrics here ----

				for i := 0; i < len(run.Applications); i++ {
					languageCounts := make(map[string]int)
					languages := []string{"java", "php", "python", "dotnet"}
					for j := 0; j < len(run.Applications[i].Rules); j++ {
						for _, lang := range languages {
							if strings.Contains(run.Applications[i].Rules[j].Name, lang) {
								languageCounts[lang]++
							}
						}
						maxCount := 0
						maxLanguage := ""
						for lang, count := range languageCounts {
							if count > maxCount {
								maxCount = count
								maxLanguage = lang
							}
						}
						run.Applications[i].PrimaryLangCount = maxCount
						run.Applications[i].PrimaryLanguage = maxLanguage

						// --- flag applications with 5 or fewer primary language findings and a score of 9 or higher
						if run.Applications[i].PrimaryLangCount <= 5 && run.Applications[i].Score >= 9 {
							// --- the scoring model is reapplied in the UI, so we need this raw score to yield a zero
							//     in the formula: 10 - log10(rawScore)
							run.Applications[i].Score = 0
							run.Applications[i].RawScore = 10000000000
						}
					}
				}

				//---- End language specific metrics ----

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
		run.StopActivityLF("analysis", fmt.Sprintf("A8nalyzing...%s", msg), false, true)
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
		run.StopActivityLF("analysis", fmt.Sprintf("A9nalyzing...%s", msg), false, true)
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
