/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package natural

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/fatih/camelcase"
	"csa-app/csa"
	"csa-app/db"
	"csa-app/model"
	"csa-app/util"
)

type NaturalLangService struct {
	reSplitLine       *regexp.Regexp
	removeAllNonAlpha *regexp.Regexp
	runRepository     db.RunRepository
}

func NewNaturalLanguageService(runRepository db.RunRepository) *NaturalLangService {
	reSplitLine, _ := regexp.Compile(`[ .(<>=:]`)
	removeAllNonAlpha, _ := regexp.Compile("[^a-zA-Z]")
	return &NaturalLangService{
		reSplitLine:       reSplitLine,
		removeAllNonAlpha: removeAllNonAlpha,
		runRepository:     runRepository,
	}
}

type LangTarget struct {
	run     *model.Run
	file    util.FileInfo
	english *hashset.Set
	java    *hashset.Set
	results []string
}

func (naturalLangService *NaturalLangService) LangProcess(run *model.Run) {

	naturalLangService.runRepository.StartRun(run)

	run.StartActivity("gathering")
	fileUtil := util.NewFileUtil()
	path, _ := csa.CheckForAndPrepareArchiveTarget(run, run.Target, fileUtil)
	files, _ := csa.GetFileListExcludingDirs(path, fileUtil)
	var filesCnt = len(files)
	run.Files = filesCnt

	if filesCnt < 0 {
		fmt.Printf("No files found for processing at path[%s]!", path)
		naturalLangService.runRepository.StopRun(run)
		return
	}

	english := GetEnglishDictionary()
	java := GetJavaDictionary()
	run.StopActivity("gathering", "Githering Files...done!", true)
	if(!*util.Xtract) {
		fmt.Printf("Found [%d] Files to Analyze\n", filesCnt)
	}

	run.StartActivity("processing")

	waitGroup := sync.WaitGroup{}

	for i := 0; i < filesCnt; i++ {
		var file = files[i]

		waitGroup.Add(1)
		go func(idx int) {
			util.WriteLog("Processing", "processing...   Filename: %s\n", file.FQN)
			results := naturalLangService.java2eng(file, &english, &java)
			newFileName := strings.Replace(file.Name, file.Ext, "", -1)
			util.SliceToFile(results, getOutputDir()+file.TargetPath, newFileName+"-eng", "txt")
			run.Findings += len(results)
			util.WriteLog("Processing", "Filename: %s\n done!", file.FQN)
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()

	run.StopActivity("processing", "Processing...done!", true)

	naturalLangService.runRepository.StopRun(run)
}

//--- given path to java file, produce a derivative string slice that
//    contains a natural language equivalent derived from comments and
//    java methods/variable names that have been decamelized.

func (naturalLangService *NaturalLangService) java2eng(file util.FileInfo, english *hashset.Set, java *hashset.Set) []string {

	var englishTokens []string

	reImportStatement, _ := regexp.Compile("^import ")

	//--- open file and scan
	inJavaFile, _ := os.Open(file.FQN)
	javaFileScanner := bufio.NewScanner(inJavaFile)
	lineNum := 0
	for javaFileScanner.Scan() {
		lineNum++
		curLine := fmt.Sprintf("%s", javaFileScanner.Text())
		curLine = naturalLangService.removeAllNonAlpha.ReplaceAllString(curLine, " ")

		//--- ignore import statements
		if reImportStatement.FindString(curLine) != "" {
			continue
		}

		if curLine != "" {

			nlOutputLine := ""

			//--- split code line on spaces, dots, or parens
			words := naturalLangService.reSplitLine.Split(curLine, -1)

			//--- process each word
			for _, word := range words {
				if len(word) > 3 {
					//--- split out camel case fragments
					tokens := camelcase.Split(word)
					if tokens != nil {
						for _, token := range tokens {
							token = strings.ToLower(token)
							if len(token) > 3 {
								if english.Contains(token) && !java.Contains(token) {
									nlOutputLine += token + " "
								}
							}
						}
					}
				}
			}

			if len(nlOutputLine) > 0 {
				//nlOutputLine += "\n"
				englishTokens = append(englishTokens, nlOutputLine)
			}
		}
	}
	return englishTokens
}

func getOutputDir() string {

	dir := *util.OutputDir

	if !strings.HasSuffix(dir, util.PathSeparator) {
		dir = dir + util.PathSeparator
	}

	return dir + "natural" + util.PathSeparator

}
