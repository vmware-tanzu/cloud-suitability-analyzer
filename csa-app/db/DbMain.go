/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"csa-app/model"
	"csa-app/util"
)

const sqllite_driver string = "sqlite3"
const postgres_driver string = "postgres"

type Repositories struct {
	Rules    RuleRepository
	Findings FindingRepository
	Run      RunRepository
	Sloc     SlocRepository
	Reports  ReportDataRepository
	Bins     BinRepository
	Scoring  ScoringRepository
}

type OrmRepository struct {
	dbconn   *gorm.DB
	fileUtil *util.FileUtil
}

var (
	//Used for reads...
	database        *gorm.DB
	dbConnectString string
	driver          string
)

func OpenDB(run *model.Run) *gorm.DB {

	setConnectionString(run)

	if *util.Verbose {
		fmt.Printf("Connecting to %s/%s using %s\n", *util.DB, *util.DBName, dbConnectString)
	}

	DB, err := gorm.Open(driver, dbConnectString)
	CheckDBError(true, "startup", fmt.Sprintf("Error opening %s Database %s.", *util.DB, *util.DBName), err)
	database = DB
	err = createSchema(DB)

	if err != nil {
		CheckDBError(true, "startup", fmt.Sprintf("Error migrating database schema for %s [%s]!", *util.DB, *util.DBName), err)
	}

	rows, err := database.DB().Query(getDBVersion())

	CheckDBError(true, "startup", fmt.Sprintf("Error obtaining version details for %s Database %s.", *util.DB, *util.DBName), err)

	var version string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&version)
		if (!*util.Xtract) {
			fmt.Printf("CSA: %s DBEngine: %s-%s\tDBName: %s\n", util.App.Model().Version, *util.DB, version, *util.DBName)
		}
	}

	return DB
}

func StartGormTransaction() *gorm.DB {
	return database.Begin()
}

func Exec(sql string) (sql.Result, error) {
	return database.DB().Exec(sql)
}

func setConnectionString(run *model.Run) {

	switch *util.DB {
	case util.SQLITE:
		checkAndCreateDBDir()
		if *util.DBDriverFlags == "" {
			*util.DBDriverFlags = util.Sqlite_driverFlags
		}

		if !strings.HasSuffix(*util.DBName, ".db") {
			*util.DBName = *util.DBName + ".db"
		}

		run.DbPath = *util.DbDir + util.PathSeparator + *util.DBName

		dbConnectString = fmt.Sprintf("%s/%s", *util.DbDir, *util.DBName)
		if *util.DBDriverFlags != "" {
			dbConnectString = fmt.Sprintf("%s/%s?%s", *util.DbDir, *util.DBName, *util.DBDriverFlags)
		}

		driver = sqllite_driver

		//In case we are coming from a command other than analyze!
		if *util.MaxSaveWorkers <= 0 {
			*util.MaxSaveWorkers = 1
		}

		if *util.MaxSaveWorkers != 1 {
			fmt.Println("SQLite does not support multiple workers! Using only 1!")
			*util.MaxSaveWorkers = 1
		}
	case util.POSTGRES:
		if *util.DBDriverFlags == "" {
			*util.DBDriverFlags = util.Postgres_driverFlags
		}

		if !strings.Contains(*util.DBDriverFlags, "dbname") {
			dbConnectString = fmt.Sprintf("dbname=%s %s", *util.DBName, *util.DBDriverFlags)
		} else {
			if *util.DBName != util.DEFAULT_DB_NAME {
				dbNameExp, _ := regexp.Compile("dbname=.*\\s")
				*util.DBDriverFlags = dbNameExp.ReplaceAllString(*util.DBDriverFlags, fmt.Sprintf("dbname=%s ", *util.DBName))
			}
			dbConnectString = *util.DBDriverFlags
		}

		//In case we are coming from a command other than analyze!
		if *util.MaxSaveWorkers <= 0 {
			*util.MaxSaveWorkers = util.DEFAULT_MAX_POSTGRES_WORKERS
		}
		driver = postgres_driver
	}

}

func getDBVersion() string {

	switch *util.DB {
	case util.SQLITE:
		return "select sqlite_version() as version"
	case util.POSTGRES:
		return "select version() as version"
	}
	return ""
}

