/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package report

import (
	"fmt"
	"os"
	"os/exec"



    "csa-app/db"
	"csa-app/model"
	"csa-app/util"
)

type GitReportService struct {
	runRepository db.RunRepository
}

func NewGitReportService(runRepository db.RunRepository) *GitReportService {
	return &GitReportService{
		runRepository: runRepository,
	}
}

func (gitReportService GitReportService) RunGitReports(run *model.Run) {

	gitReportService.runRepository.StartRun(run)
	//Verfiy Path contains .git!
	if _, err := os.Stat(fmt.Sprintf("%s/.git", run.Target)); err != nil {
		//Repository Path does not contain a .git dir
		util.App.Errorf("%s\n", fmt.Sprintf("Path [%s] does not contain a GIT Repo!", run.Target))
		util.App.Usage(os.Args)
		os.Exit(1)
	}

	if *util.StartDate == "" {
		*util.StartDate = "1950-01-01"
	}

	for _, reportToRun := range run.Reports {
		gitReportService.executeGitReport(run.ID, run.Target, reportToRun, *util.StartDate, *util.OutputDir, util.Verbose)
	}

	gitReportService.runRepository.StopRun(run)
}
func (gitReportService GitReportService) executeGitReport(runId uint, path string, reportNumber int, sinceDate string, reportPath string, debug *bool) {

	checkAndCreateReportDir(reportPath)
	var reportName string
	var gitCommand string
	var extension = "csv"

	switch reportNumber {
	case 1:
		fmt.Println("Git Forensics Report...")
		gitCommand = fmt.Sprintf("git -C %s log --since=%s --pretty=format:'[%%h] %%an %%ad %%s' --date=short --numstat", path, sinceDate)
		reportName = "git-forensics"
		extension = "txt"
	case 2:
		fmt.Println("Git Activity (detailed)...")
		gitWithPath := fmt.Sprintf("git -C %s log --since=%s ", path, sinceDate)
		gitCommand = gitWithPath + `--numstat --all | egrep "^[0-9]" | awk '{ print $3 "," $2 "," $1 }' | sort`
		reportName = "git-activity-detailed"
	case 3:
		fmt.Println("Git Activity (summary)...")
		gitWithPath := fmt.Sprintf("git -C %s log --since=%s ", path, sinceDate)
		gitCommand = gitWithPath + `--numstat --all | egrep "^[0-9]" | awk '{ print $3 }' | sort -u`
		reportName = "git-activity-summary"
	}

	out, _ := exec.Command("bash", "-c", gitCommand).Output()

	file := createReportFile(runId, reportName, extension, reportPath)

	output := string(out)

	if *debug {
		fmt.Printf("%s\n", output)
	}
	fmt.Fprintf(file, "%s\n", output)

}
