package csa

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"test/csa-app/model"
	"test/csa-app/util"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/antchfx/xmlquery"
)

//-----------------------------------------------------------------------
//Setup functions below
//-----------------------------------------------------------------------
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
			//fmt.Println("Rule Name => ", rule.Name)
			//fmt.Println("Target => ", rule.Target)

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

func extractRule(decoder util.FileDecoder) (model.Rule, error) {
	var rule model.Rule
	err := decoder.Decode(&rule)

	if err == nil {
		_, err = rule.IsValid()
	}

	return rule, err
}

func RuleByName(t *testing.T, rules []model.Rule, ruleName string) (rule model.Rule) {
	for _, r := range rules {
		if ruleName == r.Name {
			rule = r
		}
	}
	return rule
}

func AnalyzeFile(rule model.Rule, file util.FileInfo) (findingCnt int, value string) {

	var rulesForFile []model.Rule
	var rulesUsed []string

	findings := 0
	//wasAnalyzed := false
	//fileNameAnalyzed := false
	hasContentRules := false

	if rule.Applies(file.GetCleanedExt(), file.Name) {
		rulesUsed = append(rulesUsed, rule.Name)

		//Rule applies to this file!
		//if *util.Verbose {
		//fmt.Println("Analyzing Rule applies to file \n"+file.Name, rule.Name, file.Name, file.Ext, file.FQN)
		//}

		if rule.Target == model.LINE_TARGET {
			rulesForFile = append(rulesForFile, rule)
			//wasAnalyzed = true
		} else if rule.Target == model.CONTENTS_TARGET {
			rulesForFile = append(rulesForFile, rule)
			hasContentRules = true
			//wasAnalyzed = true
		} else {
			//Check for name hit!
			//fmt.Sprintf("Check for name hit!")file, 0, contents, rules[i]
			findings += processPatterns(&file, 0, filepath.Base(file.Name), rule)

			//fileNameAnalyzed = true
		}

	} else {
		fmt.Println("Rule does not apply for file " + rule.Name + " " + rule.Name + " " + file.Name + " " + file.Ext + " " + file.FQN)
	}

	//fmt.Printf("wasAnalyzed %v", wasAnalyzed)
	//fmt.Printf("fileNameAnalyzed %v", fileNameAnalyzed)
	//fmt.Printf("hasContentRules %v", hasContentRules)

	fileFindings, value := processFile(&file, rulesForFile, hasContentRules)
	fileFindings += findings
	//fmt.Printf("Target ========> %v", rule.Target)
	//fmt.Printf("fileFindings %v", fileFindings)
	//fmt.Printf("value *************=> %v", value)

	return fileFindings, value
}

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
									//fmt.Println(curLine)
									//fmt.Println(findingExclude)
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
						//if *util.Verbose {
						//util.WriteLog("Analyzing", "### Rule: %s Hit: %d times on File: %s Line: %d ###\n", rules[i].Name, findingCnt, file.Name, line)
						//}
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
					//if *util.Verbose {
					//util.WriteLog("Analyzing", "### Rule: %s Hit: %d times on File: %s  ###\n", rules[i].Name, findingCnt, file.Name)
					//}
				}
			}
		}

		//Create an info finding for each file with SLOC info!
		/*fileFinding := model.Finding{
		RunID:       run.ID,
		Filename:    file.Name,
		Fqn:         file.FQN,
		Ext:         file.Ext,
		Category:    model.SLOC_CATEGORY,
		Pattern:     model.FILE_SLOC_PATTERN,
		Value:       fmt.Sprintf("%d", sloc),
		Effort:      0,
		Readiness:   0,
		Criticality: "none",
		Application: file.Dir}*/

		//fileFinding.AddTag(model.INFO_FINDING)
		//fileFinding.AddTag(model.SLOC_CATEGORY)

		//Send "file" finding to save worker
		//output <- fileFinding

		//run.AddFindings(1)
	}

	return findingCnt, value
}

func processPatterns(file *util.FileInfo, line int, target string, rule model.Rule) int {
	//if *util.Verbose {
	//fmt.Printf("Rule [%s|%s] checking against Value [%s]\n", rule.Name, rule.FileType, target)
	//}
	xmlDocs := make(map[string](*xmlquery.Node))
	yamlDocs := make(map[string](*yaml.Node))

	findings := 0

	cnt := int64(0)
	pcnt := int64(0)

	//Process Patterns against line
	for i := range rule.Patterns {

		//if *util.Verbose {
		//fmt.Printf("\tRule [%s|%s] Pattern [%s] checking against Value [%s]\n", rule.Name, rule.FileType, rule.Patterns[i].Pattern, target)
		//}

		if rule.Patterns[i].Type == model.PLUGIN_MATCH_TYPE {
			//csaService.RunPlugin(run, app, file, line, target, rule, rule.Patterns[i], output)
			// RunPlugin(run, app, file, line, target, rule, pattern, output, Value, file.FQN)
		} else {
			//rule.Patterns[i].Compiler(&rule)

			matchFunc := func() (bool, string) {
				return rule.Patterns[i].Match(target)
			}

			if rule.Patterns[i].Type == model.XPATH_MATCH_TYPE {
				//csaService.xmlMux.Lock()

				if xmlDocs[file.FQN] == nil {
					if rawData, err := ioutil.ReadFile(file.FQN); err == nil {
						if xml, err := xmlquery.Parse(bytes.NewReader(rawData)); err == nil {
							xmlDocs[file.FQN] = xml
						}
					}
				}

				//csaService.xmlMux.Unlock()

				matchFunc = func() (bool, string) {
					return rule.Patterns[i].MatchXml(xmlDocs[file.FQN])
				}
			} else if rule.Patterns[i].Type == model.YAMLPATH_MATCH_TYPE {
				//csaService.yamlMux.Lock()

				if yamlDocs[file.FQN] == nil {
					if rawData, err := ioutil.ReadFile(file.FQN); err == nil {
						var node yaml.Node
						err = yaml.Unmarshal(rawData, &node)

						if err == nil {
							yamlDocs[file.FQN] = &node
						}
					}
				}

				//csaService.yamlMux.Unlock()

				matchFunc = func() (bool, string) {
					return rule.Patterns[i].MatchYaml(yamlDocs[file.FQN])
				}
			}

			// if value, ok := path.String(root); ok
			ok, result := matchFunc()

			//fmt.Printf("Match %t", ok)
			//fmt.Printf(target)

			if ok && !rule.Negative {
				if len(result) > 0 {
					target = regexp.MustCompile(`\r?\n`).ReplaceAllString(result, " ")
				}

				//csaService.handleRuleMatched(run, app, file, line, target, rule, rule.Patterns[i], output, result, nil)
				//fmt.Printf("WE ARE HERE1 => " + target)
				findings++
				cnt++
			} else if !ok && rule.Negative {
				//csaService.handleRuleMatched(run, app, file, 0, target, rule, rule.Patterns[i], output, "", nil)
				//fmt.Printf("WE ARE HERE2")
				findings++
				cnt++
			}

		}

		pcnt++

	} //End Pattern For Loop

	//run.AddFindings(findings)
	//rule.Metric.Accumulate(pcnt, cnt, time.Since(start))

	return findings
}
