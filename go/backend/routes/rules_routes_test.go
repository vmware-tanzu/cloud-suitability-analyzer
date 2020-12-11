/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package routes_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/backend/routes"
	"github.com/vmware-tanzu/cloud-suitability-analyzer/go/db/test_support"
	"github.com/stretchr/testify/assert"
)

func TestRulesRouter(t *testing.T) {

	_, dir, database, err := db_test_support.OpenTestDb()
	defer os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
	router := routes.SetupRouter(database, false)

	req, _ := http.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
