/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"database/sql"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"csa-app/model"
	"csa-app/util"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type FindingRepository interface {
	SaveFinding(finding *model.Finding) error
	GetFindingsWithTags(runId uint) ([]model.Finding, error)
	GetFinding(id uint) (model.Finding, error)
	GetFindings(runid uint) ([]model.Finding, error)
	GetFindingsWithoutPreload(runid uint) ([]model.Finding, error)
	GetFindingsByTag(runid uint, tag string) ([]model.Finding, error)
	GetFindingsByRule(runid uint, rulename string) ([]model.Finding, error)
	GetApplicationDetailsForRun(runid uint, normalMax int, normalMin int, reverse bool) ([]model.ApplicationDetails, error)
	GetTopApisByScoreForRun(runid uint) ([]model.ApiUsage, error)
	GetApisByUsageForRunAndApplication(runid uint, application string) ([]model.ApiUsage, error)
	GetApplicationsUsingApiForRun(runid uint, api string) ([]model.ApplicationsForRun, error)
	GetScoreCards(runId uint) ([]model.AppScoreCard, error)
	GetScoreCard(runId uint, app string, tags []string, includeFF bool) (model.AppScoreCard, error)
	GetScoreCardDetails(runId uint, app string, card string) ([]model.ScoreCardDetail, error)
	GetFindingsDTOForRun(runid uint) ([]*model.FindingDTO, error)
	GetFindingsDTOForRunAppLevel(runId uint, app string, card string, tags []string, includeFF bool) ([]*model.FindingDTO, error)
	GetTagsForApp(runId uint, app string) ([]string, error)
	UpdateFindingsForRecalculate(run *model.Run, rulesToApply []model.Rule) []error
}

func NewFindingRepository(db *gorm.DB) FindingRepository {
	return &OrmRepository{
		dbconn: db,
	}
}

func NewFindingRepositoryForRun(run *model.Run) FindingRepository {
	return &OrmRepository{
		dbconn: run.DB,
	}
}

