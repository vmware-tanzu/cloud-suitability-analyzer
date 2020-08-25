// Code generated by gopkg.in/src-d/enry.v1/internal/code-generator DO NOT EDIT.
// Extracted from github/linguist commit: e4560984058b4726010ca4b8f03ed9d0f8f464db

package data

import "strings"

// LanguageByAliasMap keeps alias for different languages and use the name of the languages as an alias too.
// All the keys (alias or not) are written in lower case and the whitespaces has been replaced by underscores.
var LanguageByAliasMap = map[string]string{
	"1c_enterprise":  "1C Enterprise",
	"abap":           "ABAP",
	"abl":            "OpenEdge ABL",
	"abnf":           "ABNF",
	"abuild":         "Alpine Abuild",
	"acfm":           "Adobe Font Metrics",
	"aconf":          "ApacheConf",
	"actionscript":   "ActionScript",
	"actionscript3":  "ActionScript",
	"actionscript_3": "ActionScript",
	"ada":            "Ada",
	"ada2005":        "Ada",
	"ada95":          "Ada",
	"adobe_composite_font_metrics": "Adobe Font Metrics",
	"adobe_font_metrics":           "Adobe Font Metrics",
	"adobe_multiple_font_metrics":  "Adobe Font Metrics",
	"advpl":            "xBase",
	"afdko":            "OpenType Feature File",
	"agda":             "Agda",
	"ags":              "AGS Script",
	"ags_script":       "AGS Script",
	"ahk":              "AutoHotkey",
	"alloy":            "Alloy",
	"alpine_abuild":    "Alpine Abuild",
	"amfm":             "Adobe Font Metrics",
	"ampl":             "AMPL",
	"angelscript":      "AngelScript",
	"ant_build_system": "Ant Build System",
	"antlr":            "ANTLR",
	"apache":           "ApacheConf",
	"apacheconf":       "ApacheConf",
	"apex":             "Apex",
	"api_blueprint":    "API Blueprint",
	"apkbuild":         "Alpine Abuild",
	"apl":              "APL",
	"apollo_guidance_computer":       "Apollo Guidance Computer",
	"applescript":                    "AppleScript",
	"arc":                            "Arc",
	"arexx":                          "REXX",
	"as3":                            "ActionScript",
	"asciidoc":                       "AsciiDoc",
	"asm":                            "Assembly",
	"asn.1":                          "ASN.1",
	"asp":                            "ASP",
	"aspectj":                        "AspectJ",
	"aspx":                           "ASP",
	"aspx-vb":                        "ASP",
	"assembly":                       "Assembly",
	"asymptote":                      "Asymptote",
	"ats":                            "ATS",
	"ats2":                           "ATS",
	"au3":                            "AutoIt",
	"augeas":                         "Augeas",
	"autoconf":                       "M4Sugar",
	"autohotkey":                     "AutoHotkey",
	"autoit":                         "AutoIt",
	"autoit3":                        "AutoIt",
	"autoitscript":                   "AutoIt",
	"awk":                            "Awk",
	"b3d":                            "BlitzBasic",
	"ballerina":                      "Ballerina",
	"bash":                           "Shell",
	"bash_session":                   "ShellSession",
	"bat":                            "Batchfile",
	"batch":                          "Batchfile",
	"batchfile":                      "Batchfile",
	"befunge":                        "Befunge",
	"bison":                          "Bison",
	"bitbake":                        "BitBake",
	"blade":                          "Blade",
	"blitz3d":                        "BlitzBasic",
	"blitzbasic":                     "BlitzBasic",
	"blitzmax":                       "BlitzMax",
	"blitzplus":                      "BlitzBasic",
	"bluespec":                       "Bluespec",
	"bmax":                           "BlitzMax",
	"boo":                            "Boo",
	"bplus":                          "BlitzBasic",
	"brainfuck":                      "Brainfuck",
	"brightscript":                   "Brightscript",
	"bro":                            "Bro",
	"bsdmake":                        "Makefile",
	"byond":                          "DM",
	"c":                              "C",
	"c#":                             "C#",
	"c++":                            "C++",
	"c++-objdump":                    "Cpp-ObjDump",
	"c-objdump":                      "C-ObjDump",
	"c2hs":                           "C2hs Haskell",
	"c2hs_haskell":                   "C2hs Haskell",
	"cap'n_proto":                    "Cap'n Proto",
	"carto":                          "CartoCSS",
	"cartocss":                       "CartoCSS",
	"ceylon":                         "Ceylon",
	"cfc":                            "ColdFusion CFC",
	"cfm":                            "ColdFusion",
	"cfml":                           "ColdFusion",
	"chapel":                         "Chapel",
	"charity":                        "Charity",
	"chpl":                           "Chapel",
	"chuck":                          "ChucK",
	"cirru":                          "Cirru",
	"clarion":                        "Clarion",
	"clean":                          "Clean",
	"click":                          "Click",
	"clipper":                        "xBase",
	"clips":                          "CLIPS",
	"clojure":                        "Clojure",
	"closure_templates":              "Closure Templates",
	"cloud_firestore_security_rules": "Cloud Firestore Security Rules",
	"cmake":                    "CMake",
	"cobol":                    "COBOL",
	"coffee":                   "CoffeeScript",
	"coffee-script":            "CoffeeScript",
	"coffeescript":             "CoffeeScript",
	"coldfusion":               "ColdFusion",
	"coldfusion_cfc":           "ColdFusion CFC",
	"coldfusion_html":          "ColdFusion",
	"collada":                  "COLLADA",
	"common_lisp":              "Common Lisp",
	"common_workflow_language": "Common Workflow Language",
	"component_pascal":         "Component Pascal",
	"conll":                    "CoNLL-U",
	"conll-u":                  "CoNLL-U",
	"conll-x":                  "CoNLL-U",
	"console":                  "ShellSession",
	"cool":                     "Cool",
	"coq":                      "Coq",
	"cperl":                    "Perl",
	"cpp":                      "C++",
	"cpp-objdump":              "Cpp-ObjDump",
	"creole":                   "Creole",
	"crystal":                  "Crystal",
	"csharp":                   "C#",
	"cson":                     "CSON",
	"csound":                   "Csound",
	"csound-csd":               "Csound Document",
	"csound-orc":               "Csound",
	"csound-sco":               "Csound Score",
	"csound_document":          "Csound Document",
	"csound_score":             "Csound Score",
	"css":                      "CSS",
	"csv":                      "CSV",
	"cucumber":                 "Gherkin",
	"cuda":                     "Cuda",
	"cweb":                     "CWeb",
	"cwl":                      "Common Workflow Language",
	"cycript":                  "Cycript",
	"cython":                   "Cython",
	"d":                        "D",
	"d-objdump":                "D-ObjDump",
	"darcs_patch":              "Darcs Patch",
	"dart":                     "Dart",
	"dataweave":                "DataWeave",
	"dcl":                      "DIGITAL Command Language",
	"delphi":                   "Component Pascal",
	"desktop":                  "desktop",
	"diff":                     "Diff",
	"digital_command_language": "DIGITAL Command Language",
	"django":                   "HTML+Django",
	"dm":                       "DM",
	"dns_zone":                 "DNS Zone",
	"dockerfile":               "Dockerfile",
	"dogescript":               "Dogescript",
	"dosbatch":                 "Batchfile",
	"dosini":                   "INI",
	"dpatch":                   "Darcs Patch",
	"dtrace":                   "DTrace",
	"dtrace-script":            "DTrace",
	"dylan":                    "Dylan",
	"e":                        "E",
	"eagle":                    "Eagle",
	"easybuild":                "Easybuild",
	"ebnf":                     "EBNF",
	"ec":                       "eC",
	"ecere_projects":           "Ecere Projects",
	"ecl":                      "ECL",
	"eclipse":                  "ECLiPSe",
	"ecr":                      "HTML+ECR",
	"edje_data_collection": "Edje Data Collection",
	"edn":                 "edn",
	"eeschema_schematic":  "KiCad Schematic",
	"eex":                 "HTML+EEX",
	"eiffel":              "Eiffel",
	"ejs":                 "EJS",
	"elisp":               "Emacs Lisp",
	"elixir":              "Elixir",
	"elm":                 "Elm",
	"emacs":               "Emacs Lisp",
	"emacs_lisp":          "Emacs Lisp",
	"emberscript":         "EmberScript",
	"eml":                 "EML",
	"eq":                  "EQ",
	"erb":                 "HTML+ERB",
	"erlang":              "Erlang",
	"f#":                  "F#",
	"f*":                  "F*",
	"factor":              "Factor",
	"fancy":               "Fancy",
	"fantom":              "Fantom",
	"figfont":             "FIGlet Font",
	"figlet_font":         "FIGlet Font",
	"filebench_wml":       "Filebench WML",
	"filterscript":        "Filterscript",
	"fish":                "fish",
	"flex":                "Lex",
	"flux":                "FLUX",
	"formatted":           "Formatted",
	"forth":               "Forth",
	"fortran":             "Fortran",
	"foxpro":              "xBase",
	"freemarker":          "FreeMarker",
	"frege":               "Frege",
	"fsharp":              "F#",
	"fstar":               "F*",
	"ftl":                 "FreeMarker",
	"fundamental":         "Text",
	"g-code":              "G-code",
	"game_maker_language": "Game Maker Language",
	"gams":                "GAMS",
	"gap":                 "GAP",
	"gcc_machine_description": "GCC Machine Description",
	"gdb":             "GDB",
	"gdscript":        "GDScript",
	"genie":           "Genie",
	"genshi":          "Genshi",
	"gentoo_ebuild":   "Gentoo Ebuild",
	"gentoo_eclass":   "Gentoo Eclass",
	"gerber_image":    "Gerber Image",
	"gettext_catalog": "Gettext Catalog",
	"gf":              "Grammatical Framework",
	"gherkin":         "Gherkin",
	"git-ignore":      "Ignore List",
	"git_attributes":  "Git Attributes",
	"git_config":      "Git Config",
	"gitattributes":   "Git Attributes",
	"gitconfig":       "Git Config",
	"gitignore":       "Ignore List",
	"gitmodules":      "Git Config",
	"glsl":            "GLSL",
	"glyph":           "Glyph",
	"glyph_bitmap_distribution_format": "Glyph Bitmap Distribution Format",
	"gn":                      "GN",
	"gnuplot":                 "Gnuplot",
	"go":                      "Go",
	"golang":                  "Go",
	"golo":                    "Golo",
	"gosu":                    "Gosu",
	"grace":                   "Grace",
	"gradle":                  "Gradle",
	"grammatical_framework":   "Grammatical Framework",
	"graph_modeling_language": "Graph Modeling Language",
	"graphql":                 "GraphQL",
	"graphviz_(dot)":          "Graphviz (DOT)",
	"groff":                   "Roff",
	"groovy":                  "Groovy",
	"groovy_server_pages":     "Groovy Server Pages",
	"gsp":                      "Groovy Server Pages",
	"hack":                     "Hack",
	"haml":                     "Haml",
	"handlebars":               "Handlebars",
	"haproxy":                  "HAProxy",
	"harbour":                  "Harbour",
	"haskell":                  "Haskell",
	"haxe":                     "Haxe",
	"hbs":                      "Handlebars",
	"hcl":                      "HCL",
	"hiveql":                   "HiveQL",
	"hlsl":                     "HLSL",
	"html":                     "HTML",
	"html+django":              "HTML+Django",
	"html+django/jinja":        "HTML+Django",
	"html+ecr":                 "HTML+ECR",
	"html+eex":                 "HTML+EEX",
	"html+erb":                 "HTML+ERB",
	"html+jinja":               "HTML+Django",
	"html+php":                 "HTML+PHP",
	"html+razor":               "HTML+Razor",
	"html+ruby":                "RHTML",
	"htmlbars":                 "Handlebars",
	"htmldjango":               "HTML+Django",
	"http":                     "HTTP",
	"hxml":                     "HXML",
	"hy":                       "Hy",
	"hylang":                   "Hy",
	"hyphy":                    "HyPhy",
	"i7":                       "Inform 7",
	"idl":                      "IDL",
	"idris":                    "Idris",
	"ignore":                   "Ignore List",
	"ignore_list":              "Ignore List",
	"igor":                     "IGOR Pro",
	"igor_pro":                 "IGOR Pro",
	"igorpro":                  "IGOR Pro",
	"inc":                      "PHP",
	"inform7":                  "Inform 7",
	"inform_7":                 "Inform 7",
	"ini":                      "INI",
	"inno_setup":               "Inno Setup",
	"io":                       "Io",
	"ioke":                     "Ioke",
	"ipython_notebook":         "Jupyter Notebook",
	"irc":                      "IRC log",
	"irc_log":                  "IRC log",
	"irc_logs":                 "IRC log",
	"isabelle":                 "Isabelle",
	"isabelle_root":            "Isabelle ROOT",
	"j":                        "J",
	"jasmin":                   "Jasmin",
	"java":                     "Java",
	"java_properties":          "Java Properties",
	"java_server_page":         "Groovy Server Pages",
	"java_server_pages":        "Java Server Pages",
	"javascript":               "JavaScript",
	"jflex":                    "JFlex",
	"jison":                    "Jison",
	"jison_lex":                "Jison Lex",
	"jolie":                    "Jolie",
	"jruby":                    "Ruby",
	"js":                       "JavaScript",
	"json":                     "JSON",
	"json5":                    "JSON5",
	"json_with_comments":       "JSON with Comments",
	"jsonc":                    "JSON with Comments",
	"jsoniq":                   "JSONiq",
	"jsonld":                   "JSONLD",
	"jsp":                      "Java Server Pages",
	"jsx":                      "JSX",
	"julia":                    "Julia",
	"jupyter_notebook":         "Jupyter Notebook",
	"kicad_layout":             "KiCad Layout",
	"kicad_legacy_layout":      "KiCad Legacy Layout",
	"kicad_schematic":          "KiCad Schematic",
	"kit":                      "Kit",
	"kotlin":                   "Kotlin",
	"krl":                      "KRL",
	"labview":                  "LabVIEW",
	"lasso":                    "Lasso",
	"lassoscript":              "Lasso",
	"latex":                    "TeX",
	"latte":                    "Latte",
	"lean":                     "Lean",
	"less":                     "Less",
	"lex":                      "Lex",
	"lfe":                      "LFE",
	"lhaskell":                 "Literate Haskell",
	"lhs":                      "Literate Haskell",
	"lilypond":                 "LilyPond",
	"limbo":                    "Limbo",
	"linker_script":            "Linker Script",
	"linux_kernel_module":      "Linux Kernel Module",
	"liquid":                   "Liquid",
	"lisp":                     "Common Lisp",
	"litcoffee":                "Literate CoffeeScript",
	"literate_agda":            "Literate Agda",
	"literate_coffeescript":    "Literate CoffeeScript",
	"literate_haskell":         "Literate Haskell",
	"live-script":              "LiveScript",
	"livescript":               "LiveScript",
	"llvm":                     "LLVM",
	"logos":                    "Logos",
	"logtalk":                  "Logtalk",
	"lolcode":                  "LOLCODE",
	"lookml":                   "LookML",
	"loomscript":               "LoomScript",
	"ls":                       "LiveScript",
	"lsl":                      "LSL",
	"ltspice_symbol":           "LTspice Symbol",
	"lua":                      "Lua",
	"m":                        "M",
	"m4":                       "M4",
	"m4sugar":                  "M4Sugar",
	"macruby":                  "Ruby",
	"make":                     "Makefile",
	"makefile":                 "Makefile",
	"mako":                     "Mako",
	"man":                      "Roff",
	"man-page":                 "Roff",
	"man_page":                 "Roff",
	"manpage":                  "Roff",
	"markdown":                 "Markdown",
	"marko":                    "Marko",
	"markojs":                  "Marko",
	"mask":                     "Mask",
	"mathematica":              "Mathematica",
	"matlab":                   "MATLAB",
	"maven_pom":                "Maven POM",
	"max":                      "Max",
	"max/msp":                  "Max",
	"maxmsp":                   "Max",
	"maxscript":                "MAXScript",
	"mdoc":                     "Roff",
	"mediawiki":                "MediaWiki",
	"mercury":                  "Mercury",
	"meson":                    "Meson",
	"metal":                    "Metal",
	"mf":                       "Makefile",
	"minid":                    "MiniD",
	"mirah":                    "Mirah",
	"mma":                      "Mathematica",
	"modelica":                 "Modelica",
	"modula-2":                 "Modula-2",
	"modula-3":                 "Modula-3",
	"module_management_system": "Module Management System",
	"monkey":                   "Monkey",
	"moocode":                  "Moocode",
	"moonscript":               "MoonScript",
	"mql4":                     "MQL4",
	"mql5":                     "MQL5",
	"mtml":                     "MTML",
	"muf":                      "MUF",
	"mumps":                    "M",
	"mupad":                    "mupad",
	"myghty":                   "Myghty",
	"nanorc":                   "nanorc",
	"nasm":                     "Assembly",
	"ncl":                      "NCL",
	"nearley":                  "Nearley",
	"nemerle":                  "Nemerle",
	"nesc":                     "nesC",
	"netlinx":                  "NetLinx",
	"netlinx+erb":              "NetLinx+ERB",
	"netlogo":                  "NetLogo",
	"newlisp":                  "NewLisp",
	"nextflow":                 "Nextflow",
	"nginx":                    "Nginx",
	"nginx_configuration_file": "Nginx",
	"nim":                   "Nim",
	"ninja":                 "Ninja",
	"nit":                   "Nit",
	"nix":                   "Nix",
	"nixos":                 "Nix",
	"njk":                   "HTML+Django",
	"nl":                    "NL",
	"node":                  "JavaScript",
	"nroff":                 "Roff",
	"nsis":                  "NSIS",
	"nu":                    "Nu",
	"numpy":                 "NumPy",
	"nunjucks":              "HTML+Django",
	"nush":                  "Nu",
	"nvim":                  "Vim script",
	"obj-c":                 "Objective-C",
	"obj-c++":               "Objective-C++",
	"obj-j":                 "Objective-J",
	"objc":                  "Objective-C",
	"objc++":                "Objective-C++",
	"objdump":               "ObjDump",
	"objective-c":           "Objective-C",
	"objective-c++":         "Objective-C++",
	"objective-j":           "Objective-J",
	"objectivec":            "Objective-C",
	"objectivec++":          "Objective-C++",
	"objectivej":            "Objective-J",
	"objectpascal":          "Component Pascal",
	"objj":                  "Objective-J",
	"ocaml":                 "OCaml",
	"octave":                "MATLAB",
	"omgrofl":               "Omgrofl",
	"oncrpc":                "RPC",
	"ooc":                   "ooc",
	"opa":                   "Opa",
	"opal":                  "Opal",
	"opencl":                "OpenCL",
	"openedge":              "OpenEdge ABL",
	"openedge_abl":          "OpenEdge ABL",
	"openrc":                "OpenRC runscript",
	"openrc_runscript":      "OpenRC runscript",
	"openscad":              "OpenSCAD",
	"opentype_feature_file": "OpenType Feature File",
	"org":                            "Org",
	"osascript":                      "AppleScript",
	"ox":                             "Ox",
	"oxygene":                        "Oxygene",
	"oz":                             "Oz",
	"p4":                             "P4",
	"pan":                            "Pan",
	"pandoc":                         "Markdown",
	"papyrus":                        "Papyrus",
	"parrot":                         "Parrot",
	"parrot_assembly":                "Parrot Assembly",
	"parrot_internal_representation": "Parrot Internal Representation",
	"pascal":               "Pascal",
	"pasm":                 "Parrot Assembly",
	"pawn":                 "Pawn",
	"pcbnew":               "KiCad Layout",
	"pep8":                 "Pep8",
	"perl":                 "Perl",
	"perl6":                "Perl 6",
	"perl_6":               "Perl 6",
	"php":                  "PHP",
	"pic":                  "Pic",
	"pickle":               "Pickle",
	"picolisp":             "PicoLisp",
	"piglatin":             "PigLatin",
	"pike":                 "Pike",
	"pir":                  "Parrot Internal Representation",
	"plpgsql":              "PLpgSQL",
	"plsql":                "PLSQL",
	"pod":                  "Pod",
	"pod_6":                "Pod 6",
	"pogoscript":           "PogoScript",
	"pony":                 "Pony",
	"posh":                 "PowerShell",
	"postcss":              "PostCSS",
	"postscr":              "PostScript",
	"postscript":           "PostScript",
	"pot":                  "Gettext Catalog",
	"pov-ray":              "POV-Ray SDL",
	"pov-ray_sdl":          "POV-Ray SDL",
	"povray":               "POV-Ray SDL",
	"powerbuilder":         "PowerBuilder",
	"powershell":           "PowerShell",
	"processing":           "Processing",
	"progress":             "OpenEdge ABL",
	"prolog":               "Prolog",
	"propeller_spin":       "Propeller Spin",
	"protobuf":             "Protocol Buffer",
	"protocol_buffer":      "Protocol Buffer",
	"protocol_buffers":     "Protocol Buffer",
	"public_key":           "Public Key",
	"pug":                  "Pug",
	"puppet":               "Puppet",
	"pure_data":            "Pure Data",
	"purebasic":            "PureBasic",
	"purescript":           "PureScript",
	"pwsh":                 "PowerShell",
	"pycon":                "Python console",
	"pyrex":                "Cython",
	"python":               "Python",
	"python3":              "Python",
	"python_console":       "Python console",
	"python_traceback":     "Python traceback",
	"q":                    "q",
	"qmake":                "QMake",
	"qml":                  "QML",
	"quake":                "Quake",
	"r":                    "R",
	"racket":               "Racket",
	"ragel":                "Ragel",
	"ragel-rb":             "Ragel",
	"ragel-ruby":           "Ragel",
	"rake":                 "Ruby",
	"raml":                 "RAML",
	"rascal":               "Rascal",
	"raw":                  "Raw token data",
	"raw_token_data":       "Raw token data",
	"razor":                "HTML+Razor",
	"rb":                   "Ruby",
	"rbx":                  "Ruby",
	"rdoc":                 "RDoc",
	"realbasic":            "REALbasic",
	"reason":               "Reason",
	"rebol":                "Rebol",
	"red":                  "Red",
	"red/system":           "Red",
	"redcode":              "Redcode",
	"regex":                "Regular Expression",
	"regexp":               "Regular Expression",
	"regular_expression":   "Regular Expression",
	"ren'py":               "Ren'Py",
	"renderscript":         "RenderScript",
	"renpy":                "Ren'Py",
	"restructuredtext":     "reStructuredText",
	"rexx":                 "REXX",
	"rhtml":                "RHTML",
	"ring":                 "Ring",
	"rmarkdown":            "RMarkdown",
	"robotframework":       "RobotFramework",
	"roff":                 "Roff",
	"roff_manpage":         "Roff Manpage",
	"rouge":                "Rouge",
	"rpc":                  "RPC",
	"rpcgen":               "RPC",
	"rpm_spec":             "RPM Spec",
	"rs-274x":              "Gerber Image",
	"rscript":              "R",
	"rss":                  "XML",
	"rst":                  "reStructuredText",
	"ruby":                 "Ruby",
	"runoff":               "RUNOFF",
	"rust":                 "Rust",
	"rusthon":              "Python",
	"sage":                 "Sage",
	"salt":                 "SaltStack",
	"saltstack":            "SaltStack",
	"saltstate":            "SaltStack",
	"sas":                  "SAS",
	"sass":                 "Sass",
	"scala":                "Scala",
	"scaml":                "Scaml",
	"scheme":               "Scheme",
	"scilab":               "Scilab",
	"scss":                 "SCSS",
	"sed":                  "sed",
	"self":                 "Self",
	"sh":                   "Shell",
	"shaderlab":            "ShaderLab",
	"shell":                "Shell",
	"shell-script":         "Shell",
	"shellsession":         "ShellSession",
	"shen":                 "Shen",
	"slash":                "Slash",
	"slice":                "Slice",
	"slim":                 "Slim",
	"smali":                "Smali",
	"smalltalk":            "Smalltalk",
	"smarty":               "Smarty",
	"sml":                  "Standard ML",
	"smt":                  "SMT",
	"snippet":              "YASnippet",
	"solidity":             "Solidity",
	"sourcemod":            "SourcePawn",
	"sourcepawn":           "SourcePawn",
	"soy":                  "Closure Templates",
	"sparql":               "SPARQL",
	"specfile":             "RPM Spec",
	"spline_font_database": "Spline Font Database",
	"splus":                "R",
	"sqf":                  "SQF",
	"sql":                  "SQL",
	"sqlpl":                "SQLPL",
	"squeak":               "Smalltalk",
	"squirrel":             "Squirrel",
	"srecode_template":     "SRecode Template",
	"stan":                 "Stan",
	"standard_ml":          "Standard ML",
	"stata":                "Stata",
	"ston":                 "STON",
	"stylus":               "Stylus",
	"subrip_text":          "SubRip Text",
	"sugarss":              "SugarSS",
	"supercollider":        "SuperCollider",
	"svg":                  "SVG",
	"swift":                "Swift",
	"systemverilog":        "SystemVerilog",
	"tcl":                  "Tcl",
	"tcsh":                 "Tcsh",
	"tea":                  "Tea",
	"terra":                "Terra",
	"terraform":            "HCL",
	"tex":                  "TeX",
	"text":                 "Text",
	"textile":              "Textile",
	"thrift":               "Thrift",
	"ti_program":           "TI Program",
	"tl":                   "Type Language",
	"tla":                  "TLA",
	"toml":                 "TOML",
	"troff":                "Roff",
	"ts":                   "TypeScript",
	"turing":               "Turing",
	"turtle":               "Turtle",
	"twig":                 "Twig",
	"txl":                  "TXL",
	"type_language":        "Type Language",
	"typescript":           "TypeScript",
	"udiff":                "Diff",
	"unified_parallel_c":   "Unified Parallel C",
	"unity3d_asset":        "Unity3D Asset",
	"unix_assembly":        "Unix Assembly",
	"uno":                  "Uno",
	"unrealscript":         "UnrealScript",
	"ur":                   "UrWeb",
	"ur/web":               "UrWeb",
	"urweb":                "UrWeb",
	"vala":                 "Vala",
	"vb.net":               "Visual Basic",
	"vbnet":                "Visual Basic",
	"vcl":                  "VCL",
	"verilog":              "Verilog",
	"vhdl":                 "VHDL",
	"vim":                  "Vim script",
	"vim_script":           "Vim script",
	"viml":                 "Vim script",
	"visual_basic":         "Visual Basic",
	"volt":                 "Volt",
	"vue":                  "Vue",
	"wasm":                 "WebAssembly",
	"wast":                 "WebAssembly",
	"wavefront_material":   "Wavefront Material",
	"wavefront_object":     "Wavefront Object",
	"wdl":                  "wdl",
	"web_ontology_language":    "Web Ontology Language",
	"webassembly":              "WebAssembly",
	"webidl":                   "WebIDL",
	"winbatch":                 "Batchfile",
	"windows_registry_entries": "Windows Registry Entries",
	"wisp": "wisp",
	"world_of_warcraft_addon_data": "World of Warcraft Addon Data",
	"wsdl":                   "XML",
	"x10":                    "X10",
	"x_bitmap":               "X BitMap",
	"x_font_directory_index": "X Font Directory Index",
	"x_pixmap":               "X PixMap",
	"xbase":                  "xBase",
	"xbm":                    "X BitMap",
	"xc":                     "XC",
	"xcompose":               "XCompose",
	"xdr":                    "RPC",
	"xhtml":                  "HTML",
	"xml":                    "XML",
	"xml+genshi":             "Genshi",
	"xml+kid":                "Genshi",
	"xojo":                   "Xojo",
	"xpages":                 "XPages",
	"xpm":                    "X PixMap",
	"xproc":                  "XProc",
	"xquery":                 "XQuery",
	"xs":                     "XS",
	"xsd":                    "XML",
	"xsl":                    "XSLT",
	"xslt":                   "XSLT",
	"xten":                   "X10",
	"xtend":                  "Xtend",
	"yacc":                   "Yacc",
	"yaml":                   "YAML",
	"yang":                   "YANG",
	"yara":                   "YARA",
	"yas":                    "YASnippet",
	"yasnippet":              "YASnippet",
	"yml":                    "YAML",
	"zephir":                 "Zephir",
	"zig":                    "Zig",
	"zimpl":                  "Zimpl",
	"zsh":                    "Shell",
}

// LanguageByAlias looks up the language name by it's alias or name.
// It mirrors the logic of github linguist and is needed e.g for heuristcs.yml
// that mixes names and aliases in a language field (see XPM example).
func LanguageByAlias(langOrAlias string) (lang string, ok bool) {
	k := convertToAliasKey(langOrAlias)
	lang, ok = LanguageByAliasMap[k]
	return
}

// convertToAliasKey converts language name to a key in LanguageByAliasMap.
// Following
//  - internal.code-generator.generator.convertToAliasKey()
//  - GetLanguageByAlias()
// conventions.
// It is here to avoid dependency on "generate" and "enry" packages.
func convertToAliasKey(langName string) string {
	ak := strings.SplitN(langName, `,`, 2)[0]
	ak = strings.Replace(ak, ` `, `_`, -1)
	ak = strings.ToLower(ak)
	return ak
}