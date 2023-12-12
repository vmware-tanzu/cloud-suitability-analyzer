/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package services

import (
	"fmt"
	"sort"
	"strings"

	"csa-app/csa"
	"csa-app/model"

	"csa-app/search"
)

const APP_HEADING = "App"
const BLANK_RULE_NAME = "csa-internal"

type DataService interface {
	GetRunSlocByLang(runId uint) ([]model.SlocByLang, error)
	GetRunAPISummary(runId uint) ([]model.ApiUsage, error)
	GetRunAPIDetails(runId uint) ([]model.ApiUsageDetail, error)
	GetRunAPIByApp(runId uint) (model.ApiMatrix, error)
	GetRunRuleByApp(runId uint) (model.RuleMatrix, error)
	GetAnnotations(runId uint) ([]model.LevelDetail, error)
	GetThirdParty(runId uint) ([]model.LevelDetail, error)
	GetFindingsByCriteria(criteria model.Criteria) ([]model.Finding, error)
	SearchFindings(runId uint, request *search.SearchRequest) *search.SearchResult
	GetRunRuleMetrics(runId uint) ([]model.RuleMetric, error)
}

func (repo *RepoService) GetRunSlocByLang(runId uint) ([]model.SlocByLang, error) {

	slocCnts := []model.SlocByLang{}

	slocData, err := repo.repositoryMgr.Sloc.GetSlocForRun(runId)

	if err == nil {

		//Group it up!
		totalFiles := 0
		totalBlank := 0
		totalComment := 0
		totalCode := 0

		langTotals := make(map[string]*model.SlocByLang)

		for _, item := range slocData {

			if _, ok := langTotals[item.Lang]; !ok {
				langTotals[item.Lang] = &model.SlocByLang{Language: item.Lang, TotalFiles: item.TotalFiles,
					BlankLines:   item.BlankLines,
					CodeLines:    item.CodeLines,
					CommentLines: item.CommentLines}
			} else {
				langTotals[item.Lang].TotalFiles += item.TotalFiles
				langTotals[item.Lang].BlankLines += item.BlankLines
				langTotals[item.Lang].CommentLines += item.CommentLines
				langTotals[item.Lang].CodeLines += item.CodeLines
			}

			totalFiles += item.TotalFiles
			totalBlank += item.BlankLines
			totalComment += item.CommentLines
			totalCode += item.CodeLines
		}

		for _, item := range langTotals {
			slocCnts = append(slocCnts, *item)
		}
	}
	return slocCnts, err
}

func (repo *RepoService) GetRunAPISummary(runId uint) ([]model.ApiUsage, error) {

	apiSummary := []model.ApiUsage{}

	apiDetails, err := repo.GetRunAPIDetails(runId)

	if err == nil {
		summary := make(map[string]*model.ApiUsage)

		for _, item := range apiDetails {

			if _, ok := summary[item.Api]; !ok {
				summary[item.Api] = &model.ApiUsage{Api: item.Api}
			}
			summary[item.Api].UsageCount += 1
		}

		for _, item := range summary {
			apiSummary = append(apiSummary, *item)
		}
	}

	return apiSummary, err
}

func (repo *RepoService) GetRunAPIDetails(runId uint) ([]model.ApiUsageDetail, error) {
	apiDetails := []model.ApiUsageDetail{}

	apiFindings, err := repo.repositoryMgr.Findings.GetFindingsByTag(runId, model.API_TAG)

	if err == nil {
		AdjLevel := ""
		for _, finding := range apiFindings {

			if finding.Effort < 0 {
				AdjLevel = "Boost"
			} else {
				AdjLevel = csa.GetLevelForScore(finding.Effort)
			}

			apiDetails = append(apiDetails, model.ApiUsageDetail{
				Application: finding.Application,
				Api:         finding.Category,
				Filename:    finding.Filename,
				Pattern:     finding.Pattern,
				Line:        finding.Line,
				Value:       finding.Value,
				Effort:      finding.Effort,
				Advice:      finding.Advice,
				Level:       AdjLevel})
		}
	}

	return apiDetails, err
}

