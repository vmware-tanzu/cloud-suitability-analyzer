/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"csa-app/util"

	"github.com/sirupsen/logrus"
)

const RAW_SCORE_SCORING_TOKEN = "raw_score"
const FILES_SCORING_TOKEN = "files_cnt"
const CLOUD_SCORE_SCORING_TOKEN = "raw_container_score"
const CONTAINER_SCORE_SCORING_TOKEN = "raw_cloud_score"
const FINDINGS_SCORING_TOKEN = "findings_cnt"
const SLOC_CNT_SCORING_TOKEN = "sloc_cnt"
const BUSINESS_VALUE_SCORING_TOKEN = "bv"

type Application struct {
	ID                uint              `gorm:"primary_key" json:"appId" yaml:"appId"`
	CreatedAt         time.Time         `json:"-" yaml:"-"`
	UpdatedAt         time.Time         `json:"-" yaml:"-"`
	RunID             uint              `gorm:"index;not null" sql:"type:bigint REFERENCES runs(id) ON DELETE CASCADE" json:"runId"  yaml:"runId"`
	Name              string            `gorm:"type:text;" json:"name" yaml:"name"`
	Path              string            `gorm:"type:text;" json:"path" yaml:"path"`
	Category          string            `gorm:"index;not null" json:"category,omitempty" yaml:"category,omitempty"`
	Criticality       string            `gorm:"index;not null" json:"criticality,omitempty" yaml:"criticality,omitempty"`
	BusinessDomain    string            `gorm:"type:text;" json:"businessdomain" yaml:"businessdomain"`
	BusinessValue     float64           `gorm:"default:'-1'" json:"businessvalue" yaml:"businessvalue"`
	Findings          int               `json:"findings"`
	CIFindings        int               `json:"ciFindings"`
	PrimaryLanguage   string            `gorm:"-" json:"-" yaml:"-"`
	PrimaryLangCount  int               `gorm:"-" json:"-" yaml:"-"`
	InfoFindings      int               `json:"infoFindings"`
	RawScore          int64             `json:"rawScore"`
	RawContainerScore int               `json:"rawContainerScore"`
	RawCloudScore     int               `json:"rawCloudScore"`
	NumCrits          int               `json:"numCrits"`
	ScoringModel      string            `json:"model" yaml:"model"`
	Score             float64           `json:"score"`
	OriginalScore     float64           `gorm:"default:'-1.0'" json:"originalScore"`
	ScoreModified     bool              `json:"scoreModified"`
	Recommendation    string            `json:"recommendation"`
	SlocCnt           int               `json:"slocCnt"`
	FilesCnt          int               `json:"filesCnt"`
	FindingsRatio     float64           `json:"findingsRatio"`
	Tags              []*ApplicationTag `gorm:"foreignkey:ApplicationID" json:"tags" yaml:"tags"`
	Files             []*util.FileInfo  `gorm:"-" json:"-" yaml:"-"`
	IgnoredFiles      []*util.FileInfo  `gorm:"-" json:"-" yaml:"-"`
	FileUtil          *util.FileUtil    `gorm:"-" json:"-" yaml:"-"`
	Rules             []Rule            `gorm:"-" json:"-" yaml:"-"`
	MatchedRules      map[string]int    `gorm:"-" json:"-" yaml:"-"`
	Bins              []Bin             `gorm:"-" json:"bins" yaml:"bins"`
	Model             *ScoringModel     `gorm:"-" json:"-" yaml:"-"`
	sync.Mutex        `gorm:"-" json:"-" yaml:"-"`
}

type ApplicationTag struct {
	ID            uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt     time.Time `json:"-" yaml:"-"`
	UpdatedAt     time.Time `json:"-" yaml:"-"`
	ApplicationID uint      `gorm:"index;not null" sql:"type:bigint REFERENCES applications(id) ON DELETE CASCADE" json:"-"  yaml:"-"`
	Value         string    `gorm:"type:text;"index;not null""`
}

func NewApplication(appConfig *ApplicationConfig) *Application {
	newApp := &Application{}
	newApp.Name = appConfig.Name
	newApp.Path = appConfig.Path
	newApp.BusinessDomain = appConfig.BusinessDomain
	newApp.BusinessValue = appConfig.BusinessValue
	newApp.ScoringModel = appConfig.ScoringModel
	newApp.Files = appConfig.Files
	newApp.IgnoredFiles = appConfig.IgnoredFiles
	newApp.MatchedRules = make(map[string]int)

	return newApp
}

