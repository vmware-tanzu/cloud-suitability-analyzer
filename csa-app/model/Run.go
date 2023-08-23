/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"csa-app/util"
	"github.com/jinzhu/gorm"
	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
)

type Run struct {
	ID               uint                      `gorm:"primary_key" json:"id" yaml:"id"`
	Alias            string                    `gorm:"type:text"`
	CreatedAt        time.Time                 `json:"-" yaml:"-"`
	UpdatedAt        time.Time                 `json:"-" yaml:"-"`
	User             string                    `gorm:"type:text"`
	Command          string                    `gorm:"type:text"`
	Target           string                    `gorm:"type:text"`
	ReportsRequested string                    `gorm:"type:text"`
	Files            int                       `json:"Files" yaml:"Files"`
	Findings         int                       `json:"Findings" yaml:"Findings"`
	AnalyzedCnt      int                       `gorm:"-" json:"-" yaml:"-"`
	StartTime        time.Time                 `gorm:"-" json:"-" yaml:"-"`
	RequestDateTime  string                    `gorm:"-" json:"requestDate" yaml:"requestDate"`
	Runtime          string                    `gorm:"type:text"`
	Reports          []int                     `gorm:"-" json:"-" yaml:"-"`
	Homepath         string                    `gorm:"-"`
	Exepath          string                    `gorm:"-"`
	OutputPath       string                    `gorm:"-"`
	RulesDir         string                    `gorm:"-"`
	DbPath           string                    `gorm:"-"`
	TmpPath          string                    `gorm:"-"`
	RulesImport      bool                      `gorm:"-" json:"-" yaml:"-"`
	Function         reportFunction            `gorm:"-" json:"-" yaml:"-"`
	Applications     []*Application            `gorm:"foreignkey:RunID" json:",omitempty" yaml:",omitempty"`
	Activities       map[string]*util.Activity `gorm:"-" json:"-" yaml:"-"`
	FileUtil         *util.FileUtil            `gorm:"-" json:"-" yaml:"-"`
	DB               *gorm.DB                  `gorm:"-" json:"-" yaml:"-"`
	Rules            []Rule                    `gorm:"-" json:"-" yaml:"-"`
	UnknownExts      []string                  `gorm:"-" json:"-" yaml:"-"`
	LineBufferSize   int                       `gorm:"-" json:"-" yaml:"-"`
	sync.Mutex       `gorm:"-" json:"-" yaml:"-"`
}

type reportFunction func(run *Run)

func NewRun() *Run {
	var err error

	newRun := &Run{}
	newRun.Command, err = util.App.Parse(os.Args[1:])

	context, _ := util.App.ParseContext(os.Args[1:])

	if err != nil {
		util.App.FatalUsageContext(context, "Error parsing command arguments! %s\n", err)
		os.Exit(1)
	}

	newRun.SetupDefaults()
	newRun.Activities = make(map[string]*util.Activity)
	newRun.FileUtil = util.NewFileUtil()
	//Establish Temp Directory for Run
	newRun.TmpPath, err = newRun.FileUtil.EstablishTempDirectory("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create temp directory. Details: %v", err)
		os.Exit(1)
	}

	//Store line buffer size for file scanning
	newRun.LineBufferSize = *util.LineBuffer
	if newRun.LineBufferSize > util.MAX_LINE_BUFFER_SIZE {
		newRun.LineBufferSize = util.MAX_LINE_BUFFER_SIZE
	}

	return newRun
}

func (r *Run) Cleanup() {

	if r.TmpPath != "" && r.FileUtil != nil {
		r.FileUtil.RemoveDir(r.TmpPath)
	}

	if r.DB != nil {
		r.DB.Close()
	}

}

func (r Run) CompletionMessage() {
	if r.Files == 0 {
		if (!*util.Xtract) {
			fmt.Printf("\n******** All done! Report Run [%d] Completed in %s ****************\n", r.ID, r.Runtime)
		}
	} else {

		completed := "Completed"
		if util.HasErrors() {
			completed = "Completed w/Errors"
		}
		if (!*util.Xtract) {
			fmt.Printf("\n******** All done! [%d] Findings in [%d] Files. Report Run [%d-%s] %s in %s ****************\n", r.Findings, r.Files, r.ID, r.GetAlias(), completed, r.Runtime)
			fmt.Printf("\n\nRun: %d\nAlias: %s\nUser: %s\nCommand: %s\nTarget: %s\nFiles: %d\nFindings: %d\nRuntime: %s\nDB Path: %s\nRules-Dir: %s\nOutputPath: %s\nExe Path: %s\nTmp Path: %s\n\n",
				r.ID,
				r.GetAlias(),
				r.User,
				r.Command,
				r.Target,
				r.Files,
				r.Findings,
				r.Runtime,
				r.DbPath,
				r.RulesDir,
				r.OutputPath,
				r.Exepath,
				r.TmpPath)
		}
	}
}

