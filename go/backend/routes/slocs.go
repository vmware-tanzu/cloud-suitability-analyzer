/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vmware-samples/cloud-suitability-analyzer/go/backend/services"
	"github.com/vmware-samples/cloud-suitability-analyzer/go/db"
	"net/http"
)

type slocRoutes struct {
	slocRepo db.SlocRepository
	dataSvc  services.DataService
}

func (r *slocRoutes) getAppSlocCounts(c *gin.Context) {
	runId := getId(c)
	slocScores, _ := r.slocRepo.GetSlocSummaryByApplicationForRun(runId)
	c.JSON(http.StatusOK, slocScores)
}

func (r *slocRoutes) getRunSlocCounts(c *gin.Context) {
	runId := getId(c)
	slocScores, _ := r.slocRepo.GetSummaryFindingsForRun(runId)
	c.JSON(http.StatusOK, slocScores)
}

func (r *slocRoutes) getTopLanguagesByCodeLines(c *gin.Context) {
	runId := getId(c)
	languagesByCodeLines, _ := r.slocRepo.GetTopLanguagesByCodeLines(runId)
	c.JSON(http.StatusOK, languagesByCodeLines)
}

func (r *slocRoutes) getLanguagesForRunAndApplication(c *gin.Context) {
	runId := getId(c)
	application := c.Param("app")
	languagesByCodeLines, _ := r.slocRepo.GetLanguagesForRunAndApplication(runId, application)
	c.JSON(http.StatusOK, languagesByCodeLines)
}

func (r *slocRoutes) getApplicationsByLanguageForRun(c *gin.Context) {
	runId := getId(c)
	language := c.Query("lang")
	applicationsByLanguageForRuns, _ := r.slocRepo.GetApplicationsByLanguageForRun(runId, language)
	c.JSON(http.StatusOK, applicationsByLanguageForRuns)
}

func (r *slocRoutes) getRunSlocsByLang(c *gin.Context) {
	runId := getId(c)
	runSlocDetails, _ := r.dataSvc.GetRunSlocByLang(runId)
	c.JSON(http.StatusOK, runSlocDetails)
}
