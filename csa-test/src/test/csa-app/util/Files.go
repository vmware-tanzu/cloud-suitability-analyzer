/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/mitchellh/go-homedir"
	//"github.com/mholt/archiver/v3"
)

type FileInfo struct {
	Dir          string
	FQN          string
	Name         string
	Ext          string
	TargetPath   string
	Comment      string
	Exists       bool
	MatchedRules map[string]int
	sync.Mutex
}

type ApiInfo struct {
	Dir     string
	API     string
	Pattern string
	File    string
	LineNo  string
	Source  string
	Score   int
	Advice  string
}

type LineInfo struct {
	Dir    string
	Line   string
	File   string
	LineNo string
}

const PathSeparator = string(os.PathSeparator)

type FileUtil struct {
	IncludedFilesRegEx            *regexp.Regexp
	ExcludedDirsRegExp            *regexp.Regexp
	ExcludedFilesRegEx            *regexp.Regexp
	DecompilingExcludedDirsRegExp *regexp.Regexp
	Langs                         *DefinedLanguages
	archiveRegex                  *regexp.Regexp
	useDecompileRegex             bool
}

/***********************************************************************************************************************
														FILEUTIL API
***********************************************************************************************************************/

func NewFileInfo(dir string, fqn string, name string, ext string, targetpath string, comments string, exists bool) *FileInfo {
	return &FileInfo{
		Dir:          dir,
		FQN:          fqn,
		Name:         name,
		Ext:          ext,
		TargetPath:   targetpath,
		Comment:      comments,
		Exists:       exists,
		MatchedRules: make(map[string]int),
	}
}

func NewFileUtil() *FileUtil {
	util := &FileUtil{}

	util.Langs = NewDefinedLanguages()

	var err error

	if *IncludedFilesRegEx != ".*" {
		util.IncludedFilesRegEx, err = regexp.Compile(*IncludedFilesRegEx)
		if err != nil {
			App.Fatalf("Invalid 'included-files' regex! Details: %v", err)
		}
	}

	util.ExcludedDirsRegExp, err = regexp.Compile(*ExcludedDirsRegEx)
	if err != nil {
		App.Fatalf("Invalid 'excluded-dirs' regex! Details: %v", err)
	}

	util.ExcludedFilesRegEx, err = regexp.Compile(*ExcludedFilesRegEx)
	if err != nil {
		App.Fatalf("Invalid 'excluded-files' regex! Details: %v", err)
	}

	util.DecompilingExcludedDirsRegExp, _ = regexp.Compile(strings.Replace(*ExcludedDirsRegEx, "classes|", "", -1))
	util.archiveRegex, _ = regexp.Compile(".*([.]ear|[.]war|[.]jar)$")

	return util
}

func (f *FileInfo) GetCleanedExt() string {
	if f.Ext == "" {
		//Try to get the extension from filename
		f.Ext = filepath.Ext(f.Name)
	}

	if f.Ext == "" {
		//We have a file without an extension
		return f.Name
	}

	if len(f.Ext) > 1 && strings.HasPrefix(f.Ext, ".") {
		return f.Ext[1:]
	}

	return f.Ext
}

func (fu *FileUtil) GetFileList(searchDir string, expression string) (files []FileInfo) {

	var parentDir = strings.Count(searchDir, PathSeparator)
	var domain = "unknown"
	var domainDir = parentDir + 1

	_ = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if err == nil {
			if f.IsDir() {
				var dirCnt = strings.Count(path, PathSeparator)

				//if *Verbose {
				//fmt.Printf("Walking File List for Path [%s] PathSep: %s ParentDir count: %d Dir: %d ThisDir: %d\n", path, PathSeparator, parentDir, domainDir, dirCnt)
				//}
				if dirCnt == domainDir {
					//if *Verbose {
					fmt.Printf("******************  Found a domain named [%s]!\n", f.Name())
					//}
					domain = f.Name()
				}
			} else {
				r, err := regexp.MatchString(expression, filepath.Ext(f.Name()))
				if err == nil && r {
					//if *Verbose {
					//fmt.Printf("Found File [%s] in Application [%s]\n", path, domain)
					//}
					ext := filepath.Ext(f.Name())
					domainPath := strings.Replace(path, searchDir, "", -1)
					domainPath = strings.Replace(domainPath, f.Name(), "", -1)
					files = append(files, *NewFileInfo(domain, path, f.Name(), ext, domainPath, "", true))
				}

			}
		}
		return nil
	})
	return files

}

func GetFile(name string) FileInfo {

	file, err := GetFileIfExists(name)

	if err != nil {
		App.Fatalf("Invalid File or Directory [%s]! Details: %s", name, err)
	}

	return file
}

func GetFileIfExists(name string) (FileInfo, error) {

	if strings.Contains(name, "~") {
		dir, _ := homedir.Dir()
		name = strings.Replace(name, "~", dir, 1)
	}

	file, err := os.Stat(name)
	if err != nil {
		return FileInfo{}, err
	}

	domain := "N/A"
	return *NewFileInfo(
		domain,
		name,
		file.Name(),
		filepath.Ext(file.Name()),
		filepath.Dir(name),
		"",
		true), nil
}
