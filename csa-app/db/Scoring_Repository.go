/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"csa-app/model"
	"csa-app/util"
	"gopkg.in/yaml.v2"
)

type ScoringRepository interface {
	ExportModels() error
	ImportModels()
	SaveModels(models []*model.ScoringModel) error
	SaveModel(model *model.ScoringModel) error
	GetModels() (models []model.ScoringModel, err error)
	GetModelByName(name string) (*model.ScoringModel, error)
	DeleteModel(name string) error
	DeleteAllModels(notify bool) error
	LoadModels()
	ValidateModel(filename string, run *model.Run)
}

func NewScoringRepository(db *gorm.DB) ScoringRepository {
	return &OrmRepository{
		dbconn: db,
	}
}

func NewScoringRepositoryForRun(run *model.Run) ScoringRepository {
	return &OrmRepository{
		dbconn: run.DB,
	}
}

func (repo *OrmRepository) ImportModels() {

	dir := model.GetModelsDir(true)
	//Get Files...Only json and yaml are supported at this time!
	files := repo.fileUtil.GetFileList(dir, "(json|yaml|yml)")

	importOne := *util.ImportModelName != ""
	var newModels []*model.ScoringModel
	importCnt := 0

	if len(files) > 0 {
		for _, file := range files {

			if *util.Verbose {
				fmt.Printf("Importing Scoring Model File [%s]...", file.Name)
			}
			var decoder util.FileDecoder
			reader, err := os.Open(file.FQN)
			if err != nil {
				fmt.Printf("Error reading file. Details: %v\n", err)
				continue
			}

			if strings.HasSuffix(file.Name, util.JSON) {
				decoder = json.NewDecoder(reader)
			} else {
				decoder = yaml.NewDecoder(reader)
			}

			stop := false
			cnt := 0
			for !stop {
				cnt, stop = repo.unMarshalModel(decoder, &newModels)
				importCnt += cnt
			}
			//Close the file!
			reader.Close()
		}

		if *util.ReplaceModelsFlag && len(newModels) > 0 {
			replacedCnt := 1
			if !importOne {
				models, _ := repo.GetModels()
				replacedCnt = len(models)
			}
			fmt.Printf("Replacing [%d] existing scoring models with [%d] new ones!\n", replacedCnt, len(newModels))
			repo.DeleteAllModels(false)
			repo.SaveModels(newModels)
		}

		if importCnt == 0 && importOne {
			fmt.Printf("Scoring Model [%s] not found for import @ [%s]\n", *util.ImportModelName, dir)
		} else {
			fmt.Printf("Successfully imported [%d] Scoring Model(s) found @[%s]\n", importCnt, dir)
		}

	} else {
		fmt.Printf("Found No Scoring Model files for import @ [%s]\n", dir)
	}

}

