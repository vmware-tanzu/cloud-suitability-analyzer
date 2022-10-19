/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/jinzhu/gorm"
	"csa-app/model"
	"csa-app/util"
)

type BinRepository interface {
	ExportBins()
	ImportBins(rules []model.Rule)
	SaveBins(bins []*model.Bin) error
	SaveBin(bin *model.Bin) error
	GetBins() (bins []model.Bin, err error)
	GetBinByName(binName string) (model.Bin, error)
	DeleteBin(binName string) error
	DeleteAllBins(notify bool) error
	LoadBins(rules []model.Rule)
}

func NewBinRepository(db *gorm.DB) BinRepository {
	return &OrmRepository{
		dbconn: db,
	}
}

func NewBinRepositoryForRun(run *model.Run) BinRepository {
	return &OrmRepository{
		dbconn: run.DB,
	}
}

func (repo *OrmRepository) ImportBins(rules []model.Rule) {

	if *util.ImportBinFile == "" {
		*util.ImportBinFile = util.DEFAULT_BIN_FILE
	}

	binFile, err := util.GetFileIfExists(*util.ImportBinFile)

	if err != nil {
		fmt.Printf("Found No Bin file for import @ [%s]\n", *util.ImportBinFile)
		return
	}

	if *util.Verbose {
		fmt.Printf("Importing Bin Definition File [%s]...", binFile.Name)
	}

	var decoder util.FileDecoder
	reader, err := os.Open(binFile.FQN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file. Details: %v\n", err)
		os.Exit(1)
	}

	if strings.HasSuffix(binFile.Name, util.JSON) {
		decoder = json.NewDecoder(reader)
	} else {
		decoder = yaml.NewDecoder(reader)
	}

	var newBins []*model.Bin
	importCnt := 0

	existingBins, err := repo.GetBins()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure retrieving existing rules during Bin import. Details: %v\n", err)
		os.Exit(1)
	}

	stop := false
	cnt := 0
	for !stop {
		cnt, stop = repo.unMarshalBin(decoder, &newBins, rules)
		importCnt += cnt
	}
	//Close the file!
	reader.Close()

	if len(newBins) > 0 {
		fmt.Printf("Replacing [%d] existing bins with [%d] new ones!\n", len(existingBins), len(newBins))
		err = repo.DeleteAllBins(false)
		err = repo.SaveBins(newBins)
		if err == nil {
			fmt.Printf("Successfully imported [%d] bins(s) found @[%s]\n", importCnt, binFile.FQN)
		} else {
			fmt.Fprintf(os.Stderr, "Failed importing bins! Details: %v\n", err)
		}
	}

}

