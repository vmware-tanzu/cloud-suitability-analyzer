/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"csa-app/util"
	"gopkg.in/yaml.v2"
)

type RunConfig struct {
	Run              *Run                 `json:"-" yaml:"-"`
	Alias            string               `json:"runName" yaml:"runName"`
	Applications     []*ApplicationConfig `json:"applications,required" yaml:"applications"`
	ScoringModel     string               `json:"scoring-model" yaml:"scoring-model"`
	RuleIncludeTags  string               `json:"rule-include-tags" yaml:"rule-include-tags"`
	RuleExcludeTags  string               `json:"rule-exclude-tags" yaml:"rule-exclude-tags"`
	DirExcludeRegex  string               `json:"dir-exclude-regex" yaml:"dir-exclude-regex"`
	IncludeFileRegex string               `json:"include-file-regex" yaml:"include-file-regex"`
	ExcludeFileRegex string               `json:"exclude-file-regex" yaml:"exclude-file-regex"`
	FileUtil         *util.FileUtil       `json:"-" yaml:"-"`
	Profiles  string                      `json:"profiles" yaml:"profiles"`
}

func NewRunConfig(run *Run, fileUtil *util.FileUtil) *RunConfig {
	return &RunConfig{
		run,
		run.Alias,
		nil,
		*util.ScoringModel,
		*util.RuleIncludeTags,
		*util.RuleExcludeTags,
		*util.ExcludedDirsRegEx,
		*util.IncludedFilesRegEx,
		*util.ExcludedFilesRegEx,
		fileUtil,
		*util.Profiles,
	}
}

func (rc *RunConfig) UnMarshall() {
	if *util.ConfigFile != "" {
		//Attempt to marshall from file contents
		rc.populateFromConfigFile()
	} else {
		if *util.DiscoveryMode {
			rc.Applications = rc.portfolioDiscovery(rc.Run.Target)
			if (!*util.Xtract) {
				fmt.Printf("Discovered [%d] apps @ path [%s]\n\n", len(rc.Applications), rc.Run.Target)
			}
		} else {
			appConfig := NewApplicationConfig(rc)
			appConfig.Path = rc.Run.Target
			appConfig.Name = filepath.Base(rc.Run.Target)
			appConfig.CheckForLocalAppConfig()
			rc.Applications = append(rc.Applications, appConfig)
			if (!*util.Xtract) {
				fmt.Printf("Targeting App => %s located @ path [%s]\n\n", appConfig.Name, appConfig.Path)
			}
		}
	}

	configName := fmt.Sprintf("%d-run-config", rc.Run.ID)
	configPath := rc.Run.Homepath + util.PathSeparator + "csa"
	format := util.YAML
	if *util.OutputFormatJson {
		format = util.JSON
	}
	util.WriteStructToFile(rc, configName, configPath, format, true)

	if *util.WriteConfigsOnly {
		fmt.Printf("App Configs written to app directories and Run Config written to %s%s%s.%s\n", configPath, util.PathSeparator, configName, format)
	}

}

func (rc *RunConfig) GetApplication(name string) *ApplicationConfig {

	for _, ac := range rc.Applications {
		if ac.Name == name {
			return ac
		}
	}

	return nil
}

func (rc *RunConfig) Populate() {
	waitGroup := sync.WaitGroup{}
	for i := range rc.Applications {
		waitGroup.Add(1)
		go func(idx int) {
			defer waitGroup.Done()
			if (!*util.Xtract) {
				util.WriteLog("Gathering", "Gathering Files for App[%s]\n", rc.Applications[idx].Name)
			}
			err := rc.Applications[idx].GatherFiles()
			if err != nil {
				log.Panicf("Error Gathering files for App[%s]...Details: %v\n", rc.Applications[idx].Name, err)
			}
		}(i)
	}

	waitGroup.Wait()
}

func (rc *RunConfig) Merge(mergeConfig *RunConfig) {
	rc.Alias = mergeConfig.Alias
	rc.Applications = mergeConfig.Applications

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.SCORING_MODEL_FLAG) && mergeConfig.ScoringModel != "" {
		rc.ScoringModel = mergeConfig.ScoringModel
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.RULE_INCLUDE_FLAG) && mergeConfig.RuleIncludeTags != "" {
		rc.RuleIncludeTags = mergeConfig.RuleIncludeTags
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.RULE_EXCLUDE_FLAG) && mergeConfig.RuleExcludeTags != "" {
		rc.RuleExcludeTags = mergeConfig.RuleExcludeTags
	}

	if util.IsAppFlagDefaulted(util.EXCLUDED_DIRS_FLAG) && mergeConfig.DirExcludeRegex != "" {
		rc.DirExcludeRegex = mergeConfig.DirExcludeRegex
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.INCLUDED_FILES) && mergeConfig.IncludeFileRegex != "" {
		rc.IncludeFileRegex = mergeConfig.IncludeFileRegex
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.EXCLUDED_FILES) && mergeConfig.ExcludeFileRegex != "" {
		rc.ExcludeFileRegex = mergeConfig.ExcludeFileRegex
	}

	//Update the file utils regex definitions!
	rc.reCompileRegex()
}

/***********************************************************************************************************************
													PRIVATE API
************************************************************************************************************************/

