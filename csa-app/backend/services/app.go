/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package services

import (
	"fmt"

	"csa-app/model"

	log "github.com/sirupsen/logrus"
)

type ApplicationInfoService interface {
	GetScoreCards(runId uint) ([]model.AppScoreCard, error)
	GetScoreCard(runId uint, appName string, tags model.TagsRequest, includeFF bool) (model.AppScoreCard, error)
	GetScoreCardDetails(runId uint, appName string, cardName string) ([]model.ScoreCardDetail, error)
	GetFindings(runId uint, app string, category string, tag string, level string) ([]model.Finding, error)
	GetRunFindings(runId uint) ([]*model.FindingDTO, error)
	GetAppFindings(runId uint, appName string, cardName string, tagsRequest model.TagsRequest, includeFF bool) ([]*model.FindingDTO, error)
	GetFinding(id uint) (*model.FindingDTO, error)
	UpdateApp(app *model.Application) error
	GetAppTags(id uint, appname string) ([]string, error)
}

func (repoSvc *RepoService) GetScoreCard(runId uint, appName string, tagsRequest model.TagsRequest, includeFF bool) (model.AppScoreCard, error) {

	log.Debugf("Retrieving ScoreCard for Run[%d] App[%s] Tags[%v]\n", runId, appName, tagsRequest)

	tags := []string{}

	for _, tag := range tagsRequest.Tags {
		if tag.Selected {
			tags = append(tags, tag.Name)
		}
	}

	return repoSvc.repositoryMgr.Findings.GetScoreCard(runId, appName, tags, includeFF)
}

func (repoSvc *RepoService) GetScoreCards(runId uint) ([]model.AppScoreCard, error) {
	return repoSvc.repositoryMgr.Findings.GetScoreCards(runId)
}

func (repoSvc *RepoService) GetScoreCardDetails(runId uint, appName string, cardName string) ([]model.ScoreCardDetail, error) {
	return repoSvc.repositoryMgr.Findings.GetScoreCardDetails(runId, appName, cardName)
}

func (repoSvc *RepoService) GetRunFindings(runId uint) (findings []*model.FindingDTO, err error) {
	findings, err = repoSvc.repositoryMgr.Findings.GetFindingsDTOForRun(runId)

	if findings == nil {
		findings = []*model.FindingDTO{}
	}
	return
}

func (repoSvc *RepoService) GetAppFindings(runId uint, appName string, cardName string, tagsRequest model.TagsRequest, includeFF bool) (findings []*model.FindingDTO, err error) {

	tags := []string{}

	for _, tag := range tagsRequest.Tags {
		if tag.Selected {
			tags = append(tags, tag.Name)
		}
	}

	findings, err = repoSvc.repositoryMgr.Findings.GetFindingsDTOForRunAppLevel(runId, appName, cardName, tags, includeFF)

	if findings == nil {
		findings = []*model.FindingDTO{}
	}
	return
}

func (repoSvc *RepoService) GetFindings(runId uint, app string, category string, tag string, level string) ([]model.Finding, error) {
	findings := []model.Finding{}
	return findings, fmt.Errorf("Not Implemented!")
}

func (repoSvc *RepoService) GetFinding(id uint) (*model.FindingDTO, error) {
	newFinding := &model.FindingDTO{}
	finding, err := repoSvc.repositoryMgr.Findings.GetFinding(id)
	if err == nil {
		newFinding = finding.CreateDTO()
	}
	return newFinding, err
}

func (repoSvc *RepoService) UpdateApp(app *model.Application) error {
	return repoSvc.repositoryMgr.Run.UpdateApp(app)
}

func (reposvc *RepoService) GetAppTags(id uint, appname string) (tags []string, err error) {
	return reposvc.repositoryMgr.Findings.GetTagsForApp(id, appname)
}
