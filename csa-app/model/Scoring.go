/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

const DEFAULT_MIN_SCORE = 0.0
const DEFAULT_MAX_SCORE = 10.0

type PortfolioScore struct {
	Findings       int           `json:"findings"`
	CIFindings     int           `json:"ciFindings"`
	InfoFindings   int           `json:"infoFindings"`
	RawScore       int64         `json:"rawScore"`
	AppScores      []Application `json:"appScores"`
	Recommendation string        `json:"recommendation"`
}

type ApplicationDetails struct {
	Application       string  `json:"application"`
	Findings          int     `json:"findings"`
	CIFindings        int     `json:"ciFindings"`
	InfoFindings      int     `json:"infoFindings"`
	RawScore          int64   `json:"rawScore"`
	RawCloudScore     int     `json:"rawCloudScore"`
	RawContainerScore int     `json:"rawContainerScore"`
	NumCrits          int     `json:"numCrits"`
	SlocCnt           int     `json:"slocCnt"`
	FilesCnt          int     `json:"filesCnt"`
	FindingsRatio     float64 `json:"findingsRatio"`
}

type ApplicationsForRun struct {
	Application string `json:"application"`
	UsageCount  int    `json:"usageCount"`
}

type AppScoreCard struct {
	Application string `json:"application"`
	Info        int    `json:"info"`
	Low         int    `json:"low"`
	Medium      int    `json:"medium"`
	High        int    `json:"high"`
	Total       int    `json:"total"`
}

type ScoreCardDetail struct {
	Application string `json:"application"`
	Category    string `json:"category"`
	Pattern     string `json:"pattern"`
	Effort      int    `json:"effort"`
	Level       string `json:"level"`
	Count       int    `json:"count"`
	Total       int    `json:"total"`
}

func (p *PortfolioScore) AddApplicationScore(app Application) {
	p.Findings += app.Findings
	p.CIFindings += app.CIFindings
	p.InfoFindings += app.InfoFindings
	p.RawScore += app.RawScore
	p.AppScores = append(p.AppScores, app)
}

func (p *PortfolioScore) SlocCnt() int {
	cnt := 0

	for _, app := range p.AppScores {
		cnt += app.SlocCnt
	}

	return cnt
}

func (p *PortfolioScore) FileCnt() int {
	cnt := 0

	for _, app := range p.AppScores {
		cnt += app.FilesCnt
	}

	return cnt
}
