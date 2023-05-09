package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"test/csa-app/csa"
	"test/csa-app/model"
	"test/csa-app/util"
	"test/mat"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type Input struct {
	Categories []Category `yaml:"category"`
}

type Category struct {
	RulesDirectory string  `yaml:"rules-directory"`
	Tests          []Tests `yaml:"tests"`
}

type Tests struct {
	Name         string `yaml:"name"`
	RuleName     string `yaml:"rule-name"`
	TestFileName string `yaml:"test-filename"`
	Assert       bool   `yaml:"assert"`
	AssertCount  int    `yaml:"assert-count"`
	AssertValue  string `yaml:"assert-value"`
}

var testCount = 0
var ruleCount = 0
var ruleList []model.Rule
var rulesCovered = make(map[string]struct{})

// Code Samples
var sampleBaseDir = ""

/*************************************
// Test Rule is Valid / Compiles
// Test Rule targets the right file types
// Test Match
// Test Finding text
*************************************/

func TestRules(t *testing.T) {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	sampleBaseDir = cwd + "/test_samples"

	directoryPath := cwd + "/../../test-cases"

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println("Processing rules in" + directoryPath + "/" + file.Name())
		// Read the YAML file
		file, err := ioutil.ReadFile(directoryPath + "/" + file.Name())
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		// Parse the YAML data
		var fileInput Input
		err = yaml.Unmarshal(file, &fileInput)
		if err != nil {
			t.Fatalf("Failed to parse YAML: %v", err)
		}

		for _, category := range fileInput.Categories {
			rulesDirectory := cwd + "/../../" + category.RulesDirectory
			var ruleSet = csa.Setup(rulesDirectory)
			ruleCount += len(ruleSet)
			ruleList = append(ruleList, ruleSet...)
			for _, rule := range category.Tests {

				testRuleByName(t, rule.Name, csa.RuleByName(t, ruleSet, rule.RuleName), rule.TestFileName, rule.Assert, rule.AssertCount, rule.AssertValue)
			}
		}
	}
}

func TestCoverage(t *testing.T) {
	t.Log("\033[43m\033[30m  #Test:", testCount, "\033[0m", "\033[0m")
	t.Log("\033[43m\033[30m  Rules Covered with Test:", len(rulesCovered), "/", ruleCount, "\033[0m", "\033[0m")
}

func TestExportForMat(t *testing.T) {
	mat.Export(ruleList)

	for _, r := range ruleList {
		if rulesCovered[r.Name] != struct{}{} {
			fmt.Println("// Untested Rule: ")
			fmt.Println(r.Name)
		}
	}
}

// Easily test one rule against targeted file

func testRuleByName(t *testing.T, testName string, rule model.Rule, targetFilePath string, expectedMatch bool, expectedMatchCount int, expectedValue string) {
	var exists = struct{}{}
	rulesCovered[rule.Name] = exists
	//t.Log("Default Pattern: " + rule.DefaultPattern)
	testCount++
	t.Log("\033[94m[Test: " + testName + "] Rule: \033[33m" + rule.Name + "\033[0m || File: " + targetFilePath + "\033[0m")
	isValid, err := rule.IsValid()
	assert.Equal(t, err, nil, "The rule should not fail during validation")
	assert.Equal(t, isValid, true, "The rule should be valid")

	// Loading target file
	f := util.GetFile(sampleBaseDir + "/" + targetFilePath)

	// Does the rule apply ?
	applies := rule.Applies(f.GetCleanedExt(), f.Name)
	assert.Equal(t, applies, true, "\033[91mThe rule does not apply to this file type\033[0m")

	//Is there a match
	fileFindings, value := csa.AnalyzeFile(rule, f)

	assert.Equal(t, expectedMatch, (fileFindings > 0), "\033[91mThe match result was different than expected\033[0m")
	assert.Equal(t, expectedMatchCount, fileFindings, "\033[91mThe number of matches was different than expected\033[0m")

	if expectedValue != "null" && expectedValue != "" {
		assert.Equal(t, expectedValue, value, "\033[91mThe collected finding text did not match expectation\033[0m")
	}
}
