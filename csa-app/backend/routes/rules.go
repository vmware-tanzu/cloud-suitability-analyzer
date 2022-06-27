/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"csa-app/db"
	"net/http"
)

type ruleRoutes struct {
	rulesRepo db.RuleRepository
}

func (r *ruleRoutes) getAllRules(c *gin.Context) {
	rules, _ := r.rulesRepo.GetRules()
	c.JSON(http.StatusOK, gin.H{
		"rules": rules,
	})
}

func (r *ruleRoutes) getMetrics(c *gin.Context) {
	runId := getId(c)
	fmt.Printf("Getting rule metrics for Run[%d]\n", runId)

	metrics, err := r.rulesRepo.GetRuleMetrics(runId)

	if !CheckForError(c, err, fmt.Sprintf("Error retrieving rule-metrics for run[%d]! Details => %%s", runId)) {
		c.JSON(http.StatusOK, gin.H{
			"metrics": metrics,
		})
	}
}
