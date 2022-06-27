/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package search

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"csa-app/model"
	"csa-app/util"
	"os"
	"sync"
)

var indexCache map[string]*CsaIndex
var lock sync.Mutex

type CsaIndex struct {
	RunId   uint   `json:"runId"`
	Alias   string `json:"alias"`
	Path    string `json:"path"`
	Exists  bool   `json:"exists"`
	Error   string `json:"error"`
	NumDocs uint64 `json:"numDocs"`
	index   bleve.Index
}

func GetIndex(run *model.Run) *CsaIndex {

	var err error
	var index *CsaIndex

	runAlias := fmt.Sprintf("%d.%s", run.ID, run.GetAlias())

	if *util.Verbose {
		fmt.Printf("Looking for index with alias [%s]\n", runAlias)
	}

	lock.Lock()
	defer lock.Unlock()
	index, ok := indexCache[runAlias]

	//create the index and store it
	if !ok {

		index = &CsaIndex{Alias: runAlias, RunId: run.ID}

		if indexCache == nil {
			indexCache = make(map[string]*CsaIndex)
		}

		index.Path = fmt.Sprintf("%s%s%s.bleve", run.OutputPath, util.PathSeparator, runAlias)

		if util.Exists(index.Path) {
			if *util.Verbose {
				fmt.Printf("Attempting to open existing index @[%s]\n", index.Path)
			}
			index.index, err = bleve.Open(index.Path)
		} else {
			if *util.Verbose {
				fmt.Printf("Attempting to create new index @[%s]\n", index.Path)
			}
			mapping := bleve.NewIndexMapping()
			index.index, err = bleve.New(index.Path, mapping)
		}

		if err != nil {
			index.Exists = false
			index.Error = err.Error()
			return index
		}

		indexCache[index.Alias] = index
		index.Exists = true

	}

	return index
}

func GetExistingIndex(run *model.Run) *CsaIndex {

	var err error
	var index *CsaIndex

	runAlias := fmt.Sprintf("%d.%s", run.ID, run.GetAlias())
	if *util.Verbose {
		fmt.Printf("Looking for index with alias [%s]\n", runAlias)
	}
	lock.Lock()
	defer lock.Unlock()

	index, ok := indexCache[runAlias]

	if !ok {

		index = &CsaIndex{Alias: runAlias, RunId: run.ID}

		if indexCache == nil {
			fmt.Println("Creating new index cache!")
			indexCache = make(map[string]*CsaIndex)
		}

		index.Path = fmt.Sprintf("%s%s%s.bleve", *util.OutputDir, util.PathSeparator, runAlias)
		if *util.Verbose {
			fmt.Printf("Looking for index @[%s]\n", index.Path)
		}
		if util.Exists(index.Path) {
			index.index, err = bleve.Open(index.Path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening bleve index! Details: %v\n", err.Error())
			} else {
				index.NumDocs, err = index.index.DocCount()
				if err != nil {
					fmt.Printf("Error retrieving doc count for index [%s]\n", index.Alias)
				}
				fmt.Printf("Index [%s] opened successfully and contains [%d] documents!\n", index.Path, index.NumDocs)
			}
		} else {
			err = fmt.Errorf("Index File [%s] does not exist! Please re-analyze with --build-txt-index or tell csa where the index exists using --output-dir", index.Path)
			index.Exists = false
			index.Error = err.Error()
		}

		if err != nil {
			index.Exists = false
			index.Error = err.Error()
			return index
		}

		index.index.SetName(runAlias)
		indexCache[runAlias] = index
		index.Exists = true
	} else {
		fmt.Printf("Returning cached index [%s] @ path [%s] with [%d] indexed documents for run[%d]!\n", index.Alias, index.Path, index.NumDocs, index.RunId)
	}
	return index
}

func AddFindingToIndex(run *model.Run, finding *model.Finding) error {
	index := GetIndex(run)
	if !index.Exists {
		return fmt.Errorf("Error adding finding to index! Details: %s", index.Error)
	}

	if *util.Verbose {
		fmt.Printf("Indexing finding [%d] for index [%s]\n", finding.ID, index.index.Name())
	}

	return index.AddItemToIndex(fmt.Sprint(finding.ID), finding)
}

func (i *CsaIndex) AddItemToIndex(key string, item interface{}) error {
	return i.index.Index(key, item)
}
