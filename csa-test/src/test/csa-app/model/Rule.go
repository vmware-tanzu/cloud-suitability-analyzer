/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"test/csa-app/util"
	"time"
)

type Rule struct {
	ID              uint           `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt       time.Time      `json:"-" yaml:"-"`
	UpdatedAt       time.Time      `json:"-" yaml:"-"`
	Name            string         `gorm:"type:text;unique_index;not null"`
	FileType        string         `gorm:"type:text" json:",omitempty" yaml:",omitempty"` //Extension if empty or * then rule applies to all files. Else, rule only applies to files with this extension (sans '.')
	FileNamePattern string         `gorm:"type:text" json:",omitempty" yaml:",omitempty"`
	Target          string         `gorm:"type:text"`                                     //File, Line
	Type            string         `gorm:"type:text"`                                     //Regex, SimpleText, StartsWith, Contains, EndsWith, SimpleTextCaseInsensitive
	DefaultPattern  string         `gorm:"type:text" json:",omitempty" yaml:",omitempty"` //Pattern with a placeholder that follows standard GO fmt.Sprintf rules. I.E. "[ .]%s[ (]" uses %s for string Pattern Value substitution before compilation!
	Advice          string         `gorm:"type:text" json:",omitempty" yaml:",omitempty"`
	Effort          int            `gorm:"type:bigint; column:effort" json:",omitempty" yaml:",omitempty"`
	Impact          string         `gorm:"type:text" json:",omitempty" yaml:",omitempty"`
	Readiness       int            `gorm:"type:bigint; column:readiness" json:",omitempty" yaml:",omitempty"`
	Category        string         `json:",omitempty" yaml:",omitempty"`
	Criticality     string         `json:",omitempty" yaml:",omitempty"`
	Tags            []Tag          `json:",omitempty" yaml:",omitempty"`
	Recipes         []Recipe       `gorm:"foreignkey:RuleID" json:",omitempty" yaml:",omitempty"`
	Patterns        []Pattern      `gorm:"foreignkey:RuleID"`
	Excludepatterns []Pattern      `gorm:"foreignkey:RuleID"`
	fileNameRegex   *regexp.Regexp `gorm:"-" json:"-" yaml:"-"`
	regex           *regexp.Regexp `gorm:"-" json:"-" yaml:"-"`
	Metric          *RuleMetric    `gorm:"-" json:"-" yaml:"-"`
	overrideApplies bool           `gorm:"-" json:"-" yaml:"-"`
	Negative        bool           `gorm:"type:integer"`
	sync.Mutex      `gorm:"-" json:"-" yaml:"-"`
	//Extra Metadata
	Metadata_sno         string `gorm:"type:text"`
	Metadata_category    string `gorm:"type:text"`
	Metadata_title       string `gorm:"type:text"`
	Metadata_description string `gorm:"type:text"`
	Metadata_group       string `gorm:"type:text"`
	Metadata_criticality string `gorm:"type:text"`
}

func (r *Rule) Applies(fileExt string, fileName string) bool {
	if *util.Verbose {
		fmt.Printf("Checking to see if rule [%s] applies to extension [%s] with pattern [%s], and applies to file name [%s] with pattern [%s]\n", r.Name, fileExt, r.FileType, fileName, r.FileNamePattern)
	}

	if r.overrideApplies {
		return true
	}

	return r.regex.MatchString(fileExt) && r.fileNameRegex.MatchString(fileName)
}

func (r *Rule) IsValid() (isValid bool, err error) {

	//Name is required
	if r.Name == "" {
		return false, fmt.Errorf("Rule Name is required!")
	}

	//FileType
	if r.FileType != "" && r.FileType != "*" {
		//Check FileType is valid regex!
		_, err = regexp.Compile(r.FileType)
		if err != nil {
			return false, fmt.Errorf("FileType must be a empty (nil, not specified), or \"*\", or valid regex! FileType: %s Fails with error: %v", r.FileType, err)
		}
	}

	//FileNamePattern
	if r.FileNamePattern != "" && r.FileNamePattern != "*" {
		//Check FileNamePattern is valid regex!
		_, err = regexp.Compile(r.FileNamePattern)
		if err != nil {
			return false, fmt.Errorf("FileNamePattern must be a empty (nil, not specified), or \"*\", or valid regex! FileType: %s Fails with error: %v", r.FileNamePattern, err)
		}
	}

	//Target
	if r.Target == "" {
		return false, fmt.Errorf("Rule Target is required!")
	}

	if r.Target != FILE_TARGET && r.Target != LINE_TARGET && r.Target != CONTENTS_TARGET {
		return false, fmt.Errorf("Rule Target must be either (%s|%s|%s) but was: %s",
			FILE_TARGET, LINE_TARGET, CONTENTS_TARGET, r.Target)
	}

	//Type
	isValid, err = r.typeIsValid()
	if !isValid {
		return isValid, err
	}

	isValid, err = r.impactIsValid()
	if !isValid {
		return isValid, err
	}

	//Patterns
	if len(r.Patterns) < 1 {
		return false, fmt.Errorf("Rule must have at least one (1) pattern")
	} else {
		for _, pattern := range r.Patterns {
			err = pattern.IsValid(r)
			if err != nil {
				return false, fmt.Errorf("Rule Pattern [%s] Invalid! Details: %s", pattern.ToString(), err.Error())
			}
		}
	}

	isValid, err = r.defaultPatternIsValid()
	if !isValid {
		return isValid, err
	}

	return true, nil
}

func (r *Rule) AddPattern(pattern Pattern) {
	//Is that really all there is to it?

	r.Patterns = append(r.Patterns, pattern)
	if *util.Verbose {
		fmt.Printf("Adding Pattern [%s] to rule [%s] new pattern cnt: %d\n", pattern.Value, r.Name, len(r.Patterns))
	}
}

func (r *Rule) CompilePatterns() {
	r.Lock()
	defer r.Unlock()

	if r.FileType == "" || r.FileType == "*" {
		r.overrideApplies = true
	}

	//Compile the rule file type regex!
	r.regex = regexp.MustCompile(r.FileType)

	//Compile the rule file name regex!
	r.fileNameRegex = regexp.MustCompile(r.FileNamePattern)

	for i, _ := range r.Patterns {
		r.Patterns[i].compile(r)
	}
}

func (r *Rule) GetEscapedPattern() string {
	return util.EscapeSpecials(r.DefaultPattern)
}

func GetRulesDir(isImport bool) string {

	dir := *util.RulesDir

	//User overrode default!
	if dir != util.DEFAULT_RULES_DIR || isImport {
		return dir
	}

	dir = *util.OutputDir
	if !strings.HasSuffix(dir, util.PathSeparator) {
		dir = dir + util.PathSeparator
	}

	return dir + "rules"

}

func (r *Rule) typeIsValid() (isValid bool, err error) {
	//Type
	if r.Type == "" {
		return false, fmt.Errorf("Rule Type is required!")
	}

	if r.Type != REGEX_MATCH_TYPE &&
		r.Type != XPATH_MATCH_TYPE &&
		r.Type != YAMLPATH_MATCH_TYPE &&
		r.Type != PLUGIN_MATCH_TYPE &&
		r.Type != SIMPLE_TEXT_MATCH_TYPE &&
		r.Type != SIMPLE_TEXT_CI_MATCH_TYPE &&
		r.Type != STARTS_WITH_MATCH_TYPE &&
		r.Type != STARTS_WITH_CI_MATCH_TYPE &&
		r.Type != ENDS_WITH_MATCH_TYPE &&
		r.Type != ENDS_WITH_CI_MATCH_TYPE &&
		r.Type != CONTAINS_MATCH_TYPE &&
		r.Type != CONTAINS_CI_MATCH_TYPE {
		return false, fmt.Errorf("Rule Type must be (%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s) but was: %s",
			REGEX_MATCH_TYPE,
			XPATH_MATCH_TYPE,
			YAMLPATH_MATCH_TYPE,
			PLUGIN_MATCH_TYPE,
			SIMPLE_TEXT_CI_MATCH_TYPE,
			SIMPLE_TEXT_MATCH_TYPE,
			STARTS_WITH_CI_MATCH_TYPE,
			STARTS_WITH_MATCH_TYPE,
			ENDS_WITH_CI_MATCH_TYPE,
			ENDS_WITH_MATCH_TYPE,
			CONTAINS_CI_MATCH_TYPE,
			CONTAINS_MATCH_TYPE, r.Target)
	}

	return true, nil

}

func (r *Rule) impactIsValid() (isValid bool, err error) {
	//Impact is not required
	if r.Impact == "" {
		return true, nil
	}

	if r.Impact != EVERY_IMPACT &&
		r.Impact != FILE_IMPACT &&
		r.Impact != APP_IMPACT {
		return false, fmt.Errorf("Rule Impact must be (%s|%s|%s) but was: %s",
			EVERY_IMPACT,
			FILE_IMPACT,
			APP_IMPACT,
			r.Impact)
	}

	return true, nil

}

func (r *Rule) defaultPatternIsValid() (isValid bool, err error) {

	//Contains substitution for patterns...
	if r.DefaultPattern != "" && !strings.Contains(r.DefaultPattern, "%s") {
		return false, fmt.Errorf("If provided Rule DefaultPattern should include a substitution marker of '%%s'! Actual: %s", r.DefaultPattern)
	}

	if r.Type == REGEX_MATCH_TYPE && r.DefaultPattern != "" {

		//Substitute a value!
		pattern := strings.Replace(r.DefaultPattern, "%s", "submarker", 1)

		_, err = regexp.Compile(pattern)
		if err != nil {
			return false, fmt.Errorf("DefaultPattern for '%s' rules must be a valid perl compatible regular expression! Pattern: %s Fails with error: %v", REGEX_MATCH_TYPE, r.DefaultPattern, err)
		}

	}

	if r.DefaultPattern == "" {
		//If I don't have a default pattern I must contain at least one pattern with a pattern!
		if len(r.Patterns) < 1 {
			return false, fmt.Errorf("Rule has no default pattern and no patterns!")
		}
	}

	return true, nil
}
