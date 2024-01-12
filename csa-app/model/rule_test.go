/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model_test

import (
	"fmt"
	"testing"

	"csa-app/model"
)

func TestRuleNameRequired(t *testing.T) {

	r := getValidRule()
	r.Name = ""

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Rule name should be required but is not!")
	}
}

func TestRuleFileType(t *testing.T) {

	r := getValidRule()
	r.FileType = ""
	r.RuleType = "standard"

	valid, _ := r.IsValid()

	if !valid {
		t.Errorf("FileType Can be Empty! Submitted: %s But failed validation!", r.FileType)
	}

	r.FileType = "*"
	valid, _ = r.IsValid()

	if !valid {
		t.Errorf("FileType Can be \"*\"! Submitted: %s But failed validation!", r.FileType)
	}

	r.FileType = "[]"
	valid, _ = r.IsValid()

	if valid {
		t.Errorf("Non Empty FileType must be valid reqex! Submitted: %s But passed validation!", r.FileType)
	}

}

func TestRuleTargetRequired(t *testing.T) {

	r := getValidRule()
	r.Target = ""

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Rule Target should be a required field but is not!")
	}

}

func TestRuleTargetValuesRestricted(t *testing.T) {

	r := getValidRule()
	r.Target = "axis"

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Rule Target must be either (line|file) but was: %s and still passed validation!", r.Target)
	}

	r.Target = model.FILE_TARGET

	valid, _ = r.IsValid()

	if !valid {
		t.Errorf("Rule Target of: %s should be valid but is not!", r.Target)
	}

	r.Target = model.LINE_TARGET

	valid, _ = r.IsValid()

	if !valid {
		t.Errorf("Rule Target of: %s should be valid but is not!", r.Target)
	}

}

func TestRuleTypeRequired(t *testing.T) {

	r := getValidRule()
	r.Type = ""

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Rule Type should be a required field but is not!")
	}
}

func TestRuleTypeValuesRestricted(t *testing.T) {

	r := getValidRule()
	r.Type = "Contains Text"

	valid, _ := r.IsValid()

	errmsg := fmt.Sprintf("Rule Type must be (%s|%s|%s|%s|%s|%s|%s|%s|%s)",
		model.SIMPLE_TEXT_MATCH_TYPE,
		model.SIMPLE_TEXT_CI_MATCH_TYPE,
		model.CONTAINS_MATCH_TYPE,
		model.CONTAINS_CI_MATCH_TYPE,
		model.REGEX_MATCH_TYPE,
		model.STARTS_WITH_MATCH_TYPE,
		model.STARTS_WITH_CI_MATCH_TYPE,
		model.ENDS_WITH_MATCH_TYPE,
		model.ENDS_WITH_CI_MATCH_TYPE)

	if valid {
		t.Errorf("%s but was: %s and still passed validation!", errmsg, r.Target)
	}

	r.Type = model.SIMPLE_TEXT_CI_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.SIMPLE_TEXT_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.STARTS_WITH_CI_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.STARTS_WITH_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.ENDS_WITH_CI_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.ENDS_WITH_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.CONTAINS_CI_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.CONTAINS_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

	r.Type = model.REGEX_MATCH_TYPE
	valid, _ = r.IsValid()
	if !valid {
		t.Errorf("Rule Type of: %s should be valid but is not!", r.Type)
	}

}

func TestDefaultPatternValidation(t *testing.T) {

	r := getValidRule()
	r.DefaultPattern = "asdfasdf"

	valid, _ := r.IsValid()

	if valid {
		t.Errorf("Pattern: %s must contain substitution marker but does not and still passed validation!", r.DefaultPattern)
	}

	r.DefaultPattern = "asd%s[]fasdf"
	valid, _ = r.IsValid()
	if valid {
		t.Errorf("Pattern: %s should fail validation with a regex compile error but passed!", r.DefaultPattern)
	}

	r.DefaultPattern = "asd%s[iou]fasdf"
	valid, err := r.IsValid()
	if !valid {
		t.Errorf("Pattern: %s should pass validation but failed with error: %v", r.DefaultPattern, err)
	}

}

func getValidRule() model.Rule {
	r := model.Rule{}
	r.Name = "Test Rule"
	r.FileType = "*"
	r.Target = model.FILE_TARGET
	r.Type = model.REGEX_MATCH_TYPE
	r.DefaultPattern = "Bubba G[ui]mp %s"
	r.AddPattern(model.Pattern{Value: "test"})

	return r
}
