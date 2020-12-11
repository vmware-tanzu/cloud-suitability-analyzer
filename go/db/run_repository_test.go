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

	"github.com/jinzhu/gorm"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/db"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/db/test_support"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/model"
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

func createRun(database *gorm.DB, complete bool) (*model.Run, error) {
	startTime, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	run := &model.Run{
		StartTime: startTime,
		Command:   "analyze",
	}

	var runRepository db.RunRepository
	runRepository = db.NewRunRepository(database)

	err := runRepository.StartRun(run)

	if err == nil && complete {
		err = runRepository.StopRun(run)
	}

	return run, err

}
