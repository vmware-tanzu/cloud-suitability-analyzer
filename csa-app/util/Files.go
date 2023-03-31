/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	//"github.com/mholt/archiver/v3"
	"github.com/mitchellh/go-homedir"
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

func (fu *FileUtil) CloneFileUtil() *FileUtil {

	util := &FileUtil{}
	util.Langs = NewDefinedLanguages()

	util.IncludedFilesRegEx = fu.IncludedFilesRegEx
	util.ExcludedFilesRegEx = fu.ExcludedFilesRegEx
	util.ExcludedDirsRegExp = fu.ExcludedDirsRegExp
	util.DecompilingExcludedDirsRegExp = fu.DecompilingExcludedDirsRegExp
	util.archiveRegex = fu.archiveRegex

	return util
}

func (fu *FileUtil) DirIsExcluded(dir string) bool {

	if fu.useDecompileRegex {

		if fu.DecompilingExcludedDirsRegExp.MatchString(dir) {
			if *Verbose {
				WriteLog("Gathering", "Directory [%s] Matched Excluded Dirs [%s]...Excluding from analysis\n", dir, *ExcludedDirsRegEx)
			}
			return true
		}
		return false
	}

	if fu.ExcludedDirsRegExp.MatchString(dir) {
		if *Verbose {
			WriteLog("Gathering", "Directory [%s] Matched Excluded Dirs [%s]...Excluding from analysis\n", dir, *ExcludedDirsRegEx)
		}
		return true
	}

	return false
}

func (fu *FileUtil) FileShouldBeProcessed(name string) bool {
	//Only look for included files!
	if fu.IncludedFilesRegEx != nil {
		return fu.IncludedFilesRegEx.MatchString(name)
	} else {
		return !fu.ExcludedFilesRegEx.MatchString(name)
	}
}

func (fu *FileUtil) IsDecompilableArchive(target string) (isArchive bool) {
	return fu.archiveRegex.MatchString(target)
}

func (fu *FileUtil) CheckForArchive(path string) (finalTargetPath string, alias string, decompiled bool) {

	if fu.IsDecompilableArchive(path) {

		WriteLog("Decompiling", "...\n")

		//Get DecompilePath
		decompilePath := *TmpDirPath + "/decompile"
		if *DecompileDir != "" {
			if !Exists(*DecompileDir) {
				CreateDirIfNotExist(*DecompileDir)
				decompilePath = *DecompileDir
			}
		}

		//Target all files of type extension
		if strings.Contains(path, "*.") {

			files := fu.GetFilesWithExtension(filepath.Dir(path), filepath.Ext(path))
			var alias string

			for _, file := range files {
				if len(alias) > 0 {
					alias += "|" + file.Name
				} else {
					alias = file.Name
				}

				WriteLog("Decompiling", "...   Filename: %s\n", file.Name)
				fu.Decompile(file, decompilePath)
			}

		} else {
			file := GetFile(path)
			WriteLog("Decompiling", "...   Filename: %s\n", file.Name)
			fu.Decompile(file, decompilePath)
		}

		return decompilePath, alias, true
	}

	return path, alias, false
}

