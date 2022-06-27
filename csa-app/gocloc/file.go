/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package gocloc

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"csa-app/util"
)

type ClocFile struct {
	Code     int32  `xml:"code,attr"`
	Comments int32  `xml:"comment,attr"`
	Blanks   int32  `xml:"blank,attr"`
	Name     string `xml:"name,attr"`
	Lang     string `xml:"language,attr"`
	Dir      string `xml:"dir,attr"`
}

type ClocFiles []ClocFile

func (cf ClocFiles) Len() int {
	return len(cf)
}
func (cf ClocFiles) Swap(i, j int) {
	cf[i], cf[j] = cf[j], cf[i]
}
func (cf ClocFiles) Less(i, j int) bool {
	if cf[i].Code == cf[j].Code {
		return cf[i].Name < cf[j].Name
	}
	return cf[i].Code > cf[j].Code
}

func AnalyzeFile(file util.FileInfo, language *util.Language, opts *util.ClocOptions) *ClocFile {
	if opts.Debug {
		fmt.Printf("filename=%v\n", file.Name)
	}

	clocFile := &ClocFile{
		Name: file.Name,
		Dir:  file.Dir,
		Lang: language.Name,
	}

	fp, err := os.Open(file.FQN)
	if err != nil {
		return clocFile // ignore error
	}
	defer fp.Close()

	isFirstLine := true
	isInComments := false
	isInCommentsSame := false
	buf := util.GetByteSlice()
	defer util.PutByteSlice(buf)
	scanner := bufio.NewScanner(fp)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		lineOrg := scanner.Text()
		line := strings.TrimSpace(lineOrg)

		if len(strings.TrimSpace(line)) == 0 {
			clocFile.Blanks++
			if opts.Debug {
				fmt.Printf("[BLNK,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
					clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
			}
			continue
		}

		// shebang line is 'code'
		if isFirstLine && strings.HasPrefix(line, "#!") {
			clocFile.Code++
			isFirstLine = false
			if opts.Debug {
				fmt.Printf("[CODE,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
					clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
			}
			continue
		}

		if len(language.LineComments) > 0 {
			isSingleComment := false
			if isFirstLine {
				line = trimBOM(line)
			}
			for _, singleComment := range language.LineComments {
				if strings.HasPrefix(line, singleComment) {
					clocFile.Comments++
					isSingleComment = true
					break
				}
			}
			if isSingleComment {
				if opts.Debug {
					fmt.Printf("[COMM,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
						clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
				}
				continue
			}
		}

		if language.MultiLine != "" {
			if strings.HasPrefix(line, language.MultiLine) {
				isInComments = true
			} else if strings.HasSuffix(line, language.MultiLineEnd) {
				isInComments = true
			} else if containComments(line, language.MultiLine, language.MultiLineEnd) {
				isInComments = true
				if (language.MultiLine != language.MultiLineEnd) &&
					(strings.HasSuffix(line, language.MultiLine) || strings.HasPrefix(line, language.MultiLineEnd)) {
					clocFile.Code++
					if opts.Debug {
						fmt.Printf("[CODE,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
							clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
					}
					continue
				}
			}
		}

		if isInComments {
			if language.MultiLine == language.MultiLineEnd {
				if strings.Count(line, language.MultiLineEnd) == 2 {
					isInComments = false
					isInCommentsSame = false
				} else if strings.HasPrefix(line, language.MultiLineEnd) ||
					strings.HasSuffix(line, language.MultiLineEnd) {
					if isInCommentsSame {
						isInComments = false
					}
					isInCommentsSame = !isInCommentsSame
				}
			} else {
				if strings.Contains(line, language.MultiLineEnd) {
					isInComments = false
				}
			}
			clocFile.Comments++
			if opts.Debug {
				fmt.Printf("[COMM,cd:%d,cm:%d,bk:%d,iscm:%v,iscms:%v] %s\n",
					clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, isInCommentsSame, lineOrg)
			}
			continue
		}

		clocFile.Code++
		if opts.Debug {
			fmt.Printf("[CODE,cd:%d,cm:%d,bk:%d,iscm:%v] %s\n",
				clocFile.Code, clocFile.Comments, clocFile.Blanks, isInComments, lineOrg)
		}
	}

	return clocFile
}
