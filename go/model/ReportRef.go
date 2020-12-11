/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"sort"
	"time"
)

type ReportRef struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	ReportNum int
	Type      string `gorm:"type:text"`
	Title     string `gorm:"type:text"`
	Summary   string `gorm:"type:text;column:summary"`
	Target    string `gorm:"type:text"`
	Extension string `gorm:"type:text"`
}

func ByReportNum(p1, p2 *ReportRef) bool {
	return p1.ReportNum < p2.ReportNum
}

type SortReports func(p1, p2 *ReportRef) bool

func (BY SortReports) Sort(reports []ReportRef) {
	ps := &reportSorter{
		reports: reports,
		by:      BY, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type reportSorter struct {
	reports []ReportRef
	by      func(p1, p2 *ReportRef) bool // Closure used in the Less method.
}

func (s *reportSorter) Len() int {
	return len(s.reports)
}

func (s *reportSorter) Swap(i, j int) {
	s.reports[i], s.reports[j] = s.reports[j], s.reports[i]
}

func (s *reportSorter) Less(i, j int) bool {
	return s.by(&s.reports[i], &s.reports[j])
}
