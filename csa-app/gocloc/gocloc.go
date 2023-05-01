/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package gocloc

import (
	"csa-app/model"
	"csa-app/util"
)

type Processor struct {
	langs *util.DefinedLanguages
	opts  *util.ClocOptions
}

type Result struct {
	Total         *util.Language
	Files         map[string]*ClocFile
	Domains       map[string]map[string]*util.Language
	MaxPathLength int
	ErrorMsg      string
	UnknownExts   []string
}

func NewProcessor(langs *util.DefinedLanguages, options *util.ClocOptions) *Processor {
	return &Processor{
		langs: langs,
		opts:  options,
	}
}

func (p *Processor) Analyze(path string) (*Result, error) {
	total := util.NewLanguage("TOTAL", []string{}, "", "")
	languages, err := getAllFiles(path, p.langs, p.opts)
	if err != nil {
		return nil, err
	}
	maxPathLen := 0
	num := 0
	for _, lang := range languages {
		num += len(lang.Files)
		for _, file := range lang.Files {
			l := len(file.FQN)
			if maxPathLen < l {
				maxPathLen = l
			}
		}
	}
	clocFiles := make(map[string]*ClocFile, num)
	domainTotals := make(map[string]map[string]*util.Language)

	for _, language := range languages {
		for _, file := range language.Files {
			cf := AnalyzeFile(*file, language, p.opts)
			clocFiles[file.FQN] = cf

			if _, ok := domainTotals[cf.Dir]; !ok {
				domainTotals[cf.Dir] = make(map[string]*util.Language)
			}

			if _, ok := domainTotals[cf.Dir][cf.Lang]; !ok {
				domainTotals[cf.Dir][cf.Lang] = util.NewLanguage(
					language.Name,
					language.LineComments,
					language.MultiLine,
					language.MultiLineEnd)
			}

			domainTotals[cf.Dir][cf.Lang].Code += cf.Code
			domainTotals[cf.Dir][cf.Lang].Comments += cf.Comments
			domainTotals[cf.Dir][cf.Lang].Blanks += cf.Blanks
			domainTotals[cf.Dir][cf.Lang].Total += (cf.Code + cf.Comments + cf.Blanks)

			domainTotals[cf.Dir][cf.Lang].Files = append(domainTotals[cf.Dir][cf.Lang].Files, file)
			util.WriteLog("SLOC Analysis", "CLOC of File [%s]done!\n", file.Name)

		}

		files := int32(len(language.Files))
		if len(language.Files) <= 0 {
			continue
		}

		total.Total += files
		total.Blanks += language.Blanks
		total.Comments += language.Comments
		total.Code += language.Code
	}

	return &Result{
		Total:         total,
		Files:         clocFiles,
		Domains:       domainTotals,
		MaxPathLength: maxPathLen,
	}, nil
}

func (p *Processor) AnalyzeApp(app *model.Application) (*Result, error) {
	total := util.NewLanguage("TOTAL", []string{}, "", "")
	languages, unknowns, err := getAllLangsForFiles(app, p.langs, p.opts)
	if err != nil {
		return nil, err
	}
	maxPathLen := 0
	num := 0
	for _, lang := range languages {
		num += len(lang.Files)
		for _, file := range lang.Files {
			l := len(file.FQN)
			if maxPathLen < l {
				maxPathLen = l
			}
		}
	}
	clocFiles := make(map[string]*ClocFile, num)
	domainTotals := make(map[string]map[string]*util.Language)

	for _, language := range languages {
		for _, file := range language.Files {
			cf := AnalyzeFile(*file, language, p.opts)
			clocFiles[file.FQN] = cf

			if _, ok := domainTotals[cf.Dir]; !ok {
				domainTotals[cf.Dir] = make(map[string]*util.Language)
			}

			if _, ok := domainTotals[cf.Dir][cf.Lang]; !ok {
				domainTotals[cf.Dir][cf.Lang] = util.NewLanguage(
					language.Name,
					language.LineComments,
					language.MultiLine,
					language.MultiLineEnd)
			}

			domainTotals[cf.Dir][cf.Lang].Code += cf.Code
			domainTotals[cf.Dir][cf.Lang].Comments += cf.Comments
			domainTotals[cf.Dir][cf.Lang].Blanks += cf.Blanks
			domainTotals[cf.Dir][cf.Lang].Total += (cf.Code + cf.Comments + cf.Blanks)

			domainTotals[cf.Dir][cf.Lang].Files = append(domainTotals[cf.Dir][cf.Lang].Files, file)
			if (!*util.Xtract) {
				util.WriteLog("SLOC Analysis", "CLOC of File [%s]done!\n", file.Name)
			}

		}

		files := int32(len(language.Files))
		if len(language.Files) <= 0 {
			continue
		}

		total.Total += files
		total.Blanks += language.Blanks
		total.Comments += language.Comments
		total.Code += language.Code
	}

	return &Result{
		Total:         total,
		Files:         clocFiles,
		Domains:       domainTotals,
		MaxPathLength: maxPathLen,
		UnknownExts:   unknowns,
	}, nil
}
