/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"math"

	"github.com/Knetic/govaluate"
)

type Expression struct {
	Functions map[string]govaluate.ExpressionFunction
	Params    map[string]interface{}
}

var LOG_0_ERROR = fmt.Errorf("attempt to perform log operation on zero or negative #")

func NewExpression(parameters map[string]interface{}) *Expression {

	functions := make(map[string]govaluate.ExpressionFunction)

	functions["log"] = func(args ...interface{}) (interface{}, error) {
		base, err := getFloat(args[0])
		if err != nil {
			return math.NaN(), fmt.Errorf("invalid base[%v] for log function! details: %s", args[0], err.Error())
		}

		val, err := getFloat(args[1])
		if err != nil {
			return math.NaN(), fmt.Errorf("invalid target value [%v] for log function! details: %s", args[1], err.Error())
		}

		if val <= 0 {
			return math.NaN(), LOG_0_ERROR
		}

		if base == 10.0 {
			return math.Log10(val), nil
		} else {
			return math.Log10(val) / math.Log10(base), nil
		}
	}

	//If needed more functions can be added here...

	return &Expression{Functions: functions, Params: parameters}
}

func (ex *Expression) Calculate(target string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(target, ex.Functions)
	result, err := expression.Evaluate(ex.Params)

	if err == nil {
		return getFloat(result)
	}

	return math.NaN(), err
}

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	default:
		return math.NaN(), fmt.Errorf("cannot covert [%v] to a float", unk)
	}
}