func (repo *OrmRepository) ExportModels() error {
	var extension = util.YAML
	if *util.ExportModelsJson {
		extension = util.JSON
	}

	exportCnt := 0

	modelsDir := model.GetModelsDir(false)

	if *util.ExportModelName != "" {
		model, err := repo.GetModelByName(*util.ExportRuleName)
		if err != nil && model.Name != "" {
			util.WriteStructToFile(model, model.Name, modelsDir, extension, true)
			exportCnt++
		} else {
			msg := fmt.Sprintf("Scoring Model [%s] not found for export!\n", *util.ExportModelName)
			fmt.Printf(msg)
			return fmt.Errorf(msg)
		}
	} else {

		models, err := repo.GetModels()

		if err == nil {
			if *util.ExportModelsSingleFile {
				currentTime := time.Now()
				filename := fmt.Sprintf("csa-scoring-models-%d%02d%02d-%02d%02d%02d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())
				exportCnt = util.WriteStructsToFile(models, filename, modelsDir, extension, true)
			} else {
				for _, model := range models {

					util.WriteStructToFile(model, model.Name, modelsDir, extension, true)
					exportCnt++

				}
			}
		} else {
			msg := fmt.Sprintf("Error retrieving models for export! Details: %v", err)
			fmt.Printf(msg)
			return fmt.Errorf(msg)
		}
	}

	fmt.Printf("Successfully exported [%d] scoring-models @ [%s]\n", exportCnt, modelsDir)

	return nil
}

func (repo *OrmRepository) SaveModels(models []*model.ScoringModel) error {
	//Begin Transaction
	tx, err := repo.dbconn.DB().Begin()
	if err == nil {
		for i := range models {
			err = repo.SaveModel(models[i])
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}
	return err
}

func (repo *OrmRepository) ValidateModel(filename string, run *model.Run) {
	fmt.Printf("Validating Scoring Model File [%s]\n\n", filename)

	var decoder util.FileDecoder
	reader, err := util.OpenFileAtPath(filename, run.Homepath)
	if err != nil {
		util.App.Fatalf("Error reading file. Details: %v\n", err)
	}

	if strings.HasSuffix(filepath.Ext(filename), util.JSON) {
		decoder = json.NewDecoder(reader)
	} else {
		decoder = yaml.NewDecoder(reader)
	}

	model, err := extractModel(decoder)

	if err == nil {
		fmt.Printf("Model [%s] is VALID!\n\t# Sloc Ranges: %d\n",
			model.Name, len(model.Ranges))
		fmt.Print("\nUnmarshalled Scoring Model Dump =>\n\n")

		spew.Config.DisablePointerAddresses = true
		spew.Config.DisableMethods = true
		spew.Config.Indent = "\t"
		spew.Config.DisableCapacities = true
		spew.Config.DisablePointerMethods = true
		spew.Config.MaxDepth = 3
		spew.Config.SortKeys = true

		spew.Dump(model)
	} else {
		util.App.Errorf("%v", err)
	}

}

func (repo *OrmRepository) SaveModel(m *model.ScoringModel) (err error) {
	err = m.Validate()
	if err == nil {
		err = m.PrePersist()
		if err == nil {
			err = repo.dbconn.Save(m).Error
		}
	}
	return err
}

func (repo *OrmRepository) GetModels() ([]model.ScoringModel, error) {
	var models []model.ScoringModel
	res := repo.dbconn.
		Find(&models)

	if res.Error != nil {
		return models, res.Error
	}

	for i := range models {
		err := models[i].Hydrate()
		if err != nil {
			return models, err
		}
	}

	return models, res.Error
}

func (repo *OrmRepository) GetModelByName(name string) (*model.ScoringModel, error) {
	availableModels, err := repo.GetModels()
	if err != nil {
		return nil, fmt.Errorf("error retrieving scoring model [%s]! details: %s", name, err.Error())
	}

	if len(availableModels) == 0 {
		return nil, fmt.Errorf("error retrieving scoring model [%s]! no scoring models available", name)
	}

	//Doing it this way to ensure case-insensitivity
	for _, m := range availableModels {
		if strings.ToLower(m.Name) == strings.ToLower(name) {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("error retrieving scoring model [%s]! there are [%d] available models but none matching that name", name, len(availableModels))
}

func (repo *OrmRepository) DeleteModel(name string) error {

	existing, _ := repo.GetModelByName(name)

	if existing.Name == *util.DeleteModelName {
		fmt.Printf("Deleting Model [%s]..\n", existing.Name)

		CheckDBError(true,
			"Delete Scoring Model",
			fmt.Sprintf("Failed Deleting Scoring Model [%s]", existing.Name),
			database.Delete(&existing).Error)
		fmt.Print("done!\n")
	}
	return nil
}

func (repo *OrmRepository) DeleteAllModels(notify bool) error {
	err := repo.dbconn.Delete(&model.Bin{}).Error
	if notify {
		if err == nil {
			fmt.Printf("Successfully deleted all Scoring Models!")
		} else {
			fmt.Fprintf(os.Stderr, "Failed deteling all scoring models! Details: %v", err)
		}
	}
	return err
}

func (repo *OrmRepository) unMarshalModel(decoder util.FileDecoder, newModels *[]*model.ScoringModel) (cnt int, stop bool) {

	model, err := extractModel(decoder)

	if err != nil {
		if err != io.EOF {
			fmt.Printf("Error unmarshalling/validating Scoring Model. Details: %v\n", err)
		}

		stop = true

	} else {
		cnt = 1
		*newModels = append(*newModels, &model)
	}

	return
}

func (repo *OrmRepository) LoadModels() {
	models, _ := repo.GetModels()
	if len(models) < 1 {
		repo.ImportModels()
		//Check again to see if we found any on the file system!
		models, _ = repo.GetModels()
		if len(models) < 1 {
			cnt := 0
			for _, m := range model.BootstrapModels() {
				err := repo.SaveModel(&m)
				if err != nil {
					util.App.Fatalf("Error saving scoring model [%s]! Details: %s", m.Name, err.Error())
				}

				cnt++
			}
			fmt.Printf("Created/Converted [%d] Models\n", cnt)
		}
	}
}

func BootStrapModels(dir string, templatesDir string) ([]model.ScoringModel, error) {

	models, err := getAllResourcesBasedModels(dir)

	if err != nil {
		return models, err
	}

	if err := validateUniqueModelNames(&models); err != nil {
		return models, err
	}

	if len(models) == 0 {
		fmt.Printf("No Scoring Models Found in models-dir[%s]\n", dir)
		os.Exit(1)
	}

	fmt.Printf("Bootstraping [%d] Models\n", len(models))

	modelsTemplate := getScoringModelTemplate(templatesDir, util.SCORING_MODEL_BOOTSTRAP_TEMPLATE)

	path := "model" + util.PathSeparator
	name := "ScoringModelBootstrap"

	fmt.Printf("Name: %s, Path: %s\n", name, path)

	file := util.CreateFile(name, "go", path)

	if err = modelsTemplate.Execute(file, models); err != nil {
		return models, err
	}

	return models, nil
}

func getAllResourcesBasedModels(modelsDir string) ([]model.ScoringModel, error) {
	//Get Files...Only json and yaml are supported at this time!
	fileUtil := util.NewFileUtil()
	files := fileUtil.GetFileList(modelsDir, "(json|yaml|yml)")

	var newModels []model.ScoringModel

	var errs []error

	if len(files) > 0 {
		for _, file := range files {
			if *util.Verbose {
				fmt.Printf("Importing Model File [%s]...", file.Name)
			}

			m, err := readModelFile(file)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			newModels = append(newModels, m)
		}
		if len(errs) > 0 {

			var msg strings.Builder

			msg.WriteString("\t[\n")
			for _, e := range errs {
				msg.WriteString(fmt.Sprintf("\t\t%s", e.Error()))
			}
			msg.WriteString("\t]\n")
			return nil, fmt.Errorf("errors =>\n%s", msg.String())
		}
	} else {
		fmt.Printf("Found No Model files for import @ [%s]\n", modelsDir)
	}
	return newModels, nil
}

func validateUniqueModelNames(models *[]model.ScoringModel) error {
	var m map[string]bool
	m = make(map[string]bool)

	var errs []error
	for _, ml := range *models {
		if _, ok := m[ml.Name]; ok {
			errs = append(errs, fmt.Errorf("more than 1 model named [%s] found!", ml.Name))
			dumpModel(ml, fmt.Sprintf("Duplicate Models Named [%s] detected!", ml.Name))
		}
		m[ml.Name] = true
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

func getScoringModelTemplate(templatesDir string, filename string) *template.Template {
	name := templatesDir + util.PathSeparator + filename

	funcMap := template.FuncMap{
		"now": time.Now,
	}

	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error reading Scoring Model template [%s]! Details: %v\n", name, err)
		os.Exit(1)
	}

	return template.Must(template.New("ScoringModels").Funcs(funcMap).Parse(string(content)))

}

func extractModel(decoder util.FileDecoder) (model.ScoringModel, error) {
	m := model.ScoringModel{MaxScore: 10.0, MinScore: 0.0}
	err := decoder.Decode(&m)
	if err == nil {
		err = m.Validate()
	}

	return m, err
}

func readModelFile(file util.FileInfo) (model.ScoringModel, error) {
	var decoder util.FileDecoder

	reader, err := os.Open(file.FQN)
	if err != nil {
		return model.ScoringModel{}, fmt.Errorf("Error reading file: %s. Details: %v\n", file.Name, err)
	}

	defer reader.Close()

	if strings.HasSuffix(file.Name, util.JSON) {
		decoder = json.NewDecoder(reader)
	} else {
		decoder = yaml.NewDecoder(reader)
	}

	return extractModel(decoder)
}

func dumpModel(model model.ScoringModel, msg string) {

	fmt.Printf("\n%s\n", msg)
	fmt.Printf("Scoring Model [%s]\n\n", model.Name)

	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableMethods = true
	spew.Config.Indent = "\t"
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerMethods = true
	spew.Config.MaxDepth = 3
	spew.Config.SortKeys = true

	spew.Dump(model)
}
