/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapRulesTemplate.txt found under go/resources folder
//Created @ 2023-08-18 15:18:19.365592 -0500 CDT m=+0.260106214

func BootstrapRules() []Rule {
    var BootstrapRules = []Rule{
        
            { Name: "SNAP-ETL-import", RuleType: "fire-once", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*.*%s.*$", Advice: "Vendor specific integration implementation", Effort: 100, Readiness: 0, Impact: "", Category: "etl", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "etl",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.talend", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "talend", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.odi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle-odi", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.is", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-is", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.sf.jasper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jasper", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.pentaho", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pentaho", Recipe: "", },
             }, },
        
            { Name: "SNAP-build-Ant-Maven", RuleType: "fire-once", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Align with standard build system", Effort: 0, Readiness: 0, Impact: "", Category: "buildSystem", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "build-system",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "build.xml", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "ant", Recipe: "", },
             { Type: "", Pattern: "", Value: "pom.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "maven", Recipe: "", },
             }, },
        
            { Name: "SNAP-build-Gradle", RuleType: "standard", FileType: "gradle$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Gradle is being used", Effort: 0, Readiness: 0, Impact: "", Category: "buildSystem", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "build.gradle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-SQL-properties", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Database coupling detected consider using ORM framework", Effort: 10, Readiness: 0, Impact: "", Category: "jdbc", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jdbc",}, { Value: "sql",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "StoredProcedureQuery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CallableStatement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-SQL", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Database coupling detected consider using ORM framework", Effort: 10, Readiness: 0, Impact: "", Category: "jdbc", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jdbc",}, { Value: "sql",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "StoredProcedureQuery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CallableStatement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-package-Gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "apply\\s*plugin:\\s*['']%s['']", Advice: "Application Server coupling detected.  Consider repackaging artifact as either war or executable jar", Effort: 100, Readiness: 0, Impact: "", Category: "packaging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ear", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-package-GradleJar", RuleType: "standard", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "[{] *%s", Advice: "Executable jar packaging is used", Effort: 0, Readiness: 0, Impact: "", Category: "packaging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-package-Maven-Ant", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application Server coupling detected", Effort: 100, Readiness: 0, Impact: "", Category: "packaging", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "maven",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<packaging>ear</packaging>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<target name=\"ear\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-ver-Maven-Ant", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Legacy Java detected.  Consider upgrading.", Effort: 100, Readiness: 0, Impact: "", Category: "java-ver", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "snap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<target>1.4</target>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java-14", Recipe: "", },
             { Type: "", Pattern: "", Value: "<target>1.5</target>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java-15", Recipe: "", },
             { Type: "", Pattern: "", Value: "<target>1.6</target>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java-16", Recipe: "", },
             { Type: "", Pattern: "", Value: "<javac target=\"1.4\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<javac target=\"1.5\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<javac target=\"1.6\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-RDS-connection-string-user-password-properties", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "relational database service connection string, username or password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "relational database service", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "relational database service",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\.|\\s)datasource\\.(.*\\.)?(url|jdbc-url|u\\-r\\-l)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)datasource\\.(.*\\.)?(username|user|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)r2dbc\\.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)r2dbc\\.(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)jdbc\\.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)jdbc\\.(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "=\\s*jdbc:.+:.+", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-RDS-connection-string-user-password-yaml", RuleType: "standard", FileType: "(yaml$|yml$|json$|jsn$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "relational database service connection string, username or password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "relational database service", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "relational database service",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$..datasource.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..datasource.username", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..datasource.password", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..datasource[*][\"jdbc-url\",\"url\",\"u-r-l\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..datasource[*][\"username\",\"user\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..r2dbc[\"url\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..r2dbc[\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..jdbc[\"url\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..jdbc[\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..[?(@=~/(?i)jdbc:.+:.+/)]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-apm-dependency-maven", RuleType: "standard", FileType: "xml$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "The application has integrated an Application Performance Management (APM) tool as a dependency", Effort: 0, Readiness: 0, Impact: "", Category: "apm", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "apm",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[contains(groupId, \"applicationinsights\") or contains(artifactId, \"applicationinsights\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[contains(groupId, \"newrelic\") or contains(artifactId, \"newrelic\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[contains(groupId, \"elastic.apm\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[contains(groupId, \"dynatrace.oneagent\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mq-connection-string-user-password-properties", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Message queue connection string, username or password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "message queue", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "message queue",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\.|\\s)kafka\\.(.*\\.)?bootstrap-servers", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)kafka\\.(.*\\.)?properties\\.(bootstrap\\.servers|bootstrap-servers)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)kafka\\.(.*\\.)?properties\\.sasl\\.jaas\\.config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)kafka\\.(.*\\.)?properties\\.schema\\.registry\\.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)kafka\\.(.*\\.)?properties\\.schema\\.registry\\.basic\\.auth\\.user\\.info", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)rabbitmq\\.(.*\\.)?(addresses|host|virtual-host)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)rabbitmq\\.(.*\\.)?(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)artemis\\.(broker-url)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)artemis\\.(user|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring\\.cloud\\.azure\\.(eventhub|eventhubs|servicebus)\\.(.*\\.)?(connection-string)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring\\.cloud\\.azure\\.(eventhub|eventhubs|servicebus)\\.(.*\\.)?credential\\.(password|username)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring\\.cloud\\.azure\\.(eventhub|eventhubs|servicebus)\\.(.*\\.)?credential\\.(client-id|client-secret)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring\\.jms\\.servicebus\\.(connection-string|password|username)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring\\.cloud\\.stream\\.rocketmq\\.binder.(name-server|access-key|secret-key)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mq-connection-string-user-password-yaml", RuleType: "standard", FileType: "(yaml$|yml$|json$|jsn$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "Message queue connection string, username or password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "message queue", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "message queue",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$..kafka.bootstrap-servers", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka[*][\"bootstrap-servers\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka.properties[\"bootstrap-servers\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka[*][\"properties\"][\"bootstrap-servers\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka.properties.bootstrap.servers", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka[*][\"properties\"][\"bootstrap\"][\"servers\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka.properties.sasl.jaas.config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka[*][\"properties\"][\"sasl\"][\"jaas\"][\"config\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka.properties.schema.registry.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka[*][\"properties\"][\"schema\"][\"registry\"][\"url\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka.properties.schema.registry.basic.auth.user.info", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..kafka[*][\"properties\"][\"schema\"][\"registry\"][\"basic\"][\"auth\"][\"user\"][\"info\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..rabbitmq[\"addresses\",\"host\",\"virtual-host\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..rabbitmq[*][\"addresses\",\"host\",\"virtual-host\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..rabbitmq[\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..rabbitmq[*][\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..artemis.broker-url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..artemis[\"user\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-non-lts-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(^|\\.|\\s)targetCompatibility(\\s|=)[^\\d]+(%s)", Advice: "The application is using non-LTS version Java. JDK on LTS version is recommended, i.e., JAVA_8, JAVA_11 or JAVA_17.", Effort: 6, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "java version",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "9", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "10", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "12", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "13", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "14", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "15", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "16", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "19", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-non-lts-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[starts-with(java.version,\"%s\")]/java.version", Advice: "The application is using non-LTS version Java. JDK on LTS version is recommended, i.e., JAVA_8, JAVA_11 or JAVA_17.", Effort: 6, Readiness: 0, Impact: "", Category: "java", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "java version",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "9", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "10", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "12", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "13", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "14", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "15", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "16", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "19", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-properties-config-client-configs", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "The application is using spring config server and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "config server", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "config server",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\s)spring\\.config\\.import", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\s)spring\\.cloud\\.config\\.uri", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-properties-eureka-client-configs", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "The application is using eureka and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "eureka", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "eureka",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\s)eureka\\.client\\.(service-url|serviceUrl)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-properties-logging-console-appender", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "regex", DefaultPattern: "", Advice: "Please enable logging to console", Effort: 5, Readiness: 0, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(?i)(appender.*=.*ConsoleAppender|appender.*\\.type.*=.*Console|(org/springframework/boot/logging/logback/base.xml)|(org/springframework/boot/logging/log4j2/log4j2.xml))", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-redis-connection-string-user-password-properties", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Redis connection string, username and password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "redis", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "redis",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\.|\\s)redis\\.(.*\\.)?(url|host|nodes)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)redis\\.(.*\\.)?(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)jedis\\.(url|host)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)jedis\\.(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)lettuce\\.(url|host)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)lettuce\\.(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-redis-connection-string-user-password-yaml", RuleType: "standard", FileType: "yaml$|yml$", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "Redis connection string, username and password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "redis", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "redis",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$..redis[\"url\",\"host\",\"nodes\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..redis[\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..redis[*][\"url\",\"host\",\"nodes\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..redis[*][\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..jedis[\"url\",\"host\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..jedis[\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..lettuce[\"url\",\"host\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$..lettuce[\"username\",\"password\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-scheduled-job-annotation", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "The application has scheduled jobs such as Quartz Scheduler tasks or cron jobs. Please be aware that after migrating to the cloud and scaling out, scheduled jobs in applications may run more than once per scheduled period and lead to unintended consequences.", Effort: 6, Readiness: 0, Impact: "", Category: "scheduled job", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "scheduled job",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Scheduled", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-scheduled-job-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import\\s*%s.*", Advice: "The application has scheduled jobs such as Quartz Scheduler tasks or cron jobs. Please be aware that after migrating to the cloud and scaling out, scheduled jobs in applications may run more than once per scheduled period and lead to unintended consequences.", Effort: 6, Readiness: 0, Impact: "file", Category: "scheduled job", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "scheduled job",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java\\.util\\.concurrent\\.ScheduledExecutorService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "java\\.util\\.concurrent\\.ScheduledFuture", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.quartz", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.apache\\.commons\\.scheduler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "com\\.google\\.common\\.util\\.concurrent\\.AbstractScheduledService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "com\\.google\\.common\\.util\\.concurrent\\.ListeningScheduledExecutorService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-oss-support-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Spring boot version is out of spring boot open source support support(https://spring.io/projects/spring-boot#support). If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org\\.springframework\\.boot(.*)([\\r\\n]*)version( *[:=]? *)[''\"]?(2\\.[4-6])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i)(springBootVersion)(.*)(2\\.[4-6])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.springframework\\.boot[''\"]?( *[:=]? *)[''\"]?(2\\.[4-6])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-oss-support-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[groupId=\"org.springframework.boot\" and starts-with(version,\"%s\")]/version", Advice: "Spring boot version is out of spring boot open source support support(https://spring.io/projects/spring-boot#support). If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "2.4", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.6", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-starter", RuleType: "standard", FileType: "gradle$|kts$|xml$", Target: "contents", Type: "regex", DefaultPattern: "", Advice: "The application has spring boot/spring cloud starter dependencies", Effort: 0, Readiness: 0, Impact: "", Category: "spring boot", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-boot-admin-starter-server", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-boot-starter-web|spring-boot-starter-webflux", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-boot-starter-undertow|spring-boot-starter-tomcat|spring-boot-starter-jetty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-starter-gateway", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-config-server", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-starter-netflix-eureka-server", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-support-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Spring boot version is out of any spring boot support(https://spring.io/projects/spring-boot#support). Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org\\.springframework\\.boot(.*)([\\r\\n]*)version( *[:=]? *)[''\"]?(2\\.[0-3])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i)(springBootVersion)(.*)(2\\.[0-3])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.springframework\\.boot[''\"]?( *[:=]? *)[''\"]?(2\\.[0-3])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-support-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[groupId=\"org.springframework.boot\" and starts-with(version,\"%s\")]/version", Advice: "Spring boot version is out of any spring boot support(https://spring.io/projects/spring-boot#support). Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "2.0", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.3", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Spring boot version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org\\.springframework\\.boot(.*)([\\r\\n]*)version( *[:=]? *)[''\"]?(1\\.)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i)(springBootVersion)(.*)(1\\.)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.springframework\\.boot[''\"]?( *[:=]? *)[''\"]?(1\\.)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[groupId=\"org.springframework.boot\" and starts-with(version,\"%s\")]/version", Advice: "Spring boot version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring boot",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-oss-support-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(?i)(springCloudVersion|org\\.springframework\\.cloud)(.*)(%s)", Advice: "Spring cloud version is out of open source support support. If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring cloud",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "2020\\.|ilford", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-oss-support-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "Spring cloud version is out of open source support support. If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring cloud",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^2020\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^ilford\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^2020\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^ilford\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-support-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(?i)(springCloudVersion|org\\.springframework\\.cloud)(.*)(%s)", Advice: "Spring cloud version is out of any support. Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring cloud",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Hoxton|Greenwich|Finchley", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-support-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "Spring cloud version is out of any support. Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring cloud",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^Hoxton\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^Greenwich\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^Finchley\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^Hoxton\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^Greenwich\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^Finchley\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(?i)(springCloudVersion|org\\.springframework\\.cloud)(.*)(%s)", Advice: "Spring cloud version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring cloud",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Edgware|Dalston", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "Spring cloud version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring cloud",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^Edgware\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[groupId=\"org.springframework.cloud\" and matches(version,\"(?i)^Dalston\")]/version", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^Edgware\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "//*[matches(name(), \"spring-?cloud.version\") and matches(text(),\"(?i)^Dalston\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-version-gradle", RuleType: "standard", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(^|\\.|\\s)targetCompatibility(\\s|=)[^\\d]+(%s)", Advice: "JDK version is found to be lower than JAVA_8", Effort: 50, Readiness: 0, Impact: "", Category: "version", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "java version",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1\\.[0-7]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-version-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[starts-with(java.version,\"%s\")]/java.version", Advice: "JDK version is found to be lower than JAVA_8", Effort: 50, Readiness: 0, Impact: "", Category: "java", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "java version",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.0", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.3", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.4", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.6", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.7", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-xml-logging-console-appender", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "regex", DefaultPattern: "", Advice: "Please enable logging to console", Effort: 5, Readiness: 0, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(?i)(ConsoleAppender|type(\\s)*=(\\s)*\"console\"|</console>|org/springframework/boot/logging/logback/base.xml|org/springframework/boot/logging/log4j2/log4j2.xml)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yaml-config-client-configs", RuleType: "standard", FileType: "(yaml$|yml$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "The application is using spring config server and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "config server", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "config server",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.spring.config.import", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$.spring.cloud.config.uri", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yaml-eureka-client-configs", RuleType: "standard", FileType: "(yaml$|yml$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "The application is using eureka and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "eureka", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "eureka",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.eureka.client[\"service-url\",\"serviceUrl\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yaml-json-log4j2-fileAppender", RuleType: "standard", FileType: "yaml$|yml$|json$|jsn$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "Replace file appender with console appender", Effort: 3, Readiness: 0, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "log4j2",}, { Value: "logging",}, { Value: "log2file",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$[*][\"appenders\",\"Appenders\"][?( @.File || @.file || @.appender[?(@.type=~/(?i)file/)])]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$[*][\"appenders\",\"Appenders\"][\"Routing\"][\"Routes\"][\"Route\"][?( @.File || @.file)]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$[*][\"appenders\",\"Appenders\"][?(@.RollingFile || @.rollingFile || @.appender[?(@.type=~/(?i)RollingFile/)])]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$[*][\"appenders\",\"Appenders\"][\"Routing\"][\"Routes\"][\"Route\"][?( @.RollingFile || @.rollingFile)]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yamljson-logging-console-appender", RuleType: "standard", FileType: "yaml$|yml$|json$|jsn$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "Please enable logging to console", Effort: 5, Readiness: 0, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$[*][\"appenders\",\"Appenders\"][?(@.console || @.Console || @.appender[?(@.type=~/(?i)console/)])]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-zipkin-maven", RuleType: "standard", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "The application uses Zipkin. Update the application to use Application Insights(https://learn.microsoft.com/en-us/azure/spring-apps/how-to-application-insights) instead if migrating to Azure", Effort: 5, Readiness: 0, Impact: "", Category: "zipkin", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "azure",}, { Value: "zipkin",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[contains(groupId, \"zipkin\") or contains(artifactId, \"zipkin\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "properties-port", RuleType: "standard", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "The application is setting the server port. Please be aware of potential port reliance issues during the migration process", Effort: 0, Readiness: 0, Impact: "", Category: "port", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "port",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\s)server\\.port", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "windows-dynamic-link-library", RuleType: "standard", FileType: "dll$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "This Dynamic-Link Library is Microsoft Windows platform dependent. It needs to be replaced with a Linux-style shared library", Effort: 10, Readiness: 0, Impact: "", Category: "os", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "os",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "dll$", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "windows-file-path", RuleType: "standard", FileType: "java$|properties$|xml$|yaml$|yml$|json$|jsn$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "This file system path is Microsoft Windows platform dependent. It needs to be replaced with a Linux-style path", Effort: 4, Readiness: 0, Impact: "", Category: "os", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "os",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(\"|''|`|\\s|^)(?:[a-zA-Z]\\:|\\\\\\\\[\\w\\s\\.]+\\\\[\\w\\s\\.$]+)([\\\\\\/][^\\n\\t]+)+", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "yaml-nonstandard-port", RuleType: "standard", FileType: "(yaml$|yml$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "The application is setting the server port. Please be aware of potential port reliance issues during the migration process", Effort: 0, Readiness: 0, Impact: "", Category: "port", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "port",}, },
            Profiles:
            []Profile{  },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.server.port", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "bootCDI", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "cdi", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cdi",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javaee-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javaee-api", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.inject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-inject", Recipe: "", },
             { Type: "", Pattern: "", Value: "cdi-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cdi-api", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.0_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.1_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.2_spec", Recipe: "", },
             }, },
        
            { Name: "bootEJB", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.0_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.1_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.2_spec", Recipe: "", },
             }, },
        
            { Name: "bootJAXWS", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "jaxws", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jax-ws",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jboss-annotations-api_1.3_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-annotations-1-3", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-servlet-api_4.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-servlet-4.0", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.enterprise.concurrent-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-enterprise-concurrent", Recipe: "", },
             }, },
        
            { Name: "bootJDBC", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 0, Readiness: 5, Impact: "", Category: "persistence", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring-boot",}, { Value: "jdbc",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.springframework.jdbc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-jdbc", Recipe: "", },
             }, },
        
            { Name: "bootJSF", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "jsf", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.faces-", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-faces", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsf-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jsf", Recipe: "", },
             { Type: "", Pattern: "", Value: "myfaces-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "myfaces", Recipe: "", },
             { Type: "", Pattern: "", Value: "myfaces-impl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "myfaces", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsf-impl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jsf", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-jsf-api_2.3_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-jsf-2.3", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-jsf-api_2.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-jsf-2.2", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-jsf-api_2.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-jsf-2.1", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-jsf-api_2.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-jsf-2.0", Recipe: "", },
             }, },
        
            { Name: "bootMDB", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "mdb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-0", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-1", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-2", Recipe: "", },
             }, },
        
            { Name: "bootSTRUTS", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts2-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts-2", Recipe: "", },
             }, },
        
            { Name: "bootTXN", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "txn", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "txn",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-0", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-1", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-2", Recipe: "", },
             }, },
        
            { Name: "bootWEBSOCKET", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 0, Readiness: 2, Impact: "", Category: "WebServlet", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-socket",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "websocket-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websocket", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.websocket-client-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websocket-client", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.websocket-all", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websocket-all", Recipe: "", },
             }, },
        
            { Name: "config-dotnet-webConfig", RuleType: "standard", FileType: "config$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Upgrade to .Net Core", Effort: 0, Readiness: 0, Impact: "", Category: "Config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-config",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web.config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Web.config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-encryption", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Use of encrypted sections is problematic in cloud environments. Configurations should be externalized as environment variables instead of encrypting them in configuration file.", Effort: 300, Readiness: 0, Impact: "", Category: "security", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "security",}, { Value: "encryption",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[@configProtectionProvider][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-security", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Use of generated machine keys is problematic", Effort: 100, Readiness: 100, Impact: "", Category: "security", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "auto-gen-key",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/machineKey[contains(@validationKey, \"AutoGenerate\") or contains(@decryptionKey, \"AutoGenerate\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-sessionState", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "App should be executed as a stateless process. User session information could be store in the database (Ex https://www.c-sharpcorner.com/article/configure-sql-server-session-state-in-asp-net-core/). A better approach is to use a Redis or SQL Server distributed cache, which doesn't require sticky sessions. For more information, see Distributed caching in ASP.NET Core (https://learn.microsoft.com/en-us/aspnet/core/performance/caching/distributed?view=aspnetcore-7.0).", Effort: 200, Readiness: 2, Impact: "", Category: "stateful", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "session",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/sessionState[@mode=\"InProc\" or @mode=\"StateServer\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-dockerFile", RuleType: "standard", FileType: "$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Determine if TKG is more prescriptive and TBS or Choreographer can be used to containerize", Effort: 5, Readiness: 1000, Impact: "", Category: "docker", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "docker",}, { Value: "tkg",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Dockerfile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-non-root-user", RuleType: "standard", FileType: "$", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Shows evidence of avoiding root privledges", Effort: -100, Readiness: 1000, Impact: "", Category: "dockerSecurity", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "non-root-user",}, { Value: "docker",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RUN groupadd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RUN useradd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "USER", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-sudo", RuleType: "standard", FileType: "$", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Using root inside a container is a serious vulnerability.", Effort: 100, Readiness: 10, Impact: "", Category: "docker", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "docker",}, { Value: "sudo",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RUN install -y sudo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RUN su root", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-FileCacheModule", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Caching should be managed externally. Relying on IIS Modules won't work in Linux environment. Alternatives are distributed cache solutions like Redis, SQL Server distributed cache, NCache distributed cache.", Effort: 1000, Readiness: 5, Impact: "", Category: "caching", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/location/system.webServer/serverRuntime[@frequentHitThreshold or @frequentHitTimePeriod][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-HttpCacheModule", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Output cache profiles etc. via configuration.", Effort: 1000, Readiness: 5, Impact: "", Category: "caching", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "cache-control-mode",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/staticContent/clientCache[@cacheControlMode][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-RequestFilteringModule", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 1, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "request-filtering",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/requestFiltering/requestLimits[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-StaticCompressionModule", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuration", Effort: 100, Readiness: 5, Impact: "", Category: "unsupported-iis-module", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "static-compression",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpCompression[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-StaticFileModule", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuation", Effort: 400, Readiness: 5, Impact: "", Category: "Unsupported modules", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "static-file-module",}, { Value: "web-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/handlers/add[contains(@modules, \"StaticFileModule\")][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.webServer/handlers/add[@name=\"StaticFileModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-TokenCacheModule", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Caches windows security tokens for password based authentication schemes (anonymous authentication; basic authentication; IIS client certificate authentication). Ensure compabible configuration.", Effort: 1000, Readiness: 5, Impact: "", Category: "unsupported-iis-module", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "cache",}, { Value: "token-cache-module",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/globalModules/add[@name=\"TokenCacheModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-WindowsAuth-config", RuleType: "fire-once", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows Integrated Authentication will be problematic in the cloud as containers won't be joining Windows Domains. Use another authentication mechanism more cloud friendly like SSO, LDAP Remote Calls, Local...", Effort: 500, Readiness: 1, Impact: "", Category: "Unsupported modules", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-auth",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//system.webServer/security/authentication/windowsAuthentication[translate(@enabled, \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\") = \"TRUE\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.serviceModel/bindings//binding/security/transport[translate(@clientCredentialType, \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\")=\"WINDOWS\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.web/authentication[translate(@mode, \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\")=\"WINDOWS\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/connectionStrings/add[contains(translate(translate(@connectionString, \" \", \"\"), \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\"), \"INTEGRATEDSECURITY=TRUE\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-WindowsAuth-csvb", RuleType: "fire-once", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Integrated Authentication will be problematic in the cloud as containers won't be joining Windows Domains. Use another authentication mechanism more cloud friendly like SSO, LDAP Remote Calls, Local...", Effort: 500, Readiness: 1, Impact: "", Category: "Unsupported modules", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-auth",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "GetSection\\(system.webServer/security/authentication/windowsAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "GetSection\\(\"system.webServer/security/authentication/windowsAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-connectionstrings", RuleType: "standard", FileType: "(json|property)", Target: "line", Type: "regex", DefaultPattern: ".*(%s)", Advice: "Remove connection strings from files, use environment variables (or mount configmap into pod in a docker like environment)", Effort: 10, Readiness: 5, Impact: "", Category: "connection-string", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "connection-string",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ConnectionStrings", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-database-access", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Remove connection strings from files, use environment variables (or mount configmap into pod in a docker like environment)", Effort: 100, Readiness: 10, Impact: "", Category: "database", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/connectionStrings[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-db2-unmanaged", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*\\s*%s$", Advice: "Refer to platform documentation", Effort: 10, Readiness: 9, Impact: "", Category: "database", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "IBM.Data.DB2", Advice: "IBM.Data.DB2 can require a special procedure so that the driver's native components are deployed with the application.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-serilog-elasticsearch", RuleType: "standard", FileType: "(cs|yaml|yml|json)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Make sure to have reachable ELK stack from deployed app", Effort: 5, Readiness: 8, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "elasticsearch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "using Serilog.Sinks.Elasticsearch", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ElasticConfiguration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-file-based-config", RuleType: "standard", FileType: "cs", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Externalize configuration to environment or use ConfigMap as file-mount into a K8S pod", Effort: 10, Readiness: 3, Impact: "", Category: "env-config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "env-config",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".AddJsonFile\\(\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-fileIO-read", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s*\\(.*", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 3, Readiness: 8, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "File.Open", Advice: "Calling File.Open", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-fileIO", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*\\(.*", Advice: "Relying on the local filesystem to store data is unreliable in a cloud platform. Containers are ephemerals and data could be lost. Data could be stored in database or the application should leverage File Sharing Systems like NFS Persistent Volume (PV) in Kubernetes environment.", Effort: 100, Readiness: 8, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "File.Append", Advice: "Appending to a file (File.Append*) - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Create", Advice: "Calling File.Create - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Move", Advice: "Calling File.Move - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.OpenWrite", Advice: "Calling File.OpenWrite - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Replace", Advice: "Calling File.Replace - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Set", Advice: "Setting File Metadata (File.Set*) - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Write", Advice: "Writing to a file (File.Write*) - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "new FileStream", Advice: "Direct construction of FileStream - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileSystemWatcher", Advice: "Use of FileSystemWatcher - Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-filepath", RuleType: "standard", FileType: "(cs|yaml|yml|json)", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform. Containers are ephemerals and data could be lost. Data could be stored in database or the application should leverage File Sharing Systems like NFS Persistent Volume (PV) in Kubernetes environment.", Effort: 8, Readiness: 3, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "\\sFilePath.*=", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.FilePath.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ipv4-addresses", RuleType: "standard", FileType: "(yaml|yml|cs$|json)", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Found hard-coded IPv4s. Make configurable, put into environment variables. Leverage config maps in a kubernetes like environment.", Effort: 3, Readiness: 8, Impact: "", Category: "hard-ip", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "hard-ip",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  { Value: ".GeneratedCodeAttribute\\(.*[,]",}, { Value: "AssemblyFileVersion\\(",}, { Value: "AssemblyVersion\\(",}, { Value: "Version=",}, { Value: "\"assemblyVersion\".*[:].\".*\"",}, { Value: "\"fileVersion\".*[:].\".*\"",}, { Value: "PackageReference.*Version=\"",}, { Value: "GeneratedCode\\(",}, { Value: "([12]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01]))",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "[^_](([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-launchProcess", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Launching additional processes within a container like environment is not recommended. In a targeted docker like environment it is generally recommended that you separate areas of concern by using one service per container.", Effort: 300, Readiness: 7, Impact: "", Category: "process-launch", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "process-launch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "[\\s=]Process[.]Start[(\\s.]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".WaitForExit\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".new\\sProcess\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-logging", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Logging to the Event Log is not recommended for cloud native apps. Write to or manage logfiles. Instead, each running process should write its event stream, unbuffered, to stdout. https://docs.lacunasoftware.com/en-us/articles/amplia/on-premises/windows/enable-stdout-log.html", Effort: 100, Readiness: 3, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, { Value: "eventlog",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EventLogTraceListener", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EventLog", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-logging-config", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Logging to the Event Log is not recommended for cloud native apps. write to or manage logfiles. Instead, each running process should write its event stream, unbuffered, to stdout. https://docs.lacunasoftware.com/en-us/articles/amplia/on-premises/windows/enable-stdout-log.html", Effort: 100, Readiness: 5, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, { Value: "eventlog",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.diagnostics/trace/listeners/add[@type=\"System.Diagnostics.EventLogTraceListener\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-oracle-umanaged", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Oracle unmanaged driver requires including binaries with app", Effort: 3, Readiness: 8, Impact: "", Category: "database", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "database",}, { Value: "oracle",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Oracle.DataAccess", Advice: "Oracle unmanaged driver requires including binaries with app -- can use buildpack.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-security", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Relying on Windows certificate stores is problematic in a cloud environment. Certificate stores could be shared via shared volume mounts. Ideally you should decouple certificates and SSL from your application. SSL could also be handled at the load balancer level.", Effort: 3, Readiness: 10, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, { Value: "certificate",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "X509Store", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "x509-store", Recipe: "", },
             }, },
        
            { Name: "dotnet-serilog", RuleType: "standard", FileType: "cs$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Logging with Serilog. Make sure not to log to file. Remote logging sinks need to be reachable on network. Logging to the Event Log is not recommended for cloud native apps. Write to or manage logfiles. Instead, each running process should write its event stream, unbuffered, to stdout.", Effort: 3, Readiness: 8, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Serilog.Sinks", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "serilog", Recipe: "", },
             { Type: "", Pattern: "", Value: ".UseSerilog()", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "serilog", Recipe: "", },
             }, },
        
            { Name: "dotnet-sharepoint", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "SharePoint is not supported on CloudFoundry.", Effort: 100, Readiness: 0, Impact: "", Category: "unsupported-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "sharepoint",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.SharePoint", Advice: "SharePoint is not supported on CloudFoundry.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-transactions", RuleType: "fire-once", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Potential use of distributed transactions which are unsupported", Effort: 100, Readiness: 10, Impact: "", Category: "transaction", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "distributed-transaction",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TransactionScope", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-bindings", RuleType: "standard", FileType: "config$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Refers to documentation when using WCF, some protocols might not be supported", Effort: 0, Readiness: 0, Impact: "", Category: "wcf", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<basic.+Binding>", Advice: "Refers to documentation when using WCF, some protocols might not be supported", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<mexHttp.+Binding>", Advice: "Refers to documentation when using WCF, some protocols might not be supported", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<webHttp.+Binding>", Advice: "Refers to documentation when using WCF, some protocols might not be supported", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-protocols-ws", RuleType: "fire-once", FileType: "config$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Unsupported protocols", Effort: 1000, Readiness: 0, Impact: "", Category: "wcf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<ws.+HttpBinding>", Advice: "Many features of WS* protocols are problematic in the cloud like distributed transactions and reliable sessions", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-protocols", RuleType: "fire-once", FileType: "config$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Unsupported protocols", Effort: 500, Readiness: 0, Impact: "", Category: "wcf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<udpBinding>", Advice: "UDP protocol not supported on PCF", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "udp", Recipe: "", },
             { Type: "", Pattern: "", Value: "<.*NamedPipeBinding>", Advice: "NamedPipe protocol not supported on PCF", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<.*MsmqBinding>", Advice: "Msmq protocol not supported on PCF", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "msmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "<net.+Binding>", Advice: "Non HTTP based protocols are either unsupported or require extensive refactoring when on PCF. TCP binding would require TCP Router to be configured and app to be self hosted (TCP-IIS activation not supported)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-ssl", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "When using HTTPS, terminate SSL at load balancer", Effort: 1, Readiness: 10, Impact: "", Category: "wcf", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "ssl",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BasicHttpSecurityMode.Transport", Advice: "Disable HTTPS at the container and allow the external load balancer to terminate SSL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SecurityMode.Transport", Advice: "Transport security at the container is not supported.  Disable transport security on the service", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mode=\"Transport\"", Advice: "Transport security at the container is not supported.  Disable transport security on the service", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-application-domain", RuleType: "fire-once", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported, For .NET Core, there is exactly one AppDomain. Isolation and unloading are provided through AssemblyLoadContext. Security boundaries should be provided by process boundaries and appropriate remoting techniques", Effort: 1000, Readiness: 9, Impact: "", Category: "windows-domain", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "app-domain",}, { Value: "security",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AppDomain.CreateDomain", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-code-access-security", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "unsupported", Effort: 10, Readiness: 9, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "CodeAccessPermission", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-components", RuleType: "fire-once", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported on TAS", Effort: 1000, Readiness: 9, Impact: "", Category: "", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "com+",}, { Value: "api",}, { Value: "distributed-transaction",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "GetObjectContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.Transactions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-presentation-foundation", RuleType: "fire-once", FileType: "xaml", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Desktop technology only available on Windows platform", Effort: 500, Readiness: 1, Impact: "", Category: "windows-desktop", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-desktop",}, { Value: "windows-wpf",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<Page\\sxmlns=\".*xaml\\/presentation\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsForms", RuleType: "fire-once", FileType: "(cs$|vb$|csproj$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Forms module is not supported in the cloud. Refactor to a web application.", Effort: 1000, Readiness: 5, Impact: "", Category: "Unsupported modules", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-forms",}, { Value: "windows-desktop",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Windows.Forms", Advice: "Windows Forms module is not supported. Refactor to a web application.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsPrincipal", RuleType: "standard", FileType: "(cs$|vb$|csproj$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Operations requiring a Windows domain are not supported in a container like environment. Features relying on Windows authentication will be problematic if targeting linux environment.", Effort: 100, Readiness: 8, Impact: "", Category: "Unsupported modules", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-principal",}, { Value: "windows-domain",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Security.Principal", Advice: "Operations requiring a Windows domain are not supported", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsRegistry", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "^.*%s.*", Advice: "External configurations should be made available by environment variables or some another external service (Property Management Service).", Effort: 100, Readiness: 2, Impact: "", Category: "windows-registry", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-registry",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RegistryKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry.CurrentUser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry.CurrentConfig", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry.CurrentLocalMachine", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry.Users", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry.GetValue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry.SetValue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsServices", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Traditional Windows services will not run on linux environments. Windows services should be turned into their own deployable console application.", Effort: 0, Readiness: 4, Impact: "", Category: "windows-services", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-services",}, { Value: "process",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "[:].*ServiceBase.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*ServiceController.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*System.ServiceProcess.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*[.]UseWindowsService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*[.]ConfigureServices", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "faas-meta", RuleType: "fire-once", FileType: "meta$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "App should be started in the shortest time possible", Effort: 200, Readiness: 2, Impact: "", Category: "boottime", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "faas",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/libraries[@count>13]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "hardcode-uri", RuleType: "standard", FileType: "(java$|^vb$|py$|go$|aspx$|^c$|h$|cs$|csx$|cpp$|cob$|cfm$|cfml$|dockerfile$|jsp$|php$|^r$|^ts$|yaml$|yml$|json$)", Target: "line", Type: "regex", DefaultPattern: "(?:(?:^|[^\"])|(?:^|[^ =])\"|(?:^|[^:]) \"|(?:^|[^s])=\"|(?:^|[^\"]): \"|(?:^|[^n])s=\"|(?:^|[^d])\": \"|(?:^|[^l])ns=\"|(?:^|[^e])d\": \"|(?:^|[^m])lns=\"|(?:^|[^v])ed\": \"|(?:^|[^x])mlns=\"|(?:^|[^l])ved\": \"|(?:^|[^o])lved\": \"|(?:^|[^s])olved\": \"|(?:^|[^e])solved\": \"|(?:^|[^r])esolved\": \")(%s)\\:{1}\\/{2}[a-zA-Z0-9]", Advice: "Found hard-coded URI. Make configurable, put into environment or config map", Effort: 3, Readiness: 8, Impact: "", Category: "env-config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "hardcoded-uri",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  { Value: "xmlns=\"(http)\\:{1}\\/{2}[a-zA-Z0-9]",}, { Value: "(https)\\:{1}\\/{2}registry.npmjs.org/@angular/",}, { Value: "<meta.*content=.*",}, { Value: "<!doctype.*",}, { Value: "<!DOCTYPE.*",}, { Value: "http://activemq.apache.org",}, { Value: "http://www.apache.org",}, { Value: "<!-- .*-->",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "http", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "https", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "file", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "sftp", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ftp", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-3rdPartyImports", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 0, Readiness: 0, Impact: "", Category: "third-party", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "third-party",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ehcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caching", Recipe: "", },
             { Type: "", Pattern: "", Value: "47deg", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "advantageous", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "reflection", Recipe: "", },
             { Type: "", Pattern: "", Value: "agileclick", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "orm", Recipe: "", },
             { Type: "", Pattern: "", Value: "agilej", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "codeanalysis", Recipe: "", },
             { Type: "", Pattern: "", Value: "akka", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "multi-threaded", Recipe: "", },
             { Type: "", Pattern: "", Value: "alibaba", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring", Recipe: "", },
             { Type: "", Pattern: "", Value: "androidtransfuse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "angular", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "browser", Recipe: "", },
             { Type: "", Pattern: "", Value: "apache.commons", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "apitrary", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "orm", Recipe: "", },
             { Type: "", Pattern: "", Value: "apptik", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rest", Recipe: "", },
             { Type: "", Pattern: "", Value: "aranea-apps", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "argonaut", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "asm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "aspectj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "aop", Recipe: "", },
             { Type: "", Pattern: "", Value: "asynchttpclient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "http", Recipe: "", },
             { Type: "", Pattern: "", Value: "atmosphere", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "event-driven", Recipe: "", },
             { Type: "", Pattern: "", Value: "attoparser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "avast", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "aws-java", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "s3-storage", Recipe: "", },
             { Type: "", Pattern: "", Value: "axion", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "webservices", Recipe: "", },
             { Type: "", Pattern: "", Value: "azure", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cloud", Recipe: "", },
             { Type: "", Pattern: "", Value: "baracus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "clojure", Recipe: "", },
             { Type: "", Pattern: "", Value: "bcel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "beanshell", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "interpreter", Recipe: "", },
             { Type: "", Pattern: "", Value: "bluelinelabs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "brightify", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "bsf", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "beanscripting", Recipe: "", },
             { Type: "", Pattern: "", Value: "bytebuddy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "caffeine", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caching", Recipe: "", },
             { Type: "", Pattern: "", Value: "cassandra", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "ccil", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osgi", Recipe: "", },
             { Type: "", Pattern: "", Value: "celum", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             { Type: "", Pattern: "", Value: "cemerick", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "clojure", Recipe: "", },
             { Type: "", Pattern: "", Value: "ceylon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ceylon", Recipe: "", },
             { Type: "", Pattern: "", Value: "cglib", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "circe", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "circumflex", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "clojure", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "clojure", Recipe: "", },
             { Type: "", Pattern: "", Value: "cojen", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "codegenerator", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.m3", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "commonjava", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "maven", Recipe: "", },
             { Type: "", Pattern: "", Value: "commons-configuration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "config", Recipe: "", },
             { Type: "", Pattern: "", Value: "commons-digester", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml", Recipe: "", },
             { Type: "", Pattern: "", Value: "commons-logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "constretto", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "config", Recipe: "", },
             { Type: "", Pattern: "", Value: "critero", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commerce", Recipe: "", },
             { Type: "", Pattern: "", Value: "cyrusinnovation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commercial", Recipe: "", },
             { Type: "", Pattern: "", Value: "databene", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "datanucleus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "persistence", Recipe: "", },
             { Type: "", Pattern: "", Value: "datoin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "http", Recipe: "", },
             { Type: "", Pattern: "", Value: "modeltools", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "statistics", Recipe: "", },
             { Type: "", Pattern: "", Value: "dropwizard", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rest", Recipe: "", },
             { Type: "", Pattern: "", Value: "easygson", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "easymetrics", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "easymock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "ebay", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commerce", Recipe: "", },
             { Type: "", Pattern: "", Value: "eclipse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ide", Recipe: "", },
             { Type: "", Pattern: "", Value: "eclipsesource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rest", Recipe: "", },
             { Type: "", Pattern: "", Value: "esotericsoftware", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "serialization", Recipe: "", },
             { Type: "", Pattern: "", Value: "everit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "exparity", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "beans", Recipe: "", },
             { Type: "", Pattern: "", Value: "fasterxml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml", Recipe: "", },
             { Type: "", Pattern: "", Value: "flexjson", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "foursquare", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commercial", Recipe: "", },
             { Type: "", Pattern: "", Value: "gatling", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "github", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "repo", Recipe: "", },
             { Type: "", Pattern: "", Value: "glassfish", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jee", Recipe: "", },
             { Type: "", Pattern: "", Value: "gmock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "goggle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cloud", Recipe: "", },
             { Type: "", Pattern: "", Value: "grails", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "browser", Recipe: "", },
             { Type: "", Pattern: "", Value: "greenrobot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "event-driven", Recipe: "", },
             { Type: "", Pattern: "", Value: "groovy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "groovy", Recipe: "", },
             { Type: "", Pattern: "", Value: "grouplens", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "guava", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "h2database", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "helger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "hibernate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "orm", Recipe: "", },
             { Type: "", Pattern: "", Value: "htmlcleaner", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "http4s", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "http", Recipe: "", },
             { Type: "", Pattern: "", Value: "hydra", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "ini4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "config", Recipe: "", },
             { Type: "", Pattern: "", Value: "inject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dependency-injection", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.spray", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.wcm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "itextpdf", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pdf", Recipe: "", },
             { Type: "", Pattern: "", Value: "j256", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "orm", Recipe: "", },
             { Type: "", Pattern: "", Value: "j4Unit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "jamon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "jasonjson", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "javaclub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "orm", Recipe: "", },
             { Type: "", Pattern: "", Value: "javalite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "protocol-buffers", Recipe: "", },
             { Type: "", Pattern: "", Value: "javassist", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jee", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caching", Recipe: "", },
             { Type: "", Pattern: "", Value: "jcabi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "jcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caching", Recipe: "", },
             { Type: "", Pattern: "", Value: "jclouds", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cloud", Recipe: "", },
             { Type: "", Pattern: "", Value: "jconfig", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "config", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdev", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "maven", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmemcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caching", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmetrix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmockring", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "joda-time", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "date-time", Recipe: "", },
             { Type: "", Pattern: "", Value: "jodd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "math", Recipe: "", },
             { Type: "", Pattern: "", Value: "joestelmach", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "date-time", Recipe: "", },
             { Type: "", Pattern: "", Value: "jpox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "persistence", Recipe: "", },
             { Type: "", Pattern: "", Value: "jruby", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ruby", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsog", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "json4s", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsonbuddy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsoup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "jtidy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "jvnet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mime", Recipe: "", },
             { Type: "", Pattern: "", Value: "kasource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "kie", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bpm", Recipe: "", },
             { Type: "", Pattern: "", Value: "kodein", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "kohesive", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "kohsuke", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dcom", Recipe: "", },
             { Type: "", Pattern: "", Value: "kotlin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "ktor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "librato", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "liftweb", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "lihaoyi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "littleshoot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mulimedia", Recipe: "", },
             { Type: "", Pattern: "", Value: "log4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "log4s", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "logback", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "loopj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "lowagie", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pdf", Recipe: "", },
             { Type: "", Pattern: "", Value: "lyft", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "madgag", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "mashape", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "mchange", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             { Type: "", Pattern: "", Value: "mckoi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "metrics4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metics", Recipe: "", },
             { Type: "", Pattern: "", Value: "minidev", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockachino", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockejb", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockito", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockobjects", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockrunner", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "mongo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "mortbay", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jetty", Recipe: "", },
             { Type: "", Pattern: "", Value: "mozilla", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "browser", Recipe: "", },
             { Type: "", Pattern: "", Value: "mybatis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "nekohtml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "netflix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix", Recipe: "", },
             { Type: "", Pattern: "", Value: "ning", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "http", Recipe: "", },
             { Type: "", Pattern: "", Value: "nscala", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "objectos", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "assertions", Recipe: "", },
             { Type: "", Pattern: "", Value: "objectweb", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "ocpsoft", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "opensynphony", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jee", Recipe: "", },
             { Type: "", Pattern: "", Value: "ops4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.json", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "outr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "ow2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "p6spy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "parfait", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "perf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "picocontainer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dependency-injection", Recipe: "", },
             { Type: "", Pattern: "", Value: "plausible", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "R-language", Recipe: "", },
             { Type: "", Pattern: "", Value: "plexus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "pnuts", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scripting", Recipe: "", },
             { Type: "", Pattern: "", Value: "pojava", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "datetime", Recipe: "", },
             { Type: "", Pattern: "", Value: "polyforms", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dependency-injection", Recipe: "", },
             { Type: "", Pattern: "", Value: "powermock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "proofpoint", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rest", Recipe: "", },
             { Type: "", Pattern: "", Value: "proxool", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "rabbitmq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "messaging", Recipe: "", },
             { Type: "", Pattern: "", Value: "redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "rhq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rhq", Recipe: "", },
             { Type: "", Pattern: "", Value: "roboguice", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dependency-injection", Recipe: "", },
             { Type: "", Pattern: "", Value: "rubicon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "saliman", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "gradle", Recipe: "", },
             { Type: "", Pattern: "", Value: "scala-lang", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "scala-tools", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "scalaj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "scalamock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "scaldi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "sejda", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pdf", Recipe: "", },
             { Type: "", Pattern: "", Value: "sf.json-lib", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "sharegov", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "browser", Recipe: "", },
             { Type: "", Pattern: "", Value: "sirona", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "skinny-framework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "slf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "slf4j-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "loggin", Recipe: "", },
             { Type: "", Pattern: "", Value: "socialmetrix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "email", Recipe: "", },
             { Type: "", Pattern: "", Value: "softsmithy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "softwaremill", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commercial", Recipe: "", },
             { Type: "", Pattern: "", Value: "sorm-framework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "soywiz", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "specs2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring", Recipe: "", },
             { Type: "", Pattern: "", Value: "squareup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "http", Recipe: "", },
             { Type: "", Pattern: "", Value: "squeryl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "stackmob", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "security", Recipe: "", },
             { Type: "", Pattern: "", Value: "stickycode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring", Recipe: "", },
             { Type: "", Pattern: "", Value: "streametry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "sun", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sun", Recipe: "", },
             { Type: "", Pattern: "", Value: "sun.mail", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "email", Recipe: "", },
             { Type: "", Pattern: "", Value: "swinglabs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "swing", Recipe: "", },
             { Type: "", Pattern: "", Value: "tedhi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "datatime", Recipe: "", },
             { Type: "", Pattern: "", Value: "tehuti", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "time4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "datetime", Recipe: "", },
             { Type: "", Pattern: "", Value: "tophe", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "android", Recipe: "", },
             { Type: "", Pattern: "", Value: "tornado", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "clojure", Recipe: "", },
             { Type: "", Pattern: "", Value: "toucanpdf", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pdf", Recipe: "", },
             { Type: "", Pattern: "", Value: "trove4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "collections", Recipe: "", },
             { Type: "", Pattern: "", Value: "truecommons", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "utility", Recipe: "", },
             { Type: "", Pattern: "", Value: "twitter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "twitter", Recipe: "", },
             { Type: "", Pattern: "", Value: "typesafe", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "akka", Recipe: "", },
             { Type: "", Pattern: "", Value: "uberfire", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "ujoframework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "database", Recipe: "", },
             { Type: "", Pattern: "", Value: "ujorm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "html", Recipe: "", },
             { Type: "", Pattern: "", Value: "unboundid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ldap", Recipe: "", },
             { Type: "", Pattern: "", Value: "unimi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "collections", Recipe: "", },
             { Type: "", Pattern: "", Value: "uniscala", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "vaadin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "vertx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "vibur", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             { Type: "", Pattern: "", Value: "whirlycott", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caching", Recipe: "", },
             { Type: "", Pattern: "", Value: "wix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "wvlet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala", Recipe: "", },
             { Type: "", Pattern: "", Value: "yammer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "ymock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "testing", Recipe: "", },
             { Type: "", Pattern: "", Value: "zalando", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "zappos", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json", Recipe: "", },
             { Type: "", Pattern: "", Value: "zaxxer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             }, },
        
            { Name: "java-3rdPartySecurity-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 7, Readiness: 10, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.netegrity.policyserver.smapi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netegrity", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netegrity.sdk.apiutil", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netegrity", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.tspm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-tspm", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.oblix.access", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oblix", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netegrity.sdk.imspolicyapi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netegrity", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netegrity.sdk.policyapi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netegrity", Recipe: "", },
             { Type: "", Pattern: "", Value: "netegrity.siteminder.javaagent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netegrity", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.rsa", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rsa", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.oracle.am.asdk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle-am-asdk", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netegrity.sdk.dmsapi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netegrity", Recipe: "", },
             }, },
        
            { Name: "java-MBeans", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\:{1}", Advice: "MBean is application server specific, change to refactor using cloud friendly technology", Effort: 100, Readiness: 10, Impact: "", Category: "websphere", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "websphere",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WebSphere:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-processexit", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[.]%s[ (].*$", Advice: "Refer to IBM documentation", Effort: 9, Readiness: 10, Impact: "", Category: "process-exit", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "process-exit",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "exit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "halt", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-activemq", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: ".*[ .]%s[ (].*", Advice: "Remediate any persistence issues", Effort: 7, Readiness: 7, Impact: "", Category: "activemq", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "activemq",}, { Value: "message-queue",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "createActiveMQConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-alarmD-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Integrate with new alarmD service or migrate to Spring Boot scheduler", Effort: 5, Readiness: 0, Impact: "", Category: "ThirdParty", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "alarm-d",}, { Value: "scheduler",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.tfcci.ucs.manager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-apacheFop-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "[ .]%s[ .({]", Advice: "Usage requires configuration remdiation", Effort: 5, Readiness: 10, Impact: "", Category: "thirdParty", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "apache-fop",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FopFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Fop", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FopFactoryBuilder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FOUserAgent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setTargetResolution", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setDocumentHandlerOverride", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setFOEventHandlerOverride", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-batch", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Remove jndi provider or move to TKG", Effort: 10, Readiness: 7, Impact: "", Category: "jndi", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "batch",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "NamingEnumeration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameClassPair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameParser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DirContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BinaryRefAddr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompositeName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompoundName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InitialContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LinkRef", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameClassPair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RefAddr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StringRefAddr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BasicControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ControlFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InitialLdapContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LdapName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ManageReferralControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PagedResultsControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PagedResultsResponseControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SortControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SortKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SortResponseControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StartTlsRequest", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StartTlsResponse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnsolicitedNotificationEvent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamingEvent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamingExceptionEvent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-batchAnnotations", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Batch processing can include long running processes consider using k8s job scheduler", Effort: 10, Readiness: 7, Impact: "", Category: "batch", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "batch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BatchProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-cache-dist-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Distributed caches must be remediated to function in K8S", Effort: 50, Readiness: 10, Impact: "", Category: "distcache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "cache",}, { Value: "dist-cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.ehcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ehcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.hazelcast", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hazelcast", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.ignite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-ignite", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.infinispan", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "infinispan", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.spy.memcached", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "net-spy-memcached", Recipe: "", },
             }, },
        
            { Name: "java-cache-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Cloud readiness issue as potential state information that is not persisted to a backing service", Effort: 50, Readiness: 10, Impact: "", Category: "nondistcache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "non-dist-cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.tangosol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tangosol", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.commons-jcs-jcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcs-jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.websphere.cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websphere-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.jboss.ha.cachemanager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ha-cachemanager", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.spy.memcached", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "net-spy-memcached", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.opensymphony.oscache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "opensymphony-oscache", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.shiftone-cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "shiftone-casche", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.websphere.objectgrid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websphere-objectgrid", Recipe: "", },
             }, },
        
            { Name: "java-corba", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: ".*[ .]%s[ (].*", Advice: "Replace with cloud-friendly framework or move to TKG", Effort: 10, Readiness: 6, Impact: "", Category: "corba", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "rpc",}, { Value: "corba",}, { Value: "non-standard-protocol",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AnyHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AnySeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AnySeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BooleanHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BooleanSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BooleanSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ByteHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CharHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CharSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CharSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompletionStatus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompletionStatusHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ContextList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CurrentHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CurrentHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DefinitionKind", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DefinitionKindHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DoubleHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DoubleSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DoubleSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DynamicImplementatio", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ExceptionList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FieldNameHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FixedHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FloatHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FloatSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FloatSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "IdentifierHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "IDLTypeHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "IntHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LocalObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LongHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LongLongSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LongLongSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LongSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LongSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamedValue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameValuePair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameValuePairHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NVList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ObjectHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ObjectHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "OctetSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "OctetSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ParameterMode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ParameterModeHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ParameterModeHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PolicyErrorCodeHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PolicyErrorHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PolicyHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PolicyListHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PolicyListHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PolicyTypeHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PrincipalHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RepositoryIdHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerRequest", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceDetail", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceDetailHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceInformation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceInformationHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceInformationHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SetOverrideType", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SetOverrideTypeHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ShortHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ShortSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ShortSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StringHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StringSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StringSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StringValueHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StructMember", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StructMemberHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TCKind", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TypeCode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TypeCodeHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ULongLongSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ULongLongSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ULongSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ULongSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnionMember", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnionMemberHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnknownUserExceptionHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnknownUserExceptionHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UShortSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UShortSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ValueBaseHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ValueBaseHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ValueMember", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ValueMemberHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "VersionSpecHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "VisibilityHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WCharSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WCharSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WrongTransactionHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WrongTransactionHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WStringSeqHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WStringSeqHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WStringValueHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ehcache", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consider to externalize cache if not already", Effort: 100, Readiness: 6, Impact: "", Category: "cache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "ehcache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "net.sf.ehcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-invocation", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 10, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, { Value: "ejb",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJB\\(lookup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-mdb", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult MDB documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "mdb",}, { Value: "messaging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb.MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.ejb.ActivationConfigProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-rmi", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Removing RMI calls from client applications.", Effort: 10, Readiness: 4, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateless",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJBContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBLocalHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBLocalObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBMetaData", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EnterpriseBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HomeHandle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageDrivenBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageDrivenContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionSynchronization", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TimedObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TimerHandle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TimerService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMISecurityManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PortableRemoteObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "exportObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "unexportObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AccessException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AlreadyBoundException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectIOException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NoSuchObjectException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NotBoundException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMISecurityException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerRuntimeException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StubNotFoundException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnexpectedException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnmarshalException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationInstantiator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationMonitor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationSystem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Activatable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationDesc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroup_Stub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroupDesc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroupDesc.CommandEnvironment", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroupID", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationID", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LocateRegistry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LoaderHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteCall", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteRef", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIClientSocketFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIFailureHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIServerSocketFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LogStream", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ObjID", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteObjectInvocationHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteServer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteStub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIClassLoader", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIClassLoaderSpi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMISocketFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnicastRemoteObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateful-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateful",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb.Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateful", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Refer EJB stateful/stateless documentation", Effort: 100, Readiness: 100, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "ejb",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StatefulTimeout", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateless", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Removing RMI calls from client applications.", Effort: 10, Readiness: 4, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateless",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJBContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBLocalHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBLocalObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBMetaData", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EnterpriseBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HomeHandle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageDrivenBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageDrivenContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionSynchronization", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TimedObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TimerHandle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TimerService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMISecurityManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PortableRemoteObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "exportObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "unexportObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AccessException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AlreadyBoundException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectIOException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NoSuchObjectException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NotBoundException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMISecurityException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerRuntimeException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StubNotFoundException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnexpectedException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnmarshalException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationInstantiator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationMonitor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationSystem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Activatable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationDesc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroup_Stub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroupDesc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroupDesc.CommandEnvironment", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationGroupID", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationID", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LocateRegistry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LoaderHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteCall", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteRef", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIClientSocketFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIFailureHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIServerSocketFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LogStream", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ObjID", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteObjectInvocationHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteServer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteStub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIClassLoader", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMIClassLoaderSpi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RMISocketFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnicastRemoteObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow-Annotations", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Batch processing can include long running processes", Effort: 10, Readiness: 7, Impact: "", Category: "batch", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FlowScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Review usage to determine how the customized state of JSF flow is being used and determined if it can be externalized.", Effort: 100, Readiness: 100, Impact: "", Category: "jsf-flow", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.faces.flow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Review usage to determine how the customized state of JSF flow is being used and determined if it can be externalized.", Effort: 10, Readiness: 10, Impact: "", Category: "jsf-flow", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FlowCallNode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FlowHandlerFactoryWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-fileIO", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Move to cloud friendly alternative storage service", Effort: 8, Readiness: 8, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, { Value: "file storage",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getNextEntry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "putNextEntry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "canExecute", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "canRead", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "canWrite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "createNewFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getAbsoluteFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getAbsolutePath", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getCanonicalFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getCanonicalPath", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getFreeSpace", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getTotalSpace", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getUsableSpace", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "isDirectory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "isFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "isHidden", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "lastModified", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "listFiles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mkdir", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mkdirs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "renameTo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setExecutable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setLastModified", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setReadOnly", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setReadable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setWritable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileDescriptor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileInputStream", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileOutputStream", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FilePermission", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileReader", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileWriter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LineNumberInputStream", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LineNumberReader", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RandomAccessFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-file-system", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\:{1}\\/{2}", Advice: "Use backing storage service", Effort: 100, Readiness: 10, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "file", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-glassfish-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Refer to Weblogic documentation", Effort: 100, Readiness: 10, Impact: "", Category: "oracle", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "glassfish",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.bea.common.security.xacml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bea", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.httppubsub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bea", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bea", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.security.saml2.providers", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "saml2", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.security.saml2.providers.registry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "saml2", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp.jdbc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp.jdbc.oracle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "weblogic", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogicx.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "weblogic", Recipe: "", },
             }, },
        
            { Name: "java-handles-term", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "For containerization, the TERM signal must be handled, this pattern is a positive finding.", Effort: 1, Readiness: 6, Impact: "", Category: "term-signal", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "term",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Runtime.getRuntime().addShutdownHook", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-hardIP", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Hardcoded IP addresses are problematic in K8S", Effort: 1, Readiness: 8, Impact: "", Category: "hard-ip", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "hard-ip",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-hazelcast", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Make sure Cache instance is externalized", Effort: 100, Readiness: 6, Impact: "", Category: "cache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "hazelcast",}, { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "HazelcastInstance", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NearCacheClientSupport", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Hazelcast", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueryCacheConfig", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HazelcastClient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "newHazelcastClient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setMapEvictionPolicy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MapEvictionPolicy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "newHazelcastInstance", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-iop", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (.].*", Advice: "Remote Method Invocations create coupling between componets. Move to cloud friendly alternatives such as REST endpoints.", Effort: 3, Readiness: 6, Impact: "", Category: "iop", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "non-standard-protocol",}, { Value: "corba",}, { Value: "legacy-protocol",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import javax.rmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-rmi", Recipe: "", },
             { Type: "", Pattern: "", Value: "import org.omg.IOP", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "omg-IOP", Recipe: "", },
             }, },
        
            { Name: "java-java-fx-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Java-fx is not cloud compatible and requires the JRE on the remote device.", Effort: 1000, Readiness: 100, Impact: "", Category: "java-fx", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "java-fx",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javafx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jaxrs-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to platform documentation", Effort: 1, Readiness: 8, Impact: "", Category: "jax-rs", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jax-rs",}, { Value: "rest",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ws.rs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jaxrs", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.sun.jersey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jersey", Recipe: "", },
             }, },
        
            { Name: "java-jboss", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(import)\\s+(%s)\\.{1}", Advice: "Convert to Spring based POJOs instead of using container specific functionality", Effort: 50, Readiness: 5, Impact: "", Category: "jboss", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jboss",}, { Value: "migrate-off-legacy-server",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.jboss", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jcaAnnotations", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Convert to a cloud friendly backing service.", Effort: 100, Readiness: 7, Impact: "", Category: "jca", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "jca",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AdministeredObjectDefinition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionFactoryDefinition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionFactoryDefinitions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jcache", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Make sure its externalized and using Cloud Friendly cache implementation.", Effort: 50, Readiness: 6, Impact: "", Category: "cache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "jcache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getConfigurableCacheFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamedCache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ReflectionExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ChainedExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jersey-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Refer to 3rd party organization for cloud affinity of library", Effort: 5, Readiness: 8, Impact: "", Category: "jersey", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jersey",}, { Value: "rest",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.sun.jersey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jetty-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Use Managed Executor", Effort: 1, Readiness: 10, Impact: "", Category: "threading", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "threading",}, { Value: "jetty",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.eclipse.jetty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jks", RuleType: "standard", FileType: "jks$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "Make sure to externalize jks store", Effort: 7, Readiness: 7, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, { Value: "jks",}, { Value: "mutual-auth",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*\\.jks", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jms", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.].*", Advice: "Run embedded service broker as a JMS Provider.", Effort: 10, Readiness: 6, Impact: "", Category: "jms", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jms",}, { Value: "message-queue",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageConsumer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jms-listener", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMSException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BytesMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompletionListener", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionConsumer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionMetaData", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DeliveryMode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ExceptionListener", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMSConsumer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMSContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMSProducer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MapMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageConsumer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageListener", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageProducer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ObjectMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueueBrowser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueueConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueueConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueueReceiver", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueueSender", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "QueueSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerSessionPool", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StreamMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TemporaryQueue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TemporaryTopic", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TextMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TopicConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "topic", Recipe: "", },
             { Type: "", Pattern: "", Value: "TopicConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TopicPublisher", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TopicSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TopicSubscriber", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAJMSContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAQueueConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAQueueConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAQueueSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XASession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xa", Recipe: "", },
             { Type: "", Pattern: "", Value: "XATopicConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XATopicConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XATopicSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jndi", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Refactor jndi calls", Effort: 250, Readiness: 7, Impact: "", Category: "jndi", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "NamingEnumeration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameClassPair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameParser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DirContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BinaryRefAddr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompositeName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CompoundName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InitialContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LinkRef", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NameClassPair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RefAddr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StringRefAddr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BasicControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ControlFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InitialLdapContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LdapName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ManageReferralControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PagedResultsControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PagedResultsResponseControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SortControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SortKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SortResponseControl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StartTlsRequest", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StartTlsResponse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UnsolicitedNotificationEvent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamingEvent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamingExceptionEvent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jni", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*%s\\s*native\\s*", Advice: "A few conditions have to be met to make JNI calls", Effort: 1000, Readiness: 7, Impact: "", Category: "jni", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "native",}, { Value: "jni",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "public", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "private", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "static", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jpa", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ @.]%s[ (].*", Advice: "JPA will work inside of Cloud Native applications, make sure to use best practices to externalize connection parameters.", Effort: 2, Readiness: 6, Impact: "", Category: "jpa", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jpa",}, { Value: "web-profile",}, { Value: "entity-manager",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "createEntityManagerFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "createEntityManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getTransaction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityManagerFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MapKeyClass", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MapKeyColumn", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MapKeyEnumerated", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EntityTransaction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ElementCollection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jsf", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*$", Advice: "Consider migrating to modern frameworks that have better support for the cloud.", Effort: 5, Readiness: 4, Impact: "", Category: "jsf", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ApplicationConfigurationPopulator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ApplicationFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ApplicationWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConfigurableNavigationHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConfigurableNavigationHandlerWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FacesMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FacesMessage.Severity", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NavigationCase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NavigationCaseWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NavigationHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NavigationHandlerWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ResourceHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ResourceHandlerWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ResourceWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StateManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StateManagerWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ViewHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ViewHandlerWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ViewResource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ApplicationScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CustomScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Managemodelean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ManagedProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NoneScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Referencemodelean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RequestScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SessionScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ViewScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionSource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionSource2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ContextCallback", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EditableValueHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamingContainer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PartialStateHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StateHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StateHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransientStateHelper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransientStateHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UniqueIdVendor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ValueHolder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIColumn", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UICommand", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIComponent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIComponentBase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIData", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIForm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIGraphic", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIInput", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIMessages", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UINamingContainer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIOutcomeTarget", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIOutput", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIPanel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIParameter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UISelectBoolean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UISelectItem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UISelectItems", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UISelectMany", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UISelectOne", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIViewAction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIViewParameter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIViewParameter.Reference", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIViewRoot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jsp", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider migrating to modern UI frameworks that have better support for the cloud.", Effort: 2, Readiness: 5, Impact: "", Category: "jsp", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ui",}, { Value: "jsp",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TagSupport", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BodyContent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HttpJspPage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspApplicationContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspPage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspEngineInfo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspWriter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PageContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JspTagException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SkipPageException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jta", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (.{].*", Advice: "Distributed transactions are problematic and should be remediated.  Consider Eventual Consistency pattern.", Effort: 500, Readiness: 6, Impact: "file", Category: "jta", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "transaction",}, { Value: "jta",}, { Value: "web-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "UserTransaction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InvalidTransactionException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionRequiredException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionRollemodelackException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionSynchronizationRegistry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UserTransaction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HeuristicCommitException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HeuristicMixedException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HeuristicRollbackException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InvalidTransactionException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NotSupportedException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RollbackException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionalException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionRequiredException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionRollemodelackException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAResource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Xid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XAException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "enlistForDurableTwoPhase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "enlistForVolatileTwoPhase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jvm-runtimeConfigProps", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Do not change these properties at runtime in application code", Effort: 50, Readiness: 0, Impact: "", Category: "config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "config",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.xml.soap.SOAPFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "soap", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.parsers.DocumentBuilderFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml-parser", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.parsers.SAXParserFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml-parser", Recipe: "", },
             }, },
        
            { Name: "java-logging-file-appender", RuleType: "standard", FileType: "xml$|properties$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Replace file appender with console appender", Effort: 3, Readiness: 5, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, { Value: "log2file",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "fileappender", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "</File>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             { Type: "", Pattern: "", Value: "</RollingFile>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             { Type: "", Pattern: "", Value: "type(\\s)*=(\\s)*\"File\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             { Type: "", Pattern: "", Value: "type(\\s)*=(\\s)*\"RollingFile\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             { Type: "", Pattern: "", Value: "appender.rolling.type", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             { Type: "", Pattern: "", Value: "appender.file.type", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             }, },
        
            { Name: "java-logging-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Change to an implementation of SLF4J i.e. Logback", Effort: 3, Readiness: 5, Impact: "", Category: "logging", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java\\.util\\.logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java-util", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.apache\\.log4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.apache\\.logging\\.log4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j2", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.apache\\.commons\\.logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commons-logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.osgi\\.service\\.log", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "org-osgi", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.jboss\\.logging\\.Logger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-logging", Recipe: "", },
             }, },
        
            { Name: "java-message-driven-annotations", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "To convert a message driven bean to spring cloud stream with rabbitmq", Effort: 10, Readiness: 0, Impact: "", Category: "annotations", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "message-driven", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationConfigProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "message-driven", Recipe: "", },
             }, },
        
            { Name: "java-messageDrivenBeans", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 10, Impact: "", Category: "mdb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJB\\(lookup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-metrics", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Indicates use of a metrics collection library, which supports containerization", Effort: -10, Readiness: 7, Impact: "", Category: "metrics", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "metrics",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "io.micrometer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "micrometer", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.management", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-management", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.dropwizard", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dropwizard", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netuitive.ananke", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netuitive", Recipe: "", },
             { Type: "", Pattern: "", Value: "edu.iris.dmc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "iris-dmc", Recipe: "", },
             }, },
        
            { Name: "java-mongo-cassandra", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(import)\\s+(%s)", Advice: "Application is using a non relational database", Effort: 0, Readiness: 0, Impact: "", Category: "springFramework", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.springframework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.springframework.mongo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mongo", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.springframework.cassandra", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cassandra", Recipe: "", },
             }, },
        
            { Name: "java-mqseries", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "You need to make sure that you are using the dependencies that match the version of IBM MQ on the server", Effort: 7, Readiness: 7, Impact: "", Category: "ibm-mq", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ibm-mq",}, { Value: "message-queue",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MQQueueConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WMQConstants", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQReceiveExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQSecurityExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQSendExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQAsyncStatus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQChannelDefinition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQChannelExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQConnectionSecurityParameters", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQDestination", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQDistributionList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQDistributionListItem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQEnvironment", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQExitChain", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQExternalReceiveExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQExternalSecurityExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQExternalSendExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQExternalUserExit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQGetMessageOptions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQJavaLevel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQManagedObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQMessageTracker", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQPoolToken", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQPropertyDescriptor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQPutMessageOptions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQQueue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQQueueManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQReceiveExitChain", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQSendExitChain", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQSimpleConnectionManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQSubscription", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MQTopic", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mulesoft-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "There are several changes required to move a Mule Project to PCF", Effort: 100, Readiness: 10, Impact: "", Category: "etl", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mulesoft",}, { Value: "etl",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.mule", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mulesoft-intf", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ {].*", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "etl", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mulesoft",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AbstractMessageTransformer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AbstractTransformer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-netflix-healthcheck", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Indicates existance of healthcheck endpoint, which is positive finding", Effort: -100, Readiness: 7, Impact: "", Category: "health-check", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "health-check",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "netflix.karyon.transport.http.health.HealthCheckEndpoint", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "karyon-health-check", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.guice", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-health-check", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.api.HealthCheckAggregator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-health-check", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.api.HealthCheckStatus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-health-check", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.api.IndicatorMatcher", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-health-check", Recipe: "", },
             }, },
        
            { Name: "java-nio", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Make sure to use Cloud Friendly storage provider.", Effort: 8, Readiness: 8, Impact: "", Category: "nio", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "nio",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BufferOverflowException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BufferUnderflowException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InvalidMarkException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ReadOnlyBufferException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileSystemProvider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "GroupPrincipal", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UserPrincipal", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BasicFileAttributeView", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UserPrincipalLookupService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BasicFileAttributeView", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BasicFileAttributes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PosixFileAttributeView", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PosixFileAttributes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PosixFilePermission", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PosixFilePermissions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SocketChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CharacterCodingException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerSocketChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SelectionKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CancelledKeyException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DatagramChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StandardOpenOption", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MappemodelyteBuffer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ClosedChannelException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StandardOpenOption", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AsynchronousChannelGroup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AsynchronousFileChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AsynchronousServerSocketChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AsynchronousSocketChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DatagramChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileChannel.MapMode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileLock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MembershipKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Pipe.SinkChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Pipe.SourceChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SelectableChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SelectionKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Selector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServerSocketChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SocketChannel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-nonstandard-protocol", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Ensure cloud friendly communication protocols", Effort: 50, Readiness: 10, Impact: "", Category: "protocol", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "non-standard-protocol",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "t3:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "t3", Recipe: "", },
             { Type: "", Pattern: "", Value: "iiop:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "iiop", Recipe: "", },
             { Type: "", Pattern: "", Value: "corbaloc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "corba", Recipe: "", },
             { Type: "", Pattern: "", Value: "tcp:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tcp", Recipe: "", },
             { Type: "", Pattern: "", Value: "ssl:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "udp:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "udp", Recipe: "", },
             { Type: "", Pattern: "", Value: "imap:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "imap", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdbc:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             { Type: "", Pattern: "", Value: "ldap:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ldap", Recipe: "", },
             { Type: "", Pattern: "", Value: "ldaps:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ldaps", Recipe: "", },
             { Type: "", Pattern: "", Value: "pop2:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pop2", Recipe: "", },
             { Type: "", Pattern: "", Value: "smtp:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "smtp", Recipe: "", },
             { Type: "", Pattern: "", Value: "ssl:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "tcp:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tcp", Recipe: "", },
             { Type: "", Pattern: "", Value: "ftp:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ftp", Recipe: "", },
             { Type: "", Pattern: "", Value: "sftp:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sftp", Recipe: "", },
             }, },
        
            { Name: "java-persistence", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 5, Readiness: 5, Impact: "", Category: "persistence", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "javax-persistence",}, { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.persistence", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-portUsage", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^*.(%s?:.*):(\\d*)\\/?(.*)", Advice: "Ensure port usage is cloud-friendly or use TKG", Effort: 5, Readiness: 6, Impact: "", Category: "port-usage", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "port-usage",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "https", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "http", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-rabbitmq-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Make sure that configuration is Cloud friendly", Effort: 1, Readiness: 5, Impact: "", Category: "mq", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "message-queue",}, { Value: "rabbitmq",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.rabbitmq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-remoteEJB", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 10, Readiness: 10, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "ejb",}, { Value: "fullprofile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RemoteHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "remote", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-remoteWebService-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 10, Readiness: 10, Impact: "", Category: "web-service", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-service",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.xml.ws.AsyncHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-xml-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.ws.Service", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-xml-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.ws.Service.Mode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-xml-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.ws.WebServiceClient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-xml-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.ws.WebServiceRef", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-xml-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.ws.WebServiceRefs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-xml-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.client.wink.client", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "wink", Recipe: "", },
             }, },
        
            { Name: "java-resource-cci", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 70, Readiness: 7, Impact: "", Category: "jca", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cci",}, { Value: "resource-adapter",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "createInteraction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getEISProductName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getEISProductVersion", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getAdapterName", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getAdapterShortDescription", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getAdapterVersion", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getInteractionSpecsSupported", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getSpecVersion", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "supportsExecuteWithInputAndOutputRecord", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "supportsExecuteWithInputRecordOnly", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "supportsLocalTransactionDemarcation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-resource-spi", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 10, Readiness: 70, Impact: "", Category: "jca", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "resource-adapter",}, { Value: "spi",}, { Value: "jca",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getTransactionSynchronizationRegistry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionRequestInfo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DissociatableManagedConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "dissociateConnections", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LazyAssociatableConnectionManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "associateConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LazyEnlistableConnectionManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "lazyEnlist", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LazyEnlistableManagedConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ResourceAdapter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "endpointActivation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "endpointDeactivation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ResourceAdapterAssociation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getResourceAdapter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setResourceAdapter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ValidatingManagedConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getInvalidConnections", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "XATerminator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageEndpointFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "isDeliveryTransacted", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getMechType", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DistributableWork", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DistributableWorkManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WorkContextLifecycleListener", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "contextSetupComplete", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "contextSetupFailed", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WorkContextProvider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "scheduleWork", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "HintsContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setupSecurityContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WorkAdapter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "workAccepted", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "workCompleted", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "workRejected", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "workStarted", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WorkContextErrorCodes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getStartDuration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-restlet-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "The Restlet library appears to be dead at this point in time. Consider upgrading to a Cloud friendly UI framework.", Effort: 10, Readiness: 7, Impact: "", Category: "restlet", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "rest",}, { Value: "restlet",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.restlet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-rpc-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Adapt cloud friendly protocol", Effort: 2, Readiness: 10, Impact: "", Category: "web-service", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-service",}, { Value: "rpc",}, { Value: "soap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.apache.soap.rpc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apach-soap", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.rpc.Call", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.xml.rpc", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.rpc.Service", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.xml.rpc", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.rpc.Stub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.xml.rpc", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.rpc.ServiceFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.xml.rpc", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.rpc.ServiceException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.xml.rpc", Recipe: "", },
             }, },
        
            { Name: "java-security-annotations", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Java EE security appears to be used", Effort: 5, Readiness: 0, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DeclareRoles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DenyAll", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PermitAll", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RolesAllowed", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-security", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Application is using a java keystore", Effort: 10, Readiness: 0, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.net.ssl.keyStore", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-net-ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.net.ssl.keyStoreType", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-net-ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.net.ssl.keyStorePassword", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-net-ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.net.ssl.trustStore", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-net-ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.net.ssl.trustStoreType", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-net-ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.net.ssl.trustStorePassword", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-net-ssl", Recipe: "", },
             { Type: "", Pattern: "", Value: "java.security", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "x509", Recipe: "", },
             }, },
        
            { Name: "java-servlet-session", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "HTTP Session Java API Import. Make sure that externalized data store is used for session", Effort: 10, Readiness: 0, Impact: "", Category: "servlet", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "session",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.servlet.http.HttpSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-servlet", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Servlet Java API Import", Effort: 1, Readiness: 0, Impact: "", Category: "servlet", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.servlet.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-slf4j-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Review logging configuration and remove file appenders.", Effort: 1, Readiness: 0, Impact: "", Category: "slf4j", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "slf4j",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.slf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-soap", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.* %s.*[(].*", Advice: "Consider upgrading to modern cloud native messaging", Effort: 4, Readiness: 3, Impact: "", Category: "soap", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-service",}, { Value: "soap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "SOAPBody", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPBodyElement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPConstants", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPElement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPEnvelope", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPFault", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPFaultElement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPHeader", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPHeaderElement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AttachmentPart", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MimeHeader", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MimeHeaders", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SAAJMetaFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SAAJResult", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPElementFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPPart", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SOAPBinding", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-springboot-annotations", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Spring Boot is a positive score", Effort: -100, Readiness: 0, Impact: "", Category: "spring-boot", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring-boot",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "SpringBootApplication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-springframework", RuleType: "standard", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Presence of spring framework may indicate the app should target TAS", Effort: -100, Readiness: 0, Impact: "", Category: "springFramework", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "spring",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import org.springframework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import org.springframework.mongo", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "mongo", Recipe: "", },
             { Type: "", Pattern: "", Value: "import org.springframework.cassandra", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "cassandra", Recipe: "", },
             }, },
        
            { Name: "java-stateful-annotations", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Relies on application server for shared state and will require a rewrite to Stateless or externalize state storage", Effort: 10, Readiness: 0, Impact: "", Category: "stateful", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "stateful",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Init", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Remove", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PostActivate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PrePassivate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "stateful", Recipe: "", },
             }, },
        
            { Name: "java-stateless-annotations", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Consider rearchitecting if decisions is made to not use application server", Effort: 5, Readiness: 0, Impact: "", Category: "stateless", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "stateless",}, { Value: "ejb",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateless", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             }, },
        
            { Name: "java-struts-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 6, Readiness: 7, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "web-app",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.apache.struts", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.struts2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.opensymphony", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "opensymphony", Recipe: "", },
             }, },
        
            { Name: "java-struts", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 6, Readiness: 7, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "web-app",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "LibaryBaseAction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionForward", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionForm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RequestProcessor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DynaValidationForm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionComponent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionError", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActionMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AppendIterator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Checkbox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CheckboxList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ClosingUIBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ComboBox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ComponentUrlProvider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ContextBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DateTextField", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DoubleListUIBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DoubleSelect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FieldError", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FormButton", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "GenericUIBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "InputTransferSelect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "IteratorComponent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ListUIBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MergeIterator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "OptGroup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "OptionTransferSelect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Radio", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServletUrlRenderer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TextArea", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TextField", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UIBean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "UpDownSelect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-swing", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 500, Readiness: 100, Impact: "", Category: "Swing", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "desktop-app",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "JApplet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JButton", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JCheckBox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JCheckBoxMenuItem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JColorChooser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JComboBox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JComponent", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JDesktopPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JDialog", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JEditorPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JFileChooser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JFormattedTextField", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JFormattedTextField.AbstractFormatter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JFormattedTextField.AbstractFormatterFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JFrame", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JInternalFrame", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JInternalFrame.JDesktopIcon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JLabel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JLayer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JLayeredPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JList.DropLocation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMenu", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMenuBar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JMenuItem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JOptionPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JPanel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JPasswordField", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JPopupMenu", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JPopupMenu.Separator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JProgressBar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JRadioButton", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JRadioButtonMenuItem", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JRootPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JScrollBar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JScrollPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSeparator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSlider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSpinner", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSpinner.DateEditor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSpinner.DefaultEditor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSpinner.ListEditor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSpinner.NumberEditor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JSplitPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTabbedPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTable.DropLocation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTextArea", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTextField", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTextPane", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JToggleButton", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JToggleButton.ToggleButtonModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JToolBar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JToolBar.Separator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JToolTip", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTree", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTree.DropLocation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTree.DynamicUtilTreeNode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JTree.EmptySelectionModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JViewport", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "JWindow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-system-config", RuleType: "standard", FileType: "$", Target: "line", Type: "regex", DefaultPattern: "(System.)%s\\(", Advice: "Review usage of environment variables and system properties and externalize.", Effort: 3, Readiness: 10, Impact: "", Category: "config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "config",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getenv", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "setProperties", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-systemLoad", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(System.)%s", Advice: "Remediate to cloud friendly implentation", Effort: 1000, Readiness: 10, Impact: "", Category: "process-launch", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jni",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "loadLibrary", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "load", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-tangosol", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Convert to cloud friendly cache implementation", Effort: 10, Readiness: 6, Impact: "", Category: "cache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "tangosol",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getConfigurableCacheFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamedCache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ReflectionExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ChainedExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-threadUsage-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Use Managed Executor", Effort: 100, Readiness: 10, Impact: "", Category: "threading", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "threading",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java.lang.Thread", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "java.lang.Runnable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-tibco-jms", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Integrating with TIBCO BusinessWorks JMS queues from a Spring application requires vendor-specific implementation", Effort: 7, Readiness: 6, Impact: "", Category: "tibco", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "tibco",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TibjmsConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-transaction-annotations", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Review if distributed transcations are used and consider rearchitecting using eventual consistency", Effort: 5, Readiness: 0, Impact: "", Category: "transaction", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "transaction",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Remote", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LocalHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionManagement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionAttribute", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-transportSecurity", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Servlet data protection is used, make sure that it is supported by target platform runtime", Effort: 10, Readiness: 10, Impact: "", Category: "Security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, { Value: "transport-security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TransportGuarantee.CONFIDENTIAL", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-weblogic-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 100, Readiness: 10, Impact: "", Category: "oracle", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.bea.common.security.xacml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bea", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.httppubsub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bea", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bea", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.security.saml2.providers", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "saml2", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.security.saml2.providers.registry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "saml2", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp.jdbc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp.jdbc.oracle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "weblogic", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogicx.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "weblogic", Recipe: "", },
             }, },
        
            { Name: "java-weblogic", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 50, Readiness: 10, Impact: "", Category: "webLogic", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "weblogic.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-websockets-import", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Make sure that target platform supports websocket api", Effort: 100, Readiness: 100, Impact: "", Category: "websockets", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-socket",}, { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.websocket", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.websocket.Endpoint", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.vertx.core.http.ServerWebSocket;", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "vertx", Recipe: "", },
             }, },
        
            { Name: "java-ws2liberty-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Vendor proprietary implementation.  Consider rearchitecting if decision is made to move off application server", Effort: 10, Readiness: 10, Impact: "", Category: "webSphere", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "liberty",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.ibm.websphere", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websphere", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.axis2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-axis2", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.workplace.extension", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-workplace", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.isc.api.platform", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-isc", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.portal", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-portal", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.wsspi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "wsspi", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.tuscany.sca", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-tuscany", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.oasisopen.sca", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oasisopen", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.osoa.sca", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osao-sca", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.uddi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-uddi", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.ws", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-ws", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.ejs.ras", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-ejs-ras", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.ffdc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-ffdc", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.ras", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-ras", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.servlet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-servlet", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.etools", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-etools", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.jca", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-jca", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.webtools", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-webtools", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.wsdl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-wsdl", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.wsif", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-wsif", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.xmlsoap.schemas.wsdl.wsadie.messages", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "soap-wsdl", Recipe: "", },
             }, },
        
            { Name: "java-ws2liberty-methods", RuleType: "standard", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Vendor proprietary implementation.  Consider rearchitecting if decision is made to move off application server", Effort: 10, Readiness: 10, Impact: "", Category: "webSphere", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "websphere",}, { Value: "liberty",}, { Value: "ltpa",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getCallerList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getFirstCaller", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getFirstServer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getServerList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "addPropagationAttribute", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getPropagationAttributes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "convertCookieStringToBytes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "revokeSSOCookies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "revokeSSOCookiesForPortlets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getLTPACookieFromSSOToken", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-bool-comparison", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(!|=)=\\\\s+%s", Advice: "Boolean literals should be avoided in comparison expressions == and != to improve code readability.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "bool-compare",}, { Value: "readability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "true", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "false", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-cache", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Using caching in memory will not scale when used in a multi nodes environment. Use distributed caching like redis instead.", Effort: 20, Readiness: 1000, Impact: "", Category: "session_management", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*(cache)\\s*\\[.*\\]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*(CACHE)\\s*\\[.*\\]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*cache.get\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*cache.set\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-cipher-algorithms-should-be-robust", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s+.+,+)", Advice: "Weak cipher algorithms are used. A general recommendation is to only use cipher algorithms intensively tested and promoted by the cryptographic community.", Effort: 10, Readiness: 1000, Impact: "", Category: "vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DES", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DES-EDE", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DES-EDE3", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RC2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RC4", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-continue", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(\\\\s+)?%s(\\\\s+)?;", Advice: "continue is an unstructured control flow statement. It makes code less testable, less readable and less maintainable. Structured control flow statements such as if should be used instead.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "testability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "continue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-document-write.yaml", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "The use of document.write where native DOM alternatives such as document.createElement are more appropriate. document.write has been grossly misused over the years and has quite a few disadvantages, including that if it’s executed after the page has been loaded, it can actually overwrite the page. Opting for more DOM-friendly methods such as document.createElement is more favourable", Effort: 5, Readiness: 1000, Impact: "", Category: "code_smell", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "code_smell",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "document.write", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-fileIO", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".%s", Advice: "Cloud-native applications should not rely on a local filesystem to store information, configuration or logs. Some alternatives could be cloud storage (buckets), storing data in a shared database or leveraging shared mounted volumes in a kubernetes like environment.", Effort: 30, Readiness: 8, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "appendFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "writeFile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-function", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\\\s+[A-Z]", Advice: "Shared naming conventions allow teams to collaborate efficiently. This rule checks that all function names match a provided regular expression.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "function",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "function", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-jwt-signed-verify-with-strong-cipher-algorithms", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".*\\s*{\\s(%s\\s*:+\\s*.none.\\s*}|.*\\s*(%ss:+\\s*\\[\\s*.*,+\\s*.*none.*\\s*]+))", Advice: "If a JSON Web Token (JWT) is not signed with a strong cipher algorithm (or not signed at all) an attacker can forge it and impersonate user identities. Do not use none algorithm to sign or verify the validity of a token. Do not use a token without verifying its signature before.", Effort: 10, Readiness: 1000, Impact: "", Category: "vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "algorithm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-stdout", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\s*\\(.*\\)", Advice: "Debug statements are always useful during development. But include them in production code - particularly in code that runs client-side - and you run the risk of inadvertently exposing sensitive information, slowing down the browser, or even erroring-out the site for some users.", Effort: 2, Readiness: 4, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "console.log", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-symbol", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "new\\\\s+%s", Advice: "Symbol is a primitive type introduced in ECMAScript2015. Its instances are mainly used as unique property keys. An instance can only be created by using Symbol as a function. Using Symbol with the new operator will raise a TypeError.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "type-error",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Symbol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-throw-literal", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\\\s+(\\\\d|''.+''|\\\".+\\\")", Advice: "It is a bad practice to throw something that's not derived at some level from Error. If you can't find an existing Error type that suitably conveys what you need to convey, then you should extend Error to create one.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "exception",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "throw", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-md5-sha1-noncompliant", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*(1.1)+", Advice: "TLS1.1 is not secure", Effort: 5, Readiness: 1000, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "tls", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TLS", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-var", RuleType: "standard", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(\\\\s+)?%s\\\\s+", Advice: "ECMAScript 2015 introduced the let and const keywords for block-scope variable declaration. Using const creates a read-only (constant) variable.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "read-only",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "var", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "log2file-import", RuleType: "standard", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Logging should be to console", Effort: 1, Readiness: 8, Impact: "", Category: "log2file", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "log2file",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java.util.logging.FileHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java.util.logging", Recipe: "", },
             }, },
        
            { Name: "log4j-properties", RuleType: "standard", FileType: "log4j2.properties$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Refer to platform documentation", Effort: 1, Readiness: 0, Impact: "", Category: "log2file", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "log2file",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "property.filename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "log4j-xml", RuleType: "standard", FileType: "log4j2.xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Refer to platform documentation", Effort: 1, Readiness: 0, Impact: "", Category: "log2file", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "log2file",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "property.filename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ISAPI-Filters-config", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 1000, Readiness: 10, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/isapiFilters[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ISAPI-Filters-vbcs", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 1000, Readiness: 0, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "isapi-filter",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "GetSection\\(system.webServer/isapiFilters\\)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "GetSection\\(\"system.webServer/isapiFilters\"\\)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ip-address", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported in .netCore, Use IHttpContextAccessor instead", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "debt",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Request.UserHostAddress", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-MSMQ-vbcs", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "MSMQ is a proprietary Windows server-based messaging solution that is not supported by Linux based containers. Leverage bridging solutions like https://kubemq.io/helping-net-msmq-based-services-to-migrate-to-kubernetes/ on kubernetes", Effort: 10, Readiness: 1, Impact: "", Category: "unsupported-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "msmq",}, { Value: "messaging",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageQueue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.Messaging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceHost", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NetMsmqBinding", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-UriCacheModule", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux. Implements a generic cache for URL-specific server state; such as configuration. With this module; the server only reads configuration for the first request for a particular URL. And reuse it on subsequent requests until it changes.", Effort: 1000, Readiness: 5, Impact: "", Category: "Unsupported IIS modules", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "uri-cache",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/globalModules/add[@name=\"UriCacheModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-child-action", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*\\[%s\\].*", Advice: "Not supported, replace with a new View Component feature https://learn.microsoft.com/en-us/aspnet/core/mvc/views/view-components?view=aspnetcore-7.0", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "mvc",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ChildActionOnly", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-classic-2-0", RuleType: "standard", FileType: "(asp$)", Target: "line", Type: "regex", DefaultPattern: "\\s%s.*", Advice: "Not supported by many PAS. ASP .net pages will have to be replaced by ASP .NetCore", Effort: 10, Readiness: 9, Impact: "", Category: "Unsupported", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "asp-classic",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "runat=\"server\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "runat=\"Server\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-machine-key", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Not supported in .NetCore. Will have to be replaced https://learn.microsoft.com/en-us/aspnet/core/security/data-protection/compatibility/replacing-machinekey?view=aspnetcore-7.0", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "encryption",}, { Value: "config",}, { Value: "asp-classic",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/machineKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-membership", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Not supported, replaced by ASP.NET Core Identity. Manage users in Database.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "authentication",}, { Value: "asp",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/membership", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-mvc-form-collection", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Not supported, Replace with IFormCollection", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mvc",}, { Value: "asp",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*Create.*\\(FormCollection.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-mvc-model-update", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*\\s%s<.*", Advice: "Not supported in .NetCore, Replace with TryUpdateModelAsync", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mvc",}, { Value: "asp",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TryUpdateModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-mvc", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported on .netCore. Migrate to ASP .NetCore and replace with Microsoft.AspNetCore.Mvc", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mvc",}, { Value: "asp",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Web.Mvc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-session-context", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported, Replace with IHttpContextAccessor to manage session data", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "session",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "HttpContextBase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-web-form", RuleType: "standard", FileType: "(aspx$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Unsupported-netcore. Not supported by many PAS. ASP .net pages will have to be replaced by ASP .NetCore. See https://learn.microsoft.com/en-us/dotnet/architecture/porting-existing-aspnet-apps/migrate-web-forms", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore.", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "asp-web-form",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(Page Language=)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(CodeBehind=.*.aspx)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-classic-asp-net", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported-netcore. Not supported by many PAS. ASP .net pages will have to be replaced by ASP .NetCore. See https://learn.microsoft.com/en-us/dotnet/architecture/porting-existing-aspnet-apps/migrate-web-forms", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "asp-classic",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Web.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-delegate-begininvoke", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "BeginInvoke and EndInvoke methods on delegates are supported on .NET Core. See https://devblogs.microsoft.com/dotnet/migrating-delegate-begininvoke-calls-for-net-core/", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "netcore",}, { Value: "mvc",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BeginInvoke\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EndInvoke\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-identity-foundation", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Identity Foundation (WIF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "WIF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.IdentityModel.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-AnonymousAuthentication", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/authentication/anonymousAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Authentication", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux. Use Microsoft.AspNetCore.Authentication.Negotiate instead.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  { Value: ".authentication.*mode=\"None\"",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/authentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Authorization", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/authorization", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-CertificateMappingAuthentication", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/authentication/clientCertificateMappingAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-DefaultDocument", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/defaultDocument", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-DigestAuthentication", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/authentication/digestAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-DirectoryListing", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/directoryBrowse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-HttpErrors", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpErrors", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-HttpLogging", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpLogging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-HttpRedirection", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpRedirect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-IisClientCertificateMappingAuthentication", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/iisClientCertificateMappingAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-IpSecurity", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/ipSecurity", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-IsapiCgiRestriction", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/isapiCgiRestriction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-OutputCache", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/caching/outputCache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-ProtocolSupport", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpProtocol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-ServerSideInclude", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/serverSideInclude", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Tracing", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/tracing", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpTracing", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Validation", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "iis-module",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/validation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-javascript-serializer", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Use JSON.NET instead.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "JavaScriptSerializer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-managed-addin", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Managed Addin Framework (MAF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.AddIn", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-app-metrics-extensions-hosting", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "App.Metrics.Extensions.Hosting", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-mvc5", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.Mvc5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-owin", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.Owin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-webapi2-owin", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.WebApi2.Owin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-webapi2", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.WebApi2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-dotnetopenauth-core", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DotNetOpenAuth.Core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-fluentvalidation-mvc5", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FluentValidation.MVC5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-fluentvalidation-webapi", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FluentValidation.WebAPI", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-glimpse", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Glimpse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-imageprocessor-web", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ImageProcessor.Web", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-imageprocessor", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ImageProcessor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-imageresizer", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ImageResizer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-friendlyurls-core", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.FriendlyUrls.Core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-identity-entityframework", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.Identity.EntityFramework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-mvc", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.Mvc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-signalr", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.SignalR", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-webapi-versioning", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.WebApi.Versioning", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-azure-common-dependencies", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Azure.Common.Dependencies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-build-locator", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Build.Locator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-build-tasks-git", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Build.Tasks.Git", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-codeanalysis-workspaces-msbuild", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CodeAnalysis.Workspaces.MSBuild", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-coreassemblies", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.CoreAssemblies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-deployment", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.Deployment", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-workflow", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.Workflow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-xrmtooling-coreassembly", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.XrmTooling.CoreAssembly", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-data-sqlclient-sni", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Data.SqlClient.SNI", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-diagnostics-tracing-eventsource-redist", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Diagnostics.Tracing.EventSource.Redist", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net45", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net45", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net461", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net461", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net462", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net462", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net472", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net472", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net48", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net48", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-cors", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Cors", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-diagnostics", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Diagnostics", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-filesystems", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.FileSystems", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-activedirectory", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.ActiveDirectory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-oauth", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.OAuth", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-openidconnect", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.OpenIdConnect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-wsfederation", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.WsFederation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-selfhost", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.SelfHost", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-staticfiles", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.StaticFiles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-testing", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Testing", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-servicefabric-aspnetcore-abstractions", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.ServiceFabric.AspNetCore.Abstractions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-servicefabric-aspnetcore-kestrel", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.ServiceFabric.AspNetCore.Kestrel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-visualstudio-azure-containers-tools-targets", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.VisualStudio.Azure.Containers.Tools.Targets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-visualstudio-azure-fabric-msbuild", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.VisualStudio.Azure.Fabric.MSBuild", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-visualstudio-slowcheetah", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.VisualStudio.SlowCheetah", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-web-redissessionstateprovider", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Web.RedisSessionStateProvider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-miniprofiler", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MiniProfiler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-ninject-mvc5", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Ninject.MVC5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-ninject-web-common-webhost", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Ninject.Web.Common.WebHost", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-ninject-web-webapi", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Ninject.Web.WebApi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-nlog-extended", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "NLog.Extended", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-nlog-web", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "NLog.Web", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-portable-dataannotations", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Portable.DataAnnotations", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-razorengine", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RazorEngine", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-runtime-aot-system-reflection-primitives", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "runtime.aot.System.Reflection.Primitives", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-sqlitepclraw-provider-e_sqlite3-net45", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "SQLitePCLRaw.provider.e_sqlite3.net45", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-stub-system-data-sqlite-core-netframework", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stub.System.Data.SQLite.Core.NetFramework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-system-data-sqlite-linq", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Data.SQLite.Linq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-system-net-http-formatting-extension", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Net.Http.Formatting.Extension", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-system-reactive-windows-threading", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Reactive.Windows.Threading", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-telerik-dataaccess-core", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Telerik.DataAccess.Core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-telerik-dataaccess-web", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Telerik.DataAccess.Web", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-unity-aspnet-webapi", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Unity.AspNet.WebApi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-unity-mvc", RuleType: "standard", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Unity.Mvc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-reflection-assembly", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Use AssemblyLoadContext. Evidence is not supported in .NET Core.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Assembly.Load\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-script-blocks", RuleType: "standard", FileType: "(xsl)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Script blocks are supported only in .NET Framework. They are not supported on .NET Core or .NET 5 or later.", Effort: 10, Readiness: 10, Impact: "", Category: "WCF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "msxsl:script", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-serialization", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported in .netCore", Effort: 10, Readiness: 1, Impact: "", Category: "unsupported-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "serialization",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Runtime.Serialization", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-system-speech", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "System.Speech won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Speech.Recognition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-system-web-services", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Upgrade to CoreWCF https://devblogs.microsoft.com/dotnet/upgrading-a-wcf-service-to-dotnet-6/ or more modern rest api", Effort: 10, Readiness: 10, Impact: "", Category: "WCF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-web-services",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Web.Services", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-vb-razor-view-engine", RuleType: "standard", FileType: "(vbhtml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Not supported in .netCore https://aka.ms/vb-angular-and-web-api", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore.", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "asp-web-form",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "@Code", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-service-model-config", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Upgrade to CoreWCF https://devblogs.microsoft.com/dotnet/upgrading-a-wcf-service-to-dotnet-6/", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.serviceModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-service-model", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Upgrade to CoreWCF https://devblogs.microsoft.com/dotnet/upgrading-a-wcf-service-to-dotnet-6/", Effort: 10, Readiness: 10, Impact: "", Category: "WCF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.ServiceModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wf-activity", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "WF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Activities.Presentation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wf-integration", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "WF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "XamlIntegration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wf-presentation", RuleType: "standard", FileType: "(xaml$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 500, Readiness: 1, Impact: "", Category: "WF", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wf",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Activities.Core.Presentation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-enterprise-services", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported on .Net Core", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "com+",}, { Value: "api",}, { Value: "distributed-transaction",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.EnterpriseService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-remoting-code", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported, consider inter-process communication (IPC) System.IO.Pipes class or the MemoryMappedFile class.Also StreamJsonRpc or ASP.NET Core (either using gRPC or RESTful Web API services).", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "remoting",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Runtime.Remoting.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-remoting", RuleType: "standard", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Unsupported, consider inter-process communication (IPC) System.IO.Pipes class or the MemoryMappedFile class.Also StreamJsonRpc or ASP.NET Core (either using gRPC or RESTful Web API services).", Effort: 500, Readiness: 5, Impact: "", Category: "Unsupported", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "remoting",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.runtime.remoting", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-workflow-foundation", RuleType: "standard", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Use alternatives, see CoreWF and CoreWCF, runs on windows only and would need to be ran by a console application", Effort: 10, Readiness: 1, Impact: "", Category: "unsupported-netcore", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "windows-workflow-service",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WorkflowService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-workflow", RuleType: "standard", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Workflow.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-allow-url-in-config", RuleType: "standard", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=1$", Advice: "allow_url_fopen and allow_url_include allow code to be read into a script from URL’s. The ability to suck in executable code from outside your site, coupled with imperfect input cleansing could lay your site bare to attackers explicitly disable allow_url_fopen and allow_url_include'", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "allow_url_fopen", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "allow_url_include", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-cache", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\[.*\\]", Advice: "However, the default configuration is to store session data in a temporary file on the local disk. Again, this will not work if you’re using multiple nodes", Effort: 20, Readiness: 1000, Impact: "", Category: "session_management", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CACHE", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-cgi-force-redirect-enabled", RuleType: "standard", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+0$", Advice: "The cgi.force_redirect php.ini configuration is on by default, and it prevents unauthenticated access to scripts when PHP is running as a CGI.", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "cgi.force_redirect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-deprecated-feature-parse_str", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "([^_]%s)\\s*\\(.*[^,]\\)", Advice: "parse_str() without second argument is deprecated", Effort: 3, Readiness: 1000, Impact: "", Category: "deprecation", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "deprecated",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "parse_str", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-deprecated-feature-should-not-be-used", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "([^_]%s)\\s*\\(+.*\\)", Advice: "Deprecated language features are those that have been retained temporarily for backward compatibility, but which will eventually be removed from the language. In effect, deprecation announces a grace period to allow the smooth transition from the old features to the new ones. In that period, no use of the deprecated features should be added to the code. Refactor or upgrade to use Php 7+", Effort: 7, Readiness: 1000, Impact: "", Category: "deprecation", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "deprecated",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "call_user_method", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "call_user_method_array", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "define_syslog_variables", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "dl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ereg", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ereg_replace", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "eregi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "eregi_replace", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "set_magic_quotes_runtime", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "magic_quotes_runtime", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "session_register", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "session_unregister", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "session_is_registered", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "set_socket_blocking", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "split", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spliti", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "sql_regcase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mysql_db_query", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mysql_escape_string", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "create_function", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "gmp_random", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-disabled-enable-dl", RuleType: "standard", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+1$", Advice: "enable_dl is on by default and allows open_basedir restrictions, which limit the files a script can access, to be ignored enable_dl should be explicitly turned off in php.ini", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "enable_dl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-disabled-file-uploads", RuleType: "standard", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+1$", Advice: "file_uploads is an on-by-default PHP configuration that allows files to be uploaded to your site.", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "file_uploads", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-enable-session-use-trans-sid", RuleType: "standard", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+1$", Advice: "PHP’s session.use_trans_sid automatically appends the user’s session id to urls when cookies are disabled.", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "vulnerability",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "session.use_trans_sid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-file-system-manipulation", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Filesystem manipulation is not encouraged in cloud-native applications. Keep your content off the filesystem", Effort: 50, Readiness: 1000, Impact: "", Category: "filesystem", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "file_put_contents", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "filemtime", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fputcsv", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fputs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fopen", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fsync", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ftell", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ftruncate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fwrite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "is_writable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mkdir", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "move_uploaded_file", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "readfile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "rmdir", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "touch", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "chown", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "copy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fflush", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tmpfile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tempnam", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-function-method-naming-convention", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "^(%s)\\s+[a-z]+[a-zA-Z0-9]*\\(.*\\)", Advice: "Shared naming conventions allow teams to collaborate efficiently. All function names match a provided regular expression.", Effort: 2, Readiness: 1000, Impact: "", Category: "naming_convention", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "code-smell",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "function", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-global-variable", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(\\$%s)\\s*\\[.*\\]", Advice: "The default configuration is to store session data in a temporary file on the local disk. Again, this will not work if you’re using multiple nodes", Effort: 20, Readiness: 1000, Impact: "", Category: "session_management", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "session",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "_SESSION", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "_COOKIE", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-goto-stmt-should-not-be-used", RuleType: "standard", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s\\s*.+)", Advice: "goto is an unstructured control flow statement. It makes code less readable and maintainable. Structured control flow statements such as if, for, while, continue or break should be used instead.", Effort: 3, Readiness: 1000, Impact: "", Category: "code_smell", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "unstructured",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "goto", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-md5-sha1-noncompliant", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s\\(.*\\))", Advice: "Cryptographic hash algorithms such as MD2, MD4, MD5, MD6, HAVAL-128, HMAC-MD5, DSA (which uses SHA-1), RIPEMD, RIPEMD-128, RIPEMD-160, HMACRIPEMD160 and SHA-1 are no longer considered secure, because it is possible to have collisions (little computational effort is enough to find two or more different inputs that produce the same hash). Safer alternatives, such as SHA-256, SHA-512, SHA-3 are recommended, and for password hashing, it is even better to use algorithms that do not compute too quickly, like bcrypt, scrypt, argon2 or pbkdf2 because it slows down brute force attacks.", Effort: 10, Readiness: 1000, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "md5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "sha1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "md2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "sha1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "md4", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "haval128", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "dsa", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-reading-std-Input-security-sensitive", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Reading Standard Input is security-sensitive. It has led in the past to the following vulnerabilities:CVE-2005-2337,CVE-2017-11449", Effort: 50, Readiness: 1000, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, { Value: "filesystem",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "stream_get_line", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_copy_to_stream", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "file_get_contents", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "readfile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "fopen", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_socket_server", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_socket_client", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_socket_pair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-references-should-not-be-passed-to-func-call", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*\\s*\\(%s+.*\\)", Advice: "Passing a reference to a function parameter means that any modifications the method makes to the parameter will be made to the original value as well, since references have the effect of pointing two variables at the same memory space. This feature can be difficult to use correctly, particularly if the callee is not expecting a reference, and the improper use of references in function calls can make code less efficient rather than more efficient.", Effort: 3, Readiness: 1000, Impact: "", Category: "code_smell", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "code_smell",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "&", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-socket-security-sensitive", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Using sockets is security-sensitive. It has led in the past to the following vulnerabilities:CVE-2011-178,CVE-2017-5645,CVE-2018-6597", Effort: 50, Readiness: 1000, Impact: "", Category: "anti-pattern", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "anti-pattern",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "socket_create", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "socket_create_listen", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "socket_addrinfo_bind", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "socket_addrinfo_connect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "socket_create_pair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_socket_server", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_socket_client", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "stream_socket_pair", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-php-url-not-hardcoded", RuleType: "standard", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s://).*", Advice: "Url Hardcoding in code is anti-pattern", Effort: 2, Readiness: 1000, Impact: "", Category: "anti-pattern", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "code_smell",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "http", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "https", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "plaintext-creds", RuleType: "standard", FileType: "", Target: "line", Type: "regex", DefaultPattern: "(%s)\\s*=+\\s*[\"''][^${[(\\s][^\"${[(\\s].*", Advice: "Found plain text crdentials. Never save passwords or login information in files. Make configurable, put into environment variables or config maps in a kubernetes like environment. Credentials could also be set and retrieved from cloud friendly crdential vaults.", Effort: 0, Readiness: 9, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  { Value: "<Control.*Property.*Password=\\\"yes",}, { Value: "(Login)\\s*=+\\s*\"true",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Password", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "User", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "User Id", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "username", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Username", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Login", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Loginname", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "login", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Loginname", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "password", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-cf", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Check for cloud foundry support.", Effort: -10, Readiness: 10, Impact: "", Category: "cloud-foundry", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cloud-ready",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "load_from_vcap_services", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-db-peewee", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Check target platform has support for this library", Effort: 1, Readiness: 10, Impact: "", Category: "database", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "peewee",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "peewee", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-db-redis", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Redis is being used", Effort: 1, Readiness: 10, Impact: "", Category: "database", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "redis",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-fileIO", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform. Since containers are immutable, restarts will lose any written changes. Refactor this logic to use an external service to store state.", Effort: 8, Readiness: 8, Impact: "", Category: "io", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "io",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "open", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.open", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.rename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.path", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.mkdir", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.makedirs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.chown", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.chmod", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.remove", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.removedirs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.renames", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.replace", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.rmdir", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.symlink", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "os.unlink", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "chmod", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-rabbitmq", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "RabbitMq messaging is used", Effort: 1, Readiness: 10, Impact: "", Category: "services", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "rabbitmq",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "pika", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "aio_pika", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-redis", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "(import)\\s+(%s)", Advice: "Redis is being used", Effort: 1, Readiness: 10, Impact: "", Category: "services", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "redis",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "aioredis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-sqlite", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider migration to an external database.", Effort: 6, Readiness: 8, Impact: "", Category: "DatabaseAccess", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "database",}, { Value: "sqlite",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "sqlite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SqliteDatabase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "sqlserver-ssis", RuleType: "fire-once", FileType: "(dtsx$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "SSIS is not supported on CloudFoundry.", Effort: 100, Readiness: 0, Impact: "", Category: "Unsupported modules", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "ssis",}, { Value: "etl",}, { Value: "sql",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DTS", Advice: "SSIS is not supported on CloudFoundry. Consider leaving the packages in an external SQL Server deployment or rewrite them in a cloud native ETL Framework like Spring Cloud Data Flow.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-cfenv-gradle", RuleType: "standard", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "Application appears to be using cf env library to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-cloud-cloudfoundry-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-cfenv-maven", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(>.*%s<)", Advice: "Application appears to be using cf env library to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java-cfenv-boot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-connector-gradle", RuleType: "standard", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "Application appears to be using cloud connectors library to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-cloud-cloudfoundry-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-connector-maven", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(>.*%s<)", Advice: "Application appears to be using cloud connectors library to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-cloud-cloudfoundry-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-connector-python", RuleType: "standard", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "Application appears to be using Cloudfoundry python client to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "CloudFoundryClient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "weblogic-cluster-config", RuleType: "fire-once", FileType: "conf$", Target: "line", Type: "regex", DefaultPattern: "^%s.*", Advice: "Weblogic clusters cannot run in K8S", Effort: 1, Readiness: 0, Impact: "", Category: "wlcluster", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "wl-cluster",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WebLogicCluster", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "websphere-cluster-jacl", RuleType: "standard", FileType: "jacl$", Target: "line", Type: "regex", DefaultPattern: "^ *%s.*", Advice: "Websphere clusters cannot run in K8S", Effort: 1, Readiness: 0, Impact: "", Category: "wscluster", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ws-cluster",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "_CLUSTERS", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "wsdl-soap", RuleType: "standard", FileType: "wsdl$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "Consider upgrading to cloud friendly communication protocol", Effort: 4, Readiness: 3, Impact: "", Category: "soap", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "wsdl",}, { Value: "web-service",}, { Value: "soap",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*\\.wsdl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-activeMQ", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Make sure that Active MQ broker is available on target platform", Effort: 5, Readiness: 10, Impact: "", Category: "mq", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "amqp",}, { Value: "message-queue",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "amq:broker", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-clientCert", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Avoid reliance on SSL", Effort: 5, Readiness: 10, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "certificate",}, { Value: "auth",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<auth-method>CLIENT-CERT", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ehcache-Config", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider external cache provider", Effort: 50, Readiness: 10, Impact: "", Category: "cache", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "ehcache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ehcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-2-1", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 10, Readiness: 10, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "javaee",}, { Value: "full-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_2_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-0", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "javaee",}, { Value: "ejb-lite",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-1", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb-lite",}, { Value: "javaee",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-2", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb-lite",}, { Value: "javaee",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-mdb-import", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "MDB", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "fullprofile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mdb", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "MDB", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "fullprofile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-remote", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling.  Consider rearchitecting to use cloud friendly communication protocol", Effort: 100, Readiness: 100, Impact: "", Category: "ejb", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "remote-ejb",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "full-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<remote>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-resource-mgr-aut", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Security implementation might need to be refactored", Effort: 10, Readiness: 10, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<res-auth>Container</res-auth>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-facelets", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 2, Readiness: 2, Impact: "", Category: "Facelets", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "facelets",}, { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jsf/facelets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jboss", RuleType: "standard", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Application Server coupling", Effort: 50, Readiness: 5, Impact: "", Category: "Jboss", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jboss",}, { Value: "javaee",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jaws.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "entity-beans", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jbosscmp-jmodelc.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "fullprofile", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-service.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "config", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-web.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "webapp", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-wsse-server.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "security", Recipe: "", },
             }, },
        
            { Name: "xml-jee", RuleType: "standard", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Convert to Spring based application configuration or use importResource", Effort: 20, Readiness: 8, Impact: "", Category: "Config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "full-profile",}, { Value: "javaee",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "application.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ear", Recipe: "", },
             { Type: "", Pattern: "", Value: "application-client.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "application-client", Recipe: "", },
             { Type: "", Pattern: "", Value: "ejb-jar.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "ra.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "resourceadapter", Recipe: "", },
             { Type: "", Pattern: "", Value: "webservices.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-service", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-1-0", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-1-1", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-1-2", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-0", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-1", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-2", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-3", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_3.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-localJNDI", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "WebSphere does not allow local JNDI, refer to documentation", Effort: 10, Readiness: 10, Impact: "", Category: "jndi", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "local",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<local-jndi-name>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-messageDrivenBeans", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 10, Readiness: 10, Impact: "", Category: "mdb", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "mdb",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<ejb-ref>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-myfaces", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 5, Readiness: 2, Impact: "", Category: "myfaces", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.apache.myfaces", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-portlet-1-0", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 500, Readiness: 700, Impact: "", Category: "version-portlet", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "portlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "portlet-app_1_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-3", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 700, Readiness: 700, Impact: "", Category: "servlet", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_3.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-4", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 700, Readiness: 700, Impact: "", Category: "servlet", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_4.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-5", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 400, Readiness: 500, Impact: "", Category: "servlet", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_5.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-3-0", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 8, Readiness: 10, Impact: "", Category: "servlet", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_3_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-3-1", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 5, Readiness: 5, Impact: "", Category: "servlet", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_3_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-session-scoped-beans", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Relies on application server for shared state and will require a rewrite to Stateless or externalize state storage", Effort: 100, Readiness: 100, Impact: "", Category: "session", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "session",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "scope=\"session\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-statefulEJB", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Relies on application server for shared state and will require a rewrite to Stateless or externalize state storage", Effort: 1000, Readiness: 1000, Impact: "", Category: "ejb", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateful",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<session-type>Stateful</session-type>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-1-1", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 10, Readiness: 10, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-config_1_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-2", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.2.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-3", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.3.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-4", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.4.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-5", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.5.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-tiles", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 5, Readiness: 2, Impact: "", Category: "tiles", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, { Value: "tiles",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts.apache.org/tags-tiles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_2_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_2_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_3_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-tomahawk", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 600, Readiness: 200, Impact: "", Category: "tomahawk", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-profile",}, { Value: "tomahawk",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "myfaces.apache.org/tomahawk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-tomcat", RuleType: "fire-once", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Web Server coupling", Effort: 100, Readiness: 0, Impact: "", Category: "Tomcat", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "tomcat",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "configure.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "context.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "server.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-transportSecurity", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application Server coupling", Effort: 2, Readiness: 10, Impact: "", Category: "security", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "security",}, { Value: "web-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<transport-guarantee>CONFIDENTIAL", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-trinidad", RuleType: "standard", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 10, Readiness: 7, Impact: "", Category: "trinidad", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-profile",}, { Value: "trinidad",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "myfaces.apache.org/trinidad/config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-2", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.2/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-3", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.3/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-4", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.4/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-5", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.5/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-weblogic-1-6", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 600, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.6/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-7", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 600, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.7/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-8", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 400, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.8/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-9", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 400, Readiness: 700, Impact: "", Category: "weblogic-webapp", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "web-app",}, { Value: "weblogic",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.9/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-weblogic", RuleType: "standard", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 100, Readiness: 100, Impact: "", Category: "webLogic", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "weblogic",}, { Value: "full-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "weblogic.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-cmp-rmodelms-jar.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-ejb-jar.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-ra.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "persistence-configuration.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-webservices.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-wsee-clientHandlerChain.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "webservice-policy-ref.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-wsee-standaloneclient.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogic-application.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webprofile", RuleType: "standard", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 50, Readiness: 1000, Impact: "", Category: "config", Container: 50, Cloud: 50,
            Tags:
            []Tag{  { Value: "web-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "persistence.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jpa", Recipe: "", },
             { Type: "", Pattern: "", Value: "web.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "webapp", Recipe: "", },
             { Type: "", Pattern: "", Value: "applicationContext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "file-system-xml-application-context", Recipe: "", },
             }, },
        
            { Name: "xml-websphere", RuleType: "fire-once", FileType: "xm[li]$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 500, Readiness: 1000, Impact: "", Category: "webSphere", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "websphere",}, { Value: "full-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "client-resource.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-bnd.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-bnd.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-client-bnd.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-client-bnd.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-client-ext.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-client-ext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-ext.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-application-ext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-access-bean.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-jar-bnd.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-jar-bnd.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-jar-ext.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-jar-ext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-jar-ext-pme.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-ejb-jar-ext-pme.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-webservices-bnd.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-webservices-ext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-web-bnd.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-web-bnd.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-web-ext.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-web-ext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-web-ext-pme.xmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ibm-web-ext-pme.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "j2c_plugin.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-xa-dataSource", RuleType: "fire-once", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 1000, Readiness: 10, Impact: "", Category: "datasource", Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "transaction",}, { Value: "jta",}, { Value: "web-profile",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<xa-datasource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "yaml-test", RuleType: "standard", FileType: "(yaml|yml)$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "", Effort: 1000, Readiness: 10, Impact: "", Category: "", Container: 50, Cloud: 50,
            Tags:
            []Tag{  },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.orchestration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
    }
    return BootstrapRules
}