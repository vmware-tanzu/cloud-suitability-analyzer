/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model_test

import (
	"testing"

	"csa-app/model"
)

func TestPatternValueRequired(t *testing.T) {

	r := getValidRulePatternNoValue()

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Pattern value should be required but is not!")
	}
}

func TestPatternRegexValid(t *testing.T) {

	r := getValidRulePatternInvalidRegex()

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Pattern regex should be invalid and is not!")
	}
}

func getValidRulePatternNoValue() model.Rule {
	r := model.Rule{}
	r.Name = "Test Rule"
	r.FileType = "*"
	r.RuleType = "standard"
	r.Target = model.FILE_TARGET
	r.Type = model.REGEX_MATCH_TYPE
	r.DefaultPattern = "Bubba G[ui]mp %s"
	r.AddPattern(model.Pattern{Effort: 10})

	return r
}

func getValidRulePatternInvalidRegex() model.Rule {
	r := model.Rule{}
	r.Name = "Test Rule"
	r.FileType = "*"
	r.Target = model.FILE_TARGET
	r.Type = model.REGEX_MATCH_TYPE
	r.DefaultPattern = "Bubba G[ui]mp %s"
	r.AddPattern(model.Pattern{Value: "test", Pattern: "asd%s[]fasdf"})

	return r
}
