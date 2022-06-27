/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package csa

import "csa-app/model"

func GetLevelForScore(score int) string {

	if score >= model.Low_criticality_low_score && score <= model.Low_criticality_high_score {
		return model.Low_criticality
	} else if score >= model.Medium_criticality_low_score && score <= model.Medium_criticality_high_score {
		return model.Medium_criticality
	} else if score >= model.High_criticality_low_score {
		return model.High_criticality
	} else {
		return model.Info_criticality
	}
}
