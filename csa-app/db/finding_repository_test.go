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
	"github.com/stretchr/testify/assert"
)

func TestSaveAndRetrieveFindings(t *testing.T) {
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	target := &model.Finding{
		Rule:        "some rule",
		Pattern:     "some pattern",
		Advice:      "some advice",
		Category:    "some category",
		CreatedAt:   createdAt,
		Criticality: "HIGH",
		Application: "some domain",
		Ext:         "some ext",
		Filename:    "myfilename",
		Fqn:         "myfqn",
		Line:        100,
		Recipes: []model.FindingRecipe{
			{
				URI: "someuri1",
			},
			{
				URI: "someuri2",
			},
		},
		RunID:  10,
		Effort: 5,
		Tags: []model.FindingTag{
			{
				Value: "some tag1",
			},
		},
		UpdatedAt: createdAt,
		Value:     "some value",
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(target)

	findingsFromDb, _ := findingRepository.GetFindings(10)

	assert.Equal(t, 1, len(findingsFromDb))

}

func TestGetApplicationScores(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 1, "some category"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 2, "some category"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 0, "some category"))

	findingRepository.SaveFinding(createASampleFinding(21, "app-2", 3, "some category"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-2", 2, "some category"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-2", 95, "some category"))

	findingsFromDb, _ := findingRepository.GetFindings(21)
	assert.Equal(t, 6, len(findingsFromDb))

	appScores, _ := findingRepository.GetApplicationDetailsForRun(21, 10, 1, false)

	assert.Equal(t, 2, len(appScores))

	assert.Equal(t, "app-2", appScores[0].Application)
	assert.Equal(t, "app-1", appScores[1].Application)

	assert.Equal(t, 100, appScores[0].RawScore)
	assert.Equal(t, 3, appScores[1].RawScore)

}

func TestTopApisInFindings(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	finding1 := createASampleFindingWithPatternAndTag(21, "app-1", 1, "api1", "", "api")
	findingRepository.SaveFinding(finding1)
	finding2 := createASampleFindingWithPatternAndTag(21, "app-1", 1, "api1", "", "api")
	findingRepository.SaveFinding(finding2)
	finding3 := createASampleFindingWithPatternAndTag(21, "app-1", 2, "api2", "", "api")
	findingRepository.SaveFinding(finding3)

	finding5 := createASampleFindingWithPatternAndTag(21, "app-2", 3, "api3", "", "api")
	findingRepository.SaveFinding(finding5)
	finding6 := createASampleFindingWithPatternAndTag(21, "app-2", 2, "api3", "", "api")
	findingRepository.SaveFinding(finding6)
	finding7 := createASampleFindingWithPatternAndTag(21, "app-2", 95, "api3", "", "api")
	findingRepository.SaveFinding(finding7)

	topApisForRun, _ := findingRepository.GetTopApisByScoreForRun(21)
	assert.Equal(t, 3, len(topApisForRun))
	assert.Equal(t, model.ApiUsage{Api: "api3", UsageCount: 3}, topApisForRun[0])
	assert.Equal(t, model.ApiUsage{Api: "api1", UsageCount: 2}, topApisForRun[1])
	assert.Equal(t, model.ApiUsage{Api: "api2", UsageCount: 1}, topApisForRun[2])
}

func TestGetApisForRunAndApplication(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	finding1 := createASampleFindingWithPatternAndTag(21, "app-1", 1, "api1", "", "api")
	findingRepository.SaveFinding(finding1)
	finding2 := createASampleFindingWithPatternAndTag(21, "app-1", 1, "api1", "", "api")
	findingRepository.SaveFinding(finding2)
	finding3 := createASampleFindingWithPatternAndTag(21, "app-1", 2, "api2", "", "api")
	findingRepository.SaveFinding(finding3)
	finding4 := createASampleFindingWithPatternAndTag(21, "app-1", 1, "api2", "", "api")
	findingRepository.SaveFinding(finding4)

	finding5 := createASampleFindingWithPatternAndTag(21, "app-2", 3, "api3", "", "api")
	findingRepository.SaveFinding(finding5)
	finding6 := createASampleFindingWithPatternAndTag(21, "app-2", 2, "api3", "", "api")
	findingRepository.SaveFinding(finding6)
	finding7 := createASampleFindingWithPatternAndTag(21, "app-2", 95, "api3", "", "api")
	findingRepository.SaveFinding(finding7)

	topApisForRun, _ := findingRepository.GetApisByUsageForRunAndApplication(21, "app-1")
	assert.Equal(t, 2, len(topApisForRun))
	assert.ElementsMatch(t, []model.ApiUsage{{Api: "api1", UsageCount: 2}, {Api: "api2", UsageCount: 2}}, topApisForRun)
}

