/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"sort"
	"strings"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func SortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

type ByApiFileLine []ApiInfo

func (apiData ByApiFileLine) Len() int      { return len(apiData) }
func (apiData ByApiFileLine) Swap(i, j int) { apiData[i], apiData[j] = apiData[j], apiData[i] }
func (apiData ByApiFileLine) Less(i, j int) bool {
	f1 := apiData[i]
	f2 := apiData[j]
	if f1.API == f2.API {
		if f1.File == f2.File {
			return f1.LineNo < f2.LineNo
		}
		return f1.File < f2.File
	}
	return f1.API < f2.API
}

type StringsCaseInsensitive []string

func (s StringsCaseInsensitive) Len() int      { return len(s) }
func (s StringsCaseInsensitive) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s StringsCaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(s[i]) < strings.ToLower(s[j])
}
