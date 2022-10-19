/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"sync"
)

var errors = make(map[string][]error)
var errLock sync.Mutex

func TrackError(source string, err error) {
	errLock.Lock()
	defer errLock.Unlock()

	if _, found := errors[source]; !found {
		errors[source] = []error{}
	}
	errors[source] = append(errors[source], err)
}

func HasErrors() bool {
	return len(errors) > 0
}

func ProcessHadErrors(name string) bool {
	_, found := errors[name]
	return found
}

func DumpErrors(runId uint) {
	fmt.Println("\n*************************************************************************************************")
	fmt.Printf("\t\t\t\tRun [%d] encountered errors!\n", runId)
	fmt.Printf("*************************************************************************************************\n\n")

	for k, v := range errors {
		fmt.Printf("%s\n", k)
		for _, e := range v {
			fmt.Printf("\t%s\n", e.Error())
		}
		fmt.Printf("\n")
	}
	fmt.Println("*************************************************************************************************")
}
