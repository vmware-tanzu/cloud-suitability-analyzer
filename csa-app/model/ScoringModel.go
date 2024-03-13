/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"csa-app/util"
)

const MAX_INT = "int.max"
const MIN_INT = "int.min"
const MAX_FLT = "flt.max"
const MIN_FLT = "flt.min"
const MAX_SCORE = "max_score"
const MIN_SCORE = "min_score"

const SLOC_BKT_TYPE = "sloc"
const RAW_BKT_TYPE = "raw"
const BV_BKT_TYPE = "bv"

const FLT_MIN = math.MaxFloat64 * -1

type ScoringModel struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	Name      string    `gorm:"type:text;" json:"name" yaml:"name"`
	MaxScore  float64   `gorm:"-"json:"max-score,omitempty" yaml:"max-score,omitempty"`
	MinScore  float64   `gorm:"-" json:"min-score,omitempty" yaml:"min-score,omitempty"`
	Ranges    []*Bucket `gorm:"-" json:"ranges" yaml:"ranges"`
	Json      string    `gorm:"serializedModel" json:"-" yaml:"-"`
}

type Bucket struct {
	Type    string    `gorm:"-" json:"type" yaml:"type"`
	Start   string    `json:"start" yaml:"start"`
	End     string    `json:"end" yaml:"end"`
	Ranges  []*Bucket `gorm:"-" json:"ranges,omitempty" yaml:"ranges,omitempty"`
	Outcome *Outcome  `gorm:"-" json:"outcome,omitempty" yaml:"outcome,omitempty"`
}

type Outcome struct {
	Calculate      bool    `gorm:"-" json:"calculate,omitempty" yaml:"calculate,omitempty"`
	Expression     string  `gorm:"-" json:"expression,omitempty" yaml:"expression,omitempty"`
	Score          float64 `gorm:"-" json:"score,omitempty" yaml:"score,omitempty"`
	Recommendation string  `gorm:"-" json:"recommendation" yaml:"recommendation"`
}

func (sm *ScoringModel) PrePersist() error {
	out, err := json.Marshal(sm)

	if err != nil {
		return err
	}

	sm.Json = string(out)

	return nil
}

func (sm *ScoringModel) Hydrate() (err error) {
	hModel := &ScoringModel{}
	err = json.Unmarshal([]byte(sm.Json), hModel)
	if err == nil {
		sm.MinScore = hModel.MinScore
		sm.MaxScore = hModel.MaxScore
		sm.Ranges = hModel.Ranges
	}
	return
}

func (sm *ScoringModel) AddRange(bktType string, start interface{}, end interface{}) *Bucket {
	var bkt *Bucket
	if bktType == BV_BKT_TYPE {
		bkt = &Bucket{Type: bktType, Start: getStringValFromFloat(start.(float64)), End: getStringValFromFloat(end.(float64))}
	} else {
		bkt = &Bucket{Type: bktType, Start: getStringVal(start.(int)), End: getStringVal(end.(int))}
	}
	sm.Ranges = append(sm.Ranges, bkt)
	return bkt
}

func (sm *ScoringModel) AddRangeWithOutcome(bktType string, start interface{}, end interface{}, score float64, outcome string) *Bucket {

	var bkt *Bucket
	if bktType == BV_BKT_TYPE {
		bkt = &Bucket{Type: bktType, Start: getStringValFromFloat(start.(float64)), End: getStringValFromFloat(end.(float64))}
	} else {
		bkt = &Bucket{Type: bktType, Start: getStringVal(start.(int)), End: getStringVal(end.(int))}
	}

	bkt.Outcome = &Outcome{Score: score, Recommendation: outcome}
	sm.Ranges = append(sm.Ranges, bkt)
	return bkt
}

func (sm *ScoringModel) AddRangeWithCalculatedScore(bktType string, start interface{}, end interface{}, expression string, outcome string) *Bucket {

	var bkt *Bucket
	if bktType == BV_BKT_TYPE {
		bkt = &Bucket{Type: bktType, Start: getStringValFromFloat(start.(float64)), End: getStringValFromFloat(end.(float64))}
	} else {
		bkt = &Bucket{Type: bktType, Start: getStringVal(start.(int)), End: getStringVal(end.(int))}
	}

	bkt.Outcome = &Outcome{Calculate: true, Expression: expression, Recommendation: outcome}
	sm.Ranges = append(sm.Ranges, bkt)
	return bkt
}

