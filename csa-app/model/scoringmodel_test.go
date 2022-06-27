/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model_test

import (
	"fmt"
	"log"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"csa-app/db"
	db_test_support "csa-app/db/test_support"
	"csa-app/model"
	"csa-app/util"
)

func TestScoreModelValidatesRangesExist(t *testing.T) {
	testModel := model.ScoringModel{Name: "testModel1"}
	var ranges []*model.Bucket
	err := testModel.Validate()
	assert.Contains(t, err.Error(), "a valid score model must contain 1 or more ranges", "Validation Should Fail but did not")

	testModel.Ranges = ranges
	err = testModel.Validate()
	assert.Contains(t, err.Error(), "a valid score model must contain 1 or more ranges", "Validation Should Fail but did not")

}

func TestScoreModelValidatesRangesCoverZeroToMaxInt(t *testing.T) {
	testModel := model.ScoringModel{Name: "testModel1"}
	var ranges []*model.Bucket
	bkt := model.Bucket{Type: model.SLOC_BKT_TYPE}

	bkt.Start = "1"
	bkt.End = model.MAX_INT
	ranges = append(ranges, &bkt)
	bkt.AddRangeWithOutcome(model.RAW_BKT_TYPE, math.MinInt64, math.MaxInt64, 0, "test-outcome")

	testModel.Ranges = ranges

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt.Start = "0"
	bkt.End = "1000"

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt.End = model.MAX_INT

	assert.Nil(t, testModel.Validate(), "Validation should have succeeded but did not!")

}

func TestScoreModelValidatesRangesDoNotOverlap(t *testing.T) {
	testModel := model.ScoringModel{Name: "testModel1"}
	var ranges []*model.Bucket
	bkt := model.Bucket{Type: model.SLOC_BKT_TYPE}
	bkt2 := model.Bucket{Type: model.SLOC_BKT_TYPE}
	bkt.Start = "0"
	bkt.End = "5"
	bkt.AddRangeWithOutcome(model.RAW_BKT_TYPE, math.MinInt64, math.MaxInt64, 10, "test-outcome")
	ranges = append(ranges, &bkt)
	bkt2.Start = "4"
	bkt2.End = model.MAX_INT
	bkt2.AddRangeWithOutcome(model.RAW_BKT_TYPE, math.MinInt64, math.MaxInt64, 0, "test-outcome")

	ranges = append(ranges, &bkt2)
	testModel.Ranges = ranges

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt2.Start = "5"

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt2.Start = "6"
	assert.Nil(t, testModel.Validate(), "Validation should have succeeded but did not!")

}

func TestScoreModelValidatesRangesAreContiguous(t *testing.T) {
	testModel := model.ScoringModel{Name: "testModel1"}
	var ranges []*model.Bucket
	bkt := model.Bucket{Type: model.SLOC_BKT_TYPE}
	bkt2 := model.Bucket{Type: model.SLOC_BKT_TYPE}
	bkt.Start = "0"
	bkt.End = "5"
	bkt.AddRangeWithOutcome(model.RAW_BKT_TYPE, math.MinInt64, math.MaxInt64, 0, "test-outcome")

	ranges = append(ranges, &bkt)
	bkt2.Start = "7"
	bkt2.End = model.MAX_INT
	bkt2.AddRangeWithOutcome(model.RAW_BKT_TYPE, math.MinInt64, math.MaxInt64, 0, "test-outcome")
	ranges = append(ranges, &bkt2)
	testModel.Ranges = ranges

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt2.Start = "27"

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt2.Start = "6"
	assert.Nil(t, testModel.Validate(), "Validation should have succeeded but did not!")

}

func TestBucketValidatesRangesExist(t *testing.T) {
	testModel := model.Bucket{Type: model.SLOC_BKT_TYPE}
	testModel.Start = "0"
	testModel.End = model.MAX_INT
	var ranges []*model.Bucket
	err := testModel.Validate()
	assert.Contains(t, err.Error(), "has no child ranges and no outcome", "Validation Should Fail but did not")

	testModel.Ranges = ranges
	err = testModel.Validate()
	assert.Contains(t, err.Error(), "has no child ranges and no outcome", "Validation Should Fail but did not")

	testModel.Outcome = &model.Outcome{Score: 10, Recommendation: "test-outcome"}
	err = testModel.Validate()
	assert.Nil(t, err, "Bucket should validate without ranges if it has an outcome!")

	testModel.Outcome = nil
	testModel.AddRangeWithOutcome(model.RAW_BKT_TYPE, math.MinInt64, math.MaxInt64, 0, "test-outcome")
	err = testModel.Validate()
	assert.Nil(t, err, "Bucket should validate without outcome if it has ranges!")

}

func TestBucketValidatesRangesCoverMinToMax(t *testing.T) {
	testModel := model.Bucket{Type: model.SLOC_BKT_TYPE}
	testModel.Start = "0"
	testModel.End = model.MAX_INT
	var ranges []*model.Bucket
	bkt := model.Bucket{Type: model.RAW_BKT_TYPE}

	bkt.Start = "0"
	bkt.End = model.MAX_INT
	bkt.AddRangeWithOutcome(model.BV_BKT_TYPE, model.FLT_MIN, math.MaxFloat64, 0, "test")

	ranges = append(ranges, &bkt)
	testModel.Ranges = ranges

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt.Start = model.MIN_INT
	bkt.End = "1000"

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt.End = model.MAX_INT

	assert.Nil(t, testModel.Validate(), "Validation should have succeeded but did not!")

}

func TestBucketValidatesRangesDoNotOverlap(t *testing.T) {
	testModel := model.Bucket{Type: model.SLOC_BKT_TYPE}
	testModel.Start = "0"
	testModel.End = model.MAX_INT
	var ranges []*model.Bucket
	bkt := model.Bucket{Type: model.RAW_BKT_TYPE}
	bkt.AddRangeWithOutcome(model.BV_BKT_TYPE, model.FLT_MIN, math.MaxFloat64, 0, "test")

	bkt2 := model.Bucket{Type: model.RAW_BKT_TYPE}

	bkt2.AddRangeWithOutcome(model.BV_BKT_TYPE, model.FLT_MIN, math.MaxFloat64, 0, "test")

	bkt.Start = model.MIN_INT
	bkt.End = "5"
	ranges = append(ranges, &bkt)
	bkt2.Start = "4"
	bkt2.End = model.MAX_INT
	ranges = append(ranges, &bkt2)
	testModel.Ranges = ranges

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt2.Start = "5"

	assert.NotNil(t, testModel.Validate(), "Validation should have failed but did not!")

	bkt2.Start = "6"
	assert.Nil(t, testModel.Validate(), "Validation should have succeeded but did not!")

}

func TestScoreModelReturnsCorrectOutcome(t *testing.T) {
	sm := createValidScoringModel("Test")

	app := &model.Application{SlocCnt: 800, RawScore: 800, BusinessValue: -1.0}
	err := app.CalculateScore(sm)
	assert.Nil(t, err)
	assert.Equal(t, "Rehost to TKG", app.Recommendation)
	assert.Equal(t, 7.1, app.Score)

	app.BusinessValue = 6.0
	err = app.CalculateScore(sm)
	assert.Nil(t, err)
	assert.Equal(t, "Refactor to TAS", app.Recommendation)

	app.SlocCnt = 48324
	app.RawScore = 78987987
	app.BusinessValue = 0
	err = app.CalculateScore(sm)
	assert.Nil(t, err)
	assert.Equal(t, "Rehost to TKG", app.Recommendation)

	app.BusinessValue = 10
	err = app.CalculateScore(sm)
	assert.Nil(t, err)
	assert.Equal(t, "Refactor to TAS", app.Recommendation)

}

func TestExport(t *testing.T) {

	fu := util.NewFileUtil()
	//Establish Temp Directory for Run
	tmpPath, err := fu.EstablishTempDirectory("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create temp directory. Details: %v", err)
		assert.Fail(t, "Failed to create temp directory!")
	} else {
		fmt.Printf("Created Tmp Dir => %s\n", tmpPath)
	}

	defer fu.RemoveDir(tmpPath)

	*util.ModelsDir = util.DEFAULT_MODELS_DIR
	*util.OutputDir = tmpPath
	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	repo := db.NewScoringRepository(database)
	err = repo.SaveModel(createValidScoringModel("Test-Model"))
	assert.Nil(t, err, "Save of scoring model failed!")

	assert.Nil(t, repo.ExportModels())

}

func createValidScoringModel(name string) *model.ScoringModel {
	logExp := fmt.Sprintf("%s - log(10,%s)", model.MAX_SCORE, model.RAW_SCORE_SCORING_TOKEN)
	newModel := &model.ScoringModel{Name: name, MaxScore: 10, MinScore: 0}

	bkt1 := newModel.AddRange(model.SLOC_BKT_TYPE, 0, math.MaxInt64)

	rBkt := bkt1.AddRange(model.RAW_BKT_TYPE, math.MinInt64, 100)
	rBkt.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, model.FLT_MIN, math.MaxFloat64, logExp, "Deploy to TAS")

	rBkt2 := bkt1.AddRange(model.RAW_BKT_TYPE, 101, 10000)
	rBkt2.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, model.FLT_MIN, 5.0, logExp, "Rehost to TKG")
	rBkt2.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, 5.01, math.MaxFloat64, logExp, "Refactor to TAS")

	rBkt3 := bkt1.AddRange(model.RAW_BKT_TYPE, 10001, 10000000)
	rBkt3.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, model.FLT_MIN, 5.0, logExp, "Rehost to TKG")
	rBkt3.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, 5.01, math.MaxFloat64, logExp, "Refactor to TAS")

	rBkt4 := bkt1.AddRange(model.RAW_BKT_TYPE, 10000001, math.MaxInt64)
	rBkt4.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, model.FLT_MIN, 5.0, logExp, "Rehost to TKG")
	rBkt4.AddRangeWithCalculatedScore(model.BV_BKT_TYPE, 5.01, math.MaxFloat64, logExp, "Refactor to TAS")

	return newModel

}
