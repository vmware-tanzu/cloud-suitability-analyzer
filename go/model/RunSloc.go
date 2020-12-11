/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import "time"

type RunSloc struct {
	ID           uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt    time.Time `json:"-" yaml:"-"`
	UpdatedAt    time.Time `json:"-" yaml:"-"`
	RunID        uint
	Application  string
	Lang         string
	TotalFiles   int
	BlankLines   int
	CommentLines int
	CodeLines    int
}
