/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	App = kingpin.New(APP_NAME, "CSA is used to analyze & collect data related to the cloud readiness of an application based on it's source-code.").Action(
		func(context *kingpin.ParseContext) error {
			var err error
			switch context.SelectedCommand.Model().Name {
			case RECALCULCATE_CMD:
				// Handle Default Value for Alias flag
				alias := getElementValue(context.Elements, "alias")
				runId := getElementValue(context.Elements, "run-id")
				if alias == "" {
					alias = fmt.Sprintf("Recalculate Run %s", runId)
				}
				alias = cleanElementValue(alias)
				err = context.SelectedCommand.GetFlag("alias").Model().Value.Set(alias)
				break
			}
			return err
		})
	Verbose   = App.Flag("verbose", "enable verbose mode.").Short('v').Bool()
	Efd       = App.Flag("efd", "exclude finding details").Bool()
	Profile   = App.Flag("profile", "enables profiling (cpu|mem)").Enum("cpu", "mem")
	Xtract    = App.Flag("xtract", "extract minimum output for pipeline usage.").Short('x').Bool()
	Zap       = App.Flag("zap", "zap (purge) the database before running").Short('z').Bool()
	RulesDir  = App.Flag("rules-dir", "directory where csa rules are. Rules found in this directory will be automatically imported on tool startup. This will also be the default directory for `rules` import").Default(DEFAULT_RULES_DIR).String()
	ModelsDir = App.Flag("models-dir", "directory where csa scoring models are. Scoring Models found in this directory will be automatically imported on tool startup. This will also be the default directory for `scoring-models` import").Default(DEFAULT_MODELS_DIR).String()
	OutputDir = App.Flag("output-dir", "directory path where csa results will be output").Default(DEFAULT_OUTPUT_DIR).String()

	//Export Results Configurations
	ExportFormats  = App.Flag("export", "list of expected formats divided by commas that will be used to export findings, ex: csv or csv,html").String()
	ExportDir      = App.Flag("export-dir", "directory path where csa finding exports will be written").Default(DEFAULT_EXPORT_DIR).String()
	ExportFileName = App.Flag("export-file-name", "base name of the \"export\" file, ex: \"csa-export\". Proper extensions will be appended based on \"--export\" command formats requested.").Default(DEFAULT_EXPORT_FILE_NAME).String()

	ExcludedDirsRegEx = App.Flag(EXCLUDED_DIRS_FLAG, "regex pattern of directories not to be included in analysis").Default("^([.].*|target|bin|test|node_modules|eclipse|out|vendors|obj)$").String()
	DB                = App.Flag("db", "which database engine to use (sqlite|postgres)").Default(SQLITE).Enum(SQLITE, POSTGRES)
	DBName            = App.Flag("db-name", "name of database").Default(DEFAULT_DB_NAME).String()
	DbDir             = App.Flag("database-dir", "directory path where database can be found or created. (defaults to csa executable directory)").String()
	DBDriverFlags     = App.Flag("db-driver-flags", "flags to configure the database driver (Default: sqlite: "+Sqlite_driverFlags+" postgres: "+Postgres_driverFlags).String()
	ReportsFlag       = App.Flag("report", "comma delimited list of report(s) to run. (for example \"-r1,3,4\". 0=All)").Default("0").Short('r').String()
	TmpDirPath        = App.Flag("temp-dir", "The root path where files created by csa will be placed. Defaults to OS specific temp path/run-id").Short('t').String()

	//Get Build Info
	BuildInfoCmd = App.Command("info", "Get full build details of this csa executable")

	//csa ui
	CsaCmd  = App.Command("ui", "Launch the CSA UI")
	CsaPort = CsaCmd.Flag("port", "port to start ui on").Default("3001").Int()

	//List reports Command
	ShowReports = App.Command("list", "list available reports to run")
	ReportType  = ShowReports.Flag("report-type", "show reports available for a particular type (all|csa|git)").Default(APP_NAME).String()

	//Recalculate Command
	RecalculateCmd      = App.Command(RECALCULCATE_CMD, "Recalculate the scores for an app within a given run.")
	RecalculateRunId    = RecalculateCmd.Flag("run-id", "ID of the run to recalculate").Required().Uint()
	RecalculateRunLabel = RecalculateCmd.Flag("alias", "the name or alias for this run. (Recalculate Run [run-id])").Short('a').String()

	//Analyze Command
	AnalyzeCmd            = App.Command(ANALYZE_CMD, "analyze the code on the provided path").Default()
	Path                  = AnalyzeCmd.Arg("path", "Path to source or jar/war/ear file").Default(".").String()
	Alias                 = AnalyzeCmd.Flag("alias", "the name or alias for this run. (defaults to target dir or archive name if not provided)").String()
	DecompileDir          = AnalyzeCmd.Flag("decompile-dir", "destination of fernflower decompile (defaults to <jar dir>/decompile)").String()
	AnalyzeArchives       = AnalyzeCmd.Flag("analyze-archives", "this tells csa to decompile archives (jar|ear|war) that is finds under the target path (including within other archives)").Short('a').Default("false").Hidden().Bool()
	OutputReports         = AnalyzeCmd.Flag("output-reports", "create the original csv reports").Bool()
	TxtIndexingEnabled    = AnalyzeCmd.Flag("enable-txt-index", "index the run for free form text searching").Bool()
	FernLocation          = AnalyzeCmd.Flag("fern-jar-path", "Location where fernflower jar can be found. (defaults to csa executable directory)").String()
	DomainFlag            = AnalyzeCmd.Flag("domain-dir", "include domain/1st sub-directory in target path in report(s)").Short('d').Default("false").Hidden().Bool()
	IncludedFilesRegEx    = AnalyzeCmd.Flag(INCLUDED_FILES, "regex pattern of file(s) to include in analysis. Note: if this is set/modified it will override the excluded-files switch").Default(".*").String()
	ExcludedFilesRegEx    = AnalyzeCmd.Flag(EXCLUDED_FILES, "regex pattern of file(s) to exclude from analysis").Default("^(.*[.](exe|png|tiff|tif|gif|jpg|jpeg|bmp|dmg|mpeg|class)|[.].*|csa-config[.](yaml|yml|json))$").String()
	MaxBuffer             = AnalyzeCmd.Flag("max-buffer", "size of buffer for channel(s). Note: this will affect memory utilization and speed").Default("100000").Int()
	MaxSaveWorkers        = AnalyzeCmd.Flag("max-save-workers", "maximum number of workers to utilize for finding save channel. Note: this will affect memory utilization and speed ("+SQLITE+"=1 "+POSTGRES+"=10").Int()
	MaxIndexWorkers       = AnalyzeCmd.Flag("max-idx-workers", "maximum number of workers to utilize for finding index channel. Note: this will affect memory utilization and speed").Default("1").Hidden().Int()
	DumpRuleMetrics       = AnalyzeCmd.Flag("display-rule-metrics", "show rule metrics on std out").Short('m').Bool()
	DisplayUnknownExts    = AnalyzeCmd.Flag("display-unknown-exts", "show unknown extensions on std out").Short('u').Bool()
	DisplayIgnoredFiles   = AnalyzeCmd.Flag("display-ignored-files", "show ignored files on std out").Bool()
	ConfigFile            = AnalyzeCmd.Flag("config-file", "File containing run configuration. RunName and Applications to analyze.").String()
	RuleIncludeTags       = AnalyzeCmd.Flag("rule-include-tags", "comma delimited string of rule tags to determine which rules are used for analysis.").String()
	RuleExcludeTags       = AnalyzeCmd.Flag("rule-exclude-tags", "comma delimited string of rule tags to determine which rules will not be used for analysis").String()
	DiscoveryMode         = AnalyzeCmd.Flag("enable-portfolio-discovery", "tell csa to treat first level dir of target path as application").Short('p').Bool()
	SerialAppAnalysis     = AnalyzeCmd.Flag("enable-serial-app-analysis", "tell csa to analyze one application at a time").Short('c').Bool()
	DisableIgnoreComments = AnalyzeCmd.Flag("disable-ignore-comments", "tell csa to read analyze comments for findings").Short('i').Bool()
	WriteAppConfig        = AnalyzeCmd.Flag("write-app-config", "csa will write an application configuration file to the apps root directory").Short('w').Bool()
	LineBuffer            = AnalyzeCmd.Flag("line-buffer", "size of line buffer used when reading files. This will affect memory utilization and speed").Default(strconv.Itoa(DEFAULT_LINE_BUFFER_SIZE)).Int()
	FailFast              = AnalyzeCmd.Flag("fail-fast", "tell csa to immediately stop the run on any failure. Note: there is a risk of corrupting the csa datastore.").Short('f').Bool()
	DisplayErrors         = AnalyzeCmd.Flag("display-errors", "show errors at end of run").Short('e').Bool()
	WriteConfigsOnly      = AnalyzeCmd.Flag("write-configs-only", "tell csa to only generate config files instead of performing a full run").Short('o').Bool()
	OutputFormatJson      = AnalyzeCmd.Flag("json", "write config files in json format. Default is yaml").Short('j').Bool()
	ScoringModel          = AnalyzeCmd.Flag(SCORING_MODEL_FLAG, "the name of the scoring model to use for scoring applications").Short('s').Default("default").String()
	MaxProcs              = AnalyzeCmd.Flag("max-procs", "Set the max concurrency from a processor perspective. Defaults to system processor count.").Int()
	MaxThreads            = AnalyzeCmd.Flag("max-threads", "Set the max OS threads that csa can utilize. Default is '20000'").Default(strconv.Itoa(20000)).Int()
	Profiles              = AnalyzeCmd.Flag("profiles", "The list of rule profiles that will be selected for the scan. Multiple profiles can separated by commas.").String()
	//Search Command
	SearchCmd = App.Command("search", "search full text index for findings based on query")

	//Git Command
	GitCmd    = App.Command("git", "perform git forensics analysis")
	StartDate = GitCmd.Flag("startDate", "start date (defaults to entire repository)").Short('s').String()
	GitPath   = GitCmd.Arg("path", "path to git repository directory").Required().String()

	//Rules Cmd(s)
	RulesCmd          = App.Command("rules", "modify (import/export) csa rules")
	ExportRulesCmd    = RulesCmd.Command("export", "export rule(s) from the database")
	ExportJson        = ExportRulesCmd.Flag("json", "export rules in json format. Default is yaml").Bool()
	ExportSingleFile  = ExportRulesCmd.Flag("single-file", "export all rules to one file").Bool()
	ExportRuleName    = ExportRulesCmd.Arg("name", "name of rule to export. Default is all").String()
	ImportRulesCmd    = RulesCmd.Command("import", "import rule(s) into the database. By default rules will be added/updated rather than replace existing")
	ReplaceRulesFlag  = ImportRulesCmd.Flag("over-write-rules", "flag indicating that existing rules should replaced by the new rules being loaded").Short('w').Bool()
	ImportRuleName    = ImportRulesCmd.Arg("name", "name of rule to import").String()
	DeleteRulesCmd    = RulesCmd.Command("delete", "delete a rule in the database")
	DeleteRulesName   = DeleteRulesCmd.Arg("name", "name of rule to be deleted").Required().String()
	DeleteAllRulesCmd = RulesCmd.Command("delete-all", "delete all rules in the database")
	ValidateRuleCmd   = RulesCmd.Command("validate", "validate a rule")
	ValidateRuleName  = ValidateRuleCmd.Arg("file", "rule file to validate").Required().String()

	//Bins Cmd(s)
	BinsCmd             = App.Command("bins", "modify (import/export) Bin definition(s)")
	ExportBinsCmd       = BinsCmd.Command("export", "export bins(s) from the database")
	ExportBinsJson      = ExportBinsCmd.Flag("json", "export bins in json format. Default is yaml").Bool()
	ExportBinsTimeStamp = ExportBinsCmd.Flag("timestamp", "export bins with time-stamped filename").Bool()
	ImportBinsCmd       = BinsCmd.Command("import", "import bin(s) into the database.")
	ImportBinFile       = ImportBinsCmd.Arg("bin-file", "file containing bin definitions to import").Default(DEFAULT_BIN_FILE).String()
	DeleteBinsCmd       = BinsCmd.Command("delete", "delete a bin in the database")
	DeleteBinName       = DeleteBinsCmd.Arg("name", "name of bin to be deleted").Required().String()
	DeleteAllBinsCmd    = BinsCmd.Command("delete-all", "delete all rules in the database")
	ValidateBinsCmd     = BinsCmd.Command("validate", "validate a bin(s) file")
	ValidateBinName     = ValidateBinsCmd.Arg("file", "file containing bin definitions to validate").Required().String()

	//Rules Cmd(s)
	ModelsCmd              = App.Command("score-models", "modify (import/export) csa scoring models")
	ExportModelsCmd        = ModelsCmd.Command("export", "export scoring model(s) from the database")
	ExportModelsJson       = ExportModelsCmd.Flag("json", "export scoring models in json format. Default is yaml").Bool()
	ExportModelsSingleFile = ExportModelsCmd.Flag("single-file", "export all scoring models to one file").Bool()
	ExportModelName        = ExportModelsCmd.Arg("name", "name of scoring model to export. Default is all").String()
	ImportModelsCmd        = ModelsCmd.Command("import", "import scoring model(s) into the database. By default models will be added/updated rather than replace existing")
	ReplaceModelsFlag      = ImportModelsCmd.Flag("over-write-models", "flag indicating that models rules should replaced by the new models being loaded").Short('w').Bool()
	ImportModelName        = ImportModelsCmd.Arg("name", "name of scoring model to import").String()
	DeleteModelsCmd        = ModelsCmd.Command("delete", "delete a scoring model in the database")
	DeleteModelName        = DeleteModelsCmd.Arg("name", "name of scoring model to be deleted").Required().String()
	DeleteAllModelsCmd     = ModelsCmd.Command("delete-all", "delete all scoring models in the database")
	ValidateModelCmd       = ModelsCmd.Command("validate", "validate a scoring model")
	ValidateModelName      = ValidateModelCmd.Arg("file", "scoring model file to validate").Required().String()

	//Naturalize Command
	NatLangCmd      = App.Command("naturalize", "process files and create naturual lang files for each code file found")
	NatLangPath     = NatLangCmd.Arg("path", "Path to source or jar/war/ear file").Default(".").String()
	NatLangExension = NatLangCmd.Flag("nat-extensions", "regex pattern of file extensions to include in analysis").Default("(java)").String()

	//Misc
	Lock = &sync.Mutex{}
)

