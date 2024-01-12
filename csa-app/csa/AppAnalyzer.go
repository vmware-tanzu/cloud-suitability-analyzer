/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"csa-app/db"
	"csa-app/model"
	"csa-app/search"
	"csa-app/util"
)

func (csaService *CsaService) analyzeApp(run *model.Run, app *model.Application, output chan<- interface{}) (errors []error) {

	run.StartActivity(fmt.Sprintf("%s-analysis", app.Name))
	waitGroup := sync.WaitGroup{}

	util.InitializeSpinners(len(run.Applications))

	modCnt := util.GetModCount(run.Files)

	if *util.Verbose {
		fmt.Printf("Got Mod Count [%d]\n", modCnt)
	}

	for i := range app.Files {
		waitGroup.Add(1)
		go func(idx int) {
			defer waitGroup.Done()
			if *util.Verbose {
				util.WriteLog(fmt.Sprintf("Analyzing - %s", app.Name), "Scanning Files...   Filename: %s\n", app.Files[idx].FQN)
			}

			err := csaService.analyzeFile(run, app, app.Files[idx], csaService.saveChan)

			if err != nil {
				if *util.FailFast {
					util.WriteLog("Analyzing...error!", "Error occurred during analysis! Details: %s", err.Error())
					csaService.stopRun(run)
					os.Exit(2)
				} else {
					errors = append(errors, err)
				}
			} else {
				run.FileAnalyzed()
				modResult := run.AnalyzedCnt % modCnt
				if modResult == 0 {
					util.WriteLogWithToken("Analyzing!", fmt.Sprintf("%2.f%%", float64(run.AnalyzedCnt)/float64(run.Files)*100), "Filename: %s...done\n!", app.Files[idx].FQN)
				}
			}
		}(i)
	}

	waitGroup.Wait()

	//Ensure we always get a percent complete message even if we have very few files in an app!
	util.WriteLogWithToken("Analyzing", fmt.Sprintf("%2.f%%", float64(run.AnalyzedCnt)/float64(run.Files)*100), "App: %s...done\n!", app.Name)

	if !*util.Xtract {
		run.StopActivity(fmt.Sprintf("%s-analysis", app.Name), fmt.Sprintf("Analyzing - %s...done!", app.Name), true)
	} else {
		run.StopActivity(fmt.Sprintf("%s-analysis", app.Name), "", false)
	}
	return
}

func (csaService *CsaService) analyzeFile(run *model.Run, app *model.Application, file *util.FileInfo, output chan<- interface{}) error {

	var rulesForFile []model.Rule
	var rulesUsed []string

	findings := 0
	wasAnalyzed := false
	fileNameAnalyzed := false
	hasContentRules := false

	for i := range app.Rules {

		if app.Rules[i].Applies(file.GetCleanedExt(), file.Name) {
			rulesUsed = append(rulesUsed, app.Rules[i].Name)
			//Rule applies to this file!
			if *util.Verbose {
				util.WriteLog("Analyzing", "Rule [%s] applies to file [%s|%s|%s]\n", app.Rules[i].Name, file.Name, file.Ext, file.FQN)
			}
			if app.Rules[i].Target == model.LINE_TARGET {
				rulesForFile = append(rulesForFile, app.Rules[i])
				wasAnalyzed = true
			} else if app.Rules[i].Target == model.CONTENTS_TARGET {
				rulesForFile = append(rulesForFile, app.Rules[i])
				hasContentRules = true
				wasAnalyzed = true
			} else {
				//Check for name hit!
				findingCnt := 0
				findingCnt, app.Rules[i] = csaService.processPatterns(run, app, file, 0, filepath.Base(file.Name), app.Rules[i], output)
				findings += findingCnt
				fileNameAnalyzed = true
			}

		} else {
			if *util.Verbose {
				util.WriteLog("Analyzing", "Rule [%s] does not apply to file [%s|%s|%s]\n", app.Rules[i].Name, file.Name, file.Ext, file.FQN)
			}
		}
	}

	_, fileFindings, err := csaService.processFile(run, app, file, rulesForFile, hasContentRules, output)

	if err != nil {
		return err
	}

	findings += fileFindings

	if wasAnalyzed || fileNameAnalyzed {
		var msg string
		var rulesUsedMsg strings.Builder

		cnt := 0
		for _, rule := range rulesUsed {
			if cnt > 0 {
				rulesUsedMsg.WriteString(",")
			}
			rulesUsedMsg.WriteString(rule)
			cnt++
		}

		if wasAnalyzed {
			msg = fmt.Sprintf("File read and analyzed fully! [%d] findings were documented! Rules Used: %s", findings, rulesUsedMsg.String())
		} else {
			msg = fmt.Sprintf("Filename was analyzed! [%d] findings were documented! Rules Used: %s", findings, rulesUsedMsg.String())
		}

		//Create an info finding for each file analyzed regardless of if it matched a rule!
		fileFinding := model.Finding{
			RunID:       run.ID,
			Filename:    file.Name,
			Fqn:         file.FQN,
			Ext:         file.Ext,
			Category:    model.FILE_ANALYZED_CATEGORY,
			Pattern:     model.ANALYZED_FILE_PATTERN,
			Effort:      0,
			Readiness:   0,
			Criticality: "none",
			Application: file.Dir}

		fileFinding.SetValue(msg)
		fileFinding.AddTag(model.INFO_FINDING)
		fileFinding.AddTag(model.FILE_FINDING)

		//Send "file" finding to save worker
		output <- fileFinding

		run.AddFindings(1)
	}

	if *util.Verbose {
		util.WriteLog("Analyzing", "************ FILE [%s] FINDINGS [%d] ***************\n", file.Name, findings)
	}

	return nil
}

