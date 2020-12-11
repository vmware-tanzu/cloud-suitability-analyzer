/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

//Criterion and Groups at this level are always logically ANDed!
type Criteria struct {
	criterion       []Criterion
	criterionGroups []CriterionGroup
}

//An individual criterion
type Criterion struct {
	key       string //Allowed values: Tag, Category, RuleName, Effort, RunID, Level, Criticality, Application
	value     string
	matchType string //Exact, CI-Exact, Like
}

//A grouping of criterion that can be ANDed or ORed
type CriterionGroup struct {
	criterion []Criterion
	operation string //And or Or
}
