/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db_test

import (
	"log"
	"os"

	//"os"
	"testing"
	"time"

	"csa-app/db"
	"csa-app/db/test_support"
	"csa-app/model"
	"github.com/stretchr/testify/assert"
)

func TestSaveAndRetrieveReportData(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	time1, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	reportData := &model.ReportData{
		RunID:     1,
		Data1:     "data1",
		Data2:     "data2",
		Data3:     "data3",
		Data4:     "data4",
		Data5:     "data5",
		Data6:     "data6",
		Data7:     "data7",
		Data8:     "data8",
		Data9:     "data9",
		Data10:    "data10",
		CreatedAt: time1,
		UpdatedAt: time1,
		ReportID:  2,
	}

	reportDataRepository := db.NewReportDataRepository(database)
	err = reportDataRepository.SaveReportData(reportData)

	assert.Nil(t, err)

}
