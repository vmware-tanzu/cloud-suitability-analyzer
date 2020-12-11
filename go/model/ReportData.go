/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import "time"

type ReportData struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	RunID     uint
	ReportID  int
	Data1     string `gorm:"type:text;column:data_1"`
	Data2     string `gorm:"type:text;column:data_2"`
	Data3     string `gorm:"type:text;column:data_3"`
	Data4     string `gorm:"type:text;column:data_4"`
	Data5     string `gorm:"type:text;column:data_5"`
	Data6     string `gorm:"type:text;column:data_6"`
	Data7     string `gorm:"type:text;column:data_7"`
	Data8     string `gorm:"type:text;column:data_8"`
	Data9     string `gorm:"type:text;column:data_9"`
	Data10    string `gorm:"type:text;column:data_10"`
}