func (sm *ScoringModel) ProcessScore(app Application) *Outcome {
	for _, bkt := range sm.Ranges {
		outcome := bkt.GetOutcome(app)
		if outcome != nil {
			return outcome
		}
	}

	return nil
}

func (sm *ScoringModel) Validate() error {
	if sm.Name == "" {
		return fmt.Errorf("a valid scoring model must have a unique non-null name")
	}

	if sm.Ranges == nil || len(sm.Ranges) == 0 {
		return fmt.Errorf("a valid score model must contain 1 or more ranges of the same type (sloc, raw, or bv) that are contiguous and cover required range of values for tye range-type")
	}

	return validateBuckets(sm.Ranges)
}

func (b *Bucket) AddRange(bktType string, start interface{}, end interface{}) *Bucket {
	var bkt *Bucket
	if bktType == BV_BKT_TYPE {
		bkt = &Bucket{Type: bktType, Start: getStringValFromFloat(start.(float64)), End: getStringValFromFloat(end.(float64))}
	} else {
		bkt = &Bucket{Type: bktType, Start: getStringVal(start.(int)), End: getStringVal(end.(int))}
	}
	b.Ranges = append(b.Ranges, bkt)
	return bkt
}

func (b *Bucket) AddRangeWithOutcome(bktType string, start interface{}, end interface{}, score float64, outcome string) *Bucket {

	var bkt *Bucket
	if bktType == BV_BKT_TYPE {
		bkt = &Bucket{Type: bktType, Start: getStringValFromFloat(start.(float64)), End: getStringValFromFloat(end.(float64))}
	} else {
		bkt = &Bucket{Type: bktType, Start: getStringVal(start.(int)), End: getStringVal(end.(int))}
	}

	bkt.Outcome = &Outcome{Score: score, Recommendation: outcome}
	b.Ranges = append(b.Ranges, bkt)
	return bkt
}

func (b *Bucket) AddRangeWithCalculatedScore(bktType string, start interface{}, end interface{}, expression string, outcome string) *Bucket {

	var bkt *Bucket
	if bktType == BV_BKT_TYPE {
		bkt = &Bucket{Type: bktType, Start: getStringValFromFloat(start.(float64)), End: getStringValFromFloat(end.(float64))}
	} else {
		bkt = &Bucket{Type: bktType, Start: getStringVal(start.(int)), End: getStringVal(end.(int))}
	}

	bkt.Outcome = &Outcome{Calculate: true, Expression: expression, Recommendation: outcome}
	b.Ranges = append(b.Ranges, bkt)
	return bkt
}

func (sb *Bucket) TheStart() interface{} {
	if sb.Type == BV_BKT_TYPE {
		v, _ := getFloatVal(sb.Start)
		return v
	} else {
		v, _ := getInt64Val(sb.Start)
		return v
	}
}

func (sb *Bucket) TheEnd() interface{} {
	if sb.Type == BV_BKT_TYPE {
		v, _ := getFloatVal(sb.End)
		return v
	} else {
		v, _ := getInt64Val(sb.End)
		return v
	}
}

func (b *Bucket) Contains(target interface{}) bool {
	if b.Type == BV_BKT_TYPE {
		return target.(float64) >= b.TheStart().(float64) && b.TheEnd().(float64) >= target.(float64)
	} else {
		return target.(int64) >= b.TheStart().(int64) && b.TheEnd().(int64) >= target.(int64)
	}
}

func (b *Bucket) GetOutcome(app Application) *Outcome {
	hit := false

	switch b.Type {
	case SLOC_BKT_TYPE:
		hit = b.Contains(int64(app.SlocCnt))
	case RAW_BKT_TYPE:
		hit = b.Contains(app.RawScore)
	case BV_BKT_TYPE:
		hit = b.Contains(app.BusinessValue)
	}
	if hit {
		if b.Ranges != nil && len(b.Ranges) > 0 {
			//Check the next level!
			for _, childBkt := range b.Ranges {
				outcome := childBkt.GetOutcome(app)
				if outcome != nil {
					return outcome
				}
			}
		} else {
			return b.Outcome
		}
	}

	return nil
}