func (findingRepository *OrmRepository) UpdateFindingsForRecalculate(run *model.Run, rulesToApply []model.Rule) []error {
	errors := []error{}
	ruleMap := map[string]*model.Rule{}
	for r := range rulesToApply {
		rule := &rulesToApply[r]
		ruleMap[rule.Name] = rule
	}

	appDetailsMap := map[string]model.ApplicationDetails{}
	findings, _ := findingRepository.GetFindings(run.ID)

	// apply new effort scores to any findings with rule changes
	for f := range findings {
		finding := findings[f]
		_, ok := ruleMap[finding.Rule]
		if ok && finding.Effort != ruleMap[finding.Rule].Effort {
			finding.Effort = ruleMap[finding.Rule].Effort
		}

		// keep track of the sum of effort per app as it will be the new raw score
		app, hasKey := appDetailsMap[finding.Application]
		if !hasKey {
			appDetailsMap[finding.Application] = model.ApplicationDetails{
				RawScore: int64(finding.Effort),
				Findings: 1,
			}
		} else {
			app = appDetailsMap[finding.Application]
			app.RawScore += int64(finding.Effort)
			app.Findings++
			appDetailsMap[finding.Application] = app
		}

		err := findingRepository.dbconn.Save(&finding).Error
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) == 0 {
		// update the raw scores used in calculation
		for a := range run.Applications {
			_, exists := appDetailsMap[run.Applications[a].Name]
			if exists {
				app := run.Applications[a]
				app.RawScore = appDetailsMap[app.Name].RawScore
				app.Findings = appDetailsMap[app.Name].Findings
				run.Applications[a] = app

				info := &[]*util.FileInfo{}
				for i := 1; i <= app.FilesCnt; i++ {
					*info = append(*info, &util.FileInfo{})
				}
				app.Files = *info
				run.Applications[a] = app
			}
		}
		err := findingRepository.dbconn.Save(&run).Error
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func (findingRepository *OrmRepository) SaveFinding(finding *model.Finding) error {

	if *util.Efd {
		finding.SetValue("---")
	}

	result := findingRepository.dbconn.Create(finding)

	err := result.Error

	if err != nil {
		log.Errorf("Failed saving %v findings to persistence!", err)
		return err
	}
	return nil
}

func (findingRepository *OrmRepository) GetFinding(id uint) (model.Finding, error) {
	finding := model.Finding{}
	res := findingRepository.dbconn.Where(model.Finding{ID: id}).Preload("Tags").Preload("Recipes").Find(&finding)
	return finding, res.Error
}

func (findingRepository *OrmRepository) GetFindingsWithTags(runId uint) ([]model.Finding, error) {
	findings := []model.Finding{}
	res := findingRepository.dbconn.Where("run_id = ?", runId).Preload("Tags").Find(&findings)
	return findings, res.Error
}

func (findingRepository *OrmRepository) GetFindings(runid uint) ([]model.Finding, error) {
	findings := []model.Finding{}
	res := findingRepository.dbconn.Where(&model.Finding{RunID: runid}).Preload("Recipes").Preload("Tags").Find(&findings)
	return findings, res.Error
}

func (findingRepository *OrmRepository) GetFindingsWithoutPreload(runid uint) ([]model.Finding, error) {
	findings := []model.Finding{}
	res := findingRepository.dbconn.Where(&model.Finding{RunID: runid}).Find(&findings)
	return findings, res.Error
}

func (findingRepository *OrmRepository) GetFindingsDTOForRun(runid uint) (findings []*model.FindingDTO, err error) {

	/*
	 This is required because when trying to get more than a few thousand rows and preload tags/recipes I would
	 receive errors from Gorm or SQllite of 'to many SQl variables'. So, I had to hand roll this.
	*/

	rows, err := findingRepository.dbconn.Table("findings").
		Select("findings.id, findings.run_id, findings.filename, findings.fqn, findings.ext, findings.line, findings.rule, "+
			"findings.pattern, findings.value, findings.advice, findings.effort,findings.container_effort, findings.cloud_native_effort, findings.readiness, findings.category, "+
			"findings.criticality, findings.application, finding_tags.value as tag, finding_recipes.uri as recipe_uri").
		Joins("left join finding_tags on findings.id = finding_tags.finding_id left join finding_recipes on findings.id = finding_recipes.finding_id").
		Where("run_id = ?", runid).Order("findings.id asc").Rows()

	lastFinding := &model.FindingDTO{ID: 0}
	var tagList []string
	var rcpList []string

	if err != nil {
		log.Errorf("Error retrieving findingDTOs! Details: %v", err)
		return
	}

	for rows.Next() {
		var id, run uint
		var filename, fqn, ext, rule, pattern, value, advice, cat, crit, app, tag, recipe string
		var line, effort, readiness int
		var tagExists, rcpExists bool
		var cloud_native_effort int
		var container_effort int
		rows.Scan(&id, &run, &filename, &fqn, &ext, &line, &rule,
			&pattern, &value, &advice, &effort, &container_effort, &cloud_native_effort,
			&readiness, &cat, &crit, &app, &tag, &recipe)

		if lastFinding.ID == id {
			if tag != "" {
				for idx := range tagList {
					if tagList[idx] == tag {
						tagExists = true
						break
					}
				}

				if !tagExists {
					tagList = append(tagList, tag)
					lastFinding.AddTag(tag)
				}
			}

			if recipe != "" {
				for idx := range rcpList {
					if rcpList[idx] == recipe {
						rcpExists = true
						break
					}
				}

				if !rcpExists {
					rcpList = append(rcpList, recipe)
					lastFinding.AddRecipe(recipe)
				}
			}
		} else {

			tagList = nil
			rcpList = nil
			//new finding
			newFinding := &model.FindingDTO{ID: id, RunID: run, Filename: filename,
				Fqn: fqn, Ext: ext, Rule: rule,
				Pattern: pattern, Value: value, Line: line, Category: cat,
				Effort: effort, CloudNativeEffort: cloud_native_effort, ContainerEffort: container_effort,
				Readiness: readiness, Advice: advice, Application: app}
			findings = append(findings, newFinding)
			lastFinding = newFinding

			if tag != "" {
				tagList = append(tagList, tag)
				newFinding.AddTag(tag)
			}
			if recipe != "" {
				rcpList = append(rcpList, recipe)
				newFinding.AddRecipe(recipe)
			}
		}
	}

	sort.Sort(ByAppAndFileName(findings))
	return
}

func (findingRepository *OrmRepository) GetFindingsDTOForRunAppLevel(runId uint, app string, card string, tags []string, includeFF bool) (findings []*model.FindingDTO, err error) {

	/*
	 This is required because when trying to get more than a few thousand rows and preload tags/recipes I would
	 receive errors from Gorm or SQllite of 'to many SQl variables'. So, I had to hand roll this.
	*/
	bottom := math.MaxInt64
	top := math.MinInt64
	for _, level := range strings.Split(card, ",") {
		b, t := getEffortBand(level)
		if t > top {
			top = t
		}
		if b < bottom {
			bottom = b
		}
	}

	whereClause := "findings.run_id = ? and findings.application = ? and effort >= ? and effort <= ? and findings.pattern != 'Lines of Code' and findings.pattern != 'Analyzed File'"

	if !includeFF {
		whereClause += " and findings.category !='" + model.FILE_ANALYZED_CATEGORY + "'"
	}

	var rows *sql.Rows

	if len(tags) > 0 {

		rows, err = findingRepository.dbconn.Table("findings").
			Select(
				"findings.id, findings.run_id, findings.filename, findings.fqn, findings.ext, findings.line, "+
					"findings.rule, findings.pattern, findings.value, findings.advice, "+levelCaseFragment()+
					"findings.effort, findings.cloud_native_effort, findings.container_effort, findings.readiness, findings.note,"+
					"findings.category, findings.criticality, findings.application, "+
					"finding_tags.value as tag, finding_recipes.uri as recipe_uri").
			Joins("left join finding_tags on findings.id = finding_tags.finding_id "+
				"left join finding_recipes on findings.id = finding_recipes.finding_id").
			Where(whereClause+" and finding_tags.value in (?)", runId, app, bottom, top, tags).
			Order("findings.fqn, findings.line").
			Rows()

	} else {

		rows, err = findingRepository.dbconn.Table("findings").
			Select(
				"findings.id, findings.run_id, findings.filename, findings.fqn, findings.ext, findings.line, "+
					"findings.rule, findings.pattern, findings.value, findings.advice, "+levelCaseFragment()+
					"findings.effort, findings.cloud_native_effort, findings.container_effort, findings.readiness, findings.note, "+
					"findings.category, findings.criticality, findings.application, "+
					"finding_tags.value as tag, finding_recipes.uri as recipe_uri").
			Joins("left join finding_tags on findings.id = finding_tags.finding_id "+
				"left join finding_recipes on findings.id = finding_recipes.finding_id").
			Where(whereClause, runId, app, bottom, top).
			Order("findings.fqn, findings.line").
			Rows()

	}
	lastFinding := &model.FindingDTO{ID: 0}
	var tagList []string
	var rcpList []string

	if err != nil {
		log.Errorf("Error retrieving findingDTOs! Details: %v", err)
		return
	}

	for rows.Next() {

		var id, run uint
		var filename, fqn, ext, rule, pattern, value, advice, cat, crit, note, app, tag, recipe, level string
		var line, effort, readiness int
		var tagExists, rcpExists bool
		var cloud_native_effort, container_effort int

		rows.Scan(&id, &run, &filename, &fqn, &ext, &line,
			&rule, &pattern, &value, &advice, &level,
			&effort, &cloud_native_effort, &container_effort, &readiness, &note,
			&cat, &crit, &app, &tag, &recipe)

		if lastFinding.ID == id {
			if tag != "" {
				for idx := range tagList {
					if tagList[idx] == tag {
						tagExists = true
						break
					}
				}

				if !tagExists {
					tagList = append(tagList, tag)
					lastFinding.AddTag(tag)
				}
			}

			if recipe != "" {
				for idx := range rcpList {
					if rcpList[idx] == recipe {
						rcpExists = true
						break
					}
				}

				if !rcpExists {
					rcpList = append(rcpList, recipe)
					lastFinding.AddRecipe(recipe)
				}
			}
		} else {

			tagList = nil
			rcpList = nil

			//new finding
			newFinding := &model.FindingDTO{
				ID: id, RunID: run, Filename: filename, Fqn: fqn, Ext: ext, Rule: rule,
				Pattern: pattern, Value: value, Line: line, Category: cat, Level: level,
				CloudNativeEffort: cloud_native_effort, ContainerEffort: container_effort,
				Effort: effort, Readiness: readiness, Note: note, Advice: advice, Application: app,
			}

			findings = append(findings, newFinding)
			lastFinding = newFinding

			if tag != "" {
				tagList = append(tagList, tag)
				newFinding.AddTag(tag)
			}
			if recipe != "" {
				rcpList = append(rcpList, recipe)
				newFinding.AddRecipe(recipe)
			}
		}
	}

	sort.Sort(ByAppAndFileName(findings))
	return
}

func (findingRepository *OrmRepository) GetFindingsByTag(runid uint, tag string) ([]model.Finding, error) {

	findings := []model.Finding{}

	res := findingRepository.dbconn.Table("findings").
		Joins("inner join finding_tags on finding_tags.finding_id = findings.id").
		Where("findings.run_id = ? and finding_tags.value = ?", runid, tag).
		Find(&findings)

	return findings, res.Error
}

func (findingRepository *OrmRepository) GetFindingsByRule(runid uint, rulename string) ([]model.Finding, error) {

	findings := []model.Finding{}

	res := findingRepository.dbconn.Table("findings").
		Where(&model.Finding{RunID: runid, Rule: rulename}).
		Order("application asc").
		Find(&findings)

	return findings, res.Error
}

func (findingRepository *OrmRepository) GetFindingsByFile(runid uint, file string) ([]model.Finding, error) {

	findings := []model.Finding{}

	res := findingRepository.dbconn.Table("findings").
		Where(&model.Finding{RunID: runid, Filename: file}).
		Order("application asc").
		Find(&findings)

	return findings, res.Error
}

func (findingRepository *OrmRepository) GetFindingsByFileLike(runid uint, file string) ([]model.Finding, error) {

	findings := []model.Finding{}

	res := findingRepository.dbconn.Table("findings").
		Where("run_id = ? and filename like ?", runid, file).
		Order("application asc").
		Find(&findings)

	return findings, res.Error
}

func (findingRepository *OrmRepository) GetApplicationDetailsForRun(runid uint, normalMax int, normalMin int, reverse bool) ([]model.ApplicationDetails, error) {

	var applicationScores []model.ApplicationDetails

	res := findingRepository.dbconn.Model(&model.Finding{}).
		Where(&model.Finding{RunID: runid}).
		Select("application, count(*) as findings, sum(effort) as raw_score, sum(cloud_native_effort) as raw_cloud_score, sum(container_effort) as raw_container_score").Group("application").
		Order("raw_score desc, application asc").
		Scan(&applicationScores)

	addSlocCnt(findingRepository, runid, applicationScores)

	err := res.Error

	if err == nil {

		rows, err := findingRepository.dbconn.Model(&model.Finding{}).
			Select("application, count(*) as ciFindings").Where("run_id = ? and rule <> ''", runid).Group("application").Rows()

		if err == nil {
			for rows.Next() {
				var ciFindings int
				var app string
				err = rows.Scan(&app, &ciFindings)
				if err != nil {
					return applicationScores, err
				}
				log.Debugf("Found app [%s] with [%d] Cloud Readiness Impacting findings", app, ciFindings)
				updateFindings(applicationScores, app, ciFindings)
			}

			rows, err := findingRepository.dbconn.Model(&model.Finding{}).
				Select("application, count(*) as crits").Where("run_id = ? and effort > ?", runid, model.CRIT_SCORE_THRESHOLD).Group("application").Rows()

			if err == nil {
				for rows.Next() {
					var crits int
					var app string
					err = rows.Scan(&app, &crits)
					if err != nil {
						return applicationScores, err
					}
					log.Debugf("Found app [%s] with [%d] crit findings", app, crits)
					updateCritCount(applicationScores, app, crits)
				}
			}
		}

	}

	return applicationScores, err
}

func (findingRepository *OrmRepository) GetTopApisByScoreForRun(runid uint) ([]model.ApiUsage, error) {

	topApisByRun := []model.ApiUsage{}

	res := findingRepository.dbconn.Model(&model.Finding{}).
		Joins("JOIN finding_tags on findings.id=finding_tags.finding_id").
		Select("category as api, count(1) as usage_count").
		Where("run_id=?", runid).
		Where("finding_tags.value=?", "api").
		Group("api").
		Order("usage_count desc").
		Scan(&topApisByRun)

	return topApisByRun, res.Error
}

func (findingRepository *OrmRepository) GetApisByUsageForRunAndApplication(runid uint, application string) ([]model.ApiUsage, error) {
	apisForRunAndApplication := []model.ApiUsage{}

	res := findingRepository.dbconn.Model(&model.Finding{}).
		Joins("JOIN finding_tags on findings.id=finding_tags.finding_id").
		Select("category as api, count(1) as usage_count").
		Where("run_id=? and application=?", runid, application).
		Where("finding_tags.value=?", "api").
		Group("api").
		Order("usage_count desc").
		Scan(&apisForRunAndApplication)

	return apisForRunAndApplication, res.Error
}

func (findingRepository *OrmRepository) GetApplicationsUsingApiForRun(runid uint, api string) ([]model.ApplicationsForRun, error) {

	applicationsByApiForRun := []model.ApplicationsForRun{}
	res := findingRepository.dbconn.Model(&model.Finding{}).
		Joins("JOIN finding_tags on findings.id=finding_tags.finding_id").
		Select("application, count(1) as usage_count").
		Where(&model.Finding{RunID: runid, Category: api}).
		Where("finding_tags.value=?", "api").
		Group("application").
		Order("usage_count desc").
		Scan(&applicationsByApiForRun)

	return applicationsByApiForRun, res.Error
}

func (findingRepository *OrmRepository) GetScoreCards(runid uint) ([]model.AppScoreCard, error) {

	scorecards := []model.AppScoreCard{}

	res := findingRepository.dbconn.Model(&model.Finding{}).
		Select("application, count(*) total").
		Where(&model.Finding{RunID: runid}).
		Group("application").
		Order("total desc").
		Scan(&scorecards)

	err := res.Error

	if err == nil {
		//My preference would be to use the criticality attribute but for now using score bins
		err = addFindingsByCriticality(findingRepository, model.Info_criticality, runid, model.Info_criticality_low_score, model.Info_criticality_high_score, scorecards)
		if err == nil {
			err = addFindingsByCriticality(findingRepository, model.Low_criticality, runid, model.Low_criticality_low_score, model.Low_criticality_high_score, scorecards)
			if err == nil {
				err = addFindingsByCriticality(findingRepository, model.Medium_criticality, runid, model.Medium_criticality_low_score, model.Medium_criticality_high_score, scorecards)
				if err == nil {
					addFindingsByCriticality(findingRepository, model.High_criticality, runid, model.High_criticality_low_score, model.High_criticality_high_score, scorecards)
				}
			}
		}
	}
	return scorecards, err
}

func (findingRepository *OrmRepository) GetScoreCard(runid uint, app string, tags []string, includeFF bool) (model.AppScoreCard, error) {

	scorecard := model.AppScoreCard{}
	var res *gorm.DB
	findings := []model.Finding{}

	whereClause := "findings.run_id = ? and findings.application = ?"

	if !includeFF {
		whereClause += " and findings.category !='" + model.FILE_ANALYZED_CATEGORY + "'"
	}

	if len(tags) > 0 {
		res = findingRepository.dbconn.Table("findings").Where("findings.id in ?",
			findingRepository.dbconn.Table("findings").
				Select("distinct(findings.id)").
				Joins("inner join finding_tags on finding_tags.finding_id = findings.id").
				Where(whereClause+" and  finding_tags.value in (?)", runid, app, tags).SubQuery()).Find(&findings)
	} else {
		res = findingRepository.dbconn.Table("findings").
			Where(whereClause, runid, app).
			Find(&findings)
	}

	if res.Error == nil {

		for _, finding := range findings {

			//scorecard.Total++

			if finding.Effort == model.Info_criticality_high_score {
				scorecard.Info++
				continue
			}
			if finding.Effort <= model.Low_criticality_high_score {
				scorecard.Low++
				scorecard.Total++
				continue
			}
			if finding.Effort <= model.Medium_criticality_high_score {
				scorecard.Medium++
				scorecard.Total++
				continue
			}
			if finding.Effort <= model.High_criticality_high_score {
				scorecard.High++
				scorecard.Total++
				continue
			}

			//scorecard.High++
		}
	}

	return scorecard, res.Error
}

func (findingRepository *OrmRepository) GetFindingsByScoreRange(runid uint, bottom int, top int) ([]model.Finding, error) {
	findings := []model.Finding{}
	whereClause := "run_id = ? and effort >= ? and effort <= ?"
	findingRepository.dbconn.Where(whereClause, runid, bottom, top).Find(&findings)
	return findings, findingRepository.dbconn.Error
}

func (findingRepository *OrmRepository) GetScoreCardDetails(runId uint, app string, card string) ([]model.ScoreCardDetail, error) {

	bottom, top := getEffortBand(card)

	details := []model.ScoreCardDetail{}

	whereClause := "run_id = ? and application =  ? and effort >= ? and effort <= ?"

	selectClause := "application, category, pattern, effort, cloud_native_effort, container_effort " + levelCaseFragment() +
		"count(*) as count, effort * count(*) as total"

	res := findingRepository.dbconn.Model(&model.Finding{}).
		Select(selectClause).Where(whereClause, runId, app, bottom, top).
		Group("pattern, level").
		Order("level asc, pattern asc, total desc").
		Scan(&details)

	return details, res.Error
}

func (findingRepository *OrmRepository) GetTagsForApp(runId uint, app string) ([]string, error) {

	start := time.Now()
	var tags []string

	rows, err := findingRepository.dbconn.Table("findings").
		Select("distinct(finding_tags.value)").
		Joins("inner join finding_tags on finding_tags.finding_id = findings.id").
		Where("findings.run_id = ? and findings.application = ?", runId, app).
		Rows()

	if err == nil {
		for rows.Next() {
			var tag string
			rows.Scan(&tag)
			tags = append(tags, tag)
		}
	}

	fmt.Printf("Retrieved [%d] tags for App [%s] in [%s]\n", len(tags), app, time.Since(start))

	return tags, err
}

/*
 PRIVATE API ------------------------------------------------------------------------------------------------------------
*/

func updateFindings(scores []model.ApplicationDetails, app string, ciFindings int) {
	for idx := range scores {
		if scores[idx].Application == app {
			scores[idx].CIFindings = ciFindings
			scores[idx].InfoFindings = scores[idx].Findings - ciFindings
			if ciFindings > 0 {
				scores[idx].FindingsRatio = float64(scores[idx].RawScore) / float64(ciFindings)
			}
			return
		}
	}
}

func updateCritCount(scores []model.ApplicationDetails, app string, crits int) {
	for idx := range scores {
		if scores[idx].Application == app {
			scores[idx].NumCrits = crits
			return
		}
	}
}

// TODO Pull this stuff up into the scoring service and orchestrate accross the repos!
func addSlocCnt(findingRepository *OrmRepository, runId uint, scores []model.ApplicationDetails) {

	var slocByApplication []model.SlocByApplication
	res := findingRepository.dbconn.Model(&model.RunSloc{}).
		Where(&model.RunSloc{RunID: runId}).
		Select("application, sum(code_lines) as code_lines, sum(comment_lines) as comment_lines, sum(blank_lines) as blank_lines, sum(total_files) total_files").
		Group("application").
		Order("application asc").
		Scan(&slocByApplication)

	if res.Error == nil {
		for idx := range scores {
			for _, slocApp := range slocByApplication {
				if scores[idx].Application == slocApp.Application {
					scores[idx].SlocCnt = slocApp.CodeLines
					scores[idx].FilesCnt = slocApp.TotalFiles
					break
				}
			}
		}
	}
}

// For now this is setup to update an existing set of scorecards with additional criticality details
func addFindingsByCriticality(findingRepository *OrmRepository, criticality string, runId uint, bottomScore int, topScore int, cards []model.AppScoreCard) error {

	whereClause := "run_id = ? and effort >= ? and effort <= ?"
	rows, err := findingRepository.dbconn.Model(&model.Finding{}).
		Select("application, count(*) cnt").Where(whereClause, runId, bottomScore, topScore).Group("application").Rows()

	if err == nil {
		for rows.Next() {
			var count int
			var app string
			rows.Scan(&app, &count)
			updateCard(cards, app, count, criticality)
		}
	}

	return err
}

func updateCard(cards []model.AppScoreCard, app string, count int, criticality string) {

	for idx := range cards {
		if cards[idx].Application == app {
			switch criticality {
			case model.Low_criticality:
				cards[idx].Low = count
			case model.Medium_criticality:
				cards[idx].Medium = count
			case model.High_criticality:
				cards[idx].High = count
			default:
				cards[idx].Info = count
			}
			return
		}
	}
}

type ByAppAndFileName []*model.FindingDTO

func (dtos ByAppAndFileName) Len() int      { return len(dtos) }
func (dtos ByAppAndFileName) Swap(i, j int) { dtos[i], dtos[j] = dtos[j], dtos[i] }
func (dtos ByAppAndFileName) Less(i, j int) bool {
	r1 := dtos[i]
	r2 := dtos[j]

	if r1.Application == r2.Application {
		if r1.Filename == r2.Filename {
			return r1.Line < r2.Line
		} else {
			return r1.Filename < r2.Filename
		}
	} else {
		return r1.Application < r2.Application
	}

}

func getEffortBand(card string) (bottom int, top int) {
	//All - default case
	bottom = model.All_criticality_low_score
	top = model.All_criticality_high_score

	switch strings.ToLower(card) {
	case model.Info_criticality:
		bottom = model.Info_criticality_low_score
		top = model.Info_criticality_high_score
	case model.Low_criticality:
		bottom = model.Low_criticality_low_score
		top = model.Low_criticality_high_score
	case model.Medium_criticality:
		bottom = model.Medium_criticality_low_score
		top = model.Medium_criticality_high_score
	case model.High_criticality:
		bottom = model.High_criticality_low_score
		top = model.High_criticality_high_score
	}

	return
}

func levelCaseFragment() string {

	var level = "level"

	if util.POSTGRES == *util.DB {
		level = "as level"
	}

	return fmt.Sprintf(
		"case "+
			"when effort between %d and %d then '%s' "+
			"when effort between %d and %d then '%s' "+
			"when effort between %d and %d then '%s' "+
			"when effort between %d and %d then '%s' "+
			"end %s, ",
		model.Info_criticality_low_score,
		model.Info_criticality_high_score,
		model.Info_criticality,
		model.Low_criticality_low_score,
		model.Low_criticality_high_score,
		model.Low_criticality,
		model.Medium_criticality_low_score,
		model.Medium_criticality_high_score,
		model.Medium_criticality,
		model.High_criticality_low_score,
		model.High_criticality_high_score,
		model.High_criticality,
		level)
}
