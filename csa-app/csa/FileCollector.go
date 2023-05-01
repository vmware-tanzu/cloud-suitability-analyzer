/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"csa-app/model"
	"csa-app/util"
)

//Gathers Files for Input to CSA
var pathSep, _ = strconv.Unquote(strconv.QuoteRune(os.PathSeparator))

func GetFileListExcludingDirs(searchDir string, fileUtil *util.FileUtil) (files []util.FileInfo, err error) {

	var parentDir = strings.Count(searchDir, pathSep)
	var domain = "{root}"
	var domainDir = parentDir + 1

	err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if err == nil {
			if f.IsDir() {
				if fileUtil.DirIsExcluded(f.Name()) {
					return filepath.SkipDir
				}

				var dirCnt = strings.Count(path, pathSep)
				if dirCnt == domainDir {
					domain = f.Name()
				}
			} else {
				domainPath := strings.Replace(path, searchDir, "", -1)
				domainPath = strings.Replace(domainPath, f.Name(), "", -1)
				ext := filepath.Ext(f.Name())
				if fileUtil.FileShouldBeProcessed(ext) {
					files = append(files, *util.NewFileInfo(
						domain,
						path,
						f.Name(),
						ext,
						domainPath,
						"",
						true))
						if (!*util.Xtract) {	
							util.WriteLog("Gathering Files", "Found File [%s]\n", f.Name())
						}
				}
			}
		}
		return nil
	})
	return files, err

}

func CheckForAndPrepareArchiveTarget(run *model.Run, path string, fileUtil *util.FileUtil) (finalTargetPath string, decompiled bool) {

	if run != nil {
		run.StartActivity("decompile")
	}

	finalTargetPath, alias, decompiled := fileUtil.CheckForArchive(path)

	if decompiled && run != nil {
		run.SetAlias(alias)
	}

	if run != nil {
		run.StopActivity("decompile", "Decompiling...done!", true)
	}

	return
}
