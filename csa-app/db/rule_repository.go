/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"csa-app/model"
	"csa-app/util"
)

type RuleRepository interface {
	SaveRules(rules []model.Rule) ([]model.Rule, error)
	SaveRule(rule model.Rule) (model.Rule, error)
	GetRuleByName(name string) (model.Rule, error)
	GetRules() ([]model.Rule, error)
	GetRulesForRun(run *model.Run) ([]model.Rule, error)
	GetRulesForRunRestricted(run *model.Run, tags []string, excluded bool) ([]model.Rule, error)
	ExportRules()
	ImportRules()
	LoadRules()
	DeleteRule(ruleName string) error
	DeleteAllRules() error
	ValidateRule(filename string, run *model.Run)
	GetRuleMetrics(runId uint) ([]model.RuleMetric, error)
}

func NewRuleRepository(db *gorm.DB) RuleRepository {
	return &OrmRepository{
		dbconn:   db,
		fileUtil: util.NewFileUtil(),
	}
}

func NewRuleRepositoryForRun(run *model.Run) RuleRepository {
	return &OrmRepository{
		dbconn:   run.DB,
		fileUtil: run.FileUtil,
	}
}

func (ruleRepository *OrmRepository) GetRules() ([]model.Rule, error) {
	var rules []model.Rule
	resp := ruleRepository.dbconn.Preload("Patterns").Preload("Recipes").Preload("Tags").Preload("Excludepatterns").Find(&rules)
	return rules, resp.Error
}

func (ruleRepository *OrmRepository) SaveRules(rules []model.Rule) ([]model.Rule, error) {
	var errs []error
	var savedRules []model.Rule
	for _, rule := range rules {
		savedRule, err := ruleRepository.SaveRule(rule)
		savedRules = append(savedRules, savedRule)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return savedRules, fmt.Errorf("errors %v", errs)
	}
	return savedRules, nil
}

func (ruleRepository *OrmRepository) SaveRule(rule model.Rule) (model.Rule, error) {
	if *util.Verbose {
		fmt.Printf("Saving [%s] Rule [%s] Patterns [%d]\n", rule.Target, rule.Name, len(rule.Patterns))
	}

	res := database.Save(&rule)

	return rule, res.Error
}

func (ruleRepository *OrmRepository) GetRuleByName(name string) (model.Rule, error) {
	var rule model.Rule
	res := ruleRepository.dbconn.Where(&model.Rule{Name: name}).Preload("Patterns").Preload("Recipes").Preload("Tags").Find(&rule)
	return rule, res.Error
}

func (ruleRepository *OrmRepository) GetRulesForRun(run *model.Run) ([]model.Rule, error) {
	rules, err := ruleRepository.GetRules()
	if err != nil {
		return nil, err
	}
	for i := range rules {
		rules[i].CompilePatterns()
		rules[i].Metric = &model.RuleMetric{Rule: rules[i].Name, RunID: run.ID, RuleCriticality: rules[i].Criticality}
	}

	return rules, nil
}

func (ruleRepository *OrmRepository) GetRulesForRunRestricted(run *model.Run, tags []string, excluded bool) ([]model.Rule, error) {
	rules, err := ruleRepository.GetRules()
	var restrictedRules []model.Rule
	if err != nil {
		return nil, err
	}

	if len(tags) > 0 {
		if excluded {
			for i := range rules {
				shouldAdd := true
				for _, tag := range tags {
					//Have to check against every tag!
					if rules[i].HasTag(tag) {
						shouldAdd = false
					}
				}
				if shouldAdd {
					restrictedRules = append(restrictedRules, rules[i])
				}
			}
		} else {
			for i := range rules {
				for _, tag := range tags {
					if rules[i].HasTag(tag) {
						restrictedRules = append(restrictedRules, rules[i])
						//Only add rule for first rule once
						break
					}
				}
			}
		}
	} else {
		restrictedRules = append(restrictedRules, rules...)
	}

	//Compile the remaining rules and add metrics!
	for i := range restrictedRules {
		restrictedRules[i].CompilePatterns()
		restrictedRules[i].Metric = &model.RuleMetric{Rule: restrictedRules[i].Name, RunID: run.ID, RuleCriticality: restrictedRules[i].Criticality}
	}

	return restrictedRules, nil
}

