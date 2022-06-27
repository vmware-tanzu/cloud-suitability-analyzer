/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa

import (
	"csa-app/model"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetLevelForScore(t *testing.T) {

	score := 0
	assert.Equal(t, model.Info_criticality, GetLevelForScore(score))

	score = 1
	assert.Equal(t, model.Low_criticality, GetLevelForScore(score))

	score = 3
	assert.Equal(t, model.Low_criticality, GetLevelForScore(score))

	score = 4
	assert.Equal(t, model.Medium_criticality, GetLevelForScore(score))

	score = 6
	assert.Equal(t, model.Medium_criticality, GetLevelForScore(score))

	score = 7
	assert.Equal(t, model.High_criticality, GetLevelForScore(score))

	score = 10
	assert.Equal(t, model.High_criticality, GetLevelForScore(score))

	score = math.MaxInt64
	assert.Equal(t, model.High_criticality, GetLevelForScore(score))

}
