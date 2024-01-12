/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package routes

import (
	"csa-app/backend/services"
	"csa-app/db"
	"csa-app/model"
	"csa-app/search"
	"fmt"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type findingRoutes struct {
	findingsRepo db.FindingRepository
	appService   services.ApplicationInfoService
	dataService  services.DataService
}

func (r *findingRoutes) getTopApisByScore(c *gin.Context) {
	runId := getId(c)
	topApis, _ := r.findingsRepo.GetTopApisByScoreForRun(runId)
	c.JSON(http.StatusOK, topApis)
}

func (r *findingRoutes) getApiUsageForRunAndApplication(c *gin.Context) {
	runId := getId(c)
	application := c.Param("app")
	apiUsages, _ := r.findingsRepo.GetApisByUsageForRunAndApplication(runId, application)
	c.JSON(http.StatusOK, apiUsages)
}

func (r *findingRoutes) getApplicationsUsingApiForRun(c *gin.Context) {
	runId := getId(c)
	apiName := c.Query("api")

	appsForApi, _ := r.findingsRepo.GetApplicationsUsingApiForRun(runId, apiName)
	c.JSON(http.StatusOK, appsForApi)
}

func (r *findingRoutes) getApplicationScoreCards(c *gin.Context) {
	runId := getId(c)

	appScoreCard, _ := r.appService.GetScoreCards(runId)
	c.JSON(http.StatusOK, appScoreCard)
}

func (r *findingRoutes) getApplicationScoreCard(c *gin.Context) {
	runId := getId(c)
	appName := c.Param("app")
	includeFF, err := strconv.ParseBool(c.DefaultQuery("includeFF", "false"))
	if err != nil {
		includeFF = false
	}

	tags := &model.TagsRequest{}
	if err := c.Bind(tags); err != nil {
		msg := fmt.Sprintf("{error:\"Invalid ScoredCardrequest! Details: %v\"}", err)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	log.Debugf("Received Post request for App ScoreCard run[%d] app[%s] tag[%v]\n", runId, appName, tags)

	appScoreCard, _ := r.appService.GetScoreCard(runId, appName, *tags, includeFF)
	c.JSON(http.StatusOK, appScoreCard)
}

func (r *findingRoutes) getApplicationScoreCardDetails(c *gin.Context) {
	runId := getId(c)
	appName := c.Param("app")
	scoreCard := c.Param("card")

	scoreCardDetails, _ := r.appService.GetScoreCardDetails(runId, appName, scoreCard)
	c.JSON(http.StatusOK, scoreCardDetails)
}

func (r *findingRoutes) getApiDetailedUsage(c *gin.Context) {
	runId := getId(c)
	detailedUsage, _ := r.dataService.GetRunAPIDetails(runId)
	c.JSON(http.StatusOK, detailedUsage)
}

func (r *findingRoutes) getAppApiUsage(c *gin.Context) {
	runId := getId(c)
	detailedUsage, err := r.dataService.GetRunAPIByApp(runId)
	if !CheckForError(c, err, fmt.Sprintf("Error retrieving Api Usage By App for run[%d]! Details => %%s", runId)) {
		c.JSON(http.StatusOK, detailedUsage)
	}
}

func (r *findingRoutes) getAppRuleScoring(c *gin.Context) {
	runId := getId(c)
	detailedUsage, err := r.dataService.GetRunRuleByApp(runId)
	if !CheckForError(c, err, fmt.Sprintf("Error retrieving Rule scoring By App for run[%d]! Details => %%s", runId)) {
		c.JSON(http.StatusOK, detailedUsage)
	}
}

func (r *findingRoutes) getApiSummary(c *gin.Context) {
	runId := getId(c)
	detailedUsage, _ := r.dataService.GetRunAPISummary(runId)
	c.JSON(http.StatusOK, detailedUsage)
}

func (r *findingRoutes) getRunAnnotations(c *gin.Context) {
	runId := getId(c)
	annotations, _ := r.dataService.GetAnnotations(runId)
	c.JSON(http.StatusOK, annotations)
}

func (r *findingRoutes) getRunThirdPartyLibs(c *gin.Context) {
	runId := getId(c)
	thirdParty, _ := r.dataService.GetThirdParty(runId)
	c.JSON(http.StatusOK, thirdParty)
}

func (r *findingRoutes) getApplicationFindings(c *gin.Context) {
	runId := getId(c)
	appName := c.Param("app")

	category := c.Query("category")
	tag := c.Query("tag")
	level := c.Query("level")

	findings, _ := r.appService.GetFindings(runId, appName, category, tag, level)

	c.JSON(http.StatusOK, findings)
}

func (r *findingRoutes) getRunFindings(c *gin.Context) {
	runId := getId(c)

	findings, err := r.appService.GetRunFindings(runId)

	if !CheckForError(c, err, fmt.Sprintf("Error retrieving findings for run[%d]! Details => %%s", runId)) {
		c.JSON(http.StatusOK, findings)
	}
}

func (r *findingRoutes) getAppFindings(c *gin.Context) {
	runId := getId(c)
	appName := c.Param("app")
	scoreCard := c.Param("card")
	includeFF, err := strconv.ParseBool(c.DefaultQuery("includeFF", "false"))
	if err != nil {
		includeFF = false
	}

	tags := &model.TagsRequest{}

	if err = c.Bind(tags); err != nil {
		msg := fmt.Sprintf("{error:\"Invalid App Findings Request! Details: %v\"}", err)
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	log.Debugf("Received Post request for App Findings run[%d] app[%s] card[%s] tag[%v] includeFF[%v]\n", runId, appName, scoreCard, tags, includeFF)

	findings, err := r.appService.GetAppFindings(runId, appName, scoreCard, *tags, includeFF)

	if !CheckForError(c, err, fmt.Sprintf("Error retrieving findings for run[%d]! app[%s] card[%s] Details => %%s",
		runId, appName, scoreCard)) {
		c.JSON(http.StatusOK, findings)
	}
}

func (r *findingRoutes) getFinding(c *gin.Context) {
	id := getId(c)
	findings, err := r.appService.GetFinding(id)
	if !CheckForError(c, err, fmt.Sprintf("Error retrieving finding [%d]! Details => %%s", id)) {
		c.JSON(http.StatusOK, findings)
	}
}

func (r *findingRoutes) getFindingsByCriteria(c *gin.Context) {
	var criteria model.Criteria

	if err := c.BindJSON(&criteria); err != nil {
		msg := fmt.Sprintf("{error:\"Invalid Criteria on Findings Get Request! Details: %v\"}", err)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	findings, _ := r.dataService.GetFindingsByCriteria(criteria)
	c.JSON(http.StatusOK, findings)
}

func (r *findingRoutes) searchFindings(c *gin.Context) {
	query := c.Query("query")
	maxresults := c.Query("max")
	queryType := c.Query("type")

	maxResults, err := strconv.Atoi(maxresults)
	if err != nil {
		maxResults = 100
	}

	request := &search.SearchRequest{Query: query, MaxResults: maxResults, Type: queryType}
	r.performSearch(c, request)
}

func (r *findingRoutes) searchFindingsPost(c *gin.Context) {

	request := &search.SearchRequest{}

	if err := c.Bind(request); err != nil {
		msg := fmt.Sprintf("{error:\"Invalid request!Must specifiy `query` at a minimum! Details: %v\"}", err)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	r.performSearch(c, request)
}

func (r *findingRoutes) performSearch(c *gin.Context, request *search.SearchRequest) {
	runId := getId(c)

	if request.MaxResults == 0 {
		request.MaxResults = search.DEFAULT_MAX_RESULTS
	}

	findings := r.dataService.SearchFindings(runId, request)
	c.JSON(http.StatusOK, findings)

}

func CheckForError(c *gin.Context, err error, msgfmt string) bool {

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(msgfmt, err.Error()))
		return true
	}

	return false
}
