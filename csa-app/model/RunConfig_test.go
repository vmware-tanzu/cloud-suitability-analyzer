/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model_test

//import (
//	"fmt"
//	"os"
//	"testing"
//
//	"csa-app/util"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestBlowerGetAppsFromPath(t *testing.T) {
//
//	//b := true
//	f := false
//	util.AnalyzeArchives = &f
//	util.Verbose = &f
//
//	fu := util.NewFileUtil()
//
//	//Establish Temp Directory for Run
//	tmpPath, err := fu.EstablishTempDirectory("")
//	if err != nil {
//		fmt.Fprintf(os.Stderr,"Unable to create temp directory. Details: %v", err)
//		assert.Fail(t,"Failed to create temp directory!")
//	}else{
//		fmt.Printf("Created Tmp Dir => %s\n", tmpPath)
//	}
//
//	defer fu.RemoveDir(tmpPath)
//
//	//testpath := "/Users/jonathannewell/Downloads/sample.jar"
//	testpath := "/Users/jonathannewell/working/portfolio"
//
//
//	foundry.
//
//	apps := fu.portfolioDiscovery(testpath, ".*", "(classes|bin)")
//
//	totalFiles := 0
//
//	for _, app := range apps {
//		appFileCnt := len(app.Files)
//		fmt.Printf("Found App [%s] with [%d] files\n", app.Name, appFileCnt)
//		totalFiles += appFileCnt
//	}
//
//	fmt.Printf("Found [%d] total files!\n", totalFiles)
//	fmt.Printf("Tmp Dir => %s\n", tmpPath)
//
//}
