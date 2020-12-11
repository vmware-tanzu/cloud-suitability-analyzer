/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"runtime"
)

type Worker func(id string, in <-chan interface{}, out chan<- interface{}, jointWorker chan<- interface{}, args []interface{})
type Work interface{}
type Results interface{}

func StartWorkers(w Worker, workername string, amountOfWork int, workLimitFactor int, maxWorkers int, in <-chan interface{}, out chan<- interface{}, jointWorker chan<- interface{}, args ...interface{}) int {

	if amountOfWork <= 0 {
		return 0
	}

	if workLimitFactor <= 0 {
		workLimitFactor = 1
	}

	workerCnt := amountOfWork / workLimitFactor

	if workerCnt == 0 {
		workerCnt = runtime.NumCPU()
	}

	if workerCnt > maxWorkers {
		workerCnt = maxWorkers
	}

	for i := 1; i <= workerCnt; i++ {
		go w(fmt.Sprintf("%s-%d", workername, i), in, out, jointWorker, args)
	}

	if *Verbose {
		fmt.Printf("[%d] %s-workers started to handle/process [%d] units of work!\n", workerCnt, workername, amountOfWork)
	}

	return workerCnt
}

func StartSingleTonWorker(w Worker, workername string, amountOfWork int, in <-chan interface{}, out chan<- interface{}, jointWorker chan<- interface{}, args ...interface{}) int {
	return StartWorkers(w, workername, amountOfWork, amountOfWork, 1, in, out, jointWorker, args...)
}