func (repo *OrmRepository) ExportBins() {
	var extension = util.YAML
	if *util.ExportBinsJson {
		extension = util.JSON
	}
	exportCnt := 0
	outputDir := *util.OutputDir
	bins, err := repo.GetBins()
	if err == nil {

		for _, bin := range bins {
			bin.PrepForMarshall()
		}

		currentTime := time.Now()
		filename := "bins"
		if *util.ExportBinsTimeStamp {
			filename = fmt.Sprintf("csa-bins-%d%02d%02d-%02d%02d%02d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())
		}
		exportCnt = util.WriteStructsToFile(bins, filename, outputDir, extension, true)
		fmt.Printf("Successfully exported [%d] bins @ [%s]", exportCnt, outputDir+util.PathSeparator+filename+"."+extension)
	} else {
		fmt.Printf("Error retrieving bins for export! Details: %v", err)
		os.Exit(1)
	}

}

func (repo *OrmRepository) SaveBins(bins []*model.Bin) error {
	//Begin Transaction
	tx, err := repo.dbconn.DB().Begin()
	if err == nil {
		for i := range bins {
			err = repo.SaveBin(bins[i])
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}
	return err
}

func (repo *OrmRepository) SaveBin(bin *model.Bin) (err error) {
	err = repo.validateBin(bin)
	if err == nil {
		bin.PrepForPersist()
		err = repo.dbconn.Save(bin).Error
	}
	return err
}

func (repo *OrmRepository) GetBins() ([]model.Bin, error) {
	var bins []model.Bin
	res := repo.dbconn.Preload("Tags").Find(&bins)
	return bins, res.Error
}

func (repo *OrmRepository) GetBinByName(binName string) (model.Bin, error) {
	var bin model.Bin
	response := repo.dbconn.Where(&model.Bin{Name: binName}).Preload("Tags").Find(&bin)
	return bin, response.Error
}

func (repo *OrmRepository) DeleteBin(binName string) error {

	existingBin, _ := repo.GetBinByName(binName)

	if existingBin.Name == *util.DeleteBinName {
		fmt.Printf("Deleting Bin [%s]...", existingBin.Name)

		CheckDBError(true,
			"Delete Bin",
			fmt.Sprintf("Failed Deleting Bin [%s]", existingBin.Name),
			database.Delete(&existingBin).Error)
		fmt.Print("done!\n")
	}
	return nil
}

func (repo *OrmRepository) DeleteAllBins(notify bool) error {
	err := repo.dbconn.Delete(&model.Bin{}).Error
	if notify {
		if err == nil {
			fmt.Printf("Successfully deleted all bins!")
		} else {
			fmt.Fprintf(os.Stderr, "Failed deteling all bins! Details: %v", err)
		}
	}
	return err
}

func (repo *OrmRepository) validateBin(bin *model.Bin) error {
	for _, tag := range bin.Tags {
		var t model.Tag
		response := repo.dbconn.Where(&model.Tag{Value: tag.Name}).Find(&t)
		if response.Error != nil || response.RowsAffected == 0 {
			//Check to see if tag exists on any pattern!
			var p model.Pattern
			response := repo.dbconn.Where(&model.Pattern{Tag: tag.Name}).Find(&p)
			if response.Error != nil || response.RowsAffected == 0 {
				tag.PrepForMarshall()
				return fmt.Errorf("bin [%s] is invalid! tag: %s type: %s does not exist on any rule or pattern! details: %v", bin.Name, tag.Name, tag.Action, response.Error)
			}
		}
		tag.PrepForPersist()
	}
	return nil
}

func (repo *OrmRepository) unMarshalBin(decoder util.FileDecoder, newBins *[]*model.Bin, rules []model.Rule) (cnt int, stop bool) {

	bin, err := extractBin(decoder, rules)

	if err != nil {
		if err != io.EOF {
			fmt.Printf("Error unmarshalling/validating Bin. Details: %v\n", err)
		}

		stop = true

	} else {
		cnt = 1
		*newBins = append(*newBins, &bin)
	}

	return
}

func (repo *OrmRepository) LoadBins(rules []model.Rule) {
	bins, _ := repo.GetBins()
	if len(bins) < 1 {

		repo.ImportBins(rules)

		//Check again to see if we found any on the file system!
		rules, _ := repo.GetBins()
		if len(rules) < 1 {

			cnt := 0
			for _, bin := range model.BootstrapBins() {
				err := repo.SaveBin(&bin)
				if err != nil {
					util.App.Fatalf("Error saving bin [%s]! Details: %s", bin.Name, err.Error())
				}

				cnt++
			}

			fmt.Printf("Created/Converted [%d] Bins\n", cnt)
		}
	}
}

func BootStrapBins(dir string, templatesDir string, rules []model.Rule) (bins []model.Bin, err error) {

	exts := []string{"json", "yml", "yaml"}
	var binFile util.FileInfo

	for _, ext := range exts {
		binFile, err = util.GetFileIfExists(dir + util.PathSeparator + "bins." + ext)
		if err == nil {
			break
		}
	}

	if !binFile.Exists {
		return bins, fmt.Errorf("bin file does not exists @ path [%s] with name [%s] with any of the following extensions [json, yml, yaml]", dir, "bins")
	}

	bins, err = getAllResourcesBasedBins(binFile, rules)

	if err != nil {
		return bins, err
	}

	fmt.Printf("Bootstraping [%d] Bins\n", len(bins))

	binsTemplate := getBinTemplate(templatesDir, util.BIN_BOOTSTRAP_TEMPLATE)

	path := "model" + util.PathSeparator
	name := "BinBootstrap"

	fmt.Printf("Name: %s, Path: %s\n", name, path)

	file := util.CreateFile(name, "go", path)

	if err = binsTemplate.Execute(file, bins); err != nil {
		return bins, err
	}

	return bins, nil
}

func getAllResourcesBasedBins(file util.FileInfo, rules []model.Rule) ([]model.Bin, error) {
	var newBins []model.Bin
	//fmt.Printf("Bootstraping Bin File [%s]...", file.Name)
	var decoder util.FileDecoder
	reader, err := os.Open(file.FQN)
	if err != nil {
		return newBins, err
	}
	defer reader.Close()

	if strings.HasSuffix(file.Name, util.JSON) {
		decoder = json.NewDecoder(reader)
	} else {
		decoder = yaml.NewDecoder(reader)
	}

	stop := false
	cnt := 0
	for !stop {
		bin, err := extractBin(decoder, rules)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error unmarshalling/validating Bin. Details: %v\n", err)
				return newBins, err
			}
			stop = true
		} else {
			cnt++
			newBins = append(newBins, bin)
		}
	}
	return newBins, nil
}

func getBinTemplate(templatesDir string, filename string) *template.Template {
	name := templatesDir + util.PathSeparator + filename

	funcMap := template.FuncMap{
		"now": time.Now,
	}

	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error reading bins template [%s]! Details: %v\n", name, err)
		os.Exit(1)
	}

	return template.Must(template.New("Bins").Funcs(funcMap).Parse(string(content)))

}

func extractBin(decoder util.FileDecoder, rules []model.Rule) (model.Bin, error) {
	var bin model.Bin
	err := decoder.Decode(&bin)
	if err == nil {
		for _, tag := range bin.Tags {
			tagFound := false
			for _, rule := range rules {
				for _, rTag := range rule.Tags {
					if bin.HasTag(rTag.Value) {
						tagFound = true
						break
					}
				}
				if tagFound {
					break
				}
			}
			if !tagFound {
				err = fmt.Errorf("bin [%s] has tag [%s] that does not exist on any rule", bin.Name, tag.Name)
				return bin, err
			}
		}
	}

	bin.PrepForPersist()

	return bin, err
}