func (b *Bucket) Validate() error {

	if b.Type == BV_BKT_TYPE {

		s, err := getFloatVal(b.Start)
		min := FLT_MIN
		if err != nil {
			return fmt.Errorf("[%s] bucket has invalid start [%s]. details: %s", b.Type, b.Start, err.Error())
		}
		if s < min {
			return fmt.Errorf("[%s] bucket is invalid.  start [%s] must be greater than or equal to [%s]", b.Type, b.Start, MIN_FLT)
		}

		e, err := getFloatVal(b.End)
		if err != nil {
			return fmt.Errorf("[%s] bucket has invalid end [%s]. details: %s", b.Type, b.End, err.Error())
		}
		if s > e {
			return fmt.Errorf("[%s] bucket is invalid. end [%s] must be greater than start [%s]", b.Type, b.End, b.Start)
		}

	} else {

		s, err := getInt64Val(b.Start)
		min := int64(0)
		if b.Type == RAW_BKT_TYPE {
			min = math.MinInt64
		}
		if err != nil {
			return fmt.Errorf("[%s] bucket has invalid start [%s]. details: %s", b.Type, b.Start, err.Error())
		}
		if s < min {
			return fmt.Errorf("[%s] bucket is invalid.  start [%s] must be greater than or equal to [%d]", b.Type, b.Start, min)
		}

		e, err := getInt64Val(b.End)
		if err != nil {
			return fmt.Errorf("[%s] bucket has invalid end [%s]. details: %s", b.Type, b.End, err.Error())
		}
		if s > e {
			return fmt.Errorf("[%s] bucket is invalid. end [%s] must be greater than start [%s]", b.Type, b.End, b.Start)
		}

	}

	if b.Ranges == nil || len(b.Ranges) == 0 {
		if b.Outcome == nil {
			return fmt.Errorf("[%s] bucket starting @ [%s] ending @ [%s] has no child ranges and no outcome", b.Type, b.Start, b.End)
		} else {
			if b.Outcome.Recommendation == "" {
				return fmt.Errorf("[%s] bucket starting @ [%s] ending @ [%s] has no child ranges and an outcome with no 'recommendation'. A recommendation is required", b.Type, b.Start, b.End)
			}
		}
	} else {
		return validateBuckets(b.Ranges)
	}

	return nil
}

func (b *Bucket) StartsBefore(b2 *Bucket) bool {
	if b.Type != b2.Type {
		return true
	}
	if b.Type == BV_BKT_TYPE {
		return b.TheStart().(float64) < b2.TheStart().(float64)
	} else {
		return b.TheStart().(int64) < b2.TheStart().(int64)
	}
}

func GetModelsDir(isImport bool) string {

	dir := *util.ModelsDir

	//User overrode default!
	if dir != util.DEFAULT_MODELS_DIR || isImport {
		return dir
	}

	dir = *util.OutputDir
	if !strings.HasSuffix(dir, util.PathSeparator) {
		dir = dir + util.PathSeparator
	}

	return dir + "scoring-models"

}

func validateBuckets(bkts []*Bucket) error {
	sort.Sort(BucketsByStart(bkts))

	if bkts[0].Type == BV_BKT_TYPE {
		return validateFloatBuckets(bkts)
	} else {
		return validateIntBuckets(bkts)
	}

	return nil
}

