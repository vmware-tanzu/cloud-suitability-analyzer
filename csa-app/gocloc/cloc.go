/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package gocloc

import (
	"fmt"
	"csa-app/model"
	"csa-app/util"
	"regexp"
	"strings"
)

const OutputTypeDefault string = "default"
const OutputTypeClocXml string = "cloc-xml"
const OutputTypeSloccount string = "sloccount"

const FILE_HEADER string = "File"
const LANG_HEADER string = "Language"
const COMMON_HEADER string = "files          blank        comment           code"
const ROW string = "-------------------------------------------------------------------------" +
	"-------------------------------------------------------------------------" +
	"-------------------------------------------------------------------------"

var rowLen = 79

type CmdOptions struct {
	Byfile         bool   `long:"by-file" description:"report results for every encountered source file"`
	SortTag        string `long:"sort" default:"code" description:"sort based on a certain column"`
	OutputType     string `long:"output-type" default:"default" description:"output type [values: default,cloc-xml,sloccount]"`
	ExcludeExt     string `long:"exclude-ext" description:"exclude file name extensions (separated commas)"`
	IncludeLang    string `long:"include-lang" description:"include language name (separated commas)"`
	MatchDir       string `long:"match-d" description:"include dir name (regex)"`
	NotMatchDir    string `long:"not-match-d" description:"exclude dir name (regex)"`
	Debug          bool   `long:"debug" description:"dump debug log for developer"`
	SkipDuplicated bool   `long:"skip-duplicated" description:"skip duplicated files"`
	ShowLang       bool   `long:"show-lang" description:"print about all languages and extensions"`
}

func ClocEmbedded(path string) *Result {
	var opts CmdOptions

	//Make sure excluded directories match between csa and cloc!
	opts.NotMatchDir = *util.ExcludedDirsRegEx

	clocOpts := util.NewClocOptions()
	// parse command line options

	// value for language result
	languages := util.NewDefinedLanguages()

	// setup option for exclude extensions
	for _, ext := range strings.Split(opts.ExcludeExt, ",") {
		e, ok := util.Exts[ext]
		if ok {
			clocOpts.ExcludeExts[e] = struct{}{}
		} else {
			clocOpts.ExcludeExts[ext] = struct{}{}
		}
	}

	// setup option for not match directory
	if opts.NotMatchDir != "" {
		clocOpts.ReNotMatchDir = regexp.MustCompile(opts.NotMatchDir)
	}
	if opts.MatchDir != "" {
		clocOpts.ReMatchDir = regexp.MustCompile(opts.MatchDir)
	}

	// setup option for include languages
	for _, lang := range strings.Split(opts.IncludeLang, ",") {
		if _, ok := languages.Langs[lang]; ok {
			clocOpts.IncludeLangs[lang] = struct{}{}
		}
	}

	clocOpts.Debug = opts.Debug
	clocOpts.SkipDuplicated = opts.SkipDuplicated

	processor := NewProcessor(languages, clocOpts)
	util.WriteLog("SLOC Analysis", "Scanning Files...")

	result, err := processor.Analyze(path)
	if err != nil {
		return &Result{ErrorMsg: fmt.Sprintf("fail gocloc analyze. error: %v\n", err)}
	}

	return result

}

func ClocEmbeddedByApp(app *model.Application) *Result {
	var opts CmdOptions

	//Make sure excluded directories match between csa and cloc!
	opts.NotMatchDir = *util.ExcludedDirsRegEx

	clocOpts := util.NewClocOptions()
	// parse command line options

	// value for language result
	languages := util.NewDefinedLanguages()

	// setup option for include languages
	for _, lang := range strings.Split(opts.IncludeLang, ",") {
		if _, ok := languages.Langs[lang]; ok {
			clocOpts.IncludeLangs[lang] = struct{}{}
		}
	}

	clocOpts.Debug = opts.Debug
	clocOpts.SkipDuplicated = opts.SkipDuplicated

	processor := NewProcessor(languages, clocOpts)
	if (!*util.Xtract) {
		util.WriteLog("SLOC Analysis", "Scanning Files...")
	}

	result, err := processor.AnalyzeApp(app)
	if err != nil {
		return &Result{ErrorMsg: fmt.Sprintf("fail gocloc analyze. error: %v\n", err)}
	}

	return result

}
