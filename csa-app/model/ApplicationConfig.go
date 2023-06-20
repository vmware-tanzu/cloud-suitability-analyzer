/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"csa-app/util"

	"gopkg.in/yaml.v2"
)

const APP_CONFIG_FILENAME = "csa-config"

type ApplicationConfig struct {
	Name             string           `json:"name,required" yaml:"name"`
	Path             string           `json:"path,required" yaml:"path"`
	Category         string           `json:"category,omitempty" yaml:"category,omitempty"`
	Criticality      string           `json:"criticality,omitempty" yaml:"criticality,omitempty"`
	BusinessDomain   string           `json:"business-domain,omitempty" yaml:"business-domain,omitempty"`
	BusinessValue    float64          `json:"business-value,omitempty" yaml:"business-value,omitempty"`
	ScoringModel     string           `json:"scoring-model,omitempty" yaml:"scoring-model,omitempty"`
	RuleIncludeTags  string           `json:"rule-include-tags,omitempty" yaml:"rule-include-tags,omitempty"`
	RuleExcludeTags  string           `json:"rule-exclude-tags,omitempty" yaml:"rule-exclude-tags,omitempty"`
	DirExcludeRegex  string           `json:"dir-exclude-regex,omitempty" yaml:"dir-exclude-regex,omitempty"`
	IncludeFileRegex string           `json:"include-file-regex,omitempty" yaml:"include-file-regex,omitempty"`
	ExcludeFileRegex string           `json:"exclude-file-regex,omitempty" yaml:"exclude-file-regex,omitempty"`
	Files            []*util.FileInfo `json:"-" yaml:"-"`
	IgnoredFiles     []*util.FileInfo `json:"-" yaml:"-"`
	FileUtil         *util.FileUtil   `json:"-" yaml:"-"`
	Profiles  string                  `json:"profiles,omitempty" yaml:"profiles,omitempty"`
}

type configFileConfig struct {
	Name             string  `json:"name,required" yaml:"name"`
	Category         string  `json:"category,omitempty" yaml:"category,omitempty"`
	Criticality      string  `json:"criticality,omitempty" yaml:"criticality,omitempty"`
	BusinessDomain   string  `json:"business-domain,omitempty" yaml:"business-domain,omitempty"`
	BusinessValue    float64 `json:"business-value,omitempty" yaml:"business-value,omitempty"`
	ScoringModel     string  `json:"scoring-model,omitempty" yaml:"scoring-model,omitempty"`
	RuleIncludeTags  string  `json:"rule-include-tags" yaml:"rule-include-tags"`
	RuleExcludeTags  string  `json:"rule-exclude-tags" yaml:"rule-exclude-tags"`
	DirExcludeRegex  string  `json:"dir-exclude-regex" yaml:"dir-exclude-regex"`
	IncludeFileRegex string  `json:"include-file-regex" yaml:"include-file-regex"`
	ExcludeFileRegex string  `json:"exclude-file-regex" yaml:"exclude-file-regex"`
	Profiles  string         `json:"profiles" yaml:"profiles"`
}

func NewApplicationConfig(runConfig *RunConfig) *ApplicationConfig {
	newAC := &ApplicationConfig{BusinessValue: -1.0}
	newAC.MergeRunConfig(runConfig)
	return newAC
}

func (c *ApplicationConfig) AddFile(f os.FileInfo, path string) {

	//Run this check for every file because it is a waste to actually analyze archives/binary files
	isArchive := c.FileUtil.IsDecompilableArchive(path)
	if isArchive && *util.AnalyzeArchives {
		util.WriteLog("Gathering Files", "Found Archive [%s] @ %s within app [%s]\n", f.Name(), path, c.Name)
		decompilePath, _, decompiled := c.FileUtil.CheckForArchive(path)
		if decompiled {
			c.gatherFilesOnPath(decompilePath)
			//File is decompiled and added so return
			return
		} else {
			if (!*util.Xtract) {
				util.WriteLog("Gathering Files", "Archive [%s] @ %s within app [%s] could not be decompiled!\n", f.Name(), path, c.Name)
			}
		}
	}

	ext := filepath.Ext(f.Name())
	fInfo := util.NewFileInfo(
		c.Name,
		path,
		f.Name(),
		ext,
		filepath.Base(path),
		"",
		true)

	if !isArchive {

		if c.FileUtil.FileShouldBeProcessed(f.Name()) {
			c.Files = append(c.Files, fInfo)
			util.WriteLog("Gathering Files", "Found File [%s]\n", f.Name())
		} else {
			msg := fmt.Sprintf("Found File [%s] @ %s within app [%s]. Ignoring as it is excluded!\n", f.Name(), path, c.Name)
			if (!*util.Xtract) {
				util.WriteLog("Gathering Files", msg)
			}
			c.IgnoredFiles = append(c.IgnoredFiles, fInfo)
		}
	} else {
		msg := fmt.Sprintf("Found Archive [%s] @ %s within app [%s]. Ignoring as 'analyze-archives' flag is false.\n", f.Name(), path, c.Name)
		util.WriteLog("Gathering Files", msg)
		c.IgnoredFiles = append(c.IgnoredFiles, fInfo)
	}
}

