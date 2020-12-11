/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

type ApiUsage struct {
	Api        string `json:"api"`
	UsageCount int    `json:"usageCount"`
}

type ApiUsageDetail struct {
	Application string `json:"application"`
	Api         string `json:"api"`
	Filename    string `json:"filename"`
	Pattern     string `json:"pattern"`
	Line        int    `json:"line""`
	Value       string `json:"value"`
	Effort      int    `json:"effort"`
	Level       string `json:"level"`
	Advice      string `json:"advice"`
}

type ApiMatrix struct {
	Cols []string    `json:"cols"`
	Data []*ApiByApp `json:"data"`
}

type ApiByApp struct {
	Application string     `json:"application"`
	ApiDetails  []ApiUsage `json:"apiusage"`
}

type RuleMatrix struct {
	Cols []string     `json:"cols"`
	Data []*RuleByApp `json:"data"`
}

type RuleByApp struct {
	Application string      `json:"application"`
	RuleDetails []RuleUsage `json:"ruleusage"`
}

type RuleUsage struct {
	Rule  string `json:"rule"`
	Count int    `json:"count"`
}

type LevelDetail struct {
	Application string `json:"application"`
	Category    string `json:"category"`
	Filename    string `json:"filename"`
	Pattern     string `json:"pattern"`
	Line        int    `json:"line""`
	Value       string `json:"value"`
	Effort      int    `json:"effort"`
	Level       string `json:"level"`
	Advice      string `json:"advice"`
}

func NewLevelDetail(app string, cat string, file string, line int, val string, pattern string, effort int, level string, advice string) LevelDetail {
	return LevelDetail{
		Application: app,
		Category:    cat,
		Filename:    file,
		Line:        line,
		Value:       val,
		Effort:      effort,
		Level:       level,
		Advice:      advice,
		Pattern:     pattern,
	}
}

func (a *ApiByApp) AddAPIUsage(api string, cnt int) {
	a.ApiDetails = append(a.ApiDetails, ApiUsage{Api: api, UsageCount: cnt})
}