func (repo *RepoService) GetRunAPIByApp(runId uint) (model.ApiMatrix, error) {
	apiUsage := []*model.ApiByApp{}

	appUsageMap := make(map[string]map[string]int)
	appUsageMap[model.TOTAL_FIELD] = map[string]int{model.TOTAL_FIELD: 0}
	var apiColumns []string
	apiColumns = append(apiColumns, APP_HEADING)
	apiColumns = append(apiColumns, model.TOTAL_FIELD)

	apiFindings, err := repo.repositoryMgr.Findings.GetFindingsByTag(runId, model.API_TAG)

	if err == nil {

		apiMap := make(map[string]int)
		apiMap[model.TOTAL_FIELD] = 0
		//First get all the apps/apis setup
		for _, finding := range apiFindings {

			if _, found := appUsageMap[finding.Application]; !found {
				appUsageMap[finding.Application] = map[string]int{}
			}

			if _, found := apiMap[finding.Category]; !found {
				apiMap[finding.Category] = 0
				apiColumns = append(apiColumns, finding.Category)
			}
		}

		//Add in the total row and columns

		//Setup the apiUsages with the appMaps
		for k, _ := range appUsageMap {
			//Clone the api map
			cloneApiMap := make(map[string]int)
			for k, v := range apiMap {
				cloneApiMap[k] = v
			}
			appUsageMap[k] = cloneApiMap
		}

		for _, finding := range apiFindings {
			//Fill in the matrix
			appUsageMap[finding.Application][finding.Category] += finding.Effort
			appUsageMap[finding.Application][model.TOTAL_FIELD] += finding.Effort
			appUsageMap[model.TOTAL_FIELD][finding.Category] += finding.Effort
			appUsageMap[model.TOTAL_FIELD][model.TOTAL_FIELD] += finding.Effort
		}
	}

	for k, v := range appUsageMap {
		newAppU := &model.ApiByApp{Application: k}
		apiUsage = append(apiUsage, newAppU)
		for api, use := range v {
			newAppU.ApiDetails = append(newAppU.ApiDetails, model.ApiUsage{Api: api, UsageCount: use})
		}
	}

	sort.Sort(ApiByApp(apiUsage))
	sort.Sort(ByHeadingWithTotals(apiColumns))
	return model.ApiMatrix{Cols: apiColumns, Data: apiUsage}, err
}

func (repo *RepoService) GetRunRuleByApp(runId uint) (model.RuleMatrix, error) {
	ruleUsage := []*model.RuleByApp{}

	appUsageMap := make(map[string]map[string]int)
	appUsageMap[model.TOTAL_FIELD] = map[string]int{model.TOTAL_FIELD: 0}
	var ruleColumns []string
	ruleColumns = append(ruleColumns, APP_HEADING)
	ruleColumns = append(ruleColumns, model.TOTAL_FIELD)

	findings, err := repo.repositoryMgr.Findings.GetFindingsWithoutPreload(runId)

	if err == nil {

		ruleMap := make(map[string]int)
		ruleMap[model.TOTAL_FIELD] = 0
		//First get all the apps/rule setup
		for _, finding := range findings {

			if _, found := appUsageMap[finding.Application]; !found {
				appUsageMap[finding.Application] = map[string]int{}
			}

			rule := finding.Rule

			if rule == "" {
				rule = BLANK_RULE_NAME
			}

			if _, found := ruleMap[rule]; !found {
				ruleMap[rule] = 0
				ruleColumns = append(ruleColumns, rule)
			}
		}

		//Add in the total row and columns

		//Setup the ruleUsages with the appMaps
		for k, _ := range appUsageMap {
			//Clone the api map
			cloneRuleMap := make(map[string]int)
			for k, v := range ruleMap {
				cloneRuleMap[k] = v
			}
			appUsageMap[k] = cloneRuleMap
		}

		for _, finding := range findings {
			//Fill in the matrix
			rule := finding.Rule

			if rule == "" {
				rule = BLANK_RULE_NAME
			}

			appUsageMap[finding.Application][rule] += finding.Effort
			appUsageMap[finding.Application][model.TOTAL_FIELD] += finding.Effort
			appUsageMap[model.TOTAL_FIELD][rule] += finding.Effort
			appUsageMap[model.TOTAL_FIELD][model.TOTAL_FIELD] += finding.Effort
		}
	}

	for k, v := range appUsageMap {
		newAppU := &model.RuleByApp{Application: k}
		ruleUsage = append(ruleUsage, newAppU)
		for rule, use := range v {
			newAppU.RuleDetails = append(newAppU.RuleDetails, model.RuleUsage{Rule: rule, Count: use})
		}
	}

	sort.Sort(RuleByApp(ruleUsage))
	sort.Sort(ByHeadingWithTotals(ruleColumns))
	return model.RuleMatrix{Cols: ruleColumns, Data: ruleUsage}, err
}