func (c *ApplicationConfig) GatherFiles() error {
	return c.gatherFilesOnPath(c.Path)
}

func (c *ApplicationConfig) Validate() {
	if !util.Exists(c.Path) {
		fmt.Fprintf(os.Stderr, "\nApplication [%s] Path [%s] does not exist!\n\n", c.Name, c.Path)
		os.Exit(23)
	}

	if c.Name == "" {
		c.Name = filepath.Base(c.Path)
	}
}

func (c *ApplicationConfig) MergeRunConfig(runConfig *RunConfig) {

	c.FileUtil = runConfig.FileUtil.CloneFileUtil()

	ac := runConfig.GetApplication(c.Name)

	if ac != nil {

		if c.ScoringModel == "" {
			c.ScoringModel = ac.ScoringModel
		}

		if c.RuleIncludeTags == "" {
			c.RuleIncludeTags = ac.RuleIncludeTags
		}

		if c.Profiles == "" {
			c.Profiles = ac.Profiles
		}

		if c.RuleExcludeTags == "" {
			c.RuleExcludeTags = ac.RuleExcludeTags
		}

		if c.DirExcludeRegex == "" {
			c.DirExcludeRegex = ac.DirExcludeRegex
		}

		if c.IncludeFileRegex == "" {
			c.IncludeFileRegex = ac.IncludeFileRegex
		}

		if c.RuleIncludeTags == "" {
			c.ExcludeFileRegex = ac.ExcludeFileRegex
		}
	}

	if c.ScoringModel == "" {
		c.ScoringModel = runConfig.ScoringModel
	}

	//Check for any empty and override from base config
	if c.RuleIncludeTags == "" {
		c.RuleIncludeTags = runConfig.RuleIncludeTags
	}

	if c.Profiles == "" {
		c.Profiles = runConfig.Profiles
	}

	if c.RuleExcludeTags == "" {
		c.RuleExcludeTags = runConfig.RuleExcludeTags
	}

	if c.DirExcludeRegex == "" {
		c.DirExcludeRegex = runConfig.DirExcludeRegex
	}

	if c.IncludeFileRegex == "" {
		c.IncludeFileRegex = runConfig.IncludeFileRegex
	}

	if c.RuleIncludeTags == "" {
		c.ExcludeFileRegex = runConfig.ExcludeFileRegex
	}

	//Update the file utils regex definitions!
	c.reCompileRegex()
}

func (c *ApplicationConfig) Merge(mergeConfig *configFileConfig) {

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.SCORING_MODEL_FLAG) && mergeConfig.ScoringModel != "" {
		c.ScoringModel = mergeConfig.ScoringModel
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.RULE_INCLUDE_FLAG) && mergeConfig.RuleIncludeTags != "" {
		c.RuleIncludeTags = mergeConfig.RuleIncludeTags
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.PROFILE_FLAG) && mergeConfig.Profiles != "" {
		c.Profiles = mergeConfig.Profiles
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.RULE_EXCLUDE_FLAG) && mergeConfig.RuleExcludeTags != "" {
		c.RuleExcludeTags = mergeConfig.RuleExcludeTags
	}

	if util.IsAppFlagDefaulted(util.EXCLUDED_DIRS_FLAG) && mergeConfig.DirExcludeRegex != "" {
		c.DirExcludeRegex = mergeConfig.DirExcludeRegex
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.INCLUDED_FILES) && mergeConfig.IncludeFileRegex != "" {
		c.IncludeFileRegex = mergeConfig.IncludeFileRegex
	}

	if util.IsCmdFlagDefaulted(util.ANALYZE_CMD, util.EXCLUDED_FILES) && mergeConfig.ExcludeFileRegex != "" {
		c.ExcludeFileRegex = mergeConfig.ExcludeFileRegex
	}

	if mergeConfig.BusinessDomain != "" {
		c.BusinessDomain = mergeConfig.BusinessDomain
	}

	if mergeConfig.BusinessValue > 0 {
		c.BusinessValue = mergeConfig.BusinessValue
	}

	if mergeConfig.Category != "" {
		c.Category = mergeConfig.Category
	}

	if mergeConfig.Criticality != "" {
		c.Criticality = mergeConfig.Criticality
	}

	if mergeConfig.Name != "" {
		c.Name = mergeConfig.Name
	}

	//Update the file utils regex definitions!
	c.reCompileRegex()
}

