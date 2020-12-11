/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

func Unique(lineInfos []LineInfo) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range lineInfos {
		if _, value := keys[entry.Line]; !value {
			keys[entry.Line] = true
			list = append(list, entry.Line)
		}
	}
	return list
}
