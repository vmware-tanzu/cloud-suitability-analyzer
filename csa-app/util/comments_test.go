/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util_test

import (
	"testing"

	"csa-app/util"
	"github.com/stretchr/testify/assert"
)

func TestNoComment(t *testing.T) {

	data := "Bubba Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, data, after)
	assert.True(t, process)
	assert.False(t, midComment)
}

func TestCommentBegining(t *testing.T) {

	data := "//Bubba Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "", after)
	assert.False(t, process)
	assert.False(t, midComment)
}

func TestCommentEnd(t *testing.T) {

	data := "Bubba Gump // asdfasdfasdf"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "Bubba Gump ", after)
	assert.True(t, process)
	assert.False(t, midComment)
}

func TestCommentMiddle(t *testing.T) {

	data := "Bubba//Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "Bubba", after)
	assert.True(t, process)
	assert.False(t, midComment)
}

func TestMultiStartBeginning(t *testing.T) {

	data := "/*Bubba Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "", after)
	assert.False(t, process)
	assert.True(t, midComment)
}

func TestMultiStartEnd(t *testing.T) {

	data := "Bubba Gump/*"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "Bubba Gump", after)
	assert.True(t, process)
	assert.True(t, midComment)
}

func TestMultiLineMidComment(t *testing.T) {

	data := "Bubba Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, true, lang)

	assert.Equal(t, "", after)
	assert.False(t, process)
	assert.True(t, midComment)
}

func TestMultiStartMiddle(t *testing.T) {

	data := "Bubba/* Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "Bubba", after)
	assert.True(t, process)
	assert.True(t, midComment)
}

func TestMultiStartMiddleClosed(t *testing.T) {

	data := "Bubba/* asdfasdfasdf */ Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, false, lang)

	assert.Equal(t, "Bubba Gump", after)
	assert.True(t, process)
	assert.False(t, midComment)
}

func TestMultiEndBeginning(t *testing.T) {

	data := "*/Bubba Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, true, lang)

	assert.Equal(t, "Bubba Gump", after)
	assert.True(t, process)
	assert.False(t, midComment)
}

func TestMultiEndMiddle(t *testing.T) {

	data := "Bubba */Gump"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, true, lang)

	assert.Equal(t, "Gump", after)
	assert.True(t, process)
	assert.False(t, midComment)
}

func TestMultiEndEnd(t *testing.T) {

	data := "Bubba Gump*/"
	ext := "java"

	fu := util.NewFileUtil()

	lang, _ := fu.GetLangForFileExt(ext)

	after, process, midComment := util.HandleComments(data, true, lang)

	assert.Equal(t, "", after)
	assert.False(t, process)
	assert.False(t, midComment)
}