func checkAndCreateDBDir() {

	*util.DbDir, _ = filepath.Abs(*util.DbDir)
	path := *util.DbDir
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if !strings.Contains(path, string(os.PathSeparator)) {
			path = filepath.Join(".", path)
		}
		os.MkdirAll(path, os.ModePerm)
	}
}

func createSchema(database *gorm.DB) error {
	
	//Create Run
	db := database.AutoMigrate(model.Run{}, model.ReportRef{}, model.ReportHeader{}, model.ReportData{}, model.Rule{},
		model.Recipe{}, &model.Pattern{}, model.Tag{}, model.Finding{}, model.FindingTag{}, model.FindingRecipe{},
		model.RunSloc{}, model.RuleMetric{}, model.Application{}, model.ApplicationTag{}, model.Bin{}, model.BinTag{},
		model.ScoringModel{},model.ExcludePattern{})

	return db.Error
}

func PopulateInitialData(run *model.Run, ruleRepository RuleRepository, binRepo BinRepository, scoringRepo ScoringRepository, database *gorm.DB) {
	count := 0
	database.Find(&model.ReportRef{}).Count(&count)
	if count < 1 {
		PopulateReferenceData()
	}

	database.Find(&model.ReportHeader{}).Count(&count)
	if count < 1 {
		PopulateReportHeaders()
	}

	database.Find(&model.Rule{}).Count(&count)
	if count < 1 {
		if run.Command != util.DeleteAllRulesCmd.FullCommand() &&
			run.Command != util.DeleteRulesCmd.FullCommand() &&
			run.Command != util.ImportRulesCmd.FullCommand() {
			ruleRepository.LoadRules()
		}
	}

	rules, _ := ruleRepository.GetRules()

	database.Find(&model.Bin{}).Count(&count)
	if count < 1 {
		if run.Command != util.DeleteAllBinsCmd.FullCommand() &&
			run.Command != util.DeleteBinsCmd.FullCommand() &&
			run.Command != util.ImportBinsCmd.FullCommand() {
			binRepo.LoadBins(rules)
		}
	}

	database.Find(&model.ScoringModel{}).Count(&count)
	if count < 1 {
		if run.Command != util.DeleteAllModelsCmd.FullCommand() &&
			run.Command != util.DeleteModelsCmd.FullCommand() &&
			run.Command != util.ImportModelsCmd.FullCommand() {
			scoringRepo.LoadModels()
		}
	}
}

func CheckDBForError(shutdown bool, location string, msg string) bool {
	return CheckDBError(shutdown, location, msg, database.Error)
}

func CheckDBError(shutdown bool, location string, msg string, err error) bool {
	if err != nil {
		if msg == "" {
			msg = "Database Error Occurred!"
		}

		detail := fmt.Errorf("\n\n%s\nLocation: %s\nDetails: %s\n", msg, location, err.Error())
		util.TrackError("DB Operation", detail)

		if *util.Verbose {
			fmt.Printf("%s\nSee the CSA Readme and check your ulimit settings (max open files/handles)\n", detail.Error())
		}

		if shutdown {
			_, _ = fmt.Fprintf(os.Stderr, "Database Error occurred! \n Details: %s\nShutting down!\n\n", err.Error())
			os.Exit(2)
		}

		return true
	}

	return false
}

func ExecuteRawSQL(sql string) error {
	res := database.Exec(sql)
	return res.Error
}

func NewRepositoriesManager(db *gorm.DB) *Repositories {
	return &Repositories{
		Rules:    NewRuleRepository(db),
		Sloc:     NewSlocRepository(db),
		Findings: NewFindingRepository(db),
		Run:      NewRunRepository(db),
		Reports:  NewReportDataRepository(db),
		Bins:     NewBinRepository(db),
		Scoring:  NewScoringRepository(db),
	}
}

func NewRepositoriesManagerForRun(run *model.Run) *Repositories {

	repos := &Repositories{
		Rules:    NewRuleRepositoryForRun(run),
		Sloc:     NewSlocRepositoryForRun(run),
		Findings: NewFindingRepositoryForRun(run),
		Run:      NewRunRepositoryForRun(run),
		Reports:  NewReportDataRepositoryForRun(run),
		Bins:     NewBinRepositoryForRun(run),
		Scoring:  NewScoringRepositoryForRun(run),
	}

	PopulateInitialData(run, repos.Rules, repos.Bins, repos.Scoring, run.DB)

	return repos
}
