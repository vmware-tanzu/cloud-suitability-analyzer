/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db_test

import (
	"log"
	"os"
	"testing"
	"time"

	"csa-app/db"
	"csa-app/db/test_support"
	"csa-app/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestRunRepository(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	run, err := createRun(database, true)

	assert.Nil(t, err)
	assert.True(t, run.ID > 0)
}

func TestGetDistinctRuns(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	runRepo := db.NewRunRepository(database)

	run, _ := createRun(database, false)
	run2, _ := createRun(database, false)

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFinding(run.ID, "app-1", 1, "some category"))
	findingRepository.SaveFinding(createASampleFinding(run.ID, "app-1", 2, "some category"))
	findingRepository.SaveFinding(createASampleFinding(run.ID, "app-1", 0, "some category"))
	run.Findings = 3
	runRepo.StopRun(run)

	findingRepository.SaveFinding(createASampleFinding(run2.ID, "app-2", 3, "some category"))
	findingRepository.SaveFinding(createASampleFinding(run2.ID, "app-2", 2, "some category"))
	findingRepository.SaveFinding(createASampleFinding(run2.ID, "app-2", 1, "some category"))

	run2.Findings = 3
	runRepo.StopRun(run2)

	distinctRunsFromDb, _ := runRepo.GetRunsByCommand("analyze")
	assert.Equal(t, 2, len(distinctRunsFromDb))
	assert.Equal(t, run.ID, distinctRunsFromDb[0].ID)
	assert.Equal(t, run2.ID, distinctRunsFromDb[1].ID)

}

func TestCanCloneRunForRecalculate(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
	runRepository := db.NewRunRepository(database)

	os.Args = []string{
		"csa",
		"recalculate",
		"--run-id=1",
	}

	newRun, _ := createRun(database, true)
	clone := &model.Run{}
	_ = runRepository.CloneRun(newRun.ID, clone)
	findingRepository := db.NewFindingRepository(database)
	originalFindings, _ := findingRepository.GetFindings(newRun.ID)
	cloneFindings, _ := findingRepository.GetFindings(clone.ID)

	assert.EqualValues(t, len(originalFindings), len(cloneFindings))
	assert.Greater(t, len(cloneFindings), 0)
	assert.Greater(t, clone.ID, newRun.ID)
	assert.Equal(t, len(clone.Applications), len(newRun.Applications))
}

func createRun(database *gorm.DB, complete bool) (*model.Run, error) {
	startTime, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	run := &model.Run{
		StartTime: startTime,
		Command:   "analyze",
		Applications: []*model.Application{
			{
				Name: "app-1",
				Tags: []*model.ApplicationTag{
					{
						Value: "tag1",
					},
				},
				SlocCnt:  12345,
				FilesCnt: 321,
			},
			{
				Name: "app-2",
				Tags: []*model.ApplicationTag{
					{
						Value: "tag2",
					},
					{
						Value: "tag3",
					},
				},
				SlocCnt:  54321,
				FilesCnt: 123,
			},
		},
	}

	var runRepository db.RunRepository
	runRepository = db.NewRunRepository(database)

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(1, "app-1", 0, "api1", "pattern1", "api", "rule1"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(1, "app-1", 1, "chuck-e-cheese", "pattern1", "not-api", "rule2"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(1, "app-2", 5, "api2", "pattern2", "api", "rule1"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(1, "app-2", 8, "franks-and-beans", "pattern3", "not-api-as-well", "rule1"))

	err := runRepository.StartRun(run)

	if err == nil && complete {
		err = runRepository.StopRun(run)
	}

	return run, err

}
