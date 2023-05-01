/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"log"
	"sync"
)

type Spinner struct {
	chars []string
	index int
	sync.Mutex
}

var (
	defaultSpinner      = NewSpinner(nil)
	spinners            = make(map[int]*Spinner)
	spinnersInitialized = false
)

func NewSpinner(chars []string) *Spinner {
	newSpinner := &Spinner{}
	if len(chars) > 0 {
		newSpinner.chars = chars
	} else {
		newSpinner.chars = []string{"|", "/", "-", "\\"}
	}

	return newSpinner
}

func InitializeSpinners(cnt int) {
	if !spinnersInitialized {
		Lock.Lock()
		if !spinnersInitialized {
			for i := 0; i < cnt; i++ {
				spinners[i] = NewSpinner(nil)
			}
			spinnersInitialized = true
		}

		if *Verbose {
			fmt.Printf("Initialized [%d] spinners!\n", len(spinners))
		}
		Lock.Unlock()
	}
}

func WriteLog(action string, format string, v ...interface{}) {
	WriteLogWithToken(action, "", format, v...)
}

func WriteLogWithToken(action string, token string, format string, v ...interface{}) {
	WriteLogLineWithToken(0, action, token, format, v...)
}

func WriteLogLine(line int, action string, format string, v ...interface{}) {
	WriteLogLineWithToken(line, action, "", format, v...)
}

func WriteLogLineWithToken(line int, action string, token string, format string, v ...interface{}) {

	sp := defaultSpinner

	if spinnersInitialized {
		sp = spinners[line]
	}

	if *Verbose {
		log.Printf(format, v...)
	} else {
		symbol := token
		if token == "" {
			sp.Lock()
			if sp.index == len(sp.chars) {
				sp.index = 0
			}
			symbol = sp.chars[sp.index]
			sp.index++
			sp.Unlock()
		}
		if (!*Xtract) {
			fmt.Printf("%s...%s\r", action, symbol)
		}
	}
}

func GetModCount(cnt int) int {
	modCnt := cnt / 100
	if modCnt == 0 || modCnt == 1 {
		modCnt = 2
	}

	return modCnt
}
