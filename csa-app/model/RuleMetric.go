/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"strings"
	"sync"
	"time"
)

type RuleMetric struct {
	ID              uint          `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt       time.Time     `json:"-" yaml:"-"`
	UpdatedAt       time.Time     `json:"-" yaml:"-"`
	Rule            string        `gorm:"type:text;index;not null" json:"rule" yaml:"rule"`
	RuleCriticality string        `gorm:"type:text;index;not null" json:"-" yaml:"-"`
	RunID           uint          `gorm:"index;not null"`
	Checks          int64         `json:"checks" yaml:"checks"`
	PatternChecks   int64         `json:"patternChecks" yaml:"patternChecks"`
	Hits            int64         `json:"hits" yaml:"hits"`
	TotalTime       time.Duration `json:"-" yaml:"-"`
	Longest         time.Duration `json:"-" yaml:"-"`
	Shortest        time.Duration `json:"-" yaml:"-"`
	AvgRule         time.Duration `json:"-" yaml:"-"`
	AvgPat          time.Duration `json:"-" yaml:"-"`
	AvgHit          time.Duration `json:"-" yaml:"-"`
	TotalTimeStr    string        `json:"totalTime" yaml:"totalTime"`
	LongestStr      string        `json:"longest" yaml:"longest"`
	ShortestStr     string        `json:"shortest" yaml:"shortest"`
	AvgRuleStr      string        `json:"avgRule" yaml:"avgRule"`
	AvgPatStr       string        `json:"avgPat" yaml:"avgPat"`
	AvgHitStr       string        `json:"avgHit" yaml:"avgHit"`
	sync.Mutex      `gorm:"-" json:"-" yaml:"-"`
}

func (r *RuleMetric) PrePersist() {
	r.TotalTimeStr = r.TotalTime.String()
	r.LongestStr = r.Longest.String()
	r.ShortestStr = r.Shortest.String()
	r.AvgRuleStr = r.AvgRule.String()
	r.AvgPatStr = r.AvgPat.String()
	r.AvgHitStr = r.AvgHit.String()
}

func (r *RuleMetric) Accumulate(patternCheck int64, hitcnt int64, duration time.Duration) {
	r.Lock()
	defer r.Unlock()

	r.Hits += hitcnt
	r.PatternChecks += patternCheck
	r.Checks++
	r.TotalTime += duration

	if r.Longest < duration {
		r.Longest = duration
	} else if r.Shortest > duration || r.Shortest == 0 {
		r.Shortest = duration
	}
	if r.Hits > 0 {
		r.AvgHit = time.Duration(r.TotalTime.Nanoseconds() / r.Hits)
	}

	if r.PatternChecks > 0 {
		r.AvgPat = time.Duration(r.TotalTime.Nanoseconds() / r.PatternChecks)
	}
	r.AvgRule = time.Duration(r.TotalTime.Nanoseconds() / r.Checks)
}

func (r *RuleMetric) Merge(ruleMetric *RuleMetric) {

	r.Hits += ruleMetric.Hits
	r.PatternChecks += ruleMetric.PatternChecks
	r.Checks += ruleMetric.Checks
	r.TotalTime += ruleMetric.TotalTime

	if r.Longest < ruleMetric.Longest {
		r.Longest = ruleMetric.Longest
	} else if r.Shortest > ruleMetric.Shortest || r.Shortest == 0 {
		r.Shortest = ruleMetric.Shortest
	}
	if r.Hits > 0 {
		r.AvgHit = time.Duration(r.TotalTime.Nanoseconds() / r.Hits)
	}

	if r.PatternChecks > 0 {
		r.AvgPat = time.Duration(r.TotalTime.Nanoseconds() / r.PatternChecks)
	}

	if r.Checks > 0 {
		r.AvgRule = time.Duration(r.TotalTime.Nanoseconds() / r.Checks)
	}
}

func (r *RuleMetric) CriticalityFactor() int {

	factor := 6

	if r.Hits < 1 {
		if r.Checks < 1 {
			factor = 0
		} else {
			factor = 3
		}
	}

	switch strings.ToLower(r.RuleCriticality) {

	case "high":
		return factor + 2
	case "low":
		return factor
	default:
		return factor + 1
	}
}

type MetricByTimeAndHits []KV

func (ruleMetrics MetricByTimeAndHits) Len() int { return len(ruleMetrics) }
func (ruleMetrics MetricByTimeAndHits) Swap(i, j int) {
	ruleMetrics[i], ruleMetrics[j] = ruleMetrics[j], ruleMetrics[i]
}
func (ruleMetrics MetricByTimeAndHits) Less(i, j int) bool {
	r1 := ruleMetrics[i]
	r2 := ruleMetrics[j]
	if r1.Value.TotalTime == r2.Value.TotalTime {
		return r1.Value.Hits > r2.Value.Hits
	}
	return r1.Value.TotalTime > r2.Value.TotalTime
}

type MetricByHitsAndTime []KV

func (ruleMetrics MetricByHitsAndTime) Len() int { return len(ruleMetrics) }
func (ruleMetrics MetricByHitsAndTime) Swap(i, j int) {
	ruleMetrics[i], ruleMetrics[j] = ruleMetrics[j], ruleMetrics[i]
}
func (ruleMetrics MetricByHitsAndTime) Less(i, j int) bool {
	r1 := ruleMetrics[i]
	r2 := ruleMetrics[j]
	if r1.Value.Hits == r2.Value.Hits {
		return r1.Value.TotalTime > r2.Value.TotalTime
	}
	return r1.Value.Hits > r2.Value.Hits
}

type MetricByCriticalityHitsAndTime []KV

func (ruleMetrics MetricByCriticalityHitsAndTime) Len() int { return len(ruleMetrics) }
func (ruleMetrics MetricByCriticalityHitsAndTime) Swap(i, j int) {
	ruleMetrics[i], ruleMetrics[j] = ruleMetrics[j], ruleMetrics[i]
}
func (ruleMetrics MetricByCriticalityHitsAndTime) Less(i, j int) bool {
	r1 := ruleMetrics[i]
	r2 := ruleMetrics[j]

	if r1.Value.CriticalityFactor() == r2.Value.CriticalityFactor() {
		if r1.Value.Hits == r2.Value.Hits {
			return r1.Value.TotalTime > r2.Value.TotalTime
		}
		return r1.Value.Hits > r2.Value.Hits
	}
	return r1.Value.CriticalityFactor() > r2.Value.CriticalityFactor()

}

type KV struct {
	Key   string
	Value *RuleMetric
}