func (csaService *CsaService) saveFindingToDb(id string, work <-chan interface{}, result chan<- interface{}, jointWorker chan<- interface{}, args []interface{}) {

	if len(args) < 1 {
		log.Fatalf("Invalid Arguments for %s!", id)
	}

	run := args[0].(*model.Run)
	run.StartActivity("saving")

	tx := db.StartGormTransaction()
	defer tx.Commit()

	for w := range work {
		target := w.(model.Finding)

		if db.SaveFindingTransacted(tx, &target) {
			csaService.findingsSaved++
			if *util.TxtIndexingEnabled {
				jointWorker <- target
			}

			if csaService.analysisDone {
				modCnt := util.GetModCount(run.Findings)
				if csaService.findingsSaved%modCnt == 0 {
					util.WriteLogWithToken("Saving", fmt.Sprintf("%2.f%%", float64(csaService.findingsSaved)/float64(run.Findings)*100), "Finding: %d saved\n!", target.ID)
				}
			}
		} else {

			util.TrackError("Saving", fmt.Errorf("error saving finding for file [%s]", target.Fqn))

			if *util.FailFast {
				tx.Rollback()
				fmt.Println("Error saving finding during run! Fail Fast Enabled! Stopping Run!")
				util.WriteLog("Saving...failed!", "Failed saving finding for file [%s]", target.Fqn)
				csaService.stopRun(run)
				os.Exit(2)
			}
		}

	}
	//We are done saving!
	result <- true
}

func (csaService *CsaService) addFindingToTextIndex(id string, work <-chan interface{}, result chan<- interface{}, jointWorker chan<- interface{}, args []interface{}) {

	if len(args) < 1 {
		log.Fatalf("Invalid Arguments for %s!", id)
	}

	run := args[0].(*model.Run)
	run.StartActivity("indexing")

	index := search.GetIndex(run)

	indexAvail := true

	if !index.Exists || index.Error != "" {
		fmt.Fprintf(os.Stderr, "[%s] - Error obtaining text index. Details: %s\n", id, index.Error)
		indexAvail = false
	}

	for w := range work {
		target := w.(model.Finding)
		//Add to text index
		if *util.Verbose {
			fmt.Printf("[%s] - Indexing finding [%d]\n", id, target.ID)
		}

		if indexAvail {
			err := index.AddItemToIndex(fmt.Sprint(target.ID), target)
			if err != nil {
				util.TrackError("Indexing", err)
			} else {
				csaService.findingsIndexed++
			}
		}
		if csaService.savingDone {
			modCnt := util.GetModCount(run.Findings)
			if csaService.findingsIndexed%(modCnt) == 0 {
				util.WriteLogWithToken("Indexing", fmt.Sprintf("%2.f%%", float64(csaService.findingsIndexed)/float64(run.Findings)*100), "Finding: %d indexed\n!", target.ID)
			}
		}
	}
	//We are done indexing!
	result <- true
}
