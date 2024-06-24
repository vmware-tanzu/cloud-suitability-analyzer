package test

import (
	"bufio"
	"bytes"
	"csa-app/model"
	"csa-app/util"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/antchfx/xmlquery"
	"gopkg.in/yaml.v3"
)

// -----------------------------------------------------------------------
// Loads rules from files
// -----------------------------------------------------------------------
func Setup(ruleDir string) (rules []model.Rule) {
	fileUtil := util.NewFileUtil()
	files := fileUtil.GetFileList(ruleDir, "(json|yaml|yml)")
	if len(files) > 0 {
		for _, file := range files {
			var decoder util.FileDecoder
			reader, err := os.Open(file.FQN)
			if err != nil {
				fmt.Printf("Error reading file. Details: %v\n", err)
				continue
			}
			decoder = yaml.NewDecoder(reader)
			rule, err := extractRule(decoder)

			if rule.Name == "" {
				fmt.Printf("Empty rule name")
				continue
			}
			rule.CompilePatterns()
			rules = append(rules, rule)

			reader.Close()
		}
	}
	return rules
}

// -----------------------------------------------------------------------
// Extract Rules YAMLs
// -----------------------------------------------------------------------
func extractRule(decoder util.FileDecoder) (model.Rule, error) {
	var rule model.Rule
	err := decoder.Decode(&rule)

	if err == nil {
		_, err = rule.IsValid()
	}

	return rule, err
}

// -----------------------------------------------------------------------
// Get rule by name
// -----------------------------------------------------------------------
func RuleByName(t *testing.T, rules []model.Rule, ruleName string) (rule model.Rule) {
	for _, r := range rules {
		if ruleName == r.Name {
			rule = r
		}
	}
	return rule
}

// -----------------------------------------------------------------------
// Light weight version of AppAnalyzer.analyzeFile
// -----------------------------------------------------------------------
func AnalyzeFile(rule model.Rule, file util.FileInfo) (findingCnt int, value string) {

	var rulesForFile []model.Rule
	var rulesUsed []string

	findings := 0

	hasContentRules := false

	if rule.Applies(file.GetCleanedExt(), file.Name) {
		rulesUsed = append(rulesUsed, rule.Name)

		if rule.Target == model.LINE_TARGET {
			rulesForFile = append(rulesForFile, rule)
			//wasAnalyzed = true
		} else if rule.Target == model.CONTENTS_TARGET {
			rulesForFile = append(rulesForFile, rule)
			hasContentRules = true
		} else {
			findings += processPatterns(&file, 0, filepath.Base(file.Name), rule)
		}

	} else {
		fmt.Println("Rule does not apply for file " + rule.Name + " " + rule.Name + " " + file.Name + " " + file.Ext + " " + file.FQN)
	}

	fileFindings, value := processFile(&file, rulesForFile, hasContentRules)
	fileFindings += findings

	return fileFindings, value
}

// -----------------------------------------------------------------------
// Light weight version of FileProcessor.processFile
// -----------------------------------------------------------------------
func processFile(file *util.FileInfo, rules []model.Rule, hasContentRules bool) (findingCnt int, value string) {
	if len(rules) > 0 {

		line := 0
		sloc := 0
		inFile, err := os.Open(file.FQN)
		scanner := bufio.NewScanner(inFile)
		buf := make([]byte, util.MAX_LINE_BUFFER_SIZE)
		scanner.Buffer(buf, util.MAX_LINE_BUFFER_SIZE)

		//midComment := false
		process := true
		if err != nil {

		}

		contents := ""

		for scanner.Scan() {

			//Count line regardless if comment
			line++
			curLine := scanner.Text()

			//Only accumulate if we have content rules to process
			if hasContentRules {
				contents += curLine + "\n"
			}

			if process && len(strings.TrimSpace(curLine)) > 0 {
				sloc++
				for i := range rules {

					if rules[i].Target == model.LINE_TARGET {
						nbFindings := processPatterns(file, line, curLine, rules[i])
						if nbFindings > 0 {
							if rules[i].Excludepatterns != nil {

								for j := range rules[i].Excludepatterns {
									regex := regexp.MustCompile(rules[i].Excludepatterns[j].Value)
									findingExclude := regex.MatchString(curLine)
									if findingExclude == true {
										fmt.Println("Exclude Patterns => " + curLine)
										nbFindings = 0
									}
								}
							}
						}

						if nbFindings > 0 {
							value += curLine
						}
						findingCnt += nbFindings
					}
				}
			}
		}

		//Only reprocess if necessary
		if hasContentRules {
			for i := range rules {
				if rules[i].Target == model.CONTENTS_TARGET {
					findingCnt += processPatterns(file, 0, contents, rules[i])
					if findingCnt > 0 {
						value += contents
					}
				}
			}
		}
	}

	return findingCnt, value
}

// -----------------------------------------------------------------------
// Light weight version of FileProcessor.processPatterns
// -----------------------------------------------------------------------
func processPatterns(file *util.FileInfo, line int, target string, rule model.Rule) int {

	xmlDocs := make(map[string](*xmlquery.Node))
	yamlDocs := make(map[string](*yaml.Node))

	findings := 0

	cnt := int64(0)
	pcnt := int64(0)

	//Process Patterns against line
	for i := range rule.Patterns {

		if rule.Patterns[i].Type == model.PLUGIN_MATCH_TYPE {
			// DO NOTHING
		} else {

			matchFunc := func() (bool, string) {
				return rule.Patterns[i].Match(target)
			}

			if rule.Patterns[i].Type == model.XPATH_MATCH_TYPE {
				if xmlDocs[file.FQN] == nil {
					if rawData, err := os.ReadFile(file.FQN); err == nil {
						if xml, err := xmlquery.Parse(bytes.NewReader(rawData)); err == nil {
							xmlDocs[file.FQN] = xml
						}
					}
				}

				matchFunc = func() (bool, string) {
					return rule.Patterns[i].MatchXml(xmlDocs[file.FQN])
				}
			} else if rule.Patterns[i].Type == model.YAMLPATH_MATCH_TYPE {

				if yamlDocs[file.FQN] == nil {
					if rawData, err := os.ReadFile(file.FQN); err == nil {
						var node yaml.Node
						err = yaml.Unmarshal(rawData, &node)

						if err == nil {
							yamlDocs[file.FQN] = &node
						}
					}
				}

				matchFunc = func() (bool, string) {
					return rule.Patterns[i].MatchYaml(yamlDocs[file.FQN])
				}
			}

			ok, result := matchFunc()

			if ok && !rule.Negative {
				if len(result) > 0 {
					target = regexp.MustCompile(`\r?\n`).ReplaceAllString(result, " ")
				}

				findings++
				cnt++
			} else if !ok && rule.Negative {
				findings++
				cnt++
			}

		}

		pcnt++

	} //End Pattern For Loop

	return findings
}
