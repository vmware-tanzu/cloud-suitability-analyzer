/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

type SlocByApplication struct {
	Application  string `json:"application"`
	TotalFiles   int    `json:"totalFiles"`
	BlankLines   int    `json:"blankLines"`
	CommentLines int    `json:"commentLines"`
	CodeLines    int    `json:"codeLines"`
}

type SlocByLang struct {
	Language     string `json:"language"`
	TotalFiles   int    `json:"totalFiles"`
	BlankLines   int    `json:"blankLines"`
	CommentLines int    `json:"commentLines"`
	CodeLines    int    `json:"codeLines"`
}

type SlocByRun struct {
	RunId        uint `json:"runId"`
	Apps         int  `json:"applicationCount"`
	TotalFiles   int  `json:"totalFiles"`
	BlankLines   int  `json:"blankLines"`
	CommentLines int  `json:"commentLines"`
	CodeLines    int  `json:"codeLines"`
}

type LanguagesByCodeLines struct {
	Lang      string `json:"language"`
	CodeLines int    `json:"codeLines"`
}

type ApplicationsByLanguageForRun struct {
	Application string `json:"application"`
	CodeLines   int    `json:"codeLines"`
}