func validateIntBuckets(bkts []*Bucket) error {
	bktTypes := ""

	cnt := 0
	lastEnd := int64(0)
	rangeMin := int64(math.MinInt64)
	for _, bkt := range bkts {
		bStart := bkt.TheStart().(int64)
		bEnd := bkt.TheEnd().(int64)

		if cnt == 0 {
			bktTypes = bkt.Type
			if bkt.Type == SLOC_BKT_TYPE {
				rangeMin = 0
			}
			//First bucket should start with rangeMin!
			if bStart != rangeMin {
				return fmt.Errorf("range of %s buckets does not begin with range minimum of [%d]", bkt.Type, rangeMin)
			}
		} else if bStart < lastEnd {
			return fmt.Errorf("range of %s buckets must not overlap! bucket Ending with [%d] overlaps with bucket starting with [%d]", bkt.Type, lastEnd, bStart)
		} else if bStart != lastEnd+1 {
			return fmt.Errorf("range of %s buckets is not contigous! bucket ending with [%d] has a gap with bucket starting with [%d]", bkt.Type, lastEnd, bStart)
		}

		if bktTypes != bkt.Type {
			return fmt.Errorf("all buckets in a range must be of the same type. child buckets in range identified with different types [%s] and [%s]", bkt.Type, bktTypes)
		}

		//Validate Each Bucket!
		err := bkt.Validate()
		if err != nil {
			return err
		}

		lastEnd = bEnd
		cnt++
	}

	if lastEnd != math.MaxInt64 {
		return fmt.Errorf("range of %s buckets does not cover full required range. last bucket ends with [%d] and should end with [%d]", bktTypes, lastEnd, math.MaxInt64)
	}

	return nil
}

func validateFloatBuckets(bkts []*Bucket) error {
	bktTypes := ""
	cnt := 0
	lastEnd := 0.0
	for _, bkt := range bkts {

		bStart := bkt.TheStart().(float64)
		bEnd := bkt.TheEnd().(float64)

		if cnt == 0 {
			bktTypes = bkt.Type
			//First bucket should start with rangeMin!
			if bStart != FLT_MIN {
				return fmt.Errorf("range of %s buckets does not begin with range minimum of [%s]", bkt.Type, MIN_FLT)
			}
		} else if bStart < lastEnd {
			return fmt.Errorf("range of %s buckets must not overlap! bucket ending with [%f] overlaps with bucket starting with [%f]", bkt.Type, lastEnd, bStart)
		} else if bStart != lastEnd+.01 {
			return fmt.Errorf("range of %s buckets is not contigous! bucket ending with [%f] has a gap with bucket starting with [%f]", bkt.Type, lastEnd, bStart)
		}

		//Validate Each Bucket!
		err := bkt.Validate()
		if err != nil {
			return err
		}

		lastEnd = bEnd
		cnt++
	}

	if lastEnd != math.MaxFloat64 {
		return fmt.Errorf("range of %s buckets does not cover full required range. last bucket ends with [%f] and should end with [%f]", bktTypes, lastEnd, math.MaxFloat64)
	}

	return nil
}

func getInt64Val(target string) (int64, error) {

	if target == "" {
		return math.MinInt64, nil
	}

	lVal := strings.ToLower(target)

	if lVal == MIN_INT {
		return math.MinInt64, nil
	}

	if lVal == MAX_INT {
		return math.MaxInt64, nil
	}

	val, err := strconv.ParseInt(target, 10, 64)

	if err != nil {
		return math.MinInt64, err
	}

	return val, nil
}

func getFloatVal(target string) (float64, error) {

	if target == "" {
		return math.MaxFloat64 * -1, nil
	}

	lVal := strings.ToLower(target)

	if lVal == MIN_FLT || lVal == MIN_INT {
		return math.MaxFloat64 * -1, nil
	}

	if lVal == MAX_INT || lVal == MAX_FLT {
		return math.MaxFloat64, nil
	}

	val, err := strconv.ParseFloat(target, 64)

	if err != nil {
		return math.NaN(), err
	}

	return val, nil
}

func getStringVal(val int) string {
	if val == math.MinInt64 {
		return MIN_INT
	}
	if val == math.MaxInt64 {
		return MAX_INT
	}
	return strconv.Itoa(val)
}

func getStringValFromFloat(val float64) string {
	if val == math.MaxFloat64*-1 {
		return MIN_FLT
	}
	if val == math.MaxFloat64 {
		return MAX_FLT
	}
	return fmt.Sprintf("%3.2f", val)
}

type BucketsByStart []*Bucket

func (buckets BucketsByStart) Len() int { return len(buckets) }
func (buckets BucketsByStart) Swap(i, j int) {
	buckets[i], buckets[j] = buckets[j], buckets[i]
}
func (buckets BucketsByStart) Less(i, j int) bool {
	f1 := buckets[i]
	f2 := buckets[j]

	if f1.TheStart() == f2.TheStart() {
		return true
	}

	return f1.StartsBefore(f2)
}
