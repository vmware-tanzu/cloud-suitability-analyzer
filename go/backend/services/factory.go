/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package services

import (
	"github.com/blevesearch/bleve"
	"github.com/vmware-samples/cloud-suitability-analyzer/go/db"
)

type RepoService struct {
	repositoryMgr *db.Repositories
	textIndex     TextIndex
}

type TextIndex struct {
	index     bleve.Index
	populated bool
}

func NewAppService(repoMgr *db.Repositories) ApplicationInfoService {
	return &RepoService{
		repositoryMgr: repoMgr,
	}
}

func NewDataService(repoMgr *db.Repositories) DataService {
	return &RepoService{
		repositoryMgr: repoMgr,
	}
}

func NewScoringService(repoMgr *db.Repositories) ScoringService {
	return &RepoService{
		repositoryMgr: repoMgr,
	}
}
