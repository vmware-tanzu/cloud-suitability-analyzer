/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"fmt"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vmware-samples/cloud-suitability-analyzer/go/model"
	log "github.com/sirupsen/logrus"
)

type RunRepository interface {
	StartRun(r *model.Run) error
	StopRun(r *model.Run) error
	GetRuns() ([]model.Run, error)
	GetRunsByCommand(cmd string) ([]model.Run, error)
	GetRun(runId uint) (model.Run, error)
	GetRunApps(runId uint) ([]model.Application, error)
	UpdateApp(app *model.Application) error
	GetApp(runId uint, appName string) (*model.Application, error)
	GetAppByID(runId uint, appId uint) (*model.Application, error)
}

func NewRunRepository(db *gorm.DB) RunRepository {
	return &OrmRepository{
		dbconn: db,
	}
}

func NewRunRepositoryForRun(run *model.Run) RunRepository {
	return &OrmRepository{
		dbconn: run.DB,
	}
}

func (runRepository *OrmRepository) StartRun(r *model.Run) error {
	r.StartTime = time.Now()
	err := runRepository.dbconn.Create(r).Error
	return err
}

func (runRepository *OrmRepository) StopRun(r *model.Run) error {
	r.Runtime = fmt.Sprintf("%v", time.Since(r.StartTime))
	err := runRepository.dbconn.Save(r).Error
	return err
}

func (repo *OrmRepository) GetRuns() ([]model.Run, error) {
	var runs []model.Run

	res := repo.dbconn.Find(&runs)

	for idx := range runs {
		runs[idx].PrepForMarshal()
	}

	return runs, res.Error
}

func (repo *OrmRepository) GetRun(runId uint) (model.Run, error) {
	var run model.Run

	res := repo.dbconn.Where(&model.Run{ID: runId}).Find(&run)

	return run, res.Error
}

func (repo *OrmRepository) GetRunsByCommand(cmd string) ([]model.Run, error) {
	var runs []model.Run
	res := repo.dbconn.Where(&model.Run{Command: cmd}).Order("id asc").Find(&runs)

	for idx := range runs {
		runs[idx].PrepForMarshal()
	}

	return runs, res.Error
}

func (repo *OrmRepository) GetRunApps(runId uint) ([]model.Application, error) {
	var apps []model.Application
	start := time.Now()
	res := repo.dbconn.Where(&model.Application{RunID: runId}).Order("Name asc").Preload("Tags").Find(&apps)
	log.Debugf("Get Run Apps -Retrieve Apps Took [%s]\n", time.Since(start))

	if res.Error == nil {
		//Get the apps tags and apply them
		for i := range apps {
			//Bin the apps with tags!
			sort.Sort(model.ByValue(apps[i].Tags))
			repo.binApp(&apps[i])
		}
	}

	log.Debugf("Get Run Apps Took [%s]\n", time.Since(start))

	return apps, res.Error
}

func (repo *OrmRepository) GetApp(runId uint, appName string) (*model.Application, error) {
	app := &model.Application{}
	response := repo.dbconn.Where(&model.Application{RunID: runId, Name: appName}).Find(app)
	return app, response.Error
}

func (repo *OrmRepository) GetAppByID(runId uint, appId uint) (*model.Application, error) {
	app := &model.Application{}
	response := repo.dbconn.Where(&model.Application{RunID: runId, ID: appId}).Find(app)
	return app, response.Error
}

func (repo *OrmRepository) UpdateApp(updateRequestApp *model.Application) error {
	app, err := repo.GetAppByID(updateRequestApp.RunID, updateRequestApp.ID)
	if err == nil {
		app.MergeApp(updateRequestApp)
		db := repo.dbconn.Save(app)
		err = db.Error
	}
	return err
}

func (repo *OrmRepository) binApp(app *model.Application) {
	var bins []model.Bin
	res := repo.dbconn.Preload("Tags").Find(&bins)

	if res.Error != nil {
		fmt.Printf("Error retrieving bins for application binning! Details: %s", res.Error.Error())
		return
	}

	for _, bin := range bins {
		//For each bin see if app matches
		if bin.Matches(app.Tags) {
			app.AssociateBin(bin)
		}
	}
}
