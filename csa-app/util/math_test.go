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

func TestSimpleExpression(t *testing.T) {
	expression := "7 * 7"
	ex := util.NewExpression(nil)
	result, err := ex.Calculate(expression)

	assert.Nil(t, err)
	assert.Equal(t, 49.0, result)
}

func TestSimpleLog10Expression(t *testing.T) {
	expression := "log(10,10)"
	ex := util.NewExpression(nil)
	result, err := ex.Calculate(expression)

	assert.Nil(t, err)
	assert.Equal(t, 1.0, result)
}

func TestSimpleLogOtherBaseExpression(t *testing.T) {
	expression := "log(2,16)"
	ex := util.NewExpression(nil)
	result, err := ex.Calculate(expression)

	assert.Nil(t, err)
	assert.Equal(t, 4.0, result)
}

func TestSimpleExpressionWithParams(t *testing.T) {

	params := make(map[string]interface{})
	params["sloc"] = 7

	expression := "sloc * 7"
	ex := util.NewExpression(params)
	result, err := ex.Calculate(expression)

	assert.Nil(t, err)
	assert.Equal(t, 49.0, result)
}

func TestSimpleLog10ExpressionWithParams(t *testing.T) {

	params := make(map[string]interface{})
	params["sloc"] = 10

	expression := "log(10,sloc)"
	ex := util.NewExpression(params)
	result, err := ex.Calculate(expression)

	assert.Nil(t, err)
	assert.Equal(t, 1.0, result)
}

func TestSimpleLogOtherBaseExpressionWithParams(t *testing.T) {

	params := make(map[string]interface{})
	params["sloc"] = 16

	expression := "log(2,sloc)"
	ex := util.NewExpression(params)
	result, err := ex.Calculate(expression)

	assert.Nil(t, err)
	assert.Equal(t, 4.0, result)
}