func (app *Application) MergeDetails(details ApplicationDetails) {
	app.Findings = details.Findings
	app.FilesCnt = details.FilesCnt
	app.CIFindings = details.CIFindings
	app.NumCrits = details.NumCrits
	app.RawScore = details.RawScore
	app.RawCloudScore = details.RawCloudScore
	app.RawContainerScore = details.RawContainerScore
	app.SlocCnt = details.SlocCnt
	app.InfoFindings = details.InfoFindings
	app.FindingsRatio = details.FindingsRatio
}

func (app *Application) MergeApp(appToMerge *Application) {
	app.BusinessValue = appToMerge.BusinessValue
	app.BusinessDomain = appToMerge.BusinessDomain
	if app.OriginalScore == -1 {
		app.OriginalScore = app.Score
	}
	app.Score = appToMerge.Score
	app.ScoreModified = !(app.OriginalScore == app.Score)
}

func (app *Application) CalculateScore(scoreModel *ScoringModel) (err error) {

	logrus.Debugf("[%s] Calculating App Score => Raw Score [%d] Sloc [%d]\n",
		app.Name, app.RawScore, app.SlocCnt)

	//Allow for run time override!
	if scoreModel == nil {
		if app.Model == nil {
			return fmt.Errorf("unable to calculate score for app [%s] without scoring model", app.Name)
		}
		scoreModel = app.Model
	}

	outcome := scoreModel.ProcessScore(*app)

	if outcome != nil {
		if outcome.Calculate {
			if app.RawScore <= 0 {
				app.Score = scoreModel.MaxScore
			} else {
				app.Score, err = app.calculate(outcome.Expression, scoreModel)
			}
		} else {
			app.Score = outcome.Score
		}
		app.Recommendation = outcome.Recommendation
	} else {
		app.Score = math.NaN()
		app.Recommendation = "Scoring Failed!"
		return fmt.Errorf("scoring failed! No outcome from scoring model")
	}

	if app.Score > scoreModel.MaxScore {
		app.Score = scoreModel.MaxScore
	} else if app.Score < scoreModel.MinScore {
		app.Score = scoreModel.MinScore
	}

	logrus.Debugf("[%s] Determined Score [%.2f] & Recommendation => %s\n", app.Name, app.Score, app.Recommendation)

	return err
}

func (app *Application) AssociateBin(bin Bin) {
	app.Bins = append(app.Bins, bin)
}

func (app *Application) AssociateTag(tag ApplicationTag) {

	for _, t := range app.Tags {
		if tag.Value == t.Value {
			return
		}
	}

	app.Tags = append(app.Tags, &tag)
}

type ByFileCount []*Application

func (app ByFileCount) Len() int      { return len(app) }
func (app ByFileCount) Swap(i, j int) { app[i], app[j] = app[j], app[i] }
func (app ByFileCount) Less(i, j int) bool {
	f1 := app[i]
	f2 := app[j]
	if f1.FilesCnt == f2.FilesCnt {
		return f1.Name < f2.Name
	}
	return f1.FilesCnt < f2.FilesCnt
}

type ByValue []*ApplicationTag

func (tag ByValue) Len() int      { return len(tag) }
func (tag ByValue) Swap(i, j int) { tag[i], tag[j] = tag[j], tag[i] }
func (tag ByValue) Less(i, j int) bool {
	f1 := tag[i]
	f2 := tag[j]
	return f1.Value < f2.Value
}

func (app *Application) calculate(target string, scoreModel *ScoringModel) (float64, error) {
	parameters := make(map[string]interface{})
	parameters[RAW_SCORE_SCORING_TOKEN] = app.RawScore
	parameters[CLOUD_SCORE_SCORING_TOKEN] = app.RawCloudScore
	parameters[CONTAINER_SCORE_SCORING_TOKEN] = app.RawContainerScore
	parameters[FILES_SCORING_TOKEN] = app.FilesCnt
	parameters[FINDINGS_SCORING_TOKEN] = app.Findings
	parameters[SLOC_CNT_SCORING_TOKEN] = app.SlocCnt
	parameters[BUSINESS_VALUE_SCORING_TOKEN] = app.BusinessValue
	parameters[MAX_SCORE] = scoreModel.MaxScore
	parameters[MIN_SCORE] = scoreModel.MinScore

	expression := util.NewExpression(parameters)

	result, err := expression.Calculate(target)

	if err == nil {
		result, err = strconv.ParseFloat(fmt.Sprintf("%2.2f", result), 64)
	}

	return result, err
}
