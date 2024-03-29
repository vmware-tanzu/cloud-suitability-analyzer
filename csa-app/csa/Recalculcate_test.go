/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa_test

import (
	"csa-app/csa"
	"csa-app/db"
	"csa-app/db/test_support"
	"csa-app/model"
	"csa-app/report"
	"csa-app/util"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const LOW_SCORING_APP_NAME string = "liberty-arquillian"
const HIGH_EFFORT_RULE_NAME string = "java-jboss"
const SAMPLE_PORTFOLIO_RELATIVE_PATH string = "../frontend/samplePortfolio"

func TestCanCreateNewRunForRecalculate(t *testing.T) {
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	os.Args = []string{
		"csa",
		"recalculate",
		"--run-id=1",
	}
	runRepository := db.NewRunRepository(database)

	newRun := model.NewRun()
	assert.NotNil(t, newRun)

	err = runRepository.StartRun(newRun)
	assert.Nil(t, err)

	err = runRepository.StopRun(newRun)
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, newRun.ID)

	assert.Equal(t, "recalculate", newRun.Command)
	assert.EqualValues(t, 1, *util.RecalculateRunId)
	assert.NotEmpty(t, *util.RecalculateRunLabel)
}

func TestCanCreateNamedNewRunForRecalculate(t *testing.T) {

	os.Args = []string{
		"csa",
		"recalculate",
		"--run-id=1",
	}
	newRun := model.NewRun()
	assert.NotNil(t, newRun)
	assert.EqualValues(t, 1, *util.RecalculateRunId)
	assert.EqualValues(t, "Recalculate Run 1", *util.RecalculateRunLabel)

	os.Args = []string{
		"csa",
		"recalculate",
		"--alias=\"an alias\"",
		"--run-id=2",
	}
	newRun = model.NewRun()
	assert.NotNil(t, newRun)
	assert.EqualValues(t, 2, *util.RecalculateRunId)
	assert.EqualValues(t, "an alias", *util.RecalculateRunLabel)

	os.Args = []string{
		"csa",
		"recalculate",
		"-a:'an alias'",
		"--run-id=3",
	}
	newRun = model.NewRun()
	assert.NotNil(t, newRun)
	assert.EqualValues(t, 3, *util.RecalculateRunId)
	assert.EqualValues(t, "an alias", *util.RecalculateRunLabel)

	os.Args = []string{
		"csa",
		"recalculate",
		"-a \"an alias\"",
		"--run-id=4",
	}
	newRun = model.NewRun()
	assert.NotNil(t, newRun)
	assert.EqualValues(t, 4, *util.RecalculateRunId)
	assert.EqualValues(t, "an alias", *util.RecalculateRunLabel)
}

func TestCanRecalculateSamplePortfolio(t *testing.T) {

	run, repoMgr := ScanSamplePortfolio()
	findings, _ := repoMgr.Findings.GetFindings(run.ID)

	// validate that analysis was performed
	assert.EqualValues(t, len(findings), run.Findings)
	assert.Greater(t, run.Findings, 0)
	assert.EqualValues(t, 2, len(run.Applications))

	// assert that our test subject made it into the mix and that it has a low score
	jbossApp, _ := repoMgr.Run.GetApp(run.ID, LOW_SCORING_APP_NAME)
	assert.NotNil(t, jbossApp)
	assert.Less(t, jbossApp.Score, 8.0)

	// let's get the high effort jboss rule and lower it from 50 to 0
	jbossRule, _ := repoMgr.Rules.GetRuleByName(HIGH_EFFORT_RULE_NAME)
	assert.Greater(t, jbossRule.Effort, 0)
	jbossRule.Effort = 0
	_, _ = repoMgr.Rules.SaveRule(jbossRule)
	// verify that our changed rule effort has persisted
	checkRule, _ := repoMgr.Rules.GetRuleByName(HIGH_EFFORT_RULE_NAME)
	assert.EqualValues(t, checkRule.Effort, 0)

	// set up a new run as a request to recalculate runId 1
	os.Args = []string{
		"csa",
		"recalculate",
		"-a:'jboss rule update'",
		"--run-id=1",
	}

	newRun := model.NewRun()
	newRun.DB = testDatabase
	reportService := report.NewReportService(repoMgr.Reports, repoMgr.Sloc)
	csaService := csa.NewCsaSvc(repoMgr, reportService)
	csaService.PerformRecalculate(newRun)

	reScoredApp, _ := repoMgr.Run.GetApp(newRun.ID, LOW_SCORING_APP_NAME)
	assert.Greater(t, reScoredApp.Score, jbossApp.Score)
	assert.Less(t, reScoredApp.Score, 10.0)
	assert.NotEqualValues(t, reScoredApp.ID, jbossApp.ID)
	assert.EqualValues(t, "jboss rule update", newRun.Alias)
}

var testDatabase *gorm.DB

func ScanSamplePortfolio() (*model.Run, *db.Repositories) {
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	os.Args = []string{
		"csa",
		"analyze",
		"-p",
	}

	run := model.NewRun()
	run.SetPaths(SAMPLE_PORTFOLIO_RELATIVE_PATH)
	run.DB = database

	repoMgr := db.NewRepositoriesManagerForRun(run)
	repoMgr.Scoring.LoadModels()

	reportService := report.NewReportService(repoMgr.Reports, repoMgr.Sloc)
	csaService := csa.NewCsaSvc(repoMgr, reportService)
	run.SetRequestedReports(util.ReportsFlag, "1,2,3,4,5")
	run.ValidateRun()
	csaService.PerformAnalysis(run)

	testDatabase = database
	return run, repoMgr
}
