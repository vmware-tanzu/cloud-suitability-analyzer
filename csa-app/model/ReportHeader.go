/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"sort"
	"time"
)

type ReportHeader struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	ReportID  int
	Name      string `gorm:"type:text"`
	Position  int
}

func Position(p1, p2 *ReportHeader) bool {
	return p1.Position < p2.Position
}

func ReportAndPosition(p1, p2 *ReportHeader) bool {
	if p1.ReportID == p2.ReportID {
		return p1.Position < p2.Position
	} else {
		return p1.ReportID < p2.ReportID
	}
}

type SortHeadersBy func(p1, p2 *ReportHeader) bool

func (by SortHeadersBy) Sort(headers []ReportHeader) {
	ps := &headerSorter{
		headers: headers,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type headerSorter struct {
	headers []ReportHeader
	by      func(p1, p2 *ReportHeader) bool // Closure used in the Less method.
}

func (s *headerSorter) Len() int {
	return len(s.headers)
}

func (s *headerSorter) Swap(i, j int) {
	s.headers[i], s.headers[j] = s.headers[j], s.headers[i]
}

func (s *headerSorter) Less(i, j int) bool {
	return s.by(&s.headers[i], &s.headers[j])
}