func (r *Run) StartActivity(name string) {
	r.Lock()
	if _, exists := r.Activities[name]; !exists {
		r.Activities[name] = util.NewActivity(name)
	}
	r.Unlock()
}

func (r *Run) StopActivity(name, msg string, printMsg bool) {
	r.StopActivityLF(name, msg, true, printMsg)
}

func (r *Run) StopActivityLF(name, msg string, prelinefeed bool, printMsg bool) {

	var activity *util.Activity
	var exists bool

	r.Lock()
	if activity, exists = r.Activities[name]; !exists {
		activity = util.NewActivity(name)
		r.Activities[name] = activity
	}
	r.Unlock()
	activity.Stop()

	if printMsg {
		if (!*util.Xtract){
			if prelinefeed {
				fmt.Printf("\n%s (%s)\n", msg, fmt.Sprintf("%v", activity.GetElapsed()))
			} else {
				fmt.Printf("%s (%s)\n", msg, fmt.Sprintf("%v", activity.GetElapsed()))
			}
		}
	}
}

func (r *Run) PrepForMarshal() {
	r.RequestDateTime = r.CreatedAt.Format("01/02/2006 15:04:05")
}

func (r *Run) AddFindings(count int) {
	r.Lock()
	r.Findings += count
	r.Unlock()
}

func (r *Run) FileAnalyzed() {
	r.Lock()
	r.AnalyzedCnt++
	r.Unlock()
}

func (r *Run) SetRequestedReports(list *string, defaultList string) {
	//If running default convert Reports to Full Slice
	if *list == "0" || strings.Contains(*list, "0") {
		r.ReportsRequested = defaultList
	} else {
		r.ReportsRequested = *list
	}

	r.setRequestedReports()
}

func (r *Run) setRequestedReports() {
	var reportsList = strings.Split(r.ReportsRequested, ",")
	var reportList = make([]int, len(reportsList))

	idx := 0
	for _, r := range reportsList {
		i, _ := strconv.Atoi(r)
		reportList[idx] = i
		idx++
	}
	sort.Ints(reportList)
	r.Reports = reportList
}

func (r *Run) SetPaths(targetPath string) {
	r.Homepath, _ = homedir.Dir()
	//--- convert relative path to absolute: remove tilde
	path := strings.Replace(targetPath, "~", r.Homepath, 1)
	//--- convert relative path to absolute: remove dot
	r.Target, _ = filepath.Abs(path)
	r.SetAlias(filepath.Base(path)) //Get the directory name targeted
	r.Exepath, _ = osext.ExecutableFolder()
	r.OutputPath, _ = filepath.Abs(*util.OutputDir)

	importMode := true
	switch r.Command {
	case util.ExportRulesCmd.FullCommand():
		importMode = false
	}

	r.RulesDir, _ = filepath.Abs(GetRulesDir(importMode))

	//--- check for debug mode, exe directory will be controlled by ide
	if strings.Contains(r.Exepath, "private") {
		r.Exepath, _ = filepath.Abs(r.Homepath + "/csa")
	}

	User, err := user.Current()
	if err != nil {
		r.User = fmt.Sprintf("unknown - error retrieving user! Details: %v", err)
	} else {
		r.User = User.Username
	}
}

func (r *Run) SetAlias(alias string) {
	if *util.Alias == "" {
		r.Alias = alias
	} else {
		r.Alias = *util.Alias
	}
}

func (r Run) GetAlias() string {
	return r.Alias
}

func (r *Run) PerformAnalysis() {
	r.Function(r)
}

func (r *Run) ValidateRun() {
	//--- nothing works without a valid path
	if !strings.Contains(r.Target, "*.") && !util.Exists(r.Target) {
		util.App.FatalUsage("%s\n", fmt.Sprintf("Path [%s] does not exist!", r.Target))
	}
}

func (r *Run) SetupDefaults() {

	r.SetPaths("")

	if *util.DbDir == "" {
		util.DbDir = &r.Exepath
	}

	if *util.FernLocation == "" {
		util.FernLocation = &r.Exepath
	}
}

func (r *Run) AssociateApplication(application *Application) {
	application.RunID = r.ID
	r.Applications = append(r.Applications, application)
}

func (r *Run) AppsOrdered() []*Application {
	sort.Sort(ByFileCount(r.Applications))
	return r.Applications
}
