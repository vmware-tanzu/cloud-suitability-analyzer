/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"csa-app/frontend/resources"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/pkg/browser"

	"csa-app/backend/services"
	"csa-app/db"
	"csa-app/util"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const BROWSER_URL = "http://localhost:3001"

func StartRouter(database *gorm.DB, useHttpFS bool, port int) {
	gin.SetMode(gin.ReleaseMode)
	// Set the router as the default one shipped with Gin
	router := SetupRouter(database, useHttpFS)
	// Start and run the server
	router.Run(fmt.Sprintf(":%d", port))

}

func SetupRouter(database *gorm.DB, useHttpFS bool) *gin.Engine {
	router := gin.Default()
	// Serve frontend static files

	if useHttpFS {
		fmt.Println("Using Http FileSystem!")
		defaultHandler := static.Serve("/", BinaryFileSystem("build"))
		router.Use(defaultHandler)
		fmt.Printf("\n\nOpen Browser and goto %s\n\n", BROWSER_URL)

		//Attempt to open browser for user
		_ = browser.OpenURL(BROWSER_URL)
	} else {
		fmt.Println("Statically serving from ./build!")
		router.Use(static.Serve("/", static.LocalFile("./build", true)))
	}

	repositories := db.NewRepositoriesManager(database)
	appSvc := services.NewAppService(repositories)
	dataSvc := services.NewDataService(repositories)
	scoreSvc := services.NewScoringService(repositories)
	ruleRoutes := &ruleRoutes{repositories.Rules}
	runRoutes := &runRoutes{repositories.Run, scoreSvc, appSvc}
	findingRoutes := &findingRoutes{repositories.Findings, appSvc, dataSvc}
	slocRoutes := &slocRoutes{repositories.Sloc, dataSvc}

	api := router.Group("/api")
	{
		api.GET("/health", baseRoute)
		api.GET("/rules", ruleRoutes.getAllRules)
		api.GET("/runs", runRoutes.getRuns)
		api.GET("/version", version)
		api.GET("/analyze-runs", runRoutes.getAnalyzeRuns)
		api.GET("/findings/:id", findingRoutes.getFinding)

		run := api.Group("runs/:id")
		{
			run.GET("/index", runRoutes.getIndexStatus)
			run.GET("/scorecards", findingRoutes.getApplicationScoreCards)
			run.GET("/findings", findingRoutes.getRunFindings)
			run.GET("/apps", runRoutes.getApps)
			run.GET("/rule-metrics", ruleRoutes.getMetrics)
			run.POST("/search", findingRoutes.searchFindingsPost)
			summary := run.Group("/summary")
			{
				summary.GET("/application_scores", runRoutes.getAppScores)
				summary.GET("/application_slocs", slocRoutes.getAppSlocCounts)
				summary.GET("/run_slocs", slocRoutes.getRunSlocCounts)
				summary.GET("/top_languages_by_codelines", slocRoutes.getTopLanguagesByCodeLines)
				summary.GET("/top_apis_by_score", findingRoutes.getTopApisByScore)
				summary.GET("/top_apps_for_api", findingRoutes.getApplicationsUsingApiForRun)
				summary.GET("/apps_for_language", slocRoutes.getApplicationsByLanguageForRun)
			}
			app := run.Group("/apps/:app")
			{
				app.POST("/scorecard", findingRoutes.getApplicationScoreCard)
				app.GET("/scorecard/:card", findingRoutes.getApplicationScoreCardDetails)
				app.GET("/languages", slocRoutes.getLanguagesForRunAndApplication)
				app.GET("/apis", findingRoutes.getApiUsageForRunAndApplication)
				app.GET("/findings", findingRoutes.getApplicationFindings)
				app.POST("/findings/scorecard/:card", findingRoutes.getAppFindings)
				app.GET("/tags", runRoutes.getAppTags)
				app.POST("/", runRoutes.updateApp)
			}
			data := run.Group("/data")
			{
				data.GET("/api_detailed_usage", findingRoutes.getApiDetailedUsage)
				data.GET("/app_api_usage", findingRoutes.getAppApiUsage)
				data.GET("/app_rule_score", findingRoutes.getAppRuleScoring)
				data.GET("/api_summary", findingRoutes.getApiSummary)
				data.GET("/annotations", findingRoutes.getRunAnnotations)
				data.GET("/thirdParty", findingRoutes.getRunThirdPartyLibs)
				data.GET("/sloc", slocRoutes.getRunSlocsByLang)
			}

		}
	}
	return router
}

func baseRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("CSA Server [%s] is up!", util.App.Model().Version)})
}

func version(c *gin.Context) {
	fmt.Printf("CSA Version => %s\n", util.App.Model().Version)
	c.JSON(http.StatusOK, gin.H{"version": util.App.Model().Version})
}

func getId(c *gin.Context) uint {
	id := c.Param("id")
	idAsUint, _ := strconv.ParseUint(id, 10, 64)
	return uint(idAsUint)
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	var fs *assetfs.AssetFS
	fs = &assetfs.AssetFS{Asset: resources.Asset, AssetDir: resources.AssetDir, AssetInfo: resources.AssetInfo, Prefix: root}

	return &binaryFileSystem{
		fs,
	}
}
