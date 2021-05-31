package main

import (
	"database/sql"
	"os"

	"github.com/joho/sqltocsv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//--- setup queries

	//***** CrossTab
	qryCrossTab :=
		`select application, 
		   ejb+jaxws+jaxrs+soap+json+jpa+jdbc+servlet+struts+jsf+session+jni+io+websphere+jndi as total,
		   ejb,
		   jaxws,
		   jaxrs,
		   soap,
		   json,
		   jpa,
		   jdbc,
		   servlet,
		   struts,
		   jsf,
		   session,
		   jni,
		   io,
		   websphere,
		   jndi from
		(select application
		, ifnull(sum(case when issue = 'ejb' or issue = 'javaee' then score end),0) as ejb
		, ifnull(sum(case when issue = 'jax-ws' then score end),0) as jaxws
		, ifnull(sum(case when issue = 'jax-rs' then score end),0) as jaxrs
		, ifnull(sum(case when issue = 'soap' then score end),0) as soap
		, ifnull(sum(case when issue = 'json' then score end),0) as json
		, ifnull(sum(case when issue = 'jpa' then score end),0) as jpa
		, ifnull(sum(case when issue = 'jdbc' then score end),0) as jdbc
		, ifnull(sum(case when issue = 'servlet' then score end),0) as servlet
		, ifnull(sum(case when issue = 'struts' then score end),0) as struts
		, ifnull(sum(case when issue = 'jsf' then score end),0) as jsf
		, ifnull(sum(case when issue = 'session' then score end),0) as session
		, ifnull(sum(case when issue = 'spring' then score end),0) as spring
		, ifnull(sum(case when issue = 'jni' then score end),0) as jni
		, ifnull(sum(case when rule in  ('java-fileIO','java-file-system','java-websockets-import','java-batchAnnotations','java-nio') then score end),0) as io
		, ifnull(sum(case when issue in  ('websphere') then score end),0) as websphere
		, ifnull(sum(case when issue = 'jndi' then score end),0) as jndi
		from
		(SELECT findings.application,
			findings.run_id,
			findings.rule,
			findings.effort as score,
			finding_tags.value as issue
		FROM findings
		LEFT JOIN finding_tags ON findings.id = finding_tags.id
		WHERE findings.run_id = 1 AND findings.effort >= 0 AND finding_tags.value != "info"
		--AND finding_tags.value in ('ejb','javaee','jms','messaging','jax-ws','jax-rs','json','jpa','jdbc','servlet','struts','jsf','session','spring','jni', 'io', 'nio', 'jndi')
		ORDER BY findings.application)
		group by upper(application) )
		where total > 0
		order by total desc`

	//***** PortfolioDetail
	qryPortfolioDetail :=
		`SELECT 
	   name as application,
	   score,
	   sloc_cnt as LOC,
	   files_cnt as files
	   FROM applications
	   WHERE run_id = ?
	   ORDER BY score DESC`

	//***** PortfolioSummary

	qryPortfolioSummary :=
		`SELECT 
		count(distinct name) as apps,
		sum(sloc_cnt) as lines,
		sum(files_cnt) as files
    FROM applications
	WHERE run_id = ?
	`

	qryCodeMetrics :=
		`SELECT 
	lang,
	sum(code_lines) as code,
	sum(total_files) as files,
	sum(comment_lines) as comments
	FROM run_slocs 
	WHERE run_id = ?
	GROUP by lang
	ORDER by code DESC
	LIMIT 10;
	`

	//***** TopAPIs

	qryTopAPIs :=
		`SELECT 
	finding_tags.value as issue,
	count(finding_tags.value) as occurance
	FROM findings 
	LEFT JOIN finding_tags ON findings.id = finding_tags.id
	WHERE run_id = ? AND findings.effort >= 0 AND 
	finding_tags.value not in ('sloc', 'snap','api','buildsystem','maven','api', 'ff', 'ant', 'imports', 'info')
	GROUP BY finding_tags.value  
	ORDER BY occurance DESC
	LIMIT 10;`

	qryScoreDetailbyLine :=
		`
	SELECT application,
       effort as score,
       filename,
       line,
       rule,
       pattern as 'searchedfor',
       trim(trim(value, ' '),char(9)) as found,
       advice,
       category,
       fqn
  FROM findings
  WHERE rule != '' AND 
  run_id = ? AND
  rule not in ('SNAP-build-Ant-Maven') AND 
  line > 0 AND
  score > 0
  ORDER BY application, score DESC

	`
	qryScoreDetailbyFile :=
		`
	SELECT DISTINCT application,
       effort as score,
       filename,
       rule,
       pattern as 'targetfile',
       trim(trim(value, ' '),char(9)) as found,
       advice,
       category,
       fqn
  FROM findings
  WHERE rule != '' AND 
  run_id = ? AND
  rule not in ('SNAP-build-Ant-Maven') AND 
  line = 0 AND
  score > 0
  GROUP BY application, filename
  ORDER BY application, score DESC
	`

	if len(os.Args) < 3 {
		println("usage: PAAExport <path to appfoundry.db> <run id>")
		println("example: PAAExport ../data/appfoundry.db 1")

		os.Exit(1)
	}
	dataBasePath := os.Args[1]
	runID := os.Args[2]
	println("Extracting run " + os.Args[2] + " from: " + os.Args[1])

	database, _ := sql.Open("sqlite3", dataBasePath)

	var rows *sql.Rows
	var err error

	rows, err = database.Query(qryCrossTab, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("crosstab.csv", rows)
	if err != nil {
		panic(err)
	}

	rows, err = database.Query(qryPortfolioDetail, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("portfolio-detail.csv", rows)
	if err != nil {
		panic(err)
	}

	rows, err = database.Query(qryPortfolioSummary, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("portfolio-summary.csv", rows)
	if err != nil {
		panic(err)
	}

	rows, err = database.Query(qryCodeMetrics, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("code-metrics.csv", rows)
	if err != nil {
		panic(err)
	}

	rows, err = database.Query(qryTopAPIs, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("top-APIs.csv", rows)
	if err != nil {
		panic(err)
	}

	rows, err = database.Query(qryScoreDetailbyLine, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("score-datail-by-line.csv", rows)
	if err != nil {
		panic(err)
	}

	rows, err = database.Query(qryScoreDetailbyFile, runID)
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("score-datail-by-file.csv", rows)
	if err != nil {
		panic(err)
	}

	database.Close()
}
