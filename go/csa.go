/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

//
//          FILE:  csa
//
//
//   Summary:  Apply system of patterns to extract meta-data from source, config, git, maven assets
//
//       CREATED:  3/15/18
//      REVISION:  7/19/19
//	  REVISED BY:  Steve Woods (App Tx)
//===============================================================================

package main

import (
	"fmt"
	//"log"
	"os"
	"runtime"
	"runtime/debug"

	//"runtime/pprof"

	"github.com/pkg/profile"
	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/backend/routes"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/csa"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/db"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/model"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/natural"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/report"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/search"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/util"
)

//to generate the Bootstrap.go file
//go:generate go run scripts/generate_bootstrap.go

var (
	Version       string
	BuildDate     string
	Committer     string
	CommitSha     string
	CommitShaFull string
	CommitMsg     string
)

func main() {

	//--- trace code start
	/*
		f, err := os.Create("trace.out")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()
	*/
	//--- trace code end

	//--- profile code start
	/*
		f, err := os.Create("profile.out")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		//--- profile code end
	*/
	adminMode := false

	debug.SetMaxThreads(100000)

	util.App.HelpFlag.Short('?')
	util.App.Version(Version)
	util.App.Author("Steve Woods (VMWare)")

	var run = model.NewRun()

	procsAndThreads()

	run.DB = db.OpenDB(run)
	defer run.Cleanup()

	switch *util.Profile {
	case "cpu":
		defer profile.Start().Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	}

	repoMgr := db.NewRepositoriesManagerForRun(run)

	fmt.Printf("User: %s\nCommand: %s\nUser-Home: %s\nDB Path: %s\nRules-Dir: %s\nOutputPath: %s\nExe Path: %s\nTmp Path: %s\n\n",
		run.User,
		run.Command,
		run.Homepath,
		run.DbPath,
		run.RulesDir,
		run.OutputPath,
		run.Exepath,
		run.TmpPath)

	switch run.Command {

	case util.ShowReports.FullCommand():
		reportService := report.NewReportSvc(repoMgr)
		reportService.ListReports(util.ReportType)
		os.Exit(0)
	case util.ExportRulesCmd.FullCommand():
		repoMgr.Rules.ExportRules()
		os.Exit(0)
	case util.ImportRulesCmd.FullCommand():
		repoMgr.Rules.ImportRules()
		os.Exit(0)
	case util.DeleteRulesCmd.FullCommand():
		repoMgr.Rules.DeleteRule(*util.DeleteRulesName)
		os.Exit(0)
	case util.DeleteAllRulesCmd.FullCommand():
		repoMgr.Rules.DeleteAllRules()
		os.Exit(0)
	case util.ValidateRuleCmd.FullCommand():
		repoMgr.Rules.ValidateRule(*util.ValidateRuleName, run)
		os.Exit(0)
	case util.ExportModelsCmd.FullCommand():
		repoMgr.Scoring.ExportModels()
		os.Exit(0)
	case util.ImportModelsCmd.FullCommand():
		repoMgr.Scoring.ImportModels()
		os.Exit(0)
	case util.DeleteModelsCmd.FullCommand():
		repoMgr.Scoring.DeleteModel(*util.DeleteModelName)
		os.Exit(0)
	case util.DeleteAllModelsCmd.FullCommand():
		repoMgr.Scoring.DeleteAllModels(true)
		os.Exit(0)
	case util.ValidateModelCmd.FullCommand():
		repoMgr.Scoring.ValidateModel(*util.ValidateModelName, run)
		os.Exit(0)
	case util.ExportBinsCmd.FullCommand():
		repoMgr.Bins.ExportBins()
		os.Exit(0)
	case util.ImportBinsCmd.FullCommand():
		rules, _ := repoMgr.Rules.GetRules()
		repoMgr.Bins.ImportBins(rules)
		os.Exit(0)
	case util.DeleteBinsCmd.FullCommand():
		repoMgr.Bins.DeleteBin(*util.DeleteBinName)
		os.Exit(0)
	case util.DeleteAllBinsCmd.FullCommand():
		repoMgr.Bins.DeleteAllBins(true)
		os.Exit(0)
	case util.GitCmd.FullCommand():
		run.SetPaths(*util.GitPath)
		fmt.Printf("\nTarget of Git Forensics => %s\n\n", run.Target)
		run.SetRequestedReports(util.ReportsFlag, "1,2,3")
		run.ValidateRun()
		gitReportService := report.NewGitReportService(repoMgr.Run)
		gitReportService.RunGitReports(run)
	case util.AnalyzeCmd.FullCommand():
		run.SetPaths(*util.Path)
		run.SetRequestedReports(util.ReportsFlag, "1,2,3,4,5")
		run.ValidateRun()
		csaService := csa.NewCsaSvc(repoMgr)
		csaService.PerformAnalysis(run)
	case util.NatLangCmd.FullCommand():
		run.SetPaths(*util.NatLangPath)
		run.ValidateRun()
		naturalLanguageService := natural.NewNaturalLanguageService(repoMgr.Run)
		naturalLanguageService.LangProcess(run)
	case util.CsaCmd.FullCommand():
		adminMode = true
		port := util.CsaPort
		routes.StartRouter(run.DB, true, *port)
	case util.SearchCmd.FullCommand():
		adminMode = true
		repoMgr := db.NewRepositoriesManager(run.DB)
		for true {
			search.ExecuteCLISearch(repoMgr)
		}
	case util.BuildInfoCmd.FullCommand():
		fmt.Println("****************************************************************************************")
		fmt.Println("*                             CSA BUILD DETAILS                                      *")
		fmt.Println("****************************************************************************************")
		fmt.Printf("\nVERSION:\t%s\n", Version)
		fmt.Printf("BUILD DATE:\t%s\n", BuildDate)
		fmt.Printf("COMMITTER:\t%s\n", Committer)
		fmt.Printf("COMMIT REF:\t%s\n", CommitSha)
		fmt.Printf("COMMIT SHA:\t%s\n", CommitShaFull)
		fmt.Printf("COMMT MSG:\t%s\n\n", util.LineWrap(CommitMsg, 60))
		adminMode = true
	default:
		err := fmt.Errorf("[%s] is an unknown Cmd!\n", run.Command)
		util.App.FatalUsage("%s\n", err)
		os.Exit(1)
	}
	if !adminMode {
		run.CompletionMessage()
	}
}

func procsAndThreads() {

	mp := runtime.NumCPU()
	mt := 10000

	if *util.MaxProcs > 0 {
		runtime.GOMAXPROCS(*util.MaxProcs)
		mp = *util.MaxProcs
	}

	if *util.MaxThreads > 0 {
		debug.SetMaxThreads(*util.MaxThreads)
		mt = *util.MaxThreads
	}

	if *util.Verbose {
		fmt.Println("		 GO RUNTIME		")
		fmt.Println("-----------------------")
		fmt.Printf("MAX PROCS ==> %d\n", mp)
		fmt.Printf("MAX OS THREADS ==> %d\n\n\n", mt)
	}

	if *util.Verbose {
		logrus.StandardLogger().Level = logrus.DebugLevel
	}

}