func (repo *RepoService) GetAnnotations(runId uint) ([]model.LevelDetail, error) {

	annotations := []model.LevelDetail{}

	findings, err := repo.repositoryMgr.Findings.GetFindingsByRule(runId, "java-annotations")

	for _, finding := range findings {
		annotations = append(annotations, model.NewLevelDetail(finding.Application, finding.Category, finding.Filename, finding.Line, finding.Value,
			finding.Pattern, finding.Effort, csa.GetLevelForScore(finding.Effort), finding.Advice))
	}

	return annotations, err
}

func (repo *RepoService) GetThirdParty(runId uint) ([]model.LevelDetail, error) {

	thirdParty := []model.LevelDetail{}

	findings, err := repo.repositoryMgr.Findings.GetFindingsByRule(runId, "java-3rdPartyImports")

	for _, finding := range findings {
		thirdParty = append(thirdParty, model.NewLevelDetail(finding.Application, finding.Category, finding.Filename, finding.Line, finding.Value,
			finding.Pattern, finding.Effort, csa.GetLevelForScore(finding.Effort), finding.Advice))
	}

	return thirdParty, err
}

func (repo *RepoService) GetFindingsByCriteria(criteria model.Criteria) ([]model.Finding, error) {
	findings, _ := repo.repositoryMgr.Findings.GetFindings(123)
	return findings, fmt.Errorf("Not Implemented!")
}

func (repo *RepoService) GetRunRuleMetrics(runId uint) ([]model.RuleMetric, error) {
	metrics, err := repo.repositoryMgr.Rules.GetRuleMetrics(runId)
	return metrics, err
}

func (repo *RepoService) SearchFindings(runid uint, request *search.SearchRequest) *search.SearchResult {
	searchEngine := search.NewSearchEngine(repo.repositoryMgr.Run, repo.repositoryMgr.Findings)
	return searchEngine.SearchFindingsHighlighted(runid, request, search.HTML_HIGHLIGHT)
}

type ApiByApp []*model.ApiByApp

func (apiData ApiByApp) Len() int      { return len(apiData) }
func (apiData ApiByApp) Swap(i, j int) { apiData[i], apiData[j] = apiData[j], apiData[i] }
func (apiData ApiByApp) Less(i, j int) bool {
	f1 := apiData[i]
	f2 := apiData[j]
	if f1.Application == model.TOTAL_FIELD {
		return false
	}
	if f2.Application == model.TOTAL_FIELD {
		return true
	}
	return strings.ToLower(f1.Application) < strings.ToLower(f2.Application)
}

type ByHeadingWithTotals []string

func (apiData ByHeadingWithTotals) Len() int      { return len(apiData) }
func (apiData ByHeadingWithTotals) Swap(i, j int) { apiData[i], apiData[j] = apiData[j], apiData[i] }
func (apiData ByHeadingWithTotals) Less(i, j int) bool {
	f1 := apiData[i]
	f2 := apiData[j]
	if f1 == APP_HEADING {
		return true
	}
	if f2 == APP_HEADING {
		return false
	}
	if f1 == model.TOTAL_FIELD {
		return false
	}
	if f2 == model.TOTAL_FIELD {
		return true
	}
	return f1 < f2
}

type RuleByApp []*model.RuleByApp

func (ruleData RuleByApp) Len() int      { return len(ruleData) }
func (ruleData RuleByApp) Swap(i, j int) { ruleData[i], ruleData[j] = ruleData[j], ruleData[i] }
func (ruleData RuleByApp) Less(i, j int) bool {
	f1 := ruleData[i]
	f2 := ruleData[j]
	if f1.Application == model.TOTAL_FIELD {
		return false
	}
	if f2.Application == model.TOTAL_FIELD {
		return true
	}
	return strings.ToLower(f1.Application) < strings.ToLower(f2.Application)
}
