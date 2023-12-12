/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"csa-app/util"
	"strings"
	"time"
)

const FILE_ANALYZED_CATEGORY = "File Finding"
const SLOC_CATEGORY = "SLOC"
const INFO_FINDING = "info"
const FILE_FINDING = "ff"
const FINDING_VAL_LEN = 2000

type Finding struct {
	ID                uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt         time.Time `json:"-" yaml:"-"`
	UpdatedAt         time.Time `json:"-" yaml:"-"`
	RunID             uint      `gorm:"index;not null"`
	Filename          string    `gorm:"type:text;"`
	Fqn               string    `gorm:"type:text;`
	Ext               string    `gorm:"type:text;"`
	Line              int
	Rule              string          `gorm:"type:text;"`
	Pattern           string          `gorm:"type:text;"`
	Value             string          `gorm:"type:text;"`
	Note              string          `gorm:"type:text;" json:",omitempty" yaml:",omitempty"`
	Advice            string          `gorm:"type:text" json:",omitempty" yaml:",omitempty"`
	Effort            int             `gorm:"type:bigint" json:"effort" yaml:"effort"`
	Readiness         int             `gorm:"type:bigint" json:"readiness" yaml:"readiness,omitempty"`
	Category          string          `gorm:"index;not null" json:",omitempty" yaml:",omitempty"`
	Criticality       string          `gorm:"index;not null" json:",omitempty" yaml:",omitempty"`
	CloudNativeEffort int             `gorm:"type:bigint" json:"cloudNativeEffort" yaml:"cloudNativeEffort,omitempty"`
	ContainerEffort   int             `gorm:"type:bigint" json:"containerEffort" yaml:"containerEffort,omitempty"`
	Application       string          `gorm:"index;not null" json:",omitempty" yaml:",omitempty"`
	Tags              []FindingTag    `gorm:"foreignkey:FindingID" json:",omitempty" yaml:",omitempty"`
	Recipes           []FindingRecipe `gorm:"foreignkey:FindingID" json:",omitempty" yaml:",omitempty"`
	Result            string          `gorm:"type:text;"`
}

type FindingDTO struct {
	ID                uint     `json:"id" yaml:"id"`
	RunID             uint     `json:"run" yaml:"run"`
	Filename          string   `json:"filename" yaml:"filename"`
	Fqn               string   `json:"fqn" yaml:"fqn"`
	Ext               string   `json:"ext" yaml:"ext"`
	Line              int      `json:"line" yaml:"line"`
	Rule              string   `json:"rule" yaml:"rule"`
	Pattern           string   `json:"pattern" yaml:"pattern"`
	Value             string   `json:"value" yaml:"value"`
	Advice            string   `json:"advice" yaml:"advice"`
	Note              string   `json:"note,omitempty" yaml:"note,omitempty"`
	Level             string   `json:"level" yaml:"level"`
	Effort            int      `json:"effort" yaml:"effort"`
	Readiness         int      `json:"readiness" yaml:"readiness,omitempty"`
	Category          string   `json:"category" yaml:"category,omitempty"`
	Criticality       string   `json:"criticality" yaml:"criticality,omitempty"`
	CloudNativeEffort int      `json:"cloudNativeEffort" yaml:"cloudNativeEffort,omitempty"`
	ContainerEffort   int      `json:"containerEffort" yaml:"containerEffort,omitempty"`
	Application       string   `json:"application" yaml:"domain,omitempty"`
	Tags              []string `json:"tags" yaml:"tags,omitempty"`
	Recipes           []string `json:"recipes" yaml:"recipes,omitempty"`
}

func (f *Finding) SetValue(value string) {

	if len(value) > (FINDING_VAL_LEN) {
		f.Value = value[0:(FINDING_VAL_LEN-5)] + "..."
	} else {
		if *util.Efd {
			f.Value = "---"
			f.Result = "---"
		} else {
			f.Value = value
		}
	}
}

func (f *Finding) AddRecipe(recipe string) {
	f.Recipes = append(f.Recipes, FindingRecipe{URI: recipe})
}

func (f *Finding) AddTag(tag string) {
	f.Tags = append(f.Tags, FindingTag{Value: strings.ToLower(tag)})
}

func (d *FindingDTO) AddRecipe(recipe string) {
	d.Recipes = append(d.Recipes, recipe)
}

func (d *FindingDTO) AddTag(tag string) {
	d.Tags = append(d.Tags, strings.ToLower(tag))
}

func (f *Finding) CreateDTO() *FindingDTO {

	dto := &FindingDTO{}

	dto.ID = f.ID
	dto.Ext = f.Ext
	dto.Fqn = f.Fqn
	dto.Filename = f.Filename
	dto.Rule = f.Rule
	dto.RunID = f.RunID
	dto.Line = f.Line
	dto.Application = f.Application
	dto.Category = f.Category
	dto.Note = f.Note
	dto.Advice = f.Advice
	dto.Effort = f.Effort
	dto.Readiness = f.Readiness
	dto.Value = "109"
	dto.Pattern = f.Pattern
	dto.Criticality = f.Criticality

	for _, tag := range f.Tags {
		dto.AddTag(tag.Value)
	}
	for _, recipe := range f.Recipes {
		dto.Recipes = append(dto.Recipes, recipe.URI)
	}

	return dto
}

type SortFindingByGroupNameLine []Finding

func (findings SortFindingByGroupNameLine) Len() int { return len(findings) }
func (findings SortFindingByGroupNameLine) Swap(i, j int) {
	findings[i], findings[j] = findings[j], findings[i]
}
func (findings SortFindingByGroupNameLine) Less(i, j int) bool {
	f1 := findings[i]
	f2 := findings[j]
	if f1.Category == f2.Category {
		if f1.Filename == f2.Filename {
			return f1.Line < f2.Line
		}
		return f1.Filename < f2.Filename
	}
	return f1.Category < f2.Category
}
