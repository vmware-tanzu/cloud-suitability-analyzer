/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"csa-app/db"
	"csa-app/db/test_support"
	"csa-app/model"

	"github.com/stretchr/testify/assert"
)

func TestRulesRepository(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	var run model.Run

	var ruleRepository db.RuleRepository
	ruleRepository = db.NewRuleRepository(database)
	binRepo := db.NewBinRepository(database)
	scoreRepo := db.NewScoringRepository(database)

	rules, _ := ruleRepository.GetRules()

	assert.Equal(t, 0, len(rules))
	db.PopulateInitialData(&run, ruleRepository, binRepo, scoreRepo, database)

	rules, _ = ruleRepository.GetRules()
	assert.True(t, len(rules) > 10)
}

func TestGetRulesRestricted(t *testing.T) {

	database, dir := createDb()
	defer os.RemoveAll(dir)

	ruleCnt := 10
	ruleRepository := createNewRepo(ruleCnt, database)
	rules, _ := ruleRepository.GetRules()
	assert.Equal(t, ruleCnt, len(rules))

	rules, _ = ruleRepository.GetRulesForRunRestricted(&model.Run{ID: 1}, nil, false)

	assert.Equal(t, ruleCnt, len(rules))

	tags := []string{"1"}
	rules, _ = ruleRepository.GetRulesForRunRestricted(&model.Run{ID: 1}, tags, false)
	assert.Equal(t, 1, len(rules))

	tags = []string{"1", "2"}
	rules, _ = ruleRepository.GetRulesForRunRestricted(&model.Run{ID: 1}, tags, false)
	assert.Equal(t, 2, len(rules))

	tags = []string{"1"}
	rules, _ = ruleRepository.GetRulesForRunRestricted(&model.Run{ID: 1}, tags, true)
	assert.Equal(t, 9, len(rules))

	tags = []string{"1", "2"}
	rules, _ = ruleRepository.GetRulesForRunRestricted(&model.Run{ID: 1}, tags, true)
	assert.Equal(t, 8, len(rules))

}

func createNewRepo(numRulesToPopulate int, database *gorm.DB) db.RuleRepository {
	ruleRepository := db.NewRuleRepository(database)

	for i := 0; i < numRulesToPopulate; i++ {
		rule := model.Rule{
			Name: fmt.Sprintf("Rule-%d", i),
		}

		rule.Tags = append(rule.Tags, model.Tag{Value: fmt.Sprintf("%d", i)})
		ruleRepository.SaveRule(rule)
	}
	return ruleRepository
}

func createDb() (database *gorm.DB, dir string) {
	_, dir, database, err := db_test_support.OpenTestDb()
	if err != nil {
		log.Fatal(err)
	}

	return
}
