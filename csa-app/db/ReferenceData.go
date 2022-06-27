/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"csa-app/model"
	"csa-app/util"
)

//"CREATE TABLE IF NOT EXISTS REPORT_REF (report_id INTEGER, type varchar(20), description varchar(100), target varchar(20))"

func PopulateReferenceData() {

	newReport := model.ReportRef{Type: util.APP_NAME, ReportNum: model.THIRD_PARTY_REPORT_ID, Title: model.THIRD_PARTY_IMPORTS, Summary: model.THIRD_PARTY_IMPORTS_DESC, Extension: model.CSV_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: util.APP_NAME, ReportNum: model.API_SUMMARY_REPORT_ID, Title: model.API_SUMMARY, Summary: model.API_SUMMARY_DESC, Extension: model.CSV_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: util.APP_NAME, ReportNum: model.API_DETAILED_REPORT_ID, Title: model.API_DETAILED, Summary: model.API_DETAILED_DESC, Extension: model.CSV_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: util.APP_NAME, ReportNum: model.ANNOTATIONS_REPORT_ID, Title: model.ANNOTATIONS, Summary: model.ANNOTATIONS_DESC, Extension: model.CSV_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: util.APP_NAME, ReportNum: model.CLOC_REPORT_ID, Title: model.CLOC, Summary: model.CLOC_DESC, Extension: model.CSV_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: "git", ReportNum: model.GIT_FORENSICS_REPORT_ID, Title: model.GIT_FORENSICS, Summary: model.GIT_FORENSICS_DESC, Extension: model.TXT_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: "git", ReportNum: model.GIT_ACTIVITY_DETAILED_REPORT_ID, Title: model.GIT_ACTIVITY_DETAILED, Summary: model.GIT_ACTIVITY_DETAILED_DESC, Extension: model.TXT_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

	newReport = model.ReportRef{Type: "git", ReportNum: model.GIT_ACTIVITY_SUMMARY_REPORT_ID, Title: model.GIT_ACTIVITY_SUMMARY, Summary: model.GIT_ACTIVITY_SUMMARY_DESC, Extension: model.TXT_EXTENSION}
	database.Create(&newReport)
	CheckDBForError(true, "PopulateReferenceData", "Error populating Report Reference Data!")

}

func GetAvailableReportsByType(reportType string) []model.ReportRef {
	var availReports []model.ReportRef
	database.Where(&model.ReportRef{Type: reportType}).Find(&availReports)
	model.SortReports(model.ByReportNum).Sort(availReports)
	return availReports
}

func GetAvailableReportById(reportType string, reportId int) model.ReportRef {
	var report model.ReportRef
	database.Where(&model.ReportRef{Type: reportType, ReportNum: reportId}).Find(&report)
	return report
}

func GetReportTypes() []string {
	var availReports []model.ReportRef
	database.Select("DISTINCT(type)").Find(&availReports)

	var reportTypes []string

	for _, report := range availReports {
		reportTypes = append(reportTypes, report.Type)
	}

	return reportTypes
}

func GetAvailableReports() map[string][]model.ReportRef {

	rTypes := GetReportTypes()

	reportMap := make(map[string][]model.ReportRef)

	for _, rType := range rTypes {

		reportMap[rType] = GetAvailableReportsByType(rType)
	}

	return reportMap
}

func PopulateReportHeaders() {

	//THIRD PARTY IMPORTS
	newHeaderColumn := model.ReportHeader{ReportID: model.THIRD_PARTY_REPORT_ID, Name: model.THIRD_PARTY_HEADER, Position: 1}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for ThirdParty Report!")

	//API SUMMARY
	newHeaderColumn = model.ReportHeader{ReportID: model.API_SUMMARY_REPORT_ID, Name: model.API_SUMMARY_API_HEADER, Position: 1}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Summary Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_SUMMARY_REPORT_ID, Name: model.API_SUMMARY_API_COUNT_HEADER, Position: 2}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Summary Report!")

	//API DETAILED
	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_DOMAIN_HEADER, Position: 1}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_API_HEADER, Position: 2}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_PATTERN_HEADER, Position: model.API_DETAILED_REPORT_ID}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_FILE_HEADER, Position: 4}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_LINE_NO_HEADER, Position: 5}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_SOURCE_HEADER, Position: model.CLOC_REPORT_ID}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_SCORE_HEADER, Position: 7}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.API_DETAILED_REPORT_ID, Name: model.API_DETAILED_ADVICE_HEADER, Position: 8}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for API-Details Report!")

	//ANNOTATIONS
	newHeaderColumn = model.ReportHeader{ReportID: model.ANNOTATIONS_REPORT_ID, Name: model.ANNOTATIONS_HEADER, Position: 1}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for Annotations Report!")

	//CLOC
	newHeaderColumn = model.ReportHeader{ReportID: model.CLOC_REPORT_ID, Name: model.CLOC_LANG_HEADER, Position: 1}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for SLOC Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.CLOC_REPORT_ID, Name: model.CLOC_FILES_HEADER, Position: 2}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for SLOC Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.CLOC_REPORT_ID, Name: model.CLOC_BLANK_LINES_HEADER, Position: model.API_DETAILED_REPORT_ID}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for SLOC Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.CLOC_REPORT_ID, Name: model.CLOC_COMMENT_LINES_HEADER, Position: 4}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for SLOC Report!")

	newHeaderColumn = model.ReportHeader{ReportID: model.CLOC_REPORT_ID, Name: model.CLOC_CODE_LINES_HEADER, Position: 5}
	database.Save(&newHeaderColumn)
	CheckDBForError(true, "PopulateReportHeaders", "Error populating Report Headers for SLOC Report!")

}

func GetHeadersForReport(reportId int) []model.ReportHeader {
	var headers []model.ReportHeader
	database.Where(&model.ReportHeader{ReportID: reportId}).Find(&headers)
	model.SortHeadersBy(model.Position).Sort(headers)
	return headers
}

func GetAllHeaders() []model.ReportHeader {
	var headers []model.ReportHeader
	database.Find(&headers)
	model.SortHeadersBy(model.ReportAndPosition).Sort(headers)
	return headers
}