func TestGetApplicationsUsingApiForRun(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 1, "api1", "", "api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 2, "api2", "", "api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 1, "api2", "", "api"))

	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-2", 3, "api1", "", "api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-2", 2, "api1", "", "api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-2", 95, "api3", "", "api"))

	appsByApi, _ := findingRepository.GetApplicationsUsingApiForRun(21, "api1")

	assert.Equal(t, 2, len(appsByApi))
	assert.Equal(t, "app-2", appsByApi[0].Application)
	assert.Equal(t, "app-1", appsByApi[1].Application)
	assert.Equal(t, 2, appsByApi[0].UsageCount)
	assert.Equal(t, 1, appsByApi[1].UsageCount)
}

func TestGetScoreCards(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 0, "api1"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 2, "api2"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 5, "api2"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-1", 8, "api2"))

	findingRepository.SaveFinding(createASampleFinding(21, "app-2", 8, "api2"))
	findingRepository.SaveFinding(createASampleFinding(21, "app-2", 95, "api2"))

	scoreCards, _ := findingRepository.GetScoreCards(21)
	assert.Equal(t, 2, len(scoreCards))

	assert.Equal(t, "app-1", scoreCards[0].Application)
	assert.Equal(t, "app-2", scoreCards[1].Application)
	assert.Equal(t, 1, scoreCards[0].Info)
	assert.Equal(t, 1, scoreCards[0].Low)
	assert.Equal(t, 1, scoreCards[0].Medium)
	assert.Equal(t, 1, scoreCards[0].High)
	assert.Equal(t, 4, scoreCards[0].Total)
	assert.Equal(t, 0, scoreCards[1].Info)
	assert.Equal(t, 0, scoreCards[1].Low)
	assert.Equal(t, 0, scoreCards[1].Medium)
	assert.Equal(t, 2, scoreCards[1].High)
	assert.Equal(t, 2, scoreCards[1].Total)

}

func TestGetScoreCardDetails(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 0, "api1", "pattern1"))
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 0, "api2", "pattern1"))
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 0, "api2", "pattern2"))
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 8, "api2", "pattern2"))

	details, _ := findingRepository.GetScoreCardDetails(21, "app-1", "info")
	assert.Equal(t, 2, len(details))

	assert.Equal(t, "app-1", details[0].Application)
	assert.Equal(t, "app-1", details[1].Application)
	assert.Equal(t, "pattern1", details[0].Pattern)
	assert.Equal(t, "pattern2", details[1].Pattern)
	assert.Equal(t, 0, details[0].Effort)
	assert.Equal(t, 0, details[1].Effort)
	assert.Equal(t, 2, details[0].Count)
	assert.Equal(t, 1, details[1].Count)
	assert.Equal(t, 0, details[0].Total)
	assert.Equal(t, 0, details[1].Total)

	details, _ = findingRepository.GetScoreCardDetails(21, "app-1", "high")
	assert.Equal(t, 1, len(details))

	assert.Equal(t, "app-1", details[0].Application)
	assert.Equal(t, "pattern2", details[0].Pattern)
	assert.Equal(t, 8, details[0].Effort)
	assert.Equal(t, 1, details[0].Count)
	assert.Equal(t, 8, details[0].Total)
}

func TestGetAllScoreCardDetails(t *testing.T) {
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 0, "api1", "pattern1"))
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 0, "api2", "pattern1"))
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 0, "api2", "pattern2"))
	findingRepository.SaveFinding(createASampleFindingWithPattern(21, "app-1", 8, "api2", "pattern2"))

	details, _ := findingRepository.GetScoreCardDetails(21, "app-1", "all")
	assert.Equal(t, 3, len(details))

	assert.Equal(t, "app-1", details[0].Application)
	assert.Equal(t, "app-1", details[1].Application)
	assert.Equal(t, "pattern2", details[0].Pattern)
	assert.Equal(t, "pattern1", details[1].Pattern)
	assert.Equal(t, 8, details[0].Effort)
	assert.Equal(t, 0, details[1].Effort)
	assert.Equal(t, 0, details[2].Effort)
	assert.Equal(t, 1, details[0].Count)
	assert.Equal(t, 2, details[1].Count)
	assert.Equal(t, 1, details[2].Count)
	assert.Equal(t, 8, details[0].Total)
	assert.Equal(t, 0, details[1].Total)
}