func (fu *FileUtil) GetFileList(searchDir string, expression string) (files []FileInfo) {

	var parentDir = strings.Count(searchDir, PathSeparator)
	var domain = "unknown"
	var domainDir = parentDir + 1

	_ = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if err == nil {
			if f.IsDir() {
				var dirCnt = strings.Count(path, PathSeparator)

				if *Verbose {
					fmt.Printf("Walking File List for Path [%s] PathSep: %s ParentDir count: %d Dir: %d ThisDir: %d\n", path, PathSeparator, parentDir, domainDir, dirCnt)
				}
				if dirCnt == domainDir {
					if *Verbose {
						fmt.Printf("******************  Found a domain named [%s]!\n", f.Name())
					}
					domain = f.Name()
				}
			} else {
				r, err := regexp.MatchString(expression, filepath.Ext(f.Name()))
				if err == nil && r {
					if *Verbose {
						fmt.Printf("Found File [%s] in Application [%s]\n", path, domain)
					}
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

func (fu *FileUtil) GetFilesWithExtension(searchDir string, extension string) (matchingFiles []FileInfo) {

	files, err := ioutil.ReadDir(searchDir)

	if err != nil {
		App.Errorf("%s\n", fmt.Sprintf("Error getting files in directory [%s]! Details: %s", searchDir, err))
		App.Usage(os.Args)
		os.Exit(1)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), extension) {
			domain := "TBD"
			matchingFiles = append(matchingFiles, *NewFileInfo(
				domain,
				searchDir+PathSeparator+file.Name(),
				file.Name(),
				extension,
				searchDir,
				"",
				true))
		}
	}

	return
}

func (fu *FileUtil) Decompile(target FileInfo, basePath string) {

	fernFlowerPath := getDecompilerPath()

	filePath := strings.Replace(filepath.Base(target.Name), filepath.Ext(target.Name), "", -1)

	Lock.Lock()
	defer Lock.Unlock()

	decompileDir := basePath + "/" + filePath
	exists, err := CreateDirIfNotExist(decompileDir)
	if exists {
		//Already or in process of decompiling this jar...
		return
	}

	fmt.Printf("Decompiling [%s]...\n", target.Name)

	cmd := exec.Command("java", "-jar", fernFlowerPath, target.FQN, decompileDir)
	output, err := cmd.CombinedOutput()

	if err != nil {
		App.Fatalf("Decompile of %s failed!\nError-Details: %s\n%s\n", target.Name, err, output)
	}

	if *Verbose {
		fmt.Printf("Decompile Cmd Output => %s\n", output)
	}

	//decompiledJarPath := fmt.Sprintf("%s/%s", decompileDir, filepath.Base(target.Name))
	//err = archiver.Zip.Open(decompiledJarPath, decompileDir)
	//if err != nil {
	//	App.Fatalf("Decompress of decompiled archive [%s|%s] failed!\nError-Details: %v\n", target.Name, decompiledJarPath, err)
	//} else {
	//	DeleteFileOrDir(decompiledJarPath)
	//}
	fu.useDecompileRegex = true
}

func (fu *FileUtil) GetLangForFileExt(fileExt string) (language *Language, ok bool) {

	ext, ok := Exts[strings.ToLower(fileExt)]

	if !ok {
		return nil, false
	}

	lang, ok := fu.Langs.Langs[ext]
	return lang, ok
}

func (fu *FileUtil) EstablishTempDirectory(uniqueId string) (path string, err error) {

	//Will default to the OS specific default temp path if one is not specified via cmd line switch
	path, err = ioutil.TempDir(*TmpDirPath, uniqueId)

	if err == nil {
		_, err = CreateDirIfNotExist(path)
	}

	if err != nil {
		exePath, err := os.Executable()
		if err == nil {
			path, err = ioutil.TempDir(filepath.Dir(exePath), uniqueId)
			_, err = CreateDirIfNotExist(path)
		} else {
			homepath, err := homedir.Dir()
			if err == nil {
				path, err = ioutil.TempDir(homepath, uniqueId)
				_, err = CreateDirIfNotExist(path)
			}
		}
	}

	//Update to the final tmp path!
	*TmpDirPath = path

	return path, err

}

func (fu *FileUtil) RemoveDir(path string) {
	os.RemoveAll(path)
}

/***********************************************************************************************************************
														FILEINFO API
***********************************************************************************************************************/

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

/***********************************************************************************************************************
														PUBLIC API
***********************************************************************************************************************/

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

func CreateDirIfNotExist(dir string) (existed bool, err error) {

	_, err = os.Stat(dir)

	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, os.MkdirAll(dir, 0755)
	}
	return false, err
}

/***********************************************************************************************************************
														PRIVATE API
***********************************************************************************************************************/

func getDecompilerPath() string {

	fernPath := *FernLocation

	if fernPath == "" {
		fernPath = "."
	}

	//Determine Decompiler availability
	fernFlowerPath := fmt.Sprintf("%s/fernflower.jar", fernPath)

	if !Exists(fernFlowerPath) {
		App.Errorf("%s\n", fmt.Sprintf("Fernflower not installed or found at [%s]", fernFlowerPath))
		App.Usage(os.Args)
		os.Exit(1)
	}

	return fernFlowerPath

}