func (ruleRepository *OrmRepository) ExportRules() {
	var extension = util.YAML
	if *util.ExportJson {
		extension = util.JSON
	}

	exportCnt := 0

	rulesDir := model.GetRulesDir(false)

	if *util.ExportRuleName != "" {
		rule, err := ruleRepository.GetRuleByName(*util.ExportRuleName)
		if err != nil && rule.Name != "" {
			util.WriteStructToFile(rule, rule.Name, rulesDir, extension, true)
			exportCnt++
		} else {
			fmt.Printf("Rule [%s] not found for export!\n", *util.ExportRuleName)
		}
	} else {

		rules, err := ruleRepository.GetRules()

		if err == nil {
			if *util.ExportSingleFile {
				currentTime := time.Now()
				filename := fmt.Sprintf("csa-rules-%d%02d%02d-%02d%02d%02d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())
				exportCnt = util.WriteStructsToFile(rules, filename, rulesDir, extension, true)
			} else {
				for _, rule := range rules {

					util.WriteStructToFile(rule, rule.Name, rulesDir, extension, true)
					exportCnt++

				}
			}
		} else {
			fmt.Printf("Error retrieving rules for export! Details: %v", err)
		}
	}

	fmt.Printf("Successfully exported [%d] rules @ [%s]", exportCnt, rulesDir)
}

func (RuleRepository *OrmRepository) ValidateRule(filename string, run *model.Run) {
	fmt.Printf("Validating Rule File [%s]\n\n", filename)

	var decoder util.FileDecoder
	reader, err := util.OpenFileAtPath(filename, run.Homepath)
	if err != nil {
		util.App.Fatalf("Error reading file. Details: %v\n", err)
	}

	if strings.HasSuffix(filepath.Ext(filename), util.JSON) {
		decoder = json.NewDecoder(reader)
	} else {
		decoder = yaml.NewDecoder(reader)
	}

	rule, err := extractRule(decoder)

	if err == nil {
		fmt.Printf("Rule [%s] is VALID!\n\tTarget: %s\n\tType: %s\n\tFileType: %s\n\tPattern Cnt: %d\n",
			rule.Name, rule.Target, rule.Type, rule.FileType, len(rule.Patterns))
		fmt.Print("\nUnmarshalled Rule Dump =>\n\n")

		spew.Config.DisablePointerAddresses = true
		spew.Config.DisableMethods = true
		spew.Config.Indent = "\t"
		spew.Config.DisableCapacities = true
		spew.Config.DisablePointerMethods = true
		spew.Config.MaxDepth = 3
		spew.Config.SortKeys = true

		spew.Dump(rule)
	} else {
		util.App.Errorf("%v", err)
	}

}

func (ruleRepository *OrmRepository) ImportRules() {
	dir := model.GetRulesDir(true)
	//Get Files...Only json and yaml are supported at this time!
	files := ruleRepository.fileUtil.GetFileList(dir, "(json|yaml|yml)")

	importOne := *util.ImportRuleName != ""
	var newRules []model.Rule
	importCnt := 0

	if len(files) > 0 {
		for _, file := range files {

			if *util.Verbose {
				fmt.Printf("Importing Rule File [%s]...", file.Name)
			}
			var decoder util.FileDecoder
			reader, err := os.Open(file.FQN)
			if err != nil {
				fmt.Printf("Error reading file. Details: %v\n", err)
				continue
			}

			if strings.HasSuffix(file.Name, util.JSON) {
				decoder = json.NewDecoder(reader)
			} else {
				decoder = yaml.NewDecoder(reader)
			}

			stop := false
			cnt := 0
			for !stop {
				cnt, stop = ruleRepository.unMarshalAndSaveRule(decoder, &newRules)
				importCnt += cnt
			}
			//Close the file!
			reader.Close()
		}

		if *util.ReplaceRulesFlag && len(newRules) > 0 {
			replacedCnt := 1
			if !importOne {
				rules, _ := ruleRepository.GetRules()
				replacedCnt = len(rules)
			}
			fmt.Printf("Replacing [%d] existing rules with [%d] new ones!\n", replacedCnt, len(newRules))
			ruleRepository.DeleteAllRules()
			ruleRepository.SaveRules(newRules)
		}

		if importCnt == 0 && importOne {
			fmt.Printf("Rule [%s] not found for import @ [%s]\n", *util.ImportRuleName, dir)
		} else {
			fmt.Printf("Successfully imported [%d] rule(s) found @[%s]\n", importCnt, dir)
		}

	} else {
		fmt.Printf("Found No Rule files for import @ [%s]\n", dir)
	}
}

func (ruleRepository *OrmRepository) DeleteAllRules() error {
	return database.Delete(&model.Rule{}).Error
}

func (ruleRepository *OrmRepository) LoadRules() {
	rules, _ := ruleRepository.GetRules()
	if len(rules) < 1 {

		ruleRepository.ImportRules()

		//Check again to see if we found any on the file system!
		rules, _ := ruleRepository.GetRules()
		if len(rules) < 1 {
			fmt.Printf("Loading rules from Bootstap...\n")
			cnt := 0
			for _, rule := range model.BootstrapRules() {			
				_, err := ruleRepository.SaveRule(rule)
				if err != nil {
					util.App.Fatalf("Error saving rule [%s]! Details: %s", rule.Name, err.Error())
				}
				cnt++
			}

			fmt.Printf("Created/Converted [%d] rules\n", cnt)
		}
	}
}

func (ruleRepository *OrmRepository) DeleteRule(ruleName string) error {
	existingRule, _ := ruleRepository.GetRuleByName(ruleName)

	if existingRule.Name == *util.DeleteRulesName {
		fmt.Printf("Deleting rule [%s]...", existingRule.Name)

		CheckDBError(true,
			"Delete Rule",
			fmt.Sprintf("Failed Deleting Rule [%s]", existingRule.Name),
			database.Delete(&existingRule).Error)
		fmt.Print("done!\n")
	}
	return database.Delete(existingRule).Error
}

func (ruleRepository *OrmRepository) GetRuleMetrics(runId uint) ([]model.RuleMetric, error) {
	metrics := []model.RuleMetric{}
	resp := ruleRepository.dbconn.Where(model.RuleMetric{RunID: runId}).Find(&metrics)
	return metrics, resp.Error
}

/*
PRIVATE API -------------------------------------------------------------------------------------------------------------
*/

func (ruleRepository *OrmRepository) unMarshalAndSaveRule(decoder util.FileDecoder, newRules *[]model.Rule) (cnt int, stop bool) {

	importOne := *util.ImportRuleName != ""

	rule, err := extractRule(decoder)

	if err != nil {
		if err != io.EOF {
			fmt.Printf("Error unmarshalling/validating rule. Details: %v\n", err)
		}

		stop = true

	} else if !importOne || *util.ImportRuleName == rule.Name {
		cnt = 1
		if *util.ReplaceRulesFlag {
			*newRules = append(*newRules, rule)
		} else {
			//Check to see if rule already exists!
			existingRule, _ := ruleRepository.GetRuleByName(rule.Name)
			if existingRule.Name == rule.Name {
				if *util.Verbose {
					fmt.Printf("Rule [%s] exists! Updating!", rule.Name)
				}
				deletedPatterns, deletedRecipes, deletedTags := existingRule.UpdateRule(rule)
				DeletePatterns(deletedPatterns)
				DeleteRecipes(deletedRecipes)
				DeleteTags(deletedTags)
				ruleRepository.SaveRule(existingRule)
			} else {
				if *util.Verbose {
					fmt.Printf("Rule [%s] does not exist! Creating!", rule.Name)
				}
				ruleRepository.SaveRule(rule)
			}
		}

		if *util.Verbose {
			fmt.Print("Done!\n")
		}
		if importOne {
			stop = true
		}

	}

	return
}
