/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"bytes"
	"regexp"
	"strings"
)

var specials = map[string]string{
	"\"": "\\\"",
	"'":  "''",
}

func LineWrap(s string, length int) string {

	re := regexp.MustCompile("\r?\n")
	s = re.ReplaceAllString(s, "")

	var buffer bytes.Buffer
	var n_1 = length - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%length == n_1 && i != l_1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}

func Padd(character string, cnt int) (result string) {

	for i := 0; i < cnt; i++ {
		result = result + character
	}

	return result
}

func EscapeSpecials(target string) string {

	if target != "" {

		//Replace any escapes first...
		target = strings.Replace(target, "\\", "\\\\", -1)

		for special, replacement := range specials {
			target = strings.Replace(target, special, replacement, -1)
		}

	}

	return target
}

func Escape(value string) string {
	return strings.Replace(value, "'", "''", -1)
}

func RemoveLineEndings(value string) string {
	value = strings.Replace(value, "\n", "", -1)
	value = strings.Replace(value, "\r", "", -1)
	return strings.TrimSpace(value)
}