func cleanElementValue(flagValue string) string {
	value := strings.TrimSpace(flagValue)
	if len(value) > 0 {
		switch value[0] {
		case ':':
			value = strings.TrimPrefix(value, string(value[0]))
			break
		case '=':
			value = strings.TrimPrefix(value, string(value[0]))
			break
		}
		replacer := strings.NewReplacer("'", "", "\"", "")
		value = replacer.Replace(value)
	}
	return strings.TrimSpace(value)
}

func getElementValue(elements []*kingpin.ParseElement, elementName string) string {
	var name string
	for e := range elements {
		el := elements[e]
		switch el.Clause.(type) {
		case *kingpin.ArgClause:
			name = el.Clause.(*kingpin.ArgClause).Model().Name
			break
		case *kingpin.FlagClause:
			name = el.Clause.(*kingpin.FlagClause).Model().Name
			break
		case *kingpin.CmdClause:
			name = el.Clause.(*kingpin.CmdClause).Model().Name
			break
		}
		if name == elementName {
			return cleanElementValue(*el.Value)
		}
	}
	return ""
}

func IsAppFlagDefaulted(flag string) bool {
	return IsCmdFlagDefaulted("", flag)
}

func IsCmdFlagDefaulted(cmd string, flag string) bool {

	var defValues []string
	var value string
	var f *kingpin.FlagClause

	if cmd == "" {
		f = App.GetFlag(flag)
	} else {
		f = App.GetCommand(cmd).GetFlag(flag)
	}

	defValues = f.Model().Default
	value = f.Model().Value.String()

	for _, defVal := range defValues {
		if defVal == value {
			return true
		}
	}

	return false
}

