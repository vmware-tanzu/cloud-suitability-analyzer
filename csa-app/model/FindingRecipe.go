/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import "time"

type FindingRecipe struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	FindingID uint      `gorm:"index;not null" sql:"type:bigint REFERENCES findings(id) ON DELETE CASCADE" json:"-"  yaml:"-"`
	URI       string    `gorm:"type:text"`
}