func TestGetFindingsByTag(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 0, "api1", "pattern1", "api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 1, "chuck-e-cheese", "pattern1", "not-api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-2", 5, "api2", "pattern2", "api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-3", 8, "franks-and-beans", "pattern3", "not-api-as-well"))

	details, _ := findingRepository.GetFindingsByTag(21, "api")
	assert.Equal(t, 2, len(details))
	assert.Equal(t, "app-1", details[0].Application)
	assert.Equal(t, "app-2", details[1].Application)

	assert.Equal(t, "pattern1", details[0].Pattern)
	assert.Equal(t, "pattern2", details[1].Pattern)
	assert.Equal(t, "some advice", details[0].Advice)
	assert.Equal(t, 0, details[0].Effort)
	assert.Equal(t, 5, details[1].Effort)
}

func TestGetTagsForApp(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 0, "api1", "pattern1", "tag1"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 1, "chuck-e-cheese", "pattern1", "tag1"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 5, "api2", "pattern2", "tag2"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-3", 8, "franks-and-beans", "pattern3", "not-api-as-well"))

	details, _ := findingRepository.GetTagsForApp(21, "app-1")
	assert.Equal(t, 2, len(details))
	assert.Equal(t, "tag1", details[0])
	assert.Equal(t, "tag2", details[1])
}

func TestGetFindingsById(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)

	finding := createASampleFindingWithPatternAndTag(21, "app-1", 0, "api1", "pattern1", "api")
	findingRepository.SaveFinding(finding)

	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 0, "api2", "pattern4", "api-32"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-1", 1, "chuck-e-cheese", "pattern31", "not-api"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-2", 5, "api3", "pattern2", "api-11"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTag(21, "app-3", 8, "franks-and-beans", "pattern3", "not-api-as-well"))

	newFinding, _ := findingRepository.GetFinding(finding.ID)

	assert.Equal(t, finding.Application, newFinding.Application)
	assert.Equal(t, finding.Effort, newFinding.Effort)
	assert.Equal(t, finding.Category, newFinding.Category)

	assert.Equal(t, finding.Category, newFinding.Category)
	assert.Equal(t, finding.Tags[0].Value, newFinding.Tags[0].Value)
}

func TestGetFindingsByRule(t *testing.T) {
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	findingRepository := db.NewFindingRepository(database)
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(21, "app-1", 0, "api1", "pattern1", "api", "rule1"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(21, "app-1", 1, "chuck-e-cheese", "pattern1", "not-api", "rule2"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(21, "app-2", 5, "api2", "pattern2", "api", "rule1"))
	findingRepository.SaveFinding(createASampleFindingWithPatternAndTagAndRule(21, "app-3", 8, "franks-and-beans", "pattern3", "not-api-as-well", "rule1"))

	details, _ := findingRepository.GetFindingsByRule(21, "rule1")
	assert.Equal(t, 3, len(details))
	assert.Equal(t, "app-1", details[0].Application)
	assert.Equal(t, "api1", details[0].Category)
	assert.Equal(t, "app-2", details[1].Application)
	assert.Equal(t, "api2", details[1].Category)
	assert.Equal(t, "app-3", details[2].Application)
	assert.Equal(t, "franks-and-beans", details[2].Category)
}

func createASampleFinding(runId uint, domain string, score int, category string) *model.Finding {

	return createASampleFindingWithPattern(runId, domain, score, category, "")
}

func createASampleFindingWithPattern(runId uint, domain string, score int, category string, pattern string) *model.Finding {

	return createASampleFindingWithPatternAndTag(runId, domain, score, category, pattern, "")

}

func createASampleFindingWithPatternAndTag(runId uint, domain string, score int, category string, pattern string, tag string) *model.Finding {

	return createASampleFindingWithPatternAndTagAndRule(runId, domain, score, category, pattern, tag, "")
}

func createASampleFindingWithPatternAndTagAndRule(runId uint, domain string, score int, category string, pattern string, tag string, rule string) *model.Finding {
	someDate, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	if pattern == "" {
		pattern = "default pattern"
	}

	if tag == "" {
		tag = "default tag"
	}

	if rule == "" {
		rule = "default rule"
	}

	newFinding := &model.Finding{
		Rule:        rule,
		Pattern:     pattern,
		Advice:      "some advice",
		Category:    category,
		CreatedAt:   someDate,
		Criticality: "HIGH",
		Application: domain,
		Ext:         "some ext",
		Filename:    "myfilename",
		Fqn:         "myfqn",
		Line:        100,
		Recipes: []model.FindingRecipe{
			{
				URI: "someuri1",
			},
			{
				URI: "someuri2",
			},
		},
		RunID:     runId,
		Effort:    score,
		UpdatedAt: someDate,
		Value:     "some value",
	}

	newFinding.AddTag(tag)

	return newFinding
}
