/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import "time"
import "csa-app/util"

type ExcludePattern struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	RuleID    uint      `sql:"type:bigint REFERENCES rules(id) ON DELETE CASCADE" json:"-"  yaml:"-"`
	Rule      Rule      `gorm:"foreignkey:RuleID" json:"-"  yaml:"-"`
	Value     string    `gorm:"type:text"`
}

func (e *ExcludePattern) GetEscapedPattern() string {
	return util.EscapeSpecials(e.Value)
}
