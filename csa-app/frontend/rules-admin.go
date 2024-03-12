/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package main

import (
	"csa-app/db"
	"csa-app/model"
	"csa-app/util"

	"github.com/sirupsen/logrus"

	"csa-app/backend/routes"
	"os"
)

var (
	Version string
)

func main() {

	util.App.HelpFlag.Short('?')
	util.App.Version(Version)
	util.App.Author("VMware by Broadcom")

	// Begin DB Initialization
	var run = model.Run{}
	var err error

	run.Command, err = util.App.Parse(os.Args[1:])

	if *util.Verbose {
		logrus.StandardLogger().Level = logrus.DebugLevel
	}

	if err != nil {
		util.App.FatalUsage("Error parsing command arguments! %s\n", err)
		os.Exit(1)
	}

	//Set initial paths!
	run.SetPaths("")

	setupDefaults(run)

	//Open connection to DB for Reference Data and storing analysis
	database := db.OpenDB(&run)
	defer database.Close()
	routes.StartRouter(database, false, 3001)

}

func setupDefaults(run model.Run) {
	if *util.DbDir == "" {
		util.DbDir = &run.Exepath
	}

	if *util.FernLocation == "" {
		util.FernLocation = &run.Exepath
	}
}
