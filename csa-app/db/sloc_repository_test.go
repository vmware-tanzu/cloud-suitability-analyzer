/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db_test

import (
	"log"
	"os"
	"testing"

	"csa-app/db"
	"csa-app/db/test_support"
	"csa-app/model"
	"github.com/stretchr/testify/assert"
)

func TestSlocRepository(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	slocRepo := db.NewSlocRepository(database)

	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "java", 101, 3))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "java", 10, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "python", 99, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "python", 9, 3))

	slocs, err := slocRepo.GetSlocForRun(1)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(slocs))

	slocByApps, _ := slocRepo.GetSlocSummaryByApplicationForRun(1)
	assert.Equal(t, 2, len(slocByApps))

	assert.ElementsMatch(t, []model.SlocByApplication{{Application: "app-1", CodeLines: 200, TotalFiles: 5}, {Application: "app-2", CodeLines: 19, TotalFiles: 5}}, slocByApps)

}

func TestGetSummaryFindingsForRun(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	slocRepo := db.NewSlocRepository(database)

	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "java", 101, 3))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "java", 10, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "python", 99, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "python", 9, 3))

	summaryForRun, _ := slocRepo.GetSummaryFindingsForRun(1)

	assert.Equal(t, model.SlocByRun{RunId: 1, Apps: 2, TotalFiles: 10, CodeLines: 219}, summaryForRun)
}

func TestLanguagesByCodeLines(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	slocRepo := db.NewSlocRepository(database)

	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "java", 101, 3))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "java", 10, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "python", 99, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "python", 9, 3))

	languagesByCodeLines, _ := slocRepo.GetTopLanguagesByCodeLines(1)

	assert.ElementsMatch(t, []model.LanguagesByCodeLines{{Lang: "java", CodeLines: 111}, {Lang: "python", CodeLines: 108}}, languagesByCodeLines)
}

func TestGetApplicationsByLanguageForRun(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	slocRepo := db.NewSlocRepository(database)

	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "java", 9, 3))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "java", 10, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "python", 99, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-2", "python", 9, 3))

	applicationsByLanguageForRuns, _ := slocRepo.GetApplicationsByLanguageForRun(1, "java")

	assert.ElementsMatch(t, []model.ApplicationsByLanguageForRun{{Application: "app-2", CodeLines: 10}, {Application: "app-1", CodeLines: 9}}, applicationsByLanguageForRuns)
}

func TestGetLanguagesForRunAndApplication(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	slocRepo := db.NewSlocRepository(database)

	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "java", 9, 3))
	slocRepo.CreateSlocData(createTestSlocData(2, "app-2", "java", 10, 2))
	slocRepo.CreateSlocData(createTestSlocData(1, "app-1", "python", 99, 2))
	slocRepo.CreateSlocData(createTestSlocData(2, "app-2", "python", 9, 3))

	applicationsByLanguageForRuns, _ := slocRepo.GetLanguagesForRunAndApplication(1, "app-1")

	assert.ElementsMatch(t, []model.LanguagesByCodeLines{{Lang: "python", CodeLines: 99}, {Lang: "java", CodeLines: 9}}, applicationsByLanguageForRuns)
}

func createTestSlocData(runid uint, application string, lang string, codelines int, totalfiles int) *model.RunSloc {
	return &model.RunSloc{
		RunID:        runid,
		Application:  application,
		CodeLines:    codelines,
		BlankLines:   0,
		CommentLines: 0,
		Lang:         lang,
		TotalFiles:   totalfiles,
	}
}