const APP_NAME string = "csa"
const DEFAULT_DB_NAME string = "csa"

const Sqlite_driverFlags string = "mode=rwc&_fk=true&_mutex=full&cache=shared&_timeout=10000&_locking=NORMAL"

// const Sqlite_driverFlags string = ""
const Postgres_driverFlags string = "dbname=" + DEFAULT_DB_NAME + " sslmode=disable"
const SQLITE string = "sqlite"
const POSTGRES string = "postgres"
const DEFAULT_RULES_DIR = "./rules"
const DEFAULT_OUTPUT_DIR = "csa-reports"
const DEFAULT_REPORT_DIR = "Reports"
const DEFAULT_EXPORT_DIR = "Exports"
const DEFAULT_EXPORT_FILE_NAME = "export"
const DEFAULT_MODELS_DIR = "./scoring-models"
const RULE_BOOTSTRAP_TEMPLATE = "BootstrapRulesTemplate.txt"
const BIN_BOOTSTRAP_TEMPLATE = "BootstrapBinsTemplate.txt"
const SCORING_MODEL_BOOTSTRAP_TEMPLATE = "BootstrapScoringModelsTemplate.txt"
const DEFAULT_BIN_FILE = "bins.yaml"
const DEFAULT_LINE_BUFFER_SIZE int = 128 * 1024
const MAX_LINE_BUFFER_SIZE int = 4096 * 1024
const DEFAULT_MAX_POSTGRES_WORKERS = 10

// CMDS
const ANALYZE_CMD string = "analyze"
const RECALCULCATE_CMD string = "recalculate"
const PROFILE_FLAG string = "profiles"
const RULE_INCLUDE_FLAG string = "rule-include-tags"
const RULE_EXCLUDE_FLAG string = "rule-exclude-tags"
const EXCLUDED_DIRS_FLAG string = "excluded-dirs"
const INCLUDED_FILES string = "included-files"
const EXCLUDED_FILES string = "excluded-files"
const SCORING_MODEL_FLAG string = "scoring-model"
