/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"encoding/json"
	"csa-app/model"
	"csa-app/util"
	"gopkg.in/yaml.v2"
	"strings"
)

func BootStrapRules(rulesdir string, templatesDir string) (rules []model.Rule, err error) {
	rules, err = GetAllResourcesBasedRules(rulesdir)

	if err != nil {
		return rules, err
	}

	if err := validateUniqueRuleNames(&rules); err != nil {
		return rules, err
	}

	if len(rules) == 0 {
		fmt.Printf("No Rules Found in rules-dir[%s]\n", rulesdir)
		os.Exit(1)
	}

	fmt.Printf("Bootstraping [%d] Rules\n", len(rules))

	rulesTemplate := getTemplate(templatesDir, util.RULE_BOOTSTRAP_TEMPLATE)

	path := "model" + util.PathSeparator
	name := "Bootstrap"

	fmt.Printf("Name: %s, Path: %s\n", name, path)

	file := util.CreateFile(name, "go", path)

	if err = rulesTemplate.Execute(file, rules); err != nil {
		return rules, err
	}

	return rules, nil
}

func GetAllResourcesBasedRules(rulesdir string) ([]model.Rule, error) {
	//Get Files...Only json and yaml are supported at this time!
	fileUtil := util.NewFileUtil()
	files := fileUtil.GetFileList(rulesdir, "(json|yaml|yml)")

	var newRules []model.Rule

	var errs []error

	if len(files) > 0 {
		for _, file := range files {
			if *util.Verbose {
				fmt.Printf("Importing Rule File [%s]...", file.Name)
			}
			var decoder util.FileDecoder
			reader, err := os.Open(file.FQN)
			if err != nil {
				errs = append(errs, fmt.Errorf("Error reading file: %s. Details: %v\n", file.Name, err))
				continue
			}

			if strings.HasSuffix(file.Name, util.JSON) {
				decoder = json.NewDecoder(reader)
			} else {
				decoder = yaml.NewDecoder(reader)
			}

			rule, err := extractRule(decoder)

			if rule.Name == "" {
				errs = append(errs, fmt.Errorf("Error reading file: %s. Details: %v\n", file.Name, err))
				continue
			}

			if err != nil {
				errs = append(errs, fmt.Errorf("Rule [%s] is invalid! Reason: %s\n", rule.Name, err.Error()))
				continue
			}

			newRules = append(newRules, rule)
			reader.Close()
		}
		if len(errs) > 0 {

			var msg strings.Builder

			msg.WriteString("\t[\n")
			for _, e := range errs {
				msg.WriteString(fmt.Sprintf("\t\t%s", e.Error()))
			}
			msg.WriteString("\t]\n")
			return nil, fmt.Errorf("errors =>\n%s", msg.String())
		}
	} else {
		fmt.Printf("Found No Rule files for import @ [%s]\n", rulesdir)
	}
	return newRules, nil
}

func validateUniqueRuleNames(rules *[]model.Rule) error {
	var m map[string]bool
	m = make(map[string]bool)

	var errs []error
	for _, rule := range *rules {
		if _, ok := m[rule.Name]; ok {
			errs = append(errs, fmt.Errorf("Conflicting rule name : %s\n%v", rule.Name, rule))
		}
		m[rule.Name] = true
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

func DeletePatterns(patterns []model.Pattern) {
	for _, pattern := range patterns {
		CheckDBError(false,
			"Import Rules",
			fmt.Sprintf("Failed Deleting Pattern [%s]", pattern.Value),
			database.Delete(&pattern).Error)
	}
}

func DeleteRecipes(recipes []model.Recipe) {

	for _, recipe := range recipes {
		CheckDBError(false,
			"Import Rules",
			fmt.Sprintf("Failed Deleting Recipe [%s]", recipe.URI),
			database.Delete(&recipe).Error)
	}
}

func DeleteTags(tags []model.Tag) {

	for _, tag := range tags {
		CheckDBError(false,
			"Import Rules",
			fmt.Sprintf("Failed Deleting Tag [%s]", tag.Value),
			database.Delete(&tag).Error)
	}
}

func DeleteExcludePatterns(excludePatterns []model.ExcludePattern) {

	for _, excludePattern := range excludePatterns {
		CheckDBError(false,
			"Import Rules",
			fmt.Sprintf("Failed Deleting Exclude Pattern [%s]", excludePattern.Value),
			database.Delete(&excludePattern).Error)
	}
}

func getTemplate(templatesDir string, filename string) *template.Template {
	name := templatesDir + util.PathSeparator + filename

	funcMap := template.FuncMap{
		"now": time.Now,
	}

	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error reading rules template [%s]! Details: %v\n", name, err)
		os.Exit(1)
	}

	return template.Must(template.New("Rules").Funcs(funcMap).Parse(string(content)))

}

func extractRule(decoder util.FileDecoder) (model.Rule, error) {
	var rule model.Rule
	err := decoder.Decode(&rule)

	if err == nil {
		_, err = rule.IsValid()
	}

	return rule, err
}
