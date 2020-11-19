/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/backend/services"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/db"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/model"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/search"
	log "github.com/sirupsen/logrus"
)

type runRoutes struct {
	runsRepository db.RunRepository
	scoringSvc     services.ScoringService
	appSvc         services.ApplicationInfoService
}

func (r *runRoutes) getRuns(c *gin.Context) {
	runs, _ := r.runsRepository.GetRuns()
	c.JSON(http.StatusOK, gin.H{
		"runs": runs,
	})
}

func (r *runRoutes) getApps(c *gin.Context) {
	runId := getId(c)
	fmt.Printf("Getting apps for Run[%d]\n", runId)

	apps, err := r.runsRepository.GetRunApps(runId)

	if !CheckForError(c, err, fmt.Sprintf("Error retrieving apps for run[%d]! Details => %%s", runId)) {
		c.JSON(http.StatusOK, gin.H{
			"app": apps,
		})
	}
}

func (r *runRoutes) getAppScores(c *gin.Context) {
	var sm *model.ScoringModel
	runId := getId(c)
	scoreModelName := c.Query("model")
	start := time.Now()

	if scoreModelName != "" {
		//TODO Get Score Model to change scoring from original scoring...
	}

	log.Debugf("Retrieving app scores for Run [%d]\n", runId)
	portfolioScores, err := r.scoringSvc.GetPortfolioScore(runId, sm)

	log.Debugf("Scores =>\n\n %+v \n\n", portfolioScores)

	log.Infof("Fetching App Scores took [%s]\n", time.Since(start))

	if !CheckForError(c, err, fmt.Sprintf("Error retrieving app scores for run[%d]! Details => %%s", runId)) {
		c.JSON(http.StatusOK, gin.H{
			"scores": portfolioScores,
		})
	}
}

func (r *runRoutes) getAnalyzeRuns(c *gin.Context) {
	runs, _ := r.runsRepository.GetRunsByCommand("analyze")
	c.JSON(http.StatusOK, gin.H{
		"runs": runs,
	})
}

func (r *runRoutes) getIndexStatus(c *gin.Context) {
	runId := getId(c)
	fmt.Printf("Getting index details for Run[%d]\n", runId)

	run, err := r.runsRepository.GetRun(runId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Run [%d] does not exist!", runId))
	}

	index := search.GetExistingIndex(&run)

	log.Debugf("Index for Run[%d] exists? %v doc cnt: %d\n", runId, index.Exists, index.NumDocs)

	c.JSON(http.StatusOK, gin.H{
		"index": index,
	})
}

func (r *runRoutes) updateApp(c *gin.Context) {
	runId := getId(c)

	app := &model.Application{}

	if err := c.Bind(app); err != nil {
		msg := fmt.Sprintf("{error:\"Invalid update app request!Must specifiy a valid application! Details: %v\"}", err)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	if app.RunID != runId {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid Update request. Run Id [%d] and App Run Id [%d] don't match!", runId, app.RunID))
	}

	err := r.runsRepository.UpdateApp(app)

	if !CheckForError(c, err, fmt.Sprintf("Error updating app [%s] for run[%d]! Details => %%s", app.Name, runId)) {
		c.JSON(http.StatusAccepted, fmt.Sprintf("App [%s] successfully updated!", app.Name))
	}
}

func (r *runRoutes) getAppTags(c *gin.Context) {
	runId := getId(c)
	app := c.Param("app")

	var tags []string

	tags, err := r.appSvc.GetAppTags(runId, app)

	if !CheckForError(c, err, fmt.Sprintf("Error retrieving Tags for app [%s] for run[%d]! Details => %%s", app, runId)) {
		c.JSON(http.StatusOK, gin.H{"tags": tags})
	}
}
