/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import "time"

type FindingTag struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	FindingID uint      `gorm:"index;not null" sql:"type:bigint REFERENCES findings(id) ON DELETE CASCADE" json:"-"  yaml:"-"`
	Value     string    `gorm:"type:text;"index;not null""`
}

type TagsRequest struct {
	Tags []TagState `json:"tags"`
}

type TagState struct {
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
}
