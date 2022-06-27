/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"csa-app/db"
)

//Responsible for generating the Bootstrap.go file
func main() {

	rules, err := db.BootStrapRules("../rules/", "resources")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed During Rule Bootstrap!\n\t%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully updated Bootstap with [%d] rules\n", len(rules))

	models, err := db.BootStrapModels("../score-models/", "resources")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed During Scoring Model Bootstrap!\n\t%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully updated Bootstap with [%d] Scoring Models\n", len(models))

	bins, err := db.BootStrapBins("../bins/", "resources", rules)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed during Bin Bootstrap!\n\t%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully updated Bootstap with [%d] bins\n", len(bins))

}
