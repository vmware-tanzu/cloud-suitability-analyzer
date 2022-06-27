/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapRulesTemplate.txt found under go/resources folder
//Created @ 2022-06-27 12:44:54.431471 -0500 CDT m=+0.054735201

func BootstrapRules() []Rule {
    var BootstrapRules = []Rule{
        
            { Name: "SNAP-ETL-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*.*%s.*$", Advice: "Refer to IBM documentation", Effort: 100, Readiness: 0, Impact: "", Category: "etl", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "etl",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.talend", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "talend", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.odi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle-odi", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.is", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-is", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.sf.jasper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jasper", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.pentaho", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pentaho", Recipe: "", },
             }, },
        
            { Name: "SNAP-build-Ant-Maven", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Align with standard build system", Effort: 0, Readiness: 0, Impact: "", Category: "buildSystem", Criticality: "",
            Tags:
            []Tag{  { Value: "buildSystem",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "build.xml", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "ant", Recipe: "", },
             { Type: "", Pattern: "", Value: "pom.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "maven", Recipe: "", },
             }, },
        
            { Name: "SNAP-build-Gradle", FileType: "gradle$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Align with standard build system", Effort: 0, Readiness: 0, Impact: "", Category: "buildSystem", Criticality: "",
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "build.gradle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-SQL-properties", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Refer to platform documentation", Effort: 10, Readiness: 0, Impact: "", Category: "jdbc", Criticality: "",
            Tags:
            []Tag{  { Value: "jdbc",}, { Value: "sql",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "StoredProcedureQuery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CallableStatement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-SQL", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Refer to platform documentation", Effort: 10, Readiness: 0, Impact: "", Category: "jdbc", Criticality: "",
            Tags:
            []Tag{  { Value: "jdbc",}, { Value: "sql",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "StoredProcedureQuery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CallableStatement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-package-Gradle", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "apply\\s*plugin:\\s*['']%s['']", Advice: "Refer to platform documentation", Effort: 100, Readiness: 0, Impact: "", Category: "packaging", Criticality: "",
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ear", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-package-GradleJar", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "[{] *%s", Advice: "Refer to platform documentation", Effort: 0, Readiness: 0, Impact: "", Category: "packaging", Criticality: "",
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-package-Maven-Ant", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 0, Impact: "", Category: "packaging", Criticality: "",
            Tags:
            []Tag{  { Value: "maven",}, { Value: "snap",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<packaging>ear</packaging>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<target name=\"ear\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "SNAP-java-ver-Maven-Ant", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 0, Impact: "", Category: "java-ver", Criticality: "",
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "snap",}, },
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
        
            { Name: "bootCDI", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "cdi", Criticality: "",
            Tags:
            []Tag{  { Value: "cdi",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javaee-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javaee-api", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.inject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-ingect", Recipe: "", },
             { Type: "", Pattern: "", Value: "cdi-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cdi-api", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.0_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.1_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.2_spec", Recipe: "", },
             }, },
        
            { Name: "bootEJB", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "ejb",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb-api", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.0_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.1_spec", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-api_3.2_spec", Recipe: "", },
             }, },
        
            { Name: "bootJAXWS", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "jaxws", Criticality: "",
            Tags:
            []Tag{  { Value: "jaxws",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jboss-annotations-api_1.3_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-annotations-1-3", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-servlet-api_4.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-servlet-4.0", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.enterprise.concurrent-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-enterprise-concurrent", Recipe: "", },
             }, },
        
            { Name: "bootJDBC", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 5, Readiness: 5, Impact: "", Category: "persistence", Criticality: "",
            Tags:
            []Tag{  { Value: "spring-boot",}, { Value: "jdbc",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.springframework.jdbc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "srping-jdbc", Recipe: "", },
             }, },
        
            { Name: "bootJSF", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf",}, },
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
        
            { Name: "bootMDB", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "mdb", Criticality: "",
            Tags:
            []Tag{  { Value: "mdb",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax.ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-0", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-1", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-2", Recipe: "", },
             }, },
        
            { Name: "bootSTRUTS", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts2-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts-2", Recipe: "", },
             }, },
        
            { Name: "bootTXN", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "txn", Criticality: "",
            Tags:
            []Tag{  { Value: "txn",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.0_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-0", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.1_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-1", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-ejb-api_3.2_spec", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ejb-3-2", Recipe: "", },
             }, },
        
            { Name: "bootWEBSERVLET", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "WebServlet", Criticality: "",
            Tags:
            []Tag{  { Value: "webservlet",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "websocket-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websocket", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.websocket-client-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websocket-client", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.websocket-all", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websocket-all", Recipe: "", },
             }, },
        
            { Name: "config-dotnet-webConfig", FileType: "config$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Convert to Spring based application configuration", Effort: 0, Readiness: 0, Impact: "", Category: "Config", Criticality: "",
            Tags:
            []Tag{  { Value: "web-config",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web.config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Web.config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-encryption", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Use of encrypted sections is problematic in cloud environments", Effort: 300, Readiness: 0, Impact: "", Category: "Security", Criticality: "",
            Tags:
            []Tag{  { Value: "Protection-Provider",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[@configProtectionProvider][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-security", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Use of generated machine keys is problematic", Effort: 100, Readiness: 100, Impact: "", Category: "Security", Criticality: "",
            Tags:
            []Tag{  { Value: "Auto-Gen-Key",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/machineKey[contains(@validationKey, \"AutoGenerate\") or contains(@decryptionKey, \"AutoGenerate\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-sessionState", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "App should be executed as a stateless process", Effort: 200, Readiness: 2, Impact: "", Category: "stateful", Criticality: "",
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "TAS",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/sessionState[@mode=\"InProc\" or @mode=\"StateServer\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-dockerFile", FileType: "$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Determine if TKG is more prescriptive", Effort: 5, Readiness: 1000, Impact: "", Category: "docker", Criticality: "",
            Tags:
            []Tag{  { Value: "docker",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Dockerfile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-non-root-user", FileType: "$", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Shows evidence of avoiding root privledges", Effort: -100, Readiness: 1000, Impact: "", Category: "dockerSecurity", Criticality: "",
            Tags:
            []Tag{  { Value: "non-root-user",}, { Value: "docker",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RUN groupadd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RUN useradd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "USER", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-sudo", FileType: "$", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Using root inside a container is a serious vulnerability.", Effort: 100, Readiness: 10, Impact: "", Category: "docker", Criticality: "",
            Tags:
            []Tag{  { Value: "docker",}, { Value: "sudo",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RUN install -y sudo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RUN su root", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-FileCacheModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1000, Readiness: 5, Impact: "", Category: "caching", Criticality: "",
            Tags:
            []Tag{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/location/system.webServer/serverRuntime[@frequentHitThreshold or @frequentHitTimePeriod][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cache", Recipe: "", },
             }, },
        
            { Name: "dotnet-HttpCacheModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Output cache profiles etc. via configuration.", Effort: 1000, Readiness: 5, Impact: "", Category: "caching", Criticality: "",
            Tags:
            []Tag{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/staticContent/clientCache[@cacheControlMode][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cacheControlMode", Recipe: "", },
             }, },
        
            { Name: "dotnet-ISAPI-Filters-config", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1000, Readiness: 0, Impact: "", Category: "Unsupported IIS modules", Criticality: "",
            Tags:
            []Tag{  { Value: "isapi-Filter",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/isapiFilters[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ISAPI-Filters-vbcs", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1000, Readiness: 0, Impact: "", Category: "Unsupported IIS modules", Criticality: "",
            Tags:
            []Tag{  { Value: "isapi-Filter",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "GetSection(system.webServer/isapiFilters)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-MSMQ-vbcs", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 1, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "MessageQueue",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageQueue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "using System.Messaging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-RequestFilteringModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 1, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/requestFiltering/requestLimits[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-StaticCompressionModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuration", Effort: 100, Readiness: 5, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpCompression[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-StaticFileModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuation", Effort: 400, Readiness: 5, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "StaticFileModule",}, { Value: "webserver",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/handlers/add[contains(@modules, \"StaticFileModule\")][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-TokenCacheModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Caches windows security tokens for password based authentication schemes (anonymous authentication; basic authentication; IIS client certificate authentication). Ensure compabible configuration.", Effort: 1000, Readiness: 5, Impact: "", Category: "Unsupported IIS modules", Criticality: "",
            Tags:
            []Tag{  { Value: "TokenCacheModule",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/globalModules/add[@name=\"TokenCacheModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-UriCacheModule", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Implements a generic cache for URL-specific server state; such as configuration. With this module; the server only reads configuration for the first request for a particular URL. And reuse it on subsequent requests until it changes.", Effort: 1000, Readiness: 5, Impact: "", Category: "Unsupported IIS modules", Criticality: "",
            Tags:
            []Tag{  { Value: "UriCache",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/globalModules/add[@name=\"UriCacheModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-WindowsAuth-config", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuation", Effort: 500, Readiness: 1, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//system.webServer/security/authentication/windowsAuthentication[translate(@enabled, \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\") = \"TRUE\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.serviceModel/bindings//binding/security/transport[translate(@clientCredentialType, \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\")=\"WINDOWS\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.web/authentication[translate(@mode, \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\")=\"WINDOWS\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/connectionStrings/add[contains(translate(translate(@connectionString, \" \", \"\"), \"abcdefghijklmnopqrstuvwxyz\", \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\"), \"INTEGRATEDSECURITY=TRUE\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-WindowsAuth-csvb", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Ensure compatible configuation", Effort: 500, Readiness: 1, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "WIndows-Auth",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "GetSection(system.webServer/security/authentication/windowsAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-connectionstrings", FileType: "(yaml|yaml|json|property|config)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Remove connection strings from files, use environment variables (or mount configmap into pod)", Effort: 10, Readiness: 5, Impact: "", Category: "Connection-String", Criticality: "",
            Tags:
            []Tag{  { Value: "Connection-String",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ConnectionString", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-database-access", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Connection strings should be externally managed.", Effort: 100, Readiness: 10, Impact: "", Category: "Database", Criticality: "",
            Tags:
            []Tag{  { Value: "Database",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/connectionStrings[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-db2-unmanaged", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 9, Impact: "", Category: "Database", Criticality: "",
            Tags:
            []Tag{  { Value: "Database",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "IBM.Data.DB2", Advice: "IBM.Data.DB2 can require a special procedure so that the driver's native components are deployed with the application.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "DB2", Recipe: "", },
             }, },
        
            { Name: "dotnet-serilog-elasticsearch", FileType: "(cs|yaml|yml|json)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Make sure to have reachable ELK stack from deployed app", Effort: 5, Readiness: 8, Impact: "", Category: "logging", Criticality: "",
            Tags:
            []Tag{  { Value: "Elasticsearch",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "using Serilog.Sinks.Elasticsearch", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ElasticConfiguration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-file-based-config", FileType: "cs", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Externalize configuration to environment or use ConfigMap as file-mount into a K8S pod", Effort: 10, Readiness: 3, Impact: "", Category: "env-config", Criticality: "",
            Tags:
            []Tag{  { Value: "env-config",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".AddJsonFile(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-fileIO", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 3, Readiness: 8, Impact: "", Category: "I/O", Criticality: "",
            Tags:
            []Tag{  { Value: "I/O",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "File.Append", Advice: "Appending to a file (File.Append*)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Create", Advice: "Calling File.Create", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Move", Advice: "Calling File.Move", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Open", Advice: "Calling File.Open (investigate further to determine refactor rating)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.OpenWrite", Advice: "Calling File.OpenWrite", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Replace", Advice: "Calling File.Replace", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Set", Advice: "Setting File Metadata (File.Set*)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "File.Write", Advice: "Writing to a file (File.Write*)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "new FileStream(", Advice: "Direct construction of FileStream", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FileSystemWatcher", Advice: "Use of FileSystemWatcher", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-filepath", FileType: "(cs|yaml|yml|json)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "don't use local files, log to console or to server if logging, else store data in database", Effort: 8, Readiness: 3, Impact: "", Category: "I/O", Criticality: "",
            Tags:
            []Tag{  { Value: "I/O",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FilePath", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ipv4-addresses", FileType: "(yaml|yml|cs|json|txt)", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Found hard-coded IPv4s. Make configurable, put into environment or config map", Effort: 3, Readiness: 8, Impact: "", Category: "hard-code-ip", Criticality: "",
            Tags:
            []Tag{  { Value: "hard-code-ip",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-launchProcess", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "^ +%s[.(].*", Advice: "Launching additional processes within a container is not recommended.", Effort: 300, Readiness: 7, Impact: "", Category: "Process-Launch", Criticality: "",
            Tags:
            []Tag{  { Value: "Process-Launch",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Process", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-logging", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Logging to the Event Log is not recommended for cloud native apps.", Effort: 100, Readiness: 3, Impact: "", Category: "Logging", Criticality: "",
            Tags:
            []Tag{  { Value: "EventLog",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EventLogTraceListener", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "EventLogTraceListener", Recipe: "", },
             { Type: "", Pattern: "", Value: "EventLog", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-oracle-umanaged", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Oracle unmanaged driver requires including binaries with app", Effort: 3, Readiness: 8, Impact: "", Category: "Database", Criticality: "",
            Tags:
            []Tag{  { Value: "Database",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Oracle.DataAccess", Advice: "Oracle unmanaged driver requires including binaries with app -- can use buildpack.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "Oracle", Recipe: "", },
             }, },
        
            { Name: "dotnet-security", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Relying on Windows certificate stores is problematic in a cloud environment.", Effort: 3, Readiness: 10, Impact: "", Category: "Security", Criticality: "",
            Tags:
            []Tag{  { Value: "Security",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "X509Store", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "X509Store", Recipe: "", },
             }, },
        
            { Name: "dotnet-serilog", FileType: "cs", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Logging with Serilog. Make sure not to log to file. Remote logging sinks need to be reachable on network.", Effort: 3, Readiness: 8, Impact: "", Category: "logging", Criticality: "",
            Tags:
            []Tag{  { Value: "logging",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Serilog.Sinks", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "Serilog", Recipe: "", },
             { Type: "", Pattern: "", Value: ".UseSerilog()", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "Serilog", Recipe: "", },
             }, },
        
            { Name: "dotnet-sharepoint", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "SharePoint is not supported on CloudFoundry.", Effort: 100, Readiness: 0, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "SharePoint",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.SharePoint", Advice: "SharePoint is not supported on CloudFoundry.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-transactions", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Potential use of distributed transactions which are unsupported", Effort: 100, Readiness: 10, Impact: "", Category: "transactions", Criticality: "",
            Tags:
            []Tag{  { Value: "transactions",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TransactionScope", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-protocols-ws", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Unsupported protocols", Effort: 1000, Readiness: 0, Impact: "", Category: "WCF", Criticality: "",
            Tags:
            []Tag{  { Value: "wcf",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<ws.+HttpBinding>", Advice: "Many features of WS* protocols are problematic in the cloud like distributed transactions and reliable sessions", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-protocols", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Unsupported protocols", Effort: 500, Readiness: 0, Impact: "", Category: "WCF", Criticality: "",
            Tags:
            []Tag{  { Value: "udp",}, { Value: "wcf",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<net.+Binding>", Advice: "Non HTTP based protocols are either unsupported or require extensive refactoring when on PCF. TCP binding would require TCP Router to be configured and app to be self hosted (TCP-IIS activation not supported)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-ssl", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "When using HTTPS, terminate SSL at load balancer", Effort: 1, Readiness: 10, Impact: "", Category: "WCF", Criticality: "",
            Tags:
            []Tag{  { Value: "wcf",}, { Value: "ssl",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BasicHttpSecurityMode.Transport", Advice: "Disable HTTPS at the container and allow the external load balancer to terminate SSL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SecurityMode.Transport", Advice: "Transport security at the container is not supported.  Disable transport security on the service", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "mode=\"Transport\"", Advice: "Transport security at the container is not supported.  Disable transport security on the service", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsForms", FileType: "(cs$|vb$|csproj$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Windows Forms module is not supported. Refactor to a web application.", Effort: 10, Readiness: 5, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "Windows-Forms",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Windows.Forms", Advice: "Windows Forms module is not supported. Refactor to a web application.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsPrincipal", FileType: "(cs$|vb$|csproj$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Operations requiring a Windows domain are not supported.", Effort: 1000, Readiness: 8, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "Security-Principal",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Security.Principal.Windows", Advice: "Operations requiring a Windows domain are not supported", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsRegistry", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "^.*%s.*", Advice: "External configurations should be made available by environment variables or some other external service.", Effort: 100, Readiness: 2, Impact: "", Category: "AccessingRegistry", Criticality: "",
            Tags:
            []Tag{  { Value: "Windows-Registry",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RegistryKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Registry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsServices", FileType: "(cs$|vb$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Don't rely on Windows Services as CloudFoundry manages the lifecycle of your service.  Convert any Windows service to a console application to run in Cloud Foundry.", Effort: 0, Readiness: 4, Impact: "", Category: "Windows-Services", Criticality: "",
            Tags:
            []Tag{  { Value: "Windows-Services",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ": ServiceBase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ServiceBase", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceController", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ServiceController", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.ServiceProcess", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ServiceProcess", Recipe: "", },
             }, },
        
            { Name: "faas-meta", FileType: "meta$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "App should be started in the shortest time possible", Effort: 200, Readiness: 2, Impact: "", Category: "boottime", Criticality: "",
            Tags:
            []Tag{  { Value: "faas",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/libraries[@count>13]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "hard-coded-uri-config", FileType: "(yaml|yml|json|cs|java|vb|py|go)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Found hard-coded URI. Make configurable, put into environment or config map", Effort: 3, Readiness: 8, Impact: "", Category: "env-config", Criticality: "",
            Tags:
            []Tag{  { Value: "hardcoded-uri",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "http://", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "https://", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-3rdPartyImports", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 0, Readiness: 0, Impact: "", Category: "third-party", Criticality: "",
            Tags:
            []Tag{  { Value: "third-party",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ehcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ehcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "47deg", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "47deg", Recipe: "", },
             { Type: "", Pattern: "", Value: "advantageous", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "advantageous", Recipe: "", },
             { Type: "", Pattern: "", Value: "agileclick", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "agileclick", Recipe: "", },
             { Type: "", Pattern: "", Value: "agilej", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "agilej", Recipe: "", },
             { Type: "", Pattern: "", Value: "aicer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "aicer", Recipe: "", },
             { Type: "", Pattern: "", Value: "airlift", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "airlift", Recipe: "", },
             { Type: "", Pattern: "", Value: "akka", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "akka", Recipe: "", },
             { Type: "", Pattern: "", Value: "alibaba", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "alibaba", Recipe: "", },
             { Type: "", Pattern: "", Value: "alsocan", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "alsocan", Recipe: "", },
             { Type: "", Pattern: "", Value: "amazon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "amazon", Recipe: "", },
             { Type: "", Pattern: "", Value: "androidtransfuse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "androidtransfuse", Recipe: "", },
             { Type: "", Pattern: "", Value: "angular", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "angular", Recipe: "", },
             { Type: "", Pattern: "", Value: "apache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache", Recipe: "", },
             { Type: "", Pattern: "", Value: "apache.commons", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache.commons", Recipe: "", },
             { Type: "", Pattern: "", Value: "apitrary", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apitrary", Recipe: "", },
             { Type: "", Pattern: "", Value: "apptik", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apptik", Recipe: "", },
             { Type: "", Pattern: "", Value: "aranea-apps", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "aranea-apps", Recipe: "", },
             { Type: "", Pattern: "", Value: "argonaut", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "argonaut", Recipe: "", },
             { Type: "", Pattern: "", Value: "armemodelear", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "armemodelear", Recipe: "", },
             { Type: "", Pattern: "", Value: "arnx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "arnx", Recipe: "", },
             { Type: "", Pattern: "", Value: "asm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "asm", Recipe: "", },
             { Type: "", Pattern: "", Value: "aspectj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "aspectj", Recipe: "", },
             { Type: "", Pattern: "", Value: "asynchttpclient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "asynchttpclient", Recipe: "", },
             { Type: "", Pattern: "", Value: "atmosphere", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "atmosphere", Recipe: "", },
             { Type: "", Pattern: "", Value: "attoparser", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "attoparser", Recipe: "", },
             { Type: "", Pattern: "", Value: "avast", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "avast", Recipe: "", },
             { Type: "", Pattern: "", Value: "aws-java", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "aws-java", Recipe: "", },
             { Type: "", Pattern: "", Value: "axion", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "axion", Recipe: "", },
             { Type: "", Pattern: "", Value: "azure", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "azure", Recipe: "", },
             { Type: "", Pattern: "", Value: "baracus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "baracus", Recipe: "", },
             { Type: "", Pattern: "", Value: "bcel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bcel", Recipe: "", },
             { Type: "", Pattern: "", Value: "beangle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "beangle", Recipe: "", },
             { Type: "", Pattern: "", Value: "beanshell", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "beanshell", Recipe: "", },
             { Type: "", Pattern: "", Value: "bgee", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bgee", Recipe: "", },
             { Type: "", Pattern: "", Value: "blep", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "blep", Recipe: "", },
             { Type: "", Pattern: "", Value: "bluelinelabs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bluelinelabs", Recipe: "", },
             { Type: "", Pattern: "", Value: "bluestemsoftware", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bluestemsoftware", Recipe: "", },
             { Type: "", Pattern: "", Value: "brightify", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "brightify", Recipe: "", },
             { Type: "", Pattern: "", Value: "bsc.bean", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bsc.bean", Recipe: "", },
             { Type: "", Pattern: "", Value: "bsc.framework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bsc.framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "bsf", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bsf", Recipe: "", },
             { Type: "", Pattern: "", Value: "bytebuddy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "bytebuddy", Recipe: "", },
             { Type: "", Pattern: "", Value: "caffeine", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "caffeine", Recipe: "", },
             { Type: "", Pattern: "", Value: "cassandra", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cassandra", Recipe: "", },
             { Type: "", Pattern: "", Value: "ccil", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ccil", Recipe: "", },
             { Type: "", Pattern: "", Value: "celum", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "celum", Recipe: "", },
             { Type: "", Pattern: "", Value: "cemerick", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cemerick", Recipe: "", },
             { Type: "", Pattern: "", Value: "ceylon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ceylon", Recipe: "", },
             { Type: "", Pattern: "", Value: "cglib", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cglib", Recipe: "", },
             { Type: "", Pattern: "", Value: "circe", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "circe", Recipe: "", },
             { Type: "", Pattern: "", Value: "circumflex", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "circumflex", Recipe: "", },
             { Type: "", Pattern: "", Value: "ciris", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ciris", Recipe: "", },
             { Type: "", Pattern: "", Value: "cisco", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cisco", Recipe: "", },
             { Type: "", Pattern: "", Value: "clapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "clapper", Recipe: "", },
             { Type: "", Pattern: "", Value: "clojure", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "clojure", Recipe: "", },
             { Type: "", Pattern: "", Value: "codahale", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "codahale", Recipe: "", },
             { Type: "", Pattern: "", Value: "codehaus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "codehaus", Recipe: "", },
             { Type: "", Pattern: "", Value: "codingwell", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "codingwell", Recipe: "", },
             { Type: "", Pattern: "", Value: "cojen", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cojen", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.m3", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "com.m3", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.sun", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "com.sun", Recipe: "", },
             { Type: "", Pattern: "", Value: "commonjava", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commonjava", Recipe: "", },
             { Type: "", Pattern: "", Value: "commons-configuration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commons-configuration", Recipe: "", },
             { Type: "", Pattern: "", Value: "commons-digester", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commons-digester", Recipe: "", },
             { Type: "", Pattern: "", Value: "commons-logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commons-logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "constretto", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "constretto", Recipe: "", },
             { Type: "", Pattern: "", Value: "contaazul", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "contaazul", Recipe: "", },
             { Type: "", Pattern: "", Value: "critero", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "critero", Recipe: "", },
             { Type: "", Pattern: "", Value: "cronutils", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cronutils", Recipe: "", },
             { Type: "", Pattern: "", Value: "ctrlaltdev", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ctrlaltdev", Recipe: "", },
             { Type: "", Pattern: "", Value: "cyrusinnovation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cyrusinnovation", Recipe: "", },
             { Type: "", Pattern: "", Value: "darwinsys", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "darwinsys", Recipe: "", },
             { Type: "", Pattern: "", Value: "databene", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "databene", Recipe: "", },
             { Type: "", Pattern: "", Value: "databinder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "databinder", Recipe: "", },
             { Type: "", Pattern: "", Value: "datanucleus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "datanucleus", Recipe: "", },
             { Type: "", Pattern: "", Value: "datoin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "datoin", Recipe: "", },
             { Type: "", Pattern: "", Value: "davfx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "davfx", Recipe: "", },
             { Type: "", Pattern: "", Value: "modeltools", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "modeltools", Recipe: "", },
             { Type: "", Pattern: "", Value: "djodjo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "djodjo", Recipe: "", },
             { Type: "", Pattern: "", Value: "dreampie", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dreampie", Recipe: "", },
             { Type: "", Pattern: "", Value: "dropwizard", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dropwizard", Recipe: "", },
             { Type: "", Pattern: "", Value: "dslplatform", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dslplatform", Recipe: "", },
             { Type: "", Pattern: "", Value: "easygson", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "easygson", Recipe: "", },
             { Type: "", Pattern: "", Value: "easymetrics", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "easymetrics", Recipe: "", },
             { Type: "", Pattern: "", Value: "easymock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "easymock", Recipe: "", },
             { Type: "", Pattern: "", Value: "ebay", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ebay", Recipe: "", },
             { Type: "", Pattern: "", Value: "eclipse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "eclipse", Recipe: "", },
             { Type: "", Pattern: "", Value: "eclipsesource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "eclipsesource", Recipe: "", },
             { Type: "", Pattern: "", Value: "eed3si9n", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "eed3si9n", Recipe: "", },
             { Type: "", Pattern: "", Value: "enblom", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "enblom", Recipe: "", },
             { Type: "", Pattern: "", Value: "esotericsoftware", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "esotericsoftware", Recipe: "", },
             { Type: "", Pattern: "", Value: "everit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "everit", Recipe: "", },
             { Type: "", Pattern: "", Value: "exparity", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "exparity", Recipe: "", },
             { Type: "", Pattern: "", Value: "experlog", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "experlog", Recipe: "", },
             { Type: "", Pattern: "", Value: "factati", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "factati", Recipe: "", },
             { Type: "", Pattern: "", Value: "fasterxml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "fasterxml", Recipe: "", },
             { Type: "", Pattern: "", Value: "flexjson", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "flexjson", Recipe: "", },
             { Type: "", Pattern: "", Value: "foursquare", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "foursquare", Recipe: "", },
             { Type: "", Pattern: "", Value: "frege-lang", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "frege-lang", Recipe: "", },
             { Type: "", Pattern: "", Value: "frugalmechanic", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "frugalmechanic", Recipe: "", },
             { Type: "", Pattern: "", Value: "gatling", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "gatling", Recipe: "", },
             { Type: "", Pattern: "", Value: "gilt", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "gilt", Recipe: "", },
             { Type: "", Pattern: "", Value: "github", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "github", Recipe: "", },
             { Type: "", Pattern: "", Value: "glassfish", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "glassfish", Recipe: "", },
             { Type: "", Pattern: "", Value: "gmock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "gmock", Recipe: "", },
             { Type: "", Pattern: "", Value: "goggle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "goggle", Recipe: "", },
             { Type: "", Pattern: "", Value: "golo-lang", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "golo-lang", Recipe: "", },
             { Type: "", Pattern: "", Value: "googlecode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "googlecode", Recipe: "", },
             { Type: "", Pattern: "", Value: "gosu", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "gosu", Recipe: "", },
             { Type: "", Pattern: "", Value: "grails", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "grails", Recipe: "", },
             { Type: "", Pattern: "", Value: "greenrobot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "greenrobot", Recipe: "", },
             { Type: "", Pattern: "", Value: "groovy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "groovy", Recipe: "", },
             { Type: "", Pattern: "", Value: "grouplens", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "grouplens", Recipe: "", },
             { Type: "", Pattern: "", Value: "guava", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "guava", Recipe: "", },
             { Type: "", Pattern: "", Value: "h2database", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "h2database", Recipe: "", },
             { Type: "", Pattern: "", Value: "helger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "helger", Recipe: "", },
             { Type: "", Pattern: "", Value: "henix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "henix", Recipe: "", },
             { Type: "", Pattern: "", Value: "hibernate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hibernate", Recipe: "", },
             { Type: "", Pattern: "", Value: "higgs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "higgs", Recipe: "", },
             { Type: "", Pattern: "", Value: "hmil", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hmil", Recipe: "", },
             { Type: "", Pattern: "", Value: "hmsonline", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hmsonline", Recipe: "", },
             { Type: "", Pattern: "", Value: "hsqlmodel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hsqlmodel", Recipe: "", },
             { Type: "", Pattern: "", Value: "htmlcleaner", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "htmlcleaner", Recipe: "", },
             { Type: "", Pattern: "", Value: "http4s", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "http4s", Recipe: "", },
             { Type: "", Pattern: "", Value: "hydra", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hydra", Recipe: "", },
             { Type: "", Pattern: "", Value: "ikoskela", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ikoskela", Recipe: "", },
             { Type: "", Pattern: "", Value: "ini4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ini4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "inject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "inject", Recipe: "", },
             { Type: "", Pattern: "", Value: "inspiracio", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "inspiracio", Recipe: "", },
             { Type: "", Pattern: "", Value: "intarsys", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "intarsys", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.sari", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "io-sari", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.spray", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "io-spray", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.wcm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "io-wcm", Recipe: "", },
             { Type: "", Pattern: "", Value: "itextpdf", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "itextpdf", Recipe: "", },
             { Type: "", Pattern: "", Value: "j256", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "j256", Recipe: "", },
             { Type: "", Pattern: "", Value: "j4Unit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "j4Unit", Recipe: "", },
             { Type: "", Pattern: "", Value: "jamon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jamon", Recipe: "", },
             { Type: "", Pattern: "", Value: "jasonjson", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jasonjson", Recipe: "", },
             { Type: "", Pattern: "", Value: "javaclub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javaclub", Recipe: "", },
             { Type: "", Pattern: "", Value: "javalite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javalite", Recipe: "", },
             { Type: "", Pattern: "", Value: "javassist", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javassist", Recipe: "", },
             { Type: "", Pattern: "", Value: "jayway", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jayway", Recipe: "", },
             { Type: "", Pattern: "", Value: "jbee", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jbee", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "jcabi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcabi", Recipe: "", },
             { Type: "", Pattern: "", Value: "jcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "jclouds", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jclouds", Recipe: "", },
             { Type: "", Pattern: "", Value: "jconfig", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jconfig", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmodelc-wrappers", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jmodelc-wrappers", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdev", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdev", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdub", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdub", Recipe: "", },
             { Type: "", Pattern: "", Value: "jentz", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jentz", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmemcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jmemcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmetrix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jmetrix", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jmock", Recipe: "", },
             { Type: "", Pattern: "", Value: "jmockring", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jmockring", Recipe: "", },
             { Type: "", Pattern: "", Value: "joda-time", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "joda-time", Recipe: "", },
             { Type: "", Pattern: "", Value: "jodd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jodd", Recipe: "", },
             { Type: "", Pattern: "", Value: "joestelmach", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "joestelmach", Recipe: "", },
             { Type: "", Pattern: "", Value: "jolbox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jolbox", Recipe: "", },
             { Type: "", Pattern: "", Value: "jpattern", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jpattern", Recipe: "", },
             { Type: "", Pattern: "", Value: "jpox", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jpox", Recipe: "", },
             { Type: "", Pattern: "", Value: "jruby", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jruby", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsog", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jsog", Recipe: "", },
             { Type: "", Pattern: "", Value: "json4s", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "json4s", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsonbuddy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jsonbuddy", Recipe: "", },
             { Type: "", Pattern: "", Value: "jsoup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jsoup", Recipe: "", },
             { Type: "", Pattern: "", Value: "jtidy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jtidy", Recipe: "", },
             { Type: "", Pattern: "", Value: "jvnet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jvnet", Recipe: "", },
             { Type: "", Pattern: "", Value: "kasource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kasource", Recipe: "", },
             { Type: "", Pattern: "", Value: "kie", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kie", Recipe: "", },
             { Type: "", Pattern: "", Value: "kirgor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kirgor", Recipe: "", },
             { Type: "", Pattern: "", Value: "kodein", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kodein", Recipe: "", },
             { Type: "", Pattern: "", Value: "kohesive", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kohesive", Recipe: "", },
             { Type: "", Pattern: "", Value: "kohsuke", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kohsuke", Recipe: "", },
             { Type: "", Pattern: "", Value: "kopitubruk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kopitubruk", Recipe: "", },
             { Type: "", Pattern: "", Value: "kotlin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kotlin", Recipe: "", },
             { Type: "", Pattern: "", Value: "ktor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ktor", Recipe: "", },
             { Type: "", Pattern: "", Value: "kuali", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kuali", Recipe: "", },
             { Type: "", Pattern: "", Value: "lambdista", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "lambdista", Recipe: "", },
             { Type: "", Pattern: "", Value: "latte-lang", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "latte-lang", Recipe: "", },
             { Type: "", Pattern: "", Value: "librato", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "librato", Recipe: "", },
             { Type: "", Pattern: "", Value: "liftweb", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "liftweb", Recipe: "", },
             { Type: "", Pattern: "", Value: "lihaoyi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "lihaoyi", Recipe: "", },
             { Type: "", Pattern: "", Value: "littleshoot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "littleshoot", Recipe: "", },
             { Type: "", Pattern: "", Value: "log4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "log4s", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4s", Recipe: "", },
             { Type: "", Pattern: "", Value: "logback", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "logback", Recipe: "", },
             { Type: "", Pattern: "", Value: "loof", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "loof", Recipe: "", },
             { Type: "", Pattern: "", Value: "loopj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "loopj", Recipe: "", },
             { Type: "", Pattern: "", Value: "lowagie", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "lowagie", Recipe: "", },
             { Type: "", Pattern: "", Value: "lyft", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "lyft", Recipe: "", },
             { Type: "", Pattern: "", Value: "m-m-m", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "m-m-m", Recipe: "", },
             { Type: "", Pattern: "", Value: "madgag", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "madgag", Recipe: "", },
             { Type: "", Pattern: "", Value: "madisp", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "madisp", Recipe: "", },
             { Type: "", Pattern: "", Value: "mapmodel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mapmodel", Recipe: "", },
             { Type: "", Pattern: "", Value: "mashape", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mashape", Recipe: "", },
             { Type: "", Pattern: "", Value: "mchange", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mchange", Recipe: "", },
             { Type: "", Pattern: "", Value: "mckoi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mckoi", Recipe: "", },
             { Type: "", Pattern: "", Value: "metamx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metamx", Recipe: "", },
             { Type: "", Pattern: "", Value: "metrics4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "metrics4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "meza", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "meza", Recipe: "", },
             { Type: "", Pattern: "", Value: "minidev", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "minidev", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockachino", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mockachino", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockejb", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mockejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockito", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mockito", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mockk", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockobjects", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mockobjects", Recipe: "", },
             { Type: "", Pattern: "", Value: "mockrunner", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mockrunner", Recipe: "", },
             { Type: "", Pattern: "", Value: "mongo", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mongo", Recipe: "", },
             { Type: "", Pattern: "", Value: "mortbay", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mortbay", Recipe: "", },
             { Type: "", Pattern: "", Value: "mozilla", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mozilla", Recipe: "", },
             { Type: "", Pattern: "", Value: "moznion", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "moznion", Recipe: "", },
             { Type: "", Pattern: "", Value: "msebera", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "msebera", Recipe: "", },
             { Type: "", Pattern: "", Value: "msoliter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "msoliter", Recipe: "", },
             { Type: "", Pattern: "", Value: "mybatis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mybatis", Recipe: "", },
             { Type: "", Pattern: "", Value: "mynah", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mynah", Recipe: "", },
             { Type: "", Pattern: "", Value: "n3twork", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "m3twork", Recipe: "", },
             { Type: "", Pattern: "", Value: "nekohtml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nekohtml", Recipe: "", },
             { Type: "", Pattern: "", Value: "netflix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix", Recipe: "", },
             { Type: "", Pattern: "", Value: "nield", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nield", Recipe: "", },
             { Type: "", Pattern: "", Value: "nikedlab", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nikedlab", Recipe: "", },
             { Type: "", Pattern: "", Value: "ning", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ning", Recipe: "", },
             { Type: "", Pattern: "", Value: "nryotaro", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nryotaro", Recipe: "", },
             { Type: "", Pattern: "", Value: "nscala", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nscala", Recipe: "", },
             { Type: "", Pattern: "", Value: "objectlab", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "objectlab", Recipe: "", },
             { Type: "", Pattern: "", Value: "objectos", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "objectos", Recipe: "", },
             { Type: "", Pattern: "", Value: "objectweb", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "objectweb", Recipe: "", },
             { Type: "", Pattern: "", Value: "ocpsoft", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ocpsoft", Recipe: "", },
             { Type: "", Pattern: "", Value: "oneeyedman", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oneeyedman", Recipe: "", },
             { Type: "", Pattern: "", Value: "openbouquet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "openbouquet", Recipe: "", },
             { Type: "", Pattern: "", Value: "opensynphony", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "opensynphony", Recipe: "", },
             { Type: "", Pattern: "", Value: "ops4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ops4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.json", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "org-json", Recipe: "", },
             { Type: "", Pattern: "", Value: "outerj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "outerj", Recipe: "", },
             { Type: "", Pattern: "", Value: "outr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "outr", Recipe: "", },
             { Type: "", Pattern: "", Value: "ow2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ow2", Recipe: "", },
             { Type: "", Pattern: "", Value: "p6spy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "p6spy", Recipe: "", },
             { Type: "", Pattern: "", Value: "paralleluniverse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "paralleluniverse", Recipe: "", },
             { Type: "", Pattern: "", Value: "parfait", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "parfait", Recipe: "", },
             { Type: "", Pattern: "", Value: "perf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "perf4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "picocontainer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "picocontainer", Recipe: "", },
             { Type: "", Pattern: "", Value: "pistol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pistol", Recipe: "", },
             { Type: "", Pattern: "", Value: "plausible", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "plausible", Recipe: "", },
             { Type: "", Pattern: "", Value: "plexus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "plexus", Recipe: "", },
             { Type: "", Pattern: "", Value: "plukh", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "plukh", Recipe: "", },
             { Type: "", Pattern: "", Value: "pnuts", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pnuts", Recipe: "", },
             { Type: "", Pattern: "", Value: "pojava", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pojava", Recipe: "", },
             { Type: "", Pattern: "", Value: "pojomvcc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pojomvcc", Recipe: "", },
             { Type: "", Pattern: "", Value: "polyforms", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "polyforms", Recipe: "", },
             { Type: "", Pattern: "", Value: "polyjmodelc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "polyjmodelc", Recipe: "", },
             { Type: "", Pattern: "", Value: "powermock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "powermock", Recipe: "", },
             { Type: "", Pattern: "", Value: "proofpoint", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "proofpoint", Recipe: "", },
             { Type: "", Pattern: "", Value: "propensive", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "propensive", Recipe: "", },
             { Type: "", Pattern: "", Value: "proxool", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "proxool", Recipe: "", },
             { Type: "", Pattern: "", Value: "rabbitmq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rabbitmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "remodelrick", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "remodelrick", Recipe: "", },
             { Type: "", Pattern: "", Value: "redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "redix", Recipe: "", },
             { Type: "", Pattern: "", Value: "rhq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rhq", Recipe: "", },
             { Type: "", Pattern: "", Value: "roboguice", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "robguice", Recipe: "", },
             { Type: "", Pattern: "", Value: "rootdev", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rootdev", Recipe: "", },
             { Type: "", Pattern: "", Value: "rubicon", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rubicon", Recipe: "", },
             { Type: "", Pattern: "", Value: "saliman", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "saliman", Recipe: "", },
             { Type: "", Pattern: "", Value: "savant", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "savant", Recipe: "", },
             { Type: "", Pattern: "", Value: "scala-lang", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scacla", Recipe: "", },
             { Type: "", Pattern: "", Value: "scala-tools", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scala-tools", Recipe: "", },
             { Type: "", Pattern: "", Value: "scalaj", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scalaj", Recipe: "", },
             { Type: "", Pattern: "", Value: "scalamock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scalamock", Recipe: "", },
             { Type: "", Pattern: "", Value: "scaldi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "scaldi", Recipe: "", },
             { Type: "", Pattern: "", Value: "segoia", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "segoia", Recipe: "", },
             { Type: "", Pattern: "", Value: "sejda", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sejda", Recipe: "", },
             { Type: "", Pattern: "", Value: "senanque", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "senanque", Recipe: "", },
             { Type: "", Pattern: "", Value: "sf.corn", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sf-corn", Recipe: "", },
             { Type: "", Pattern: "", Value: "sf.json-lib", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sf-json-lib", Recipe: "", },
             { Type: "", Pattern: "", Value: "shadesmodel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "shadesmodel", Recipe: "", },
             { Type: "", Pattern: "", Value: "sharegov", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sharegov", Recipe: "", },
             { Type: "", Pattern: "", Value: "sigwned", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sigwned", Recipe: "", },
             { Type: "", Pattern: "", Value: "sirona", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sirona", Recipe: "", },
             { Type: "", Pattern: "", Value: "skife", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "skife", Recipe: "", },
             { Type: "", Pattern: "", Value: "skinny-framework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "skinny-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "slf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "slf4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "slf4j-api", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "slf4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "smacke", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "smacke", Recipe: "", },
             { Type: "", Pattern: "", Value: "snaq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "snaq", Recipe: "", },
             { Type: "", Pattern: "", Value: "socialmetrix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "socialmetrix", Recipe: "", },
             { Type: "", Pattern: "", Value: "softsmithy", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "softsmithy", Recipe: "", },
             { Type: "", Pattern: "", Value: "softwaremill", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "softwaremill", Recipe: "", },
             { Type: "", Pattern: "", Value: "solarmosaic", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "solarmosaic", Recipe: "", },
             { Type: "", Pattern: "", Value: "soliveirajr", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "soliveirajr", Recipe: "", },
             { Type: "", Pattern: "", Value: "sorm-framework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sorm-framework", Recipe: "", },
             { Type: "", Pattern: "", Value: "soulgalore", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "soulgalore", Recipe: "", },
             { Type: "", Pattern: "", Value: "soywiz", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "soywiz", Recipe: "", },
             { Type: "", Pattern: "", Value: "specs2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "specs2", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring", Recipe: "", },
             { Type: "", Pattern: "", Value: "squareup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "squareup", Recipe: "", },
             { Type: "", Pattern: "", Value: "squeryl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "squeryl", Recipe: "", },
             { Type: "", Pattern: "", Value: "stackmob", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "stackmob", Recipe: "", },
             { Type: "", Pattern: "", Value: "statemonitor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "statemonitor", Recipe: "", },
             { Type: "", Pattern: "", Value: "sterlingcode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sterlingcode", Recipe: "", },
             { Type: "", Pattern: "", Value: "stickycode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "stickycode", Recipe: "", },
             { Type: "", Pattern: "", Value: "strategicgains", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "strategicgains", Recipe: "", },
             { Type: "", Pattern: "", Value: "streametry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "streametry", Recipe: "", },
             { Type: "", Pattern: "", Value: "studiomobile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "studiomobile", Recipe: "", },
             { Type: "", Pattern: "", Value: "sun", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sun", Recipe: "", },
             { Type: "", Pattern: "", Value: "sun.mail", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sun-mail", Recipe: "", },
             { Type: "", Pattern: "", Value: "swinglabs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "swinglabs", Recipe: "", },
             { Type: "", Pattern: "", Value: "symentis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "symentis", Recipe: "", },
             { Type: "", Pattern: "", Value: "tedhi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tedhi", Recipe: "", },
             { Type: "", Pattern: "", Value: "tehuti", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tehuti", Recipe: "", },
             { Type: "", Pattern: "", Value: "time4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "time4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "tophe", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tophe", Recipe: "", },
             { Type: "", Pattern: "", Value: "tornado", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tornado", Recipe: "", },
             { Type: "", Pattern: "", Value: "toucanpdf", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "toucanpdf", Recipe: "", },
             { Type: "", Pattern: "", Value: "trove4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "trove4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "truecommons", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "truecommons", Recipe: "", },
             { Type: "", Pattern: "", Value: "truward", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "truward", Recipe: "", },
             { Type: "", Pattern: "", Value: "turbomanage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "turbomanage", Recipe: "", },
             { Type: "", Pattern: "", Value: "twitter", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "twitter", Recipe: "", },
             { Type: "", Pattern: "", Value: "tynne", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tynee", Recipe: "", },
             { Type: "", Pattern: "", Value: "typesafe", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "typesafe", Recipe: "", },
             { Type: "", Pattern: "", Value: "uberfire", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "uberfire", Recipe: "", },
             { Type: "", Pattern: "", Value: "ujoframework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ujoframework", Recipe: "", },
             { Type: "", Pattern: "", Value: "ujorm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ujorm", Recipe: "", },
             { Type: "", Pattern: "", Value: "unboundid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "unboundid", Recipe: "", },
             { Type: "", Pattern: "", Value: "unimi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "unimi", Recipe: "", },
             { Type: "", Pattern: "", Value: "uniscala", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "uniscala", Recipe: "", },
             { Type: "", Pattern: "", Value: "usikkert", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "usikkert", Recipe: "", },
             { Type: "", Pattern: "", Value: "vaadin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "vaadin", Recipe: "", },
             { Type: "", Pattern: "", Value: "vanillasoure", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "vanillasource", Recipe: "", },
             { Type: "", Pattern: "", Value: "vertx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "vertx", Recipe: "", },
             { Type: "", Pattern: "", Value: "vibur", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "vibur", Recipe: "", },
             { Type: "", Pattern: "", Value: "vvakame", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "vvakame", Recipe: "", },
             { Type: "", Pattern: "", Value: "whirlycott", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "whirlycott", Recipe: "", },
             { Type: "", Pattern: "", Value: "wix", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "wix", Recipe: "", },
             { Type: "", Pattern: "", Value: "wvlet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "wvlet", Recipe: "", },
             { Type: "", Pattern: "", Value: "yammer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "yammer", Recipe: "", },
             { Type: "", Pattern: "", Value: "ymock", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ymock", Recipe: "", },
             { Type: "", Pattern: "", Value: "zalando", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "zalando", Recipe: "", },
             { Type: "", Pattern: "", Value: "zappos", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "zappos", Recipe: "", },
             { Type: "", Pattern: "", Value: "zaxxer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "zaxxer", Recipe: "", },
             { Type: "", Pattern: "", Value: "zwitserloot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "zwitserloot", Recipe: "", },
             }, },
        
            { Name: "java-3rdPartySecurity-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 7, Readiness: 10, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, },
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
        
            { Name: "java-MBeans", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "MBean is application server specific, change to current server equivalent or consider TKG", Effort: 100, Readiness: 10, Impact: "", Category: "websphere", Criticality: "",
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "websphere",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WebSphere:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-processexit", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[.]%s[ (].*$", Advice: "Refer to IBM documentation", Effort: 9, Readiness: 10, Impact: "", Category: "process", Criticality: "",
            Tags:
            []Tag{  { Value: "processexit",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "exit", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "halt", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-activemq", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: ".*[ .]%s[ (].*", Advice: "Remediate any persistence issues", Effort: 7, Readiness: 7, Impact: "", Category: "activemq", Criticality: "",
            Tags:
            []Tag{  { Value: "activemq",}, { Value: "MessageQueue",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "createActiveMQConnection", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-alarmD-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Integrate with new alarmD service", Effort: 0, Readiness: 0, Impact: "", Category: "ThirdParty", Criticality: "",
            Tags:
            []Tag{  { Value: "alarmD",}, { Value: "scheduler",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.tfcci.ucs.manager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-apacheFop-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "[ .]%s[ .({]", Advice: "Usage requires configuration remdiation", Effort: 5, Readiness: 10, Impact: "", Category: "thirdParty", Criticality: "",
            Tags:
            []Tag{  { Value: "apache-fop",}, },
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
        
            { Name: "java-batch", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Remove jndi provider or move to TKG", Effort: 10, Readiness: 7, Impact: "", Category: "jndi", Criticality: "",
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "batch",}, { Value: "webcontainer",}, },
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
        
            { Name: "java-batchAnnotations", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Batch processing can include long running processes", Effort: 10, Readiness: 7, Impact: "", Category: "batch", Criticality: "",
            Tags:
            []Tag{  { Value: "batch",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BatchProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-cache-dist-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Distributed caches must be remediated to function in K8S", Effort: 50, Readiness: 10, Impact: "", Category: "distcache", Criticality: "",
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "cache",}, { Value: "distcache",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.ehcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ehcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.hazelcast", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hazelcast", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.ignite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-ignite", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.infinispan", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "infinispan", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.spy.memcached", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "net-spy-memcached", Recipe: "", },
             }, },
        
            { Name: "java-cache-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Cloud readiness issue as potential state information that is not persisted to a backing service", Effort: 50, Readiness: 10, Impact: "", Category: "nondistcache", Criticality: "",
            Tags:
            []Tag{  { Value: "cache",}, { Value: "nondistcache",}, },
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
        
            { Name: "java-corba", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: ".*[ .]%s[ (].*", Advice: "Replace with cloud-friendly framework or move to TKG", Effort: 10, Readiness: 6, Impact: "", Category: "corba", Criticality: "",
            Tags:
            []Tag{  { Value: "rpc",}, { Value: "corba",}, { Value: "non-standard",}, },
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
        
            { Name: "java-ehcache", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<]", Advice: "Convert to redis", Effort: 100, Readiness: 6, Impact: "", Category: "cache", Criticality: "",
            Tags:
            []Tag{  { Value: "cache",}, { Value: "ehcache",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "CacheConfigurationBuilder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "newCacheConfigurationBuilder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PersistableResourceService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ClusteredCacheIdentifier", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "newResourcePoolsBuilder", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CacheLoaderWriterProvider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "BulkCacheLoadingException", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EhcacheResult", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EhCacheConfiguration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EhCacheImpl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CacheManagerPersistenceConfiguration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-mdb", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Consult MDB documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "mdb",}, { Value: "messaging",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import javax.ejb.MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import javax.ejb.ActivationConfigProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateful-import", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateful",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import javax.ejb.Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateful", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Refer EJB stateful/stateless documentation", Effort: 100, Readiness: 100, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "ejb",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "StatefulTimeout", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateless", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Removing RMI calls from client applications.", Effort: 10, Readiness: 4, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateless",}, },
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
        
            { Name: "java-faces-flow-Annotations", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Batch processing can include long running processes", Effort: 10, Readiness: 7, Impact: "", Category: "batch", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FlowScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Review usage to determine how the customized state of JSF flow is being used and determined if it can be externalized.", Effort: 100, Readiness: 100, Impact: "", Category: "jsf-flow", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.faces.flow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Review usage to determine how the customized state of JSF flow is being used and determined if it can be externalized.", Effort: 10, Readiness: 10, Impact: "", Category: "jsf-flow", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FlowCallNode", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "FlowHandlerFactoryWrapper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-fileIO", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Move to cloud friendly alternative or TKG", Effort: 8, Readiness: 8, Impact: "", Category: "I/O", Criticality: "",
            Tags:
            []Tag{  { Value: "I/O",}, },
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
        
            { Name: "java-file-system", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Use backing service or use TKG", Effort: 100, Readiness: 10, Impact: "", Category: "I/O", Criticality: "",
            Tags:
            []Tag{  { Value: "I/O",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "file://", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-glassfish-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Refer to Weblogic documentation", Effort: 100, Readiness: 10, Impact: "", Category: "oracle", Criticality: "",
            Tags:
            []Tag{  { Value: "glassfish",}, },
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
        
            { Name: "java-handles-term", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "For containerization, the TERM signal must be handled, this pattern is a positive finding.", Effort: 1, Readiness: 6, Impact: "", Category: "termsignal", Criticality: "",
            Tags:
            []Tag{  { Value: "term",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Runtime.getRuntime().addShutdownHook", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-hardIP", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Hardcoded IP addresses are problematic in K8S", Effort: 9, Readiness: 8, Impact: "", Category: "hardip", Criticality: "",
            Tags:
            []Tag{  { Value: "hardip",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-hazelcast", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Convert to redis", Effort: 100, Readiness: 6, Impact: "", Category: "cache", Criticality: "",
            Tags:
            []Tag{  { Value: "hazelcast",}, { Value: "cache",}, },
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
        
            { Name: "java-iop", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (.].*", Advice: "Remote Method Invocations are not cloud native. Move to cloud friendly alternatives such as REST endpoints.", Effort: 10, Readiness: 6, Impact: "", Category: "iop", Criticality: "",
            Tags:
            []Tag{  { Value: "non-standard-protocol",}, { Value: "corba",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import javax.rmi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-rmi", Recipe: "", },
             { Type: "", Pattern: "", Value: "inport org.omg.IOP", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "omg-IOP", Recipe: "", },
             }, },
        
            { Name: "java-java-fx-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Java-fx is not cloud compatible and requires the JRE on the remote device.", Effort: 100, Readiness: 100, Impact: "", Category: "java-fx", Criticality: "",
            Tags:
            []Tag{  { Value: "java-fx",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javafx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jaxrs-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to platform documentation", Effort: 2, Readiness: 8, Impact: "", Category: "jax-rs", Criticality: "",
            Tags:
            []Tag{  { Value: "jaxrs",}, { Value: "rest",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ws.rs", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jaxrs", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.sun.jersey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jersey", Recipe: "", },
             }, },
        
            { Name: "java-jboss", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Convert to Spring based POJOs instead of using container specific functionality", Effort: 50, Readiness: 5, Impact: "", Category: "jboss", Criticality: "",
            Tags:
            []Tag{  { Value: "jboss",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import org.jboss.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jcaAnnotations", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Java messaging can present problems in TAS due to its emphemeral. Convert to a backing service or use TKG.", Effort: 100, Readiness: 7, Impact: "", Category: "jca", Criticality: "",
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "jca",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AdministeredObjectDefinition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionFactoryDefinition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ConnectionFactoryDefinitions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jcache", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Convert to redis or consider TKG", Effort: 50, Readiness: 6, Impact: "", Category: "cache", Criticality: "",
            Tags:
            []Tag{  { Value: "cache",}, { Value: "jcache",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getConfigurableCacheFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamedCache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ReflectionExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ChainedExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jersey-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Refer to 3rd party organization for cloud affinity of library", Effort: 5, Readiness: 8, Impact: "", Category: "jersey", Criticality: "",
            Tags:
            []Tag{  { Value: "jersey",}, { Value: "rest",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.sun.jersey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jetty-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Use Managed Executor", Effort: 100, Readiness: 10, Impact: "", Category: "threading", Criticality: "",
            Tags:
            []Tag{  { Value: "threading",}, { Value: "jetty",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.eclipse.jetty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jks", FileType: "jks$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "", Effort: 7, Readiness: 7, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, { Value: "jks",}, { Value: "mutual-auth",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*\\.jks", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jms", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.].*", Advice: "Run embedded service broker as a JMS Provider.", Effort: 10, Readiness: 6, Impact: "", Category: "jms", Criticality: "",
            Tags:
            []Tag{  { Value: "jms",}, { Value: "MessageQueue",}, },
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
        
            { Name: "java-jndi", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Remove jndi provider or move to TKG", Effort: 10, Readiness: 7, Impact: "", Category: "jndi", Criticality: "",
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "webcontainer",}, },
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
        
            { Name: "java-jni", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*%s\\s*native\\s*", Advice: "A few conditions have to be met to make JNI calls", Effort: 1000, Readiness: 7, Impact: "", Category: "jni", Criticality: "",
            Tags:
            []Tag{  { Value: "native",}, { Value: "jni",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "public", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "private", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "static", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jpa", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ @.]%s[ (].*", Advice: "JPA will work inside of Cloud Native applications, but changes may need to be made to the way JPA is configured to get connections.", Effort: 2, Readiness: 6, Impact: "", Category: "jpa", Criticality: "",
            Tags:
            []Tag{  { Value: "jpa",}, { Value: "webprofile",}, { Value: "entity-manager",}, },
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
        
            { Name: "java-jsf", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*$", Advice: "The conversion of a JSF application to a Spring Boot application requires several steps", Effort: 5, Readiness: 4, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "ui",}, },
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
        
            { Name: "java-jsp", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "You can also host legacy Java Web applications with JSPs on TAS", Effort: 2, Readiness: 5, Impact: "", Category: "jsp", Criticality: "",
            Tags:
            []Tag{  { Value: "ui",}, { Value: "jsp",}, { Value: "webcontainer",}, },
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
        
            { Name: "java-jta", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (.{].*", Advice: "Distributed transactions are problematic and should be remediated.", Effort: 50, Readiness: 6, Impact: "file", Category: "jta", Criticality: "",
            Tags:
            []Tag{  { Value: "transactions",}, { Value: "jta",}, { Value: "webprofile",}, },
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
        
            { Name: "java-jvm-runtimeConfigProps", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Do not change these properties at runtime in application code", Effort: 50, Readiness: 0, Impact: "", Category: "config", Criticality: "",
            Tags:
            []Tag{  { Value: "config",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.xml.soap.SOAPFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "soap", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.parsers.DocumentBuilderFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml-parser", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.xml.parsers.SAXParserFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml-parser", Recipe: "", },
             }, },
        
            { Name: "java-logging-file-appender", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Replace file appender with console appender", Effort: 3, Readiness: 5, Impact: "", Category: "logging", Criticality: "",
            Tags:
            []Tag{  { Value: "logging",}, { Value: "log2file",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "fileappender", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-logging-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Change to an implementation of SLF4J i.e. Logback", Effort: 3, Readiness: 5, Impact: "", Category: "logging", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "logging",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java.util.logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java-util", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.log4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "log4j", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.commons.logging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "commons-logging", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.osgi.service.log", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "org-osgi", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.jboss.logging.Logger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-logging", Recipe: "", },
             }, },
        
            { Name: "java-message-driven-annotations", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "To convert a message driven bean to spring cloud stream with rabbitmq", Effort: 5, Readiness: 0, Impact: "", Category: "annotations", Criticality: "",
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "message-driven",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "message-driven", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActivationConfigProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "message-driven", Recipe: "", },
             }, },
        
            { Name: "java-messageDrivenBeans", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 10, Impact: "", Category: "mdb", Criticality: "",
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJB\\(lookup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-metrics", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Indicates use of a metrics collection library, which supports containerization", Effort: 1, Readiness: 7, Impact: "", Category: "metrics", Criticality: "",
            Tags:
            []Tag{  { Value: "metrics",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "io.micrometer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "micrometer", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.management", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-management", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.dropwizard", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "dropwizard", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netuitive.ananke", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netuitive", Recipe: "", },
             { Type: "", Pattern: "", Value: "edu.iris.dmc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "iris-dmc", Recipe: "", },
             }, },
        
            { Name: "java-mqseries", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "You need to make sure that you are using the dependencies that match the version of IBM MQ on the server", Effort: 7, Readiness: 7, Impact: "", Category: "ibm-mq", Criticality: "",
            Tags:
            []Tag{  { Value: "ibm-mq",}, { Value: "MessageQueue",}, },
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
        
            { Name: "java-mulesoft-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "There are several changes required to move a Mule Project to PCF", Effort: 100, Readiness: 10, Impact: "", Category: "etl", Criticality: "",
            Tags:
            []Tag{  { Value: "mulesoft",}, { Value: "etl",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.mule", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mulesoft-intf", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ {].*", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "etl", Criticality: "",
            Tags:
            []Tag{  { Value: "mulesoft",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AbstractMessageTransformer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AbstractTransformer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-netflix-healthcheck", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Indicates existance of healthcheck endpoint, which is positive finding", Effort: -100, Readiness: 7, Impact: "", Category: "healthcheck", Criticality: "",
            Tags:
            []Tag{  { Value: "healthcheck",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "netflix.karyon.transport.http.health.HealthCheckEndpoint", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "karyon-healthcheck", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.guice", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-healthcheck", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.api.HealthCheckAggregator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-healthcheck", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.api.HealthCheckStatus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-healthcheck", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.netflix.runtime.health.api.IndicatorMatcher", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "netflix-healthcheck", Recipe: "", },
             }, },
        
            { Name: "java-nio", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Storage on TAS is emphemerol, so nio based actions may not behave well", Effort: 8, Readiness: 8, Impact: "", Category: "nio", Criticality: "",
            Tags:
            []Tag{  { Value: "nio",}, },
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
        
            { Name: "java-nonstandard-protocol", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Ensure protocol dependencies are cloud friendly or move to TKG", Effort: 100, Readiness: 10, Impact: "", Category: "protocol", Criticality: "",
            Tags:
            []Tag{  { Value: "nonstandard-protocol",}, { Value: "api",}, },
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
        
            { Name: "java-persistence", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 5, Readiness: 5, Impact: "", Category: "persistence", Criticality: "",
            Tags:
            []Tag{  { Value: "javax-persistence",}, { Value: "I/o",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.persistence", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-portUsage", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^*.(%s?:.*):(\\d*)\\/?(.*)", Advice: "Ensure port usage is cloud-friendly or use TKG", Effort: 5, Readiness: 6, Impact: "", Category: "portUsage", Criticality: "",
            Tags:
            []Tag{  { Value: "portUsage",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "https", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "http", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-rabbitmq-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to PCF/Spring documentation for remediation", Effort: 5, Readiness: 5, Impact: "", Category: "mq", Criticality: "",
            Tags:
            []Tag{  { Value: "MessageQueue",}, { Value: "rabbitmq",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.rabbitmq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-remoteEJB", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Refer to TASS/Spring documentation for remediation or move to TKG", Effort: 10, Readiness: 10, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "ejb",}, { Value: "fullprofile",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RemoteHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "remote", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBObject", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "EJBHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-remoteWebService-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Ensure valid configuration for TAS or use TKG", Effort: 10, Readiness: 10, Impact: "", Category: "webService", Criticality: "",
            Tags:
            []Tag{  { Value: "webservices",}, },
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
        
            { Name: "java-resource-cci", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consult with TAS documetation for configuration or use TKG", Effort: 7, Readiness: 7, Impact: "", Category: "jca", Criticality: "",
            Tags:
            []Tag{  { Value: "cci",}, { Value: "resource-adapter",}, },
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
        
            { Name: "java-resource-spi", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Migrate to backing service or move to TKG", Effort: 10, Readiness: 7, Impact: "", Category: "jca", Criticality: "",
            Tags:
            []Tag{  { Value: "resource-adapter",}, { Value: "spi",}, { Value: "jca",}, },
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
        
            { Name: "java-restlet-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "The Restlet library appears to be dead at this point in time.", Effort: 10, Readiness: 7, Impact: "", Category: "restlet", Criticality: "",
            Tags:
            []Tag{  { Value: "rest",}, { Value: "restlet",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.restlet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-rpc-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Adapt to cloud friendly protocl or use TKG", Effort: 2, Readiness: 10, Impact: "", Category: "webService", Criticality: "",
            Tags:
            []Tag{  { Value: "webservices",}, { Value: "rpc",}, { Value: "soap",}, },
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
        
            { Name: "java-security-annotations", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Replace with spring security org.springframework.security.access.annotation or move to TKG", Effort: 5, Readiness: 0, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DeclareRoles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DenyAll", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PermitAll", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RolesAllowed", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-security", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "", Effort: 10, Readiness: 0, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, },
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
        
            { Name: "java-servlet-session", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "HTTP Sesssion Java API Import. Typically remediated by exernalizing session state.", Effort: 10, Readiness: 0, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "session",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import javax.servlet.http.HttpSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-servlet", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Servlet Java API Import", Effort: 0, Readiness: 0, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import javax.servlet.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-slf4j-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Review logging configuration and remove file appenders.", Effort: 0, Readiness: 0, Impact: "", Category: "slf4j", Criticality: "",
            Tags:
            []Tag{  { Value: "slf4j",}, { Value: "logging",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.slf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-soap", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.* %s.*[(].*", Advice: "There occasionally exists use cases where a message interceptor is required.", Effort: 4, Readiness: 3, Impact: "", Category: "soap", Criticality: "",
            Tags:
            []Tag{  { Value: "webservices",}, { Value: "soap",}, },
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
        
            { Name: "java-springboot-annotations", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Spring Boot is a positive score", Effort: -100, Readiness: 0, Impact: "", Category: "spring-boot", Criticality: "",
            Tags:
            []Tag{  { Value: "spring-boot",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "SpringBootApplication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-springframework", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Presence of spring framework may indicate the app should target TAS", Effort: 0, Readiness: 0, Impact: "", Category: "springFramework", Criticality: "",
            Tags:
            []Tag{  { Value: "spring",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import org.springframework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import org.springframework.mongo", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "mongo", Recipe: "", },
             { Type: "", Pattern: "", Value: "import org.springframework.cassandra", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "cassandra", Recipe: "", },
             }, },
        
            { Name: "java-stateful-annotations", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Relies on container for shared state and will require a rewrite to Stateless or move to TKG", Effort: 5, Readiness: 0, Impact: "", Category: "stateful", Criticality: "",
            Tags:
            []Tag{  { Value: "stateful",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Init", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Remove", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PostActivate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "PrePassivate", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "Stateful", Recipe: "", },
             }, },
        
            { Name: "java-stateless-annotations", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Replace with the Spring annotation @Component or move to TKG", Effort: 5, Readiness: 0, Impact: "", Category: "stateless", Criticality: "",
            Tags:
            []Tag{  { Value: "stateless",}, { Value: "ejb",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateless", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             }, },
        
            { Name: "java-struts-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to platform documentation", Effort: 6, Readiness: 7, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, { Value: "webapp",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.apache.struts", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.struts2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.opensymphony", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "opensymphony", Recipe: "", },
             }, },
        
            { Name: "java-struts", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Refer to platform documentation", Effort: 6, Readiness: 7, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, { Value: "webapp",}, { Value: "webcontainer",}, },
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
        
            { Name: "java-swing", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Swing applications cannot move to the cloud", Effort: 100, Readiness: 100, Impact: "", Category: "Swing", Criticality: "",
            Tags:
            []Tag{  { Value: "desktop-app",}, { Value: "ui",}, },
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
        
            { Name: "java-system-config", FileType: "$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Review usage of environment variables and system properties and externalize.", Effort: 3, Readiness: 10, Impact: "", Category: "config", Criticality: "",
            Tags:
            []Tag{  { Value: "config",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.getenv(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.getProperty(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.setProperty(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-systemLoad", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Remediate to cloud friendly implentation or move to TKG", Effort: 1000, Readiness: 10, Impact: "", Category: "Process", Criticality: "",
            Tags:
            []Tag{  { Value: "jni",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.loadLibrary", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "System.load", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-tangosol", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Convert to redis", Effort: 10, Readiness: 6, Impact: "", Category: "cache", Criticality: "",
            Tags:
            []Tag{  { Value: "cache",}, { Value: "tangosol",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "getConfigurableCacheFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamedCache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ReflectionExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ChainedExtractor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-threadUsage-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Use Managed Executor", Effort: 100, Readiness: 10, Impact: "", Category: "threading", Criticality: "",
            Tags:
            []Tag{  { Value: "threading",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java.lang.Thread", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "java.lang.Runnable", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-tibco-jms", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Integrating with TIBCO BusinessWorks JMS queues from a Spring application requires vendor-specific implementation", Effort: 7, Readiness: 6, Impact: "", Category: "tibco", Criticality: "",
            Tags:
            []Tag{  { Value: "tibco",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TibjmsConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-transaction-annotations", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Distributed transcations are problematic and should be remediated or use TKG", Effort: 5, Readiness: 0, Impact: "", Category: "transaction", Criticality: "",
            Tags:
            []Tag{  { Value: "transaction",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Remote", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "LocalHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "RemoteHome", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionManagement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "TransactionAttribute", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-transportSecurity", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Not compatable with TAS, change to TransportSecurity.NONE or move to TKG", Effort: 10, Readiness: 10, Impact: "", Category: "Security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, { Value: "transportSecurity",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TransportGuarantee.CONFIDENTIAL", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-weblogic", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Move away from container specific APIs to portal APIs such as Spring", Effort: 50, Readiness: 10, Impact: "", Category: "webLogic", Criticality: "",
            Tags:
            []Tag{  { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import weblogic.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-websockets-import", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Refer to platform documentation", Effort: 100, Readiness: 100, Impact: "", Category: "websockets", Criticality: "",
            Tags:
            []Tag{  { Value: "websocket",}, { Value: "I/O",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.websocket", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.websocket.Endpoint", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.vertx.core.http.ServerWebSocket;", Advice: "", Effort: 100, Readiness: 0, Criticality: "", Category: "", Tag: "vertx", Recipe: "", },
             }, },
        
            { Name: "java-ws2liberty-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to IBM documentation", Effort: 10, Readiness: 10, Impact: "", Category: "webSphere", Criticality: "",
            Tags:
            []Tag{  { Value: "liberty",}, },
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
        
            { Name: "java-ws2liberty-methods", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Refer to IBM documentation", Effort: 10, Readiness: 10, Impact: "", Category: "webSphere", Criticality: "",
            Tags:
            []Tag{  { Value: "websphere",}, { Value: "liberty",}, { Value: "ltpa",}, { Value: "security",}, },
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
        
            { Name: "log2file-import", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Logging should be to console", Effort: 1, Readiness: 8, Impact: "", Category: "log2file", Criticality: "",
            Tags:
            []Tag{  { Value: "log2file",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java.util.logging.FileHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java.util.logging", Recipe: "", },
             }, },
        
            { Name: "log4j-properties", FileType: "log4j2.properties$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1, Readiness: 0, Impact: "", Category: "log2file", Criticality: "",
            Tags:
            []Tag{  { Value: "log2file",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "property.filename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "log4j-xml", FileType: "log4j2.xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1, Readiness: 0, Impact: "", Category: "log2file", Criticality: "",
            Tags:
            []Tag{  { Value: "log2file",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "property.filename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "plaintext-creds", FileType: "", Target: "line", Type: "contains", DefaultPattern: "", Advice: "never save passwords or login information in files", Effort: 11, Readiness: 9, Impact: "", Category: "Security", Criticality: "",
            Tags:
            []Tag{  { Value: "Security",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Password=", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "User=", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "User Id=", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-cf", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Check for cloud foundry support.", Effort: -10, Readiness: 10, Impact: "", Category: "cloud-foundry", Criticality: "",
            Tags:
            []Tag{  { Value: "cloud-ready",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "load_from_vcap_services", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-db-peewee", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Check for cloud foundry support.", Effort: 4, Readiness: 10, Impact: "", Category: "database", Criticality: "",
            Tags:
            []Tag{  { Value: "peewee",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "peewee", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-db-redis", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Check for cloud foundry support.", Effort: 4, Readiness: 10, Impact: "", Category: "database", Criticality: "",
            Tags:
            []Tag{  { Value: "redis",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-fileIO", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform. Since containers are immutable, restarts will lose any written changes. Refactor this logic to use an external service to store state.", Effort: 8, Readiness: 8, Impact: "", Category: "I/O", Criticality: "",
            Tags:
            []Tag{  { Value: "I/O",}, },
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
        
            { Name: "python-rabbitmq", FileType: "py$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Check for cloud foundry support.", Effort: 4, Readiness: 10, Impact: "", Category: "services", Criticality: "",
            Tags:
            []Tag{  { Value: "rabbitmq",}, },
            Recipes:
            []Recipe{  { URI: "https://app-transformation-cookbook-internal.cfapps.io/backing-services-port-binding/persistence/", },  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import pika", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import aio_pika", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-redis", FileType: "py$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Check for cloud foundry support.", Effort: 4, Readiness: 10, Impact: "", Category: "services", Criticality: "",
            Tags:
            []Tag{  { Value: "redis",}, },
            Recipes:
            []Recipe{  { URI: "https://app-transformation-cookbook-internal.cfapps.io/backing-services-port-binding/persistence/", },  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import aioredis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-sqlite", FileType: "py$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Migrate to an external database.", Effort: 6, Readiness: 8, Impact: "", Category: "DatabaseAccess", Criticality: "",
            Tags:
            []Tag{  { Value: "database",}, { Value: "sqlite",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "sqlite", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SqliteDatabase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "sqlserver-ssis", FileType: "(dtsx$)", Target: "line", Type: "contains", DefaultPattern: "", Advice: "SSIS is not supported on CloudFoundry.", Effort: 100, Readiness: 0, Impact: "", Category: "Unsupported modules", Criticality: "",
            Tags:
            []Tag{  { Value: "ssis",}, { Value: "etl",}, { Value: "sql",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DTS", Advice: "SSIS is not supported on CloudFoundry. Consider leaving the packages in an external SQL Server deployment or rewrite them in a cloud native ETL Framework like Spring Cloud Data Flow.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "weblogic-cluster-config", FileType: "conf$", Target: "line", Type: "regex", DefaultPattern: "^%s.*", Advice: "weblogic clusters cannot run in K8S", Effort: 1, Readiness: 0, Impact: "", Category: "wlcluster", Criticality: "",
            Tags:
            []Tag{  { Value: "wlcluster",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WebLogicCluster", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "websphere-cluster-jacl", FileType: "jacl$", Target: "line", Type: "regex", DefaultPattern: "^ *%s.*", Advice: "websphere clusters cannot run in K8S", Effort: 1, Readiness: 0, Impact: "", Category: "wscluster", Criticality: "",
            Tags:
            []Tag{  { Value: "wscluster",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "_CLUSTERS", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "wsdl-soap", FileType: "wsdl$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 4, Readiness: 3, Impact: "", Category: "soap", Criticality: "",
            Tags:
            []Tag{  { Value: "wsdl",}, { Value: "webservices",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*\\.wsdl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-activeMQ", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Convert to rabbitmq or use TKG", Effort: 5, Readiness: 10, Impact: "", Category: "mq", Criticality: "",
            Tags:
            []Tag{  { Value: "amqp",}, { Value: "MessageQueue",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "amq:broker", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-clientCert", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Avoid reliance on SSL", Effort: 5, Readiness: 10, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "certificates",}, { Value: "auth",}, { Value: "security",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<auth-method>CLIENT-CERT", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ehcache-Config", FileType: "xml$", Target: "file", Type: "contains", DefaultPattern: "", Advice: "Consult 3rd party documentation", Effort: 50, Readiness: 10, Impact: "", Category: "cache", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "cache",}, { Value: "ehcache",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ehcache.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-2-1", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "fullprofile",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_2_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-0", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "javaee",}, { Value: "ejb-lite",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-1", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "ejb-lite",}, { Value: "javaee",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-2", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "ejb-lite",}, { Value: "javaee",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-mdb-import", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "MDB", Criticality: "",
            Tags:
            []Tag{  { Value: "MessageDriven",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "fullprofile",}, { Value: "api",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mdb", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Consider stateless alternative or move to TKG", Effort: 7, Readiness: 7, Impact: "", Category: "MDB", Criticality: "",
            Tags:
            []Tag{  { Value: "MessageDriven",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "fullprofile",}, { Value: "api",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-remote", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 100, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "RemoteEJB",}, { Value: "ejb",}, { Value: "api",}, { Value: "javaee",}, { Value: "fullprofile",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<remote>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-resource-mgr-aut", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, { Value: "webapp",}, { Value: "api",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<res-auth>Container</res-auth>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-facelets", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 2, Readiness: 2, Impact: "", Category: "Facelets", Criticality: "",
            Tags:
            []Tag{  { Value: "facelets",}, { Value: "jsf",}, { Value: "webapp",}, { Value: "api",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jsf/facelets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jboss", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 50, Readiness: 5, Impact: "", Category: "Jboss", Criticality: "",
            Tags:
            []Tag{  { Value: "jboss",}, { Value: "javaee",}, { Value: "api",}, },
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
        
            { Name: "xml-jee", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Convert to Spring based application configuration or use importResource", Effort: 20, Readiness: 8, Impact: "", Category: "Config", Criticality: "",
            Tags:
            []Tag{  { Value: "file",}, { Value: "api",}, { Value: "fullprofile",}, { Value: "javaee",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "application.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ear", Recipe: "", },
             { Type: "", Pattern: "", Value: "application-client.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "application-client", Recipe: "", },
             { Type: "", Pattern: "", Value: "ejb-jar.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             { Type: "", Pattern: "", Value: "ra.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "resourceadapter", Recipe: "", },
             { Type: "", Pattern: "", Value: "webservices.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "webservices", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-1-0", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-1-1", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-1-2", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-0", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-1", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-2", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-3", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Criticality: "",
            Tags:
            []Tag{  { Value: "api",}, { Value: "jsf",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_3.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-localJNDI", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "WebSphere does not allow local JNDI, refer to documentation", Effort: 10, Readiness: 10, Impact: "", Category: "jndi", Criticality: "",
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "local",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<local-jndi-name>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-messageDrivenBeans", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "mdb", Criticality: "",
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "mdb",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<ejb-ref>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-myfaces", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 5, Readiness: 2, Impact: "", Category: "myfaces", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.apache.myfaces", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-portlet-1-0", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 500, Readiness: 700, Impact: "", Category: "version-portlet", Criticality: "",
            Tags:
            []Tag{  { Value: "portlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "portlet-app_1_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-3", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_3.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-4", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_4.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-5", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 400, Readiness: 500, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_5.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-3-0", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 8, Readiness: 10, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_3_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-3-1", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 5, Readiness: 5, Impact: "", Category: "servlet", Criticality: "",
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "webcontainer",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_3_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-session-scoped-beans", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Review usage and determine if externalizing session state will resolve.", Effort: 100, Readiness: 100, Impact: "", Category: "session", Criticality: "",
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "session",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "scope=\"session\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-statefulEJB", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Stateful EJBs may be better suited to TKG", Effort: 1000, Readiness: 1000, Impact: "", Category: "ejb", Criticality: "",
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateful",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<session-type>Stateful</session-type>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-1-1", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-config_1_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-2", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "ui",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.2.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-3", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.3.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-4", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "ui",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.4.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-2-5", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "struts", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, { Value: "webapp",}, { Value: "ui",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-2.5.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts-tiles", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 5, Readiness: 2, Impact: "", Category: "tiles", Criticality: "",
            Tags:
            []Tag{  { Value: "struts",}, { Value: "webapp",}, { Value: "ui",}, { Value: "tiles",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts.apache.org/tags-tiles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_2_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_2_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_3_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-tomahawk", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 600, Readiness: 200, Impact: "", Category: "tomahawk", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "webprofile",}, { Value: "tomahawk",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "myfaces.apache.org/tomahawk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-tomcat", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 0, Impact: "", Category: "Tomcat", Criticality: "",
            Tags:
            []Tag{  { Value: "tomcat",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "configure.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "context.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "server.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-transportSecurity", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Change to NONE", Effort: 2, Readiness: 10, Impact: "", Category: "security", Criticality: "",
            Tags:
            []Tag{  { Value: "security",}, { Value: "webapp",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<transport-guarantee>CONFIDENTIAL", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-trinidad", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 10, Readiness: 7, Impact: "", Category: "trinidad", Criticality: "",
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "webprofile",}, { Value: "trinidad",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "myfaces.apache.org/trinidad/config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-2", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.2/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-3", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.3/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-4", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.4/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-5", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 700, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.5/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-weblogic-1-6", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 600, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.6/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-7", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 600, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.7/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-8", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 400, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.8/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webLogic-1-9", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 400, Readiness: 700, Impact: "", Category: "weblogic-webapp", Criticality: "",
            Tags:
            []Tag{  { Value: "webapp",}, { Value: "weblogic",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1.9/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-weblogic", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Convert to Spring based application configuration", Effort: 100, Readiness: 100, Impact: "", Category: "webLogic", Criticality: "",
            Tags:
            []Tag{  { Value: "weblogic",}, { Value: "fullprofile",}, },
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
        
            { Name: "xml-webprofile", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Web application config file", Effort: 50, Readiness: 1000, Impact: "", Category: "config", Criticality: "",
            Tags:
            []Tag{  { Value: "webprofile",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "persistence.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jpa", Recipe: "", },
             { Type: "", Pattern: "", Value: "web.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "webapp", Recipe: "", },
             { Type: "", Pattern: "", Value: "applicationContext.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "FileSystemXmlApplicationContext", Recipe: "", },
             }, },
        
            { Name: "xml-websphere", FileType: "xm[li]$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Convert to Spring based application configuration", Effort: 500, Readiness: 1000, Impact: "", Category: "webSphere", Criticality: "",
            Tags:
            []Tag{  { Value: "websphere",}, { Value: "fullprofile",}, },
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
        
            { Name: "xml-xa-dataSource", FileType: "xml$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Refer to documetation", Effort: 1000, Readiness: 10, Impact: "", Category: "datasource", Criticality: "",
            Tags:
            []Tag{  { Value: "transactions",}, { Value: "jta",}, { Value: "webprofile",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<xa-datasource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "yaml-test", FileType: "(yaml|yml)$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "", Effort: 1000, Readiness: 10, Impact: "", Category: "", Criticality: "",
            Tags:
            []Tag{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.orchestration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
    }
    return BootstrapRules
}