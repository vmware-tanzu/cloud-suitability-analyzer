/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package search

import (
	"fmt"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	_ "github.com/blevesearch/bleve/search/highlight/format/ansi"
	_ "github.com/blevesearch/bleve/search/highlight/highlighter/ansi"
	bQuery "github.com/blevesearch/bleve/search/query"
	"csa-app/db"
	"csa-app/util"
)

const HTML_HIGHLIGHT = "html"
const CLI_HIGHLIGHT = "ansi"
const QUERY_STRING_QUERY = "query-string"
const FUZZY_QUERY = "fuzzy"
const WILCARD_QUERY = "wildcard"
const REGEX_QUERY = "regex"
const DEFAULT_MAX_RESULTS = 100

type SearchEngine struct {
	runRepo      db.RunRepository
	findingsRepo db.FindingRepository
}

type SearchRequest struct {
	Query      string `json:"query" binding:"required"`
	MaxResults int    `json:"max,omitempty"`
	Type       string `json:"type,omitempty"`
}

type SearchResult struct {
	Run        string
	Query      string
	MaxResults int
	Successful bool
	Error      string
	Totalhits  int
	TookMs     string
	Hits       []*SearchHit
}

type SearchHit struct {
	ID        string
	Fields    map[string]string
	MatchedOn []*Match
	Score     float64
}

type Match struct {
	FieldName string
	Matches   []string
}

func NewSearchEngine(runRepo db.RunRepository, findingsRepo db.FindingRepository) *SearchEngine {
	return &SearchEngine{runRepo: runRepo, findingsRepo: findingsRepo}
}

func (engine *SearchEngine) SearchFindings(runid uint, request *SearchRequest) *SearchResult {
	return engine.SearchFindingsHighlighted(runid, request, "")
}

func (engine *SearchEngine) SearchFindingsHighlighted(runid uint, searchRequest *SearchRequest, highlightStyle string) *SearchResult {

	searchResult := &SearchResult{Run: fmt.Sprint(runid), Query: searchRequest.getQuery(), MaxResults: searchRequest.MaxResults}

	run, err := engine.runRepo.GetRun(runid)

	if err != nil {
		fmt.Printf("\nError retrieving run [%d]! Details: %v\n", runid, err)
		return newErrorResult(runid, searchRequest.getQuery(), err.Error())
	}

	var blvquery bQuery.Query
	var qType string

	switch strings.ToLower(searchRequest.Type) {
	case FUZZY_QUERY:
		qType = strings.ToUpper(FUZZY_QUERY)
		blvquery = bleve.NewFuzzyQuery(strings.ToLower(searchRequest.getQuery()))
	case WILCARD_QUERY:
		qType = strings.ToUpper(WILCARD_QUERY)
		blvquery = bleve.NewWildcardQuery(searchRequest.getQuery())
	case REGEX_QUERY:
		qType = strings.ToUpper(REGEX_QUERY)
		blvquery = bleve.NewRegexpQuery(searchRequest.getQuery())
	default:
		qType = strings.ToUpper(QUERY_STRING_QUERY)
		blvquery = bleve.NewQueryStringQuery(searchRequest.getQuery())
	}

	fmt.Printf("Performing %s query of run [%s] with search term [%s] max-results: %d\n", qType, searchResult.Run, searchRequest.getQuery(), searchRequest.MaxResults)
	if *util.Verbose {
		fmt.Printf("Performing %s query of run [%s] with search term [%s] max-results: %d\n", qType, searchResult.Run, searchRequest.getQuery(), searchRequest.MaxResults)
	}

	request := bleve.NewSearchRequest(blvquery)
	request.Size = searchRequest.MaxResults

	//var fields []string
	//fields = append(fields, "Category", "Value", "Rule", "Score", "Criticality", "Pattern", "Filename", "Line", "Advice")
	//request.Fields = fields
	request.Fields = []string{"*"}

	if highlightStyle != "" {
		request.Highlight = bleve.NewHighlightWithStyle(highlightStyle)
	}

	index := GetExistingIndex(&run)
	if !index.Exists {
		fmt.Printf("\nError retrieving index for run [%d]! Details: %v\n", runid, err)
		return newErrorResult(runid, searchRequest.getQuery(), err.Error())
	}

	searchResult.performSearch(index, request)

	fmt.Printf("Search succeeded: %v Hits: %d Took: %s\n", searchResult.Successful, searchResult.Totalhits, searchResult.TookMs)

	return searchResult

}

func newErrorResult(runid uint, query string, errMsg string) *SearchResult {
	return &SearchResult{Run: fmt.Sprint(runid), Query: query, Error: errMsg, Successful: false}
}

func (r *SearchResult) performSearch(index *CsaIndex, request *bleve.SearchRequest) {

	result, err := index.index.Search(request)

	if err != nil {
		r.Successful = false
		r.Error = err.Error()
		return
	}

	r.Successful = true
	r.Totalhits = int(result.Total)
	r.TookMs = fmt.Sprintf("%v", result.Took)

	for _, hit := range result.Hits {
		r.addHit(hit)
	}
}

func (r *SearchResult) addHit(hit *search.DocumentMatch) {
	newHit := &SearchHit{ID: hit.ID, Score: hit.Score}

	r.Hits = append(r.Hits, newHit)

	for name, value := range hit.Fields {
		newHit.addField(name, fmt.Sprint(value))
	}

	for fieldName, matches := range hit.Fragments {
		newHit.addMatchedOn(fieldName, matches)
	}

}

func (h *SearchHit) addField(name string, value string) {

	if h.Fields == nil {
		h.Fields = make(map[string]string)
	}

	h.Fields[name] = value
}

func (h *SearchHit) addMatchedOn(name string, values []string) {
	match := &Match{FieldName: name}
	h.MatchedOn = append(h.MatchedOn, match)

	for _, value := range values {
		match.Matches = append(match.Matches, value)
	}
}

func (r *SearchRequest) getQuery() string {

	var q string

	if strings.Contains(r.Query, ":") {
		queryParts := strings.Split(r.Query, " ")

		for _, part := range queryParts {

			if len(q) > 0 {
				q += " "
			}

			if strings.Contains(part, ":") {
				subParts := strings.Split(part, ":")

				q += subParts[0]
				q += ":"
				q += strings.ToLower(subParts[1])
			} else {
				q += strings.ToLower(part)
			}
		}
	} else {
		q = strings.ToLower(r.Query)
	}

	return q
}
