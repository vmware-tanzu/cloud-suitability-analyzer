/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"csa-app/db"
	"csa-app/util"
)

const OUT_FMT = "%-4s  %-10s  %-20s  %-20s  %-20s  %-20s %-20s\n"

func ExecuteCLISearch(repoMgr *db.Repositories) {
	var input string

	fmt.Println("Perform Findings Search... (q to quit)")
	fmt.Printf("Supported Search Types [%s,%s,%s,%s]\n\n", QUERY_STRING_QUERY, REGEX_QUERY, FUZZY_QUERY, WILCARD_QUERY)

	fmt.Print("Enter Run Id:")
	fmt.Scanln(&input)
	handleQuitRequest(input)

	runIdAsUint, err := strconv.ParseUint(input, 10, 64)

	if !handleError(err, fmt.Sprintf("Invalid Run Id [%s]", input)) {
		return
	}

	//Get Run!
	run, err := repoMgr.Run.GetRun(uint(runIdAsUint))

	if !handleError(err, fmt.Sprintf("Failure retrieving Run [%s] details!", input)) {
		return
	}

	searchEngine := NewSearchEngine(repoMgr.Run, repoMgr.Findings)

	qType := getQueryType()
	query := getQuery()
	maxResults := getMaxResults()

	request := &SearchRequest{Query: query, MaxResults: maxResults, Type: qType}

	//Execute Search!
	results := searchEngine.SearchFindingsHighlighted(run.ID, request, CLI_HIGHLIGHT)

	if results.Successful {
		fmt.Printf("\nSearch Found [%d] matching findings in Run [%d] in %s!\n\n", results.Totalhits, run.ID, results.TookMs)
		fmt.Printf(OUT_FMT, "RANK", "ID", "Field-Matched-On", "Match-Details", "Field", "Value", "Score")
		fmt.Printf("%s\n", util.Padd("-", 110))

		cnt := 1
		for _, hit := range results.Hits {
			fmt.Printf(OUT_FMT, fmt.Sprint(cnt), hit.ID, "", "", "", "", fmt.Sprint(hit.Score))
			for _, match := range hit.MatchedOn {
				cnt := 0
				for _, matchVal := range match.Matches {
					if cnt > 0 {
						fmt.Printf(OUT_FMT, "", "", "", util.RemoveLineEndings(matchVal), "", "", "")
					} else {
						fmt.Printf(OUT_FMT, "", "", match.FieldName, util.RemoveLineEndings(matchVal), "", "", "")
					}
				}
			}
			fmt.Println("")
			for name, value := range hit.Fields {
				fmt.Printf(OUT_FMT, "", "", "", "", name, util.RemoveLineEndings(value), "")
			}
			fmt.Println("")
			fmt.Println("")
			cnt++
		}

	} else {
		fmt.Fprintf(os.Stderr, "Error performing search! Details: %s\n", results.Error)
	}
}

func handleError(err error, msg string) (ok bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s. Details: %v\n", msg, err.Error())
		return false
	}
	return true
}

func handleQuitRequest(input string) {
	if strings.ToLower(input) == "q" {
		os.Exit(0)
	}
}

func getQuery() string {
	fmt.Print("Enter Query:")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error gather query! Details: %v. Please retry...", err.Error())
	}
	line = util.RemoveLineEndings(line)

	fmt.Println("Received query request with => " + line)

	handleQuitRequest(line)
	return line
}

func getQueryType() string {
	fmt.Print("Enter Search Type[query-string]:")
	var qType string
	fmt.Scanln(&qType)
	handleQuitRequest(qType)

	if qType == "" {
		qType = QUERY_STRING_QUERY
	}

	return qType
}

func getMaxResults() int {
	fmt.Print("Enter Max Results[100]:")
	var max string
	fmt.Scanln(&max)
	handleQuitRequest(max)

	maxResults := DEFAULT_MAX_RESULTS
	if max != "" {
		fmt.Printf("Max = %s\n", max)
		maxResults, _ = strconv.Atoi(max)
	}
	return maxResults
}
