/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db_test

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"csa-app/db"
	"csa-app/db/test_support"
	"csa-app/model"
	"github.com/stretchr/testify/assert"
)

func TestSaveAndRetrieveBins(t *testing.T) {
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	PopulateRules(database, t)

	repo := db.NewBinRepository(database)

	andTags := []string{"tag1"}
	orTags := []string{"tag2"}
	xTags := []string{"tag3"}

	bin := CreateNewBin("bin1", andTags, orTags, xTags)

	err = repo.SaveBin(bin)

	assert.Nil(t, err)

	theBin, err := repo.GetBinByName("bin1")

	assert.Nil(t, err)

	assert.Equal(t, bin.Name, theBin.Name)
	assert.Equal(t, 3, len(theBin.Tags))
	assert.Equal(t, bin.GetTag("tag1").Type, model.AND)
	assert.Equal(t, bin.GetTag("tag2").Type, model.OR)
	assert.Equal(t, bin.GetTag("tag3").Type, model.EXCLUDE)

}

func PopulateRules(database *gorm.DB, t *testing.T) {
	var ruleRepository db.RuleRepository
	ruleRepository = db.NewRuleRepository(database)

	rules, _ := ruleRepository.GetRules()

	assert.Equal(t, 0, len(rules))

	rule := model.Rule{}
	rule.Name = "test1"
	rule.AddTag("tag1")
	rule.AddTag("tag2")
	rule.AddTag("tag3")
	ruleRepository.SaveRule(rule)
}

func CreateNewBin(name string, andTags []string, orTags []string, excludeTags []string) *model.Bin {

	bin := model.Bin{Name: name}

	for _, tag := range andTags {
		bin.AddTag(tag, "AND")
	}

	for _, tag := range orTags {
		bin.AddTag(tag, "OR")
	}

	for _, tag := range excludeTags {
		bin.AddTag(tag, "EXCLUDE")
	}

	return &bin

}
