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

type ScoringService interface {
	GetPortfolioScore(runId uint, scoreModel *model.ScoringModel) (*model.PortfolioScore, error)
}

func (svc *RepoService) GetPortfolioScore(runId uint, scoreModel *model.ScoringModel) (*model.PortfolioScore, error) {
	portfolioScore := &model.PortfolioScore{}

	//Get Applications from DB
	apps, err := svc.repositoryMgr.Run.GetRunApps(runId)

	var lastmodel string

	if err == nil {
		for _, app := range apps {
			if !app.ScoreModified {
				if app.OriginalScore == -1 {
					app.OriginalScore = app.Score
				}
				//Quasi Rescore App based on front-end request...doesn't change the raw score. Just how the final normalize score looks.
				//This does not update the Normalized score in the DB which is calculated using the standard 1 min 10 max reverse = false.
				log.Debugf("App [%s] has current score [%v] rawScore [%d]\n", app.Name, app.Score, app.RawScore)

				if scoreModel == nil {
					var sm *model.ScoringModel
					if lastmodel == "" || lastmodel != app.ScoringModel || sm == nil {
						sm, err = svc.repositoryMgr.Scoring.GetModelByName(app.ScoringModel)
						if err != nil {
							return portfolioScore, fmt.Errorf("unable to retrieve scoring model [%s] for app [%s]. details: %s", app.ScoringModel, app.Name, err.Error())
						}
						lastmodel = app.ScoringModel
					}
					err = app.CalculateScore(sm)
				} else {
					err = app.CalculateScore(scoreModel)
				}
			}
			portfolioScore.AddApplicationScore(app)
		}
	}

	return portfolioScore, err
}