/***********************************************************************************************************************
													PRIVATE API
************************************************************************************************************************/

func (c *ApplicationConfig) checkForAndPrepareArchiveTarget(path string) (finalTargetPath string, decompiled bool) {

	finalTargetPath, _, decompiled = c.FileUtil.CheckForArchive(path)
	return
}

func (c *ApplicationConfig) gatherFilesOnPath(targetPath string) error {

	//Check for Archives...jar,war,ear
	finalTargetPath, _ := c.checkForAndPrepareArchiveTarget(targetPath)

	return filepath.Walk(finalTargetPath, func(path string, f os.FileInfo, err error) error {
		if err == nil {
			if f.IsDir() {
				if c.FileUtil.DirIsExcluded(f.Name()) {
					return filepath.SkipDir
				}
			} else {
				c.AddFile(f, path)
			}

		}
		return nil
	})
}

func (c *ApplicationConfig) CheckForLocalAppConfig() {

	//Do I Have an App Config File?
	configFile, _ := util.GetFileIfExists(c.Path + util.PathSeparator + APP_CONFIG_FILENAME + ".yaml")

	if !configFile.Exists {
		configFile, _ = util.GetFileIfExists(c.Path + util.PathSeparator + APP_CONFIG_FILENAME + ".yml")
		if !configFile.Exists {
			configFile, _ = util.GetFileIfExists(c.Path + util.PathSeparator + APP_CONFIG_FILENAME + ".json")
		}
	}

	if configFile.Exists {
		var decoder util.FileDecoder
		reader, err := os.Open(configFile.FQN)
		if err != nil {
			util.App.Fatalf("Unable to open/read app config file @ [%s]! Details: %v\n", *util.ConfigFile, err)
		}

		if strings.HasSuffix(configFile.Name, util.JSON) {
			decoder = json.NewDecoder(reader)
		} else {
			decoder = yaml.NewDecoder(reader)
		}

		newConfig := &configFileConfig{}
		err = decoder.Decode(&newConfig)

		c.Merge(newConfig)
	}

	if *util.WriteAppConfig || *util.WriteConfigsOnly {

		nc := configFileConfig{c.Name, c.Category,
			c.Criticality, c.BusinessDomain,
			c.BusinessValue, c.ScoringModel,
			c.RuleIncludeTags, c.RuleExcludeTags,
			c.DirExcludeRegex, c.IncludeFileRegex,
			c.ExcludeFileRegex, c.Profiles}

		format := util.YAML
		if *util.OutputFormatJson {
			format = util.JSON
		}

		if *util.WriteConfigsOnly {
			fmt.Printf("Writing App Config for [%s] to %s%s%s.%s\n", c.Name, c.Path, util.PathSeparator, APP_CONFIG_FILENAME, format)
		}

		util.WriteStructToFile(nc, APP_CONFIG_FILENAME, c.Path, format, false)
	}

}

func (c *ApplicationConfig) reCompileRegex() {

	var err error

	if c.IncludeFileRegex != ".*" {
		c.FileUtil.IncludedFilesRegEx, err = regexp.Compile(c.IncludeFileRegex)
		if err != nil {
			util.App.Fatalf("Invalid 'included-files' regex in run config-file! Details: %v", err)
		}
	}

	c.FileUtil.ExcludedDirsRegExp, err = regexp.Compile(c.DirExcludeRegex)
	if err != nil {
		util.App.Fatalf("Invalid 'excluded-dirs' regex in run config-file! Details: %v", err)
	}

	c.FileUtil.ExcludedFilesRegEx, err = regexp.Compile(c.ExcludeFileRegex)
	if err != nil {
		util.App.Fatalf("Invalid 'excluded-files' regex in run config-file! Details: %v", err)
	}
}
