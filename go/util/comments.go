/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"strings"
)

func HandleComments(line string, isInComment bool, lang *Language) (result string, processPatterns bool, midComment bool) {

	processPatterns = true
	midComment = false

	singleCommentChar := ""

	if isInComment {
		//Look for Comment End
		if strings.Contains(line, lang.MultiLineEnd) {
			line = line[strings.Index(line, lang.MultiLineEnd)+len(lang.MultiLineEnd):]
			if len(strings.TrimSpace(line)) == 0 {
				return "", false, midComment
			}
		} else {
			return "", false, true
		}
	}

	//Test for multi-line comments first...
	if lang.MultiLine != "" {

		beforeCmnTxt := ""
		endCmntTxt := ""
		beginIdx := strings.Index(line, lang.MultiLine)
		if beginIdx > -1 {
			midComment = true
			beforeCmnTxt = line[0:beginIdx]
			endCmntTxt = line[beginIdx+len(lang.MultiLine):]

			if len(endCmntTxt) > 0 {
				endIdx := strings.Index(endCmntTxt, lang.MultiLineEnd)
				if endIdx > -1 {
					endCmntTxt = endCmntTxt[endIdx+len(lang.MultiLineEnd):]
					midComment = false
				} else {
					endCmntTxt = ""
				}
			}

			line = beforeCmnTxt + endCmntTxt
			if len(strings.TrimSpace(line)) == 0 {
				return "", false, midComment
			}
		}
	}

	if len(lang.LineComments) > 0 {
		for _, singleComment := range lang.LineComments {
			if strings.Contains(line, singleComment) {
				singleCommentChar = singleComment
				break
			}
		}
		if singleCommentChar != "" {
			line = line[0:strings.Index(line, singleCommentChar)]
			if len(strings.TrimSpace(line)) == 0 {
				return "", false, false
			}
		}
	}

	return line, true, midComment
}
