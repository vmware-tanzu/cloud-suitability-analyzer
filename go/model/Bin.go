/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"strings"
	"time"
)

const (
	AND = iota
	OR
	EXCLUDE
)

type Bin struct {
	ID          uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt   time.Time `json:"-" yaml:"-"`
	UpdatedAt   time.Time `json:"-" yaml:"-"`
	Name        string    `gorm:"type:text;unique_index;not null" json:"name" yaml:"name"`
	Tags        []*BinTag `json:"tags" yaml:"tags"`
	MatchedTags []string  `gorm:"-" json:"matched" yaml:"matched"`
}

type BinTag struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	BinId     uint      `sql:"type:bigint REFERENCES bins(id) ON DELETE CASCADE" json:"-"  yaml:"-"`
	Bin       Bin       `gorm:"foreignkey:BinId" json:"-"  yaml:"-"`
	Name      string    `gorm:"type:text;index;not null" json:"name" yaml:"name"`
	Type      int       `json:"-" yaml:"-"`
	Action    string    `gorm:"-" json:"type,omitempty" yaml:"type,omitempty"`
}

func NewBinTag(tagName string, tagType string) *BinTag {
	t := &BinTag{Name: tagName}
	t.Action = tagType
	t.SetTagType(tagType)
	return t
}

func (b *Bin) Matches(tags []*ApplicationTag) bool {

	excluded := false
	included := false
	andMatches := true
	for _, tag := range tags {

		bt := b.GetTag(tag.Value)
		if bt != nil {

			if bt.Name == tag.Value {
				switch bt.Type {
				case EXCLUDE:
					excluded = true
					break
				case OR:
					included = true
					b.MatchedTags = append(b.MatchedTags, bt.Name)
				default:
					//Store the tag as matching. Even though the bin may not match...
					b.MatchedTags = append(b.MatchedTags, bt.Name)
				}
			} else if bt.Type == AND {
				andMatches = false
			}
		}
	}

	return !excluded && included && andMatches
}

func (b *Bin) GetTag(tagName string) *BinTag {
	for _, tag := range b.Tags {
		if tag.Name == tagName {
			return tag
		}
	}
	return nil
}

func (b *Bin) HasTag(tagName string) bool {
	if b.GetTag(tagName) != nil {
		return true
	}
	return false
}

func (b *Bin) AddTag(tagName string, tagType string) {
	t := NewBinTag(tagName, tagType)
	if !b.HasTag(tagName) {
		b.Tags = append(b.Tags, t)
	} else {
		oldTag := b.GetTag(tagName)
		oldTag.Type = DecodeTagType(tagType)
	}
}

func (b *Bin) PrepForMarshall() {
	for _, tag := range b.Tags {
		tag.PrepForMarshall()
	}
}

func (b *Bin) PrepForPersist() {
	for _, tag := range b.Tags {
		tag.PrepForPersist()
	}
}

func (t *BinTag) SetTagType(tagType string) {
	t.Action = tagType
	t.Type = DecodeTagType(tagType)
}

func (t *BinTag) PrepForMarshall() {
	switch t.Type {
	case AND:
		t.Action = "AND"
	case OR:
		t.Action = "OR"
	case EXCLUDE:
		t.Action = "EXCLUDE"
	default:
		t.Action = "AND"
	}
}

func (t *BinTag) PrepForPersist() {
	if t.Action != "" {
		switch t.Action {
		case "AND":
			t.Type = AND
		case "OR":
			t.Type = OR
		case "EXCLUDE":
			t.Type = EXCLUDE
		default:
			t.Type = AND
		}
	}
}

func DecodeTagType(tagType string) int {
	switch strings.ToLower(tagType) {
	case "or":
		return 1
	case "exclude":
		return 2
	default:
		return 0
	}
}
