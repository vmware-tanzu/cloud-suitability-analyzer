/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"csa-app/model"
	"github.com/jinzhu/gorm"
)

func GetReportData(runId uint, reportId int) []model.ReportData {
	var data []model.ReportData
	database.Where(&model.ReportData{RunID: runId, ReportID: reportId}).Find(&data)
	return data
}

func SaveFinding(finding *model.Finding) bool {
	err := database.Save(finding).Error
	return !CheckDBError(true, "SaveFinding", "DB Error!", err)
}

func SaveFindingTransacted(db *gorm.DB, finding *model.Finding) bool {
	err := db.Save(finding).Error
	return !CheckDBError(false, "SaveFinding", "DB Error!", err)
}

func SaveMetrics(metrics []model.KV) {
	for _, metric := range metrics {
		SaveMetric(metric)
	}
}

func SaveMetric(metric model.KV) {
	metric.Value.PrePersist()

	err := database.Save(&metric.Value).Error
	CheckDBError(false, "SaveMetric", "DB Error!", err)
}

func GetFindingsByRun(id uint) []model.Finding {
	var findings []model.Finding

	CheckDBError(false,
		"GetFindsingByRun",
		"",
		database.Where(&model.Finding{RunID: id}).Preload("Recipes").Preload("Tags").Find(&findings).Error)

	return findings
}

func GetFindingsByRunAndTag(id uint, tag string) []model.Finding {
	var findings []model.Finding
	CheckDBError(false,
		"GetFindsingByRunAndTag",
		"",
		database.Joins("JOIN finding_tags on finding_tags.finding_id = findings.ID").Where("finding_tags.value = ?", tag).Where(&model.Finding{RunID: id}).Find(&findings).Error)

	return findings
}

func LoadTags(finding *model.Finding) {
	var tags []model.FindingTag
	database.Where(model.FindingTag{FindingID: finding.ID}).Find(&tags)
	finding.Tags = tags
}

func LoadRecipes(finding *model.Finding) {
	var recipes []model.FindingRecipe
	database.Where(model.FindingTag{FindingID: finding.ID}).Find(&recipes)
	finding.Recipes = recipes
}

func UniqueFinding(findings []model.Finding) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range findings {
		if _, value := keys[entry.Value]; !value {
			keys[entry.Value] = true
			list = append(list, entry.Value)
		}
	}
	return list
}
