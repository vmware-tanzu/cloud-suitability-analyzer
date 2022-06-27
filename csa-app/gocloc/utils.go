/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package gocloc

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"csa-app/model"
	"csa-app/util"
)

func trimBOM(line string) string {
	l := len(line)
	if l >= 3 {
		if line[0] == 0xef && line[1] == 0xbb && line[2] == 0xbf {
			trimLine := line[3:]
			return trimLine
		}
	}
	return line
}

func containComments(line, commentStart, commentEnd string) bool {
	inComments := 0
	for i := 0; i < len(line)-(len(commentStart)-1); i++ {
		section := line[i : i+len(commentStart)]

		if section == commentStart {
			inComments++
		} else if section == commentEnd {
			if inComments != 0 {
				inComments--
			}
		}
	}
	return inComments != 0
}

func checkMD5Sum(path string, fileCache map[string]struct{}) (ignore bool) {

	//Comparing content does not make sense. If the files are on different paths then they are not the same even if there contents are!
	//content, err := ioutil.ReadFile(path)
	//if err != nil {
	//	return true
	//}

	// calc md5sum
	hash := md5.Sum([]byte(path))
	c := fmt.Sprintf("%x", hash)
	if _, ok := fileCache[c]; ok {
		return true
	}

	fileCache[c] = struct{}{}
	return false
}

func isVCSDir(path string) bool {
	if len(path) > 1 && path[0] == os.PathSeparator {
		path = path[1:]
	}
	vcsDirs := []string{".bzr", ".cvs", ".hg", ".git", ".svn"}
	for _, dir := range vcsDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}
	return false
}

// getAllFiles return all of the files to be analyzed in paths.
func getAllFiles(root string, languages *util.DefinedLanguages, opts *util.ClocOptions) (result map[string]*util.Language, err error) {
	result = make(map[string]*util.Language, 0)
	fileCache := make(map[string]struct{})

	var parentDir = strings.Count(root, util.PathSeparator)
	var domain = ""
	var domainDir = parentDir + 1
	var fileUtil = util.NewFileUtil()

	vcsInRoot := isVCSDir(root)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if info.IsDir() {
			if fileUtil.DirIsExcluded(info.Name()) {
				if *util.Verbose {
					util.WriteLog("SLOC Analysis", "CLOC Found excluded directory [%s]\n", info.Name())
				}
				return filepath.SkipDir
			}
			var dirCnt = strings.Count(path, util.PathSeparator)
			if *util.Verbose {
				util.WriteLog("SLOC Analysis", "CLOC Walking File List for Path [%s] PathSep: %s ParentDir count: %d Dir: %d ThisDir: %d\n", path, util.PathSeparator, parentDir, domainDir, dirCnt)
			}

			if dirCnt == domainDir {
				if *util.Verbose {
					util.WriteLog("SLOC Analysis", "******************  Found a domain named [%s]!\n", info.Name())
				}
				domain = info.Name()
			}
			return nil
		}

		if !fileUtil.FileShouldBeProcessed(info.Name()) {
			return nil
		}

		if !vcsInRoot && isVCSDir(path) {
			return nil
		}

		if filepath.Dir(path) == root {
			domain = ""
		}

		if ext, ok := util.GetFileType(path, opts); ok {
			if targetExt, ok := util.Exts[ext]; ok {
				// check exclude extension
				if _, ok := opts.ExcludeExts[targetExt]; ok {
					return nil
				}

				if len(opts.IncludeLangs) != 0 {
					if _, ok = opts.IncludeLangs[targetExt]; !ok {
						return nil
					}
				}

				if !opts.SkipDuplicated {
					ignore := checkMD5Sum(path, fileCache)
					if ignore {
						if *util.Verbose {
							fmt.Printf("[ignore=%v] find same md5\n", path)
						}
						return nil
					}
				}

				if _, ok := result[targetExt]; !ok {
					result[targetExt] = util.NewLanguage(
						languages.Langs[targetExt].Name,
						languages.Langs[targetExt].LineComments,
						languages.Langs[targetExt].MultiLine,
						languages.Langs[targetExt].MultiLineEnd)
				}

				domainPath := strings.Replace(path, root, "", -1)
				domainPath = strings.Replace(domainPath, info.Name(), "", -1)
				ext := filepath.Ext(info.Name())
				result[targetExt].Files = append(result[targetExt].Files, util.NewFileInfo(
					domain,
					path,
					info.Name(),
					ext,
					domainPath,
					"",
					true))
			}
		}
		return nil
	})

	return
}

// getAllFiles return all of the files to be analyzed in paths.
func getAllLangsForFiles(app *model.Application, languages *util.DefinedLanguages, opts *util.ClocOptions) (result map[string]*util.Language, unknownExts []string, err error) {
	result = make(map[string]*util.Language, 0)
	fileCache := make(map[string]struct{})
	unknownExtsMap := make(map[string]struct{})
	for _, file := range app.Files {
		targetExt := "???" //Unknown ext

		if ext, ok := util.GetFileType(file.FQN, opts); ok {
			if target, ok := util.Exts[ext]; ok {
				targetExt = target
				// check exclude extension
				if _, ok := opts.ExcludeExts[target]; ok {
					continue
				}

				if len(opts.IncludeLangs) != 0 {
					if _, ok = opts.IncludeLangs[target]; !ok {
						continue
					}
				}

				if !opts.SkipDuplicated {
					ignore := checkMD5Sum(file.FQN, fileCache)
					if ignore {
						if *util.Verbose {
							fmt.Printf("[ignore=%v] find same md5\n", file.FQN)
						}
						continue
					}
				}
			} else {
				if _, ok := unknownExtsMap[ext]; !ok {
					unknownExtsMap[ext] = struct{}{}
				}
			}
			if targetExt != "???" || (targetExt == "???" && *util.DisplayUnknownExts) {
				if _, ok := result[targetExt]; !ok {
					result[targetExt] = util.NewLanguage(
						languages.Langs[targetExt].Name,
						languages.Langs[targetExt].LineComments,
						languages.Langs[targetExt].MultiLine,
						languages.Langs[targetExt].MultiLineEnd)
				}
				result[targetExt].Files = append(result[targetExt].Files, file)
			}
		}
	}

	for k, _ := range unknownExtsMap {
		unknownExts = append(unknownExts, k)
	}

	return
}
