/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"github.com/jinzhu/gorm"
	"csa-app/model"
)

type SlocRepository interface {
	GetSlocForRun(runId uint) ([]model.RunSloc, error)
	GetSlocSummaryByApplicationForRun(runId uint) ([]model.SlocByApplication, error)
	CreateSlocData(runSloc *model.RunSloc) error
	GetSummaryFindingsForRun(runid uint) (model.SlocByRun, error)
	GetTopLanguagesByCodeLines(runid uint) ([]model.LanguagesByCodeLines, error)
	GetLanguagesForRunAndApplication(runid uint, application string) ([]model.LanguagesByCodeLines, error)
	GetApplicationsByLanguageForRun(runid uint, language string) ([]model.ApplicationsByLanguageForRun, error)
}

func NewSlocRepository(dbconn *gorm.DB) SlocRepository {
	return &OrmRepository{
		dbconn: dbconn,
	}
}

func NewSlocRepositoryForRun(run *model.Run) SlocRepository {
	return &OrmRepository{
		dbconn: run.DB,
	}
}

func (slocRepository *OrmRepository) GetSlocForRun(runId uint) ([]model.RunSloc, error) {
	var data []model.RunSloc
	res := database.Where(&model.RunSloc{RunID: runId}).Find(&data)
	return data, res.Error
}

func (slocRepository *OrmRepository) GetSlocSummaryByApplicationForRun(runId uint) ([]model.SlocByApplication, error) {
	//SELECT domain as application, sum(code_lines) FROM run_slocs WHERE run_id = 1 GROUP BY application
	var slocByApplication []model.SlocByApplication
	res := slocRepository.dbconn.Model(&model.RunSloc{}).
		Where(&model.RunSloc{RunID: runId}).
		Select("application, sum(code_lines) as code_lines, sum(comment_lines) as comment_lines, sum(blank_lines) as blank_lines, sum(total_files) total_files").
		Group("application").
		Order("application asc").
		Scan(&slocByApplication)
	return slocByApplication, res.Error
}

func (slocRepository *OrmRepository) CreateSlocData(runSloc *model.RunSloc) error {
	createRes := database.Create(runSloc)
	return createRes.Error
}

func (slocRepository *OrmRepository) GetSummaryFindingsForRun(runId uint) (model.SlocByRun, error) {
	var slocByRun model.SlocByRun
	res := slocRepository.dbconn.Model(&model.RunSloc{}).
		Where(&model.RunSloc{RunID: runId}).
		Select("count(distinct application) as apps, sum(code_lines) as code_lines, sum(comment_lines) as comment_lines, sum(blank_lines) as blank_lines, sum(total_files) total_files").
		Scan(&slocByRun)
	slocByRun.RunId = runId

	return slocByRun, res.Error
}

func (slocRepository *OrmRepository) GetTopLanguagesByCodeLines(runid uint) ([]model.LanguagesByCodeLines, error) {
	var languagesByCodeLines []model.LanguagesByCodeLines

	res := slocRepository.dbconn.Model(&model.RunSloc{}).
		Where(&model.RunSloc{RunID: runid}).
		Select("lang, sum(code_lines) code_lines").
		Order("code_lines desc").
		Group("lang").
		Scan(&languagesByCodeLines)

	return languagesByCodeLines, res.Error
}

func (slocRepository *OrmRepository) GetApplicationsByLanguageForRun(runid uint, language string) ([]model.ApplicationsByLanguageForRun, error) {
	var applicationsByLanguageForRun []model.ApplicationsByLanguageForRun

	res := slocRepository.dbconn.Model(&model.RunSloc{}).
		Where(&model.RunSloc{RunID: runid, Lang: language}).
		Select("application, sum(code_lines) as code_lines").
		Order("code_lines desc").
		Group("application").
		Scan(&applicationsByLanguageForRun)

	return applicationsByLanguageForRun, res.Error

}

func (slocRepository *OrmRepository) GetLanguagesForRunAndApplication(runid uint, application string) ([]model.LanguagesByCodeLines, error) {
	var languagesForRunAndApplication []model.LanguagesByCodeLines

	res := slocRepository.dbconn.Model(&model.RunSloc{}).
		Where(&model.RunSloc{RunID: runid, Application: application}).
		Select("lang, sum(code_lines) code_lines").
		Order("code_lines desc").
		Group("lang").
		Scan(&languagesForRunAndApplication)

	return languagesForRunAndApplication, res.Error
}