func (rc *RunConfig) reCompileRegex() {

	var err error

	if strings.Compare(rc.IncludeFileRegex, *util.IncludedFilesRegEx) != 0 {
		rc.FileUtil.IncludedFilesRegEx, err = regexp.Compile(rc.IncludeFileRegex)
		if err != nil {
			util.App.Fatalf("Invalid 'included-files' regex in run config-file! Details: %v", err)
		}
		//Update the shared config so any subsequent file utils and cloc get these settings!
		util.IncludedFilesRegEx = &rc.IncludeFileRegex
	}

	if strings.Compare(rc.DirExcludeRegex, *util.ExcludedDirsRegEx) != 0 {
		rc.FileUtil.ExcludedDirsRegExp, err = regexp.Compile(rc.DirExcludeRegex)
		rc.FileUtil.DecompilingExcludedDirsRegExp, err = regexp.Compile(strings.Replace(rc.DirExcludeRegex, "classes|", "", -1))
		if err != nil {
			util.App.Fatalf("Invalid 'excluded-dirs' regex in run config-file! Details: %v", err)
		}
		//Update the shared config so any subsequent file utils and cloc get these settings!
		util.ExcludedDirsRegEx = &rc.DirExcludeRegex
	}

	if strings.Compare(rc.ExcludeFileRegex, *util.ExcludedFilesRegEx) != 0 {
		rc.FileUtil.ExcludedFilesRegEx, err = regexp.Compile(rc.ExcludeFileRegex)
		if err != nil {
			util.App.Fatalf("Invalid 'excluded-files' regex in run config-file! Details: %v", err)
		}
		//Update the shared config so any subsequent file utils and cloc get these settings!
		util.ExcludedFilesRegEx = &rc.ExcludeFileRegex
	}
}

func (rc *RunConfig) portfolioDiscovery(searchDir string) (appConfigs []*ApplicationConfig) {

	targetPath, _ := rc.checkForAndPrepareArchiveTarget(nil, searchDir)

	portfolio, err := os.Stat(targetPath)

	if err == nil {

		if portfolio.IsDir() {

			files, err := ioutil.ReadDir(targetPath)

			if err == nil {
				for _, f := range files {
					if f.IsDir() {
						if *util.Verbose {
							fmt.Printf("Found dir [%s] in porfolio dir [%s]\n", f.Name(), targetPath)
						}
						//Check for excluded directory
						if !rc.FileUtil.DirIsExcluded(f.Name()) {
							newAppConfig := NewApplicationConfig(rc)
							newAppConfig.Name = f.Name()
							newAppConfig.Path = targetPath + util.PathSeparator + f.Name()
							newAppConfig.CheckForLocalAppConfig()
							appConfigs = append(appConfigs, newAppConfig)
						}
					}
				}
			}

		} else {
			util.App.Fatalf("Unable to perform portfolio discovery on a non-directory target [%s]!", searchDir)
		}
	} else {
		util.App.Fatalf("Error discovering portfolio @[%s]...Details: %v", targetPath, err)
	}

	return
}

func (rc *RunConfig) checkForAndPrepareArchiveTarget(run *Run, path string) (finalTargetPath string, decompiled bool) {

	finalTargetPath, alias, decompiled := rc.FileUtil.CheckForArchive(path)

	if decompiled && run != nil {
		run.SetAlias(alias)
	}

	return
}

func (rc *RunConfig) populateFromConfigFile() {
	file := util.GetFile(*util.ConfigFile)

	fmt.Printf("Attempting to load run config @[%s]...\n", file.FQN)

	reader, err := os.Open(file.FQN)
	if err != nil {
		util.App.Fatalf("Unable to open/read config file @ [%s]! Details: %v\n", *util.ConfigFile, err)
	}

	newConfig := &RunConfig{}

	decodeConfig(file.Name, reader, newConfig)

	rc.Merge(newConfig)

	if len(rc.Applications) < 1 {
		util.App.Fatalf("Invalid RunConfig @ [%s]! Details: %s\n", *util.ConfigFile, "At least one application must be specified within the 'applications' collection")
	}

	for i := range rc.Applications {
		rc.Applications[i].MergeRunConfig(rc)
		rc.Applications[i].Validate()
	}
}

func decodeConfig(fileName string, reader *os.File, newConfig *RunConfig) {

	var decoder util.FileDecoder
	var err error

	if util.IsJson(fileName) {
		decoder = json.NewDecoder(reader)
		err = decoder.Decode(&newConfig)
	} else if util.IsYaml(fileName) {
		decoder = yaml.NewDecoder(reader)
		err = decoder.Decode(&newConfig)
	} else {

		//try json
		decoder = json.NewDecoder(reader)
		err = decoder.Decode(&newConfig)

		if err != nil {
			//try yaml
			decoder = yaml.NewDecoder(reader)
			err = decoder.Decode(&newConfig)
		}

		if err != nil {
			util.App.Fatalf("Invalid RunConfig @ [%s]! At this time only %s and %s files/formats are supported!", fileName, util.JSON, util.YAML)
		}
	}

	if err != nil {
		util.App.Fatalf("Invalid RunConfig @ [%s]! Details: %v\n", *util.ConfigFile, err)
	}
}
