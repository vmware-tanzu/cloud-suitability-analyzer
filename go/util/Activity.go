/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"sync"
	"time"
)

type Activity struct {
	name  string
	start time.Time
	stop  time.Time
	sync.Mutex
}

func NewActivity(name string) *Activity {

	activity := &Activity{}
	activity.name = name
	activity.start = time.Now()

	return activity
}

func (activity *Activity) Start() {
	activity.Lock()
	activity.start = time.Now()
	activity.Unlock()
}

func (activity *Activity) Stop() {
	activity.Lock()
	activity.stop = time.Now()
	activity.Unlock()
}

func (activity *Activity) GetElapsed() time.Duration {
	activity.Lock()
	defer activity.Unlock()
	if activity.stop.IsZero() {
		return time.Since(activity.start)
	}

	return activity.stop.Sub(activity.start)
}
