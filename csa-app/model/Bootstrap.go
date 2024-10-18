/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapRulesTemplate.txt found under go/resources folder
//Created @ 2024-10-18 10:32:29.365298 -0500 CDT m=+0.315640142

func BootstrapRules() []Rule {
    var BootstrapRules = []Rule{
        
            { Name: "java-glassfish-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Refer to Weblogic documentation", Effort: 100, Readiness: 10, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "glassfish",}, { Value: "container",}, { Value: "modernize",}, { Value: "weblogic",}, },
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
        
            { Name: "java-jboss", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(import)\\s+(%s)\\.{1}", Advice: "Convert to Spring based POJOs instead of using container specific functionality", Effort: 50, Readiness: 5, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jboss",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.jboss", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jetty-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$|xml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Use Managed Executor", Effort: 1, Readiness: 10, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "threading",}, { Value: "jetty",}, { Value: "tas",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.eclipse.jetty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.mortbay.jetty.Server.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-tomcat", RuleType: "", Level: "", FileType: "(xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Spring Boot Tomcat", Effort: 0, Readiness: 9, Impact: "", Category: "app-web-server", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "tomcat",}, { Value: "webserver",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<artifactId>spring-boot-starter-web<\\/artifactId>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<artifactId>spring-boot-starter-tomcat<\\/artifactId>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-undertow", RuleType: "", Level: "", FileType: "(java$|yml$|yaml$|scala$|xml$|properties$)", Target: "line", Type: "regex", DefaultPattern: "(?i)%s", Advice: "Usage of spring undertow", Effort: 0, Readiness: 10, Impact: "", Category: "other", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "app-web-server",}, { Value: "undertow",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*io.*undertow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-io-undertow-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.*undertow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-io-undertow-detect", Recipe: "", },
             { Type: "", Pattern: "", Value: "import.*spring.*undertow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-undertow-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring.*undertow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-undertow-detect", Recipe: "", },
             }, },
        
            { Name: "java-tomcat", RuleType: "", Level: "", FileType: "(xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Tomcat", Effort: 0, Readiness: 9, Impact: "", Category: "app-web-server", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "tomcat",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*org.apache.catalina.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.apache.tomcat.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*tomcat-dbcp.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*tomcat-jdbc.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*tomcat-embed-jasper.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*tomcat-embed-core.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*tomcat-embed-websocket.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*tomcat7-maven-plugin.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-weblogic", RuleType: "standard", Level: "", FileType: "(java$|conf$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 100, Readiness: 10, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "weblogic",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "weblogic.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WebLogicCluster", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.bea.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.ucp", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "weblogicx.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "websphere", RuleType: "standard", Level: "", FileType: "(java$|jacl$|jsp$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Vendor proprietary implementation.  Consider rearchitecting if decision is made to move off application server. Websphere clusters cannot run in K8S", Effort: 1, Readiness: 0, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "websphere",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "_CLUSTERS", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getCallerList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getFirstCaller", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getFirstServer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getServerList", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "addPropagationAttribute", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getPropagationAttributes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "convertCookieStringToBytes", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "revokeSSOCookies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "revokeSSOCookiesForPortlets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "getLTPACookieFromSSOToken", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.websphere", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websphere", Recipe: "", },
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
        
            { Name: "xml-jboss", RuleType: "standard", Level: "", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Application Server coupling", Effort: 50, Readiness: 5, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jboss",}, { Value: "javaee",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jaws.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jbosscmp-jmodelc.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-service.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-web.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jboss-wsse-server.xml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-tomcat", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Web Server coupling", Effort: 100, Readiness: 0, Impact: "", Category: "app-web-server", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "tomcat",}, { Value: "tas",}, { Value: "container",}, },
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
        
            { Name: "xml-weblogic", RuleType: "standard", Level: "", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 700, Readiness: 700, Impact: "", Category: "app-web-server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "weblogic",}, { Value: "container",}, { Value: "modernize",}, },
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
             { Type: "", Pattern: "", Value: "1.2/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.3/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.4/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.5/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.6/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.7/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.8/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "1.9/weblogic-web-app.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-websphere", RuleType: "fire-once", Level: "", FileType: "xm[li]$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 500, Readiness: 1000, Impact: "", Category: "app-web-server", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "websphere",}, { Value: "full-profile",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-apm-dependency-maven", RuleType: "standard", Level: "Info", FileType: "xml$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "The application has integrated an Application Performance Management (APM) tool as a dependency", Effort: 0, Readiness: 0, Impact: "", Category: "apm", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "apm",}, { Value: "monitoring",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-mq-connection-string-user-password-properties", RuleType: "standard", Level: "Info", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Message queue connection string, username or password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "message queue", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "message queue",}, { Value: "messaging",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-mq-connection-string-user-password-yaml", RuleType: "standard", Level: "Info", FileType: "(yaml$|yml$|json$|jsn$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "Message queue connection string, username or password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "message queue", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "message queue",}, { Value: "messaging",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-non-lts-version-gradle", RuleType: "standard", Level: "Warn", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(^|\\.|\\s)targetCompatibility(\\s|=)[^\\d]+(%s)", Advice: "The application is using non-LTS version Java. JDK on LTS version is recommended, i.e., JAVA_8, JAVA_11 or JAVA_17.", Effort: 6, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-non-lts-version-maven", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[starts-with(java.version,\"%s\")]/java.version", Advice: "The application is using non-LTS version Java. JDK on LTS version is recommended, i.e., JAVA_8, JAVA_11 or JAVA_17.", Effort: 6, Readiness: 0, Impact: "", Category: "java", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-properties-config-client-configs", RuleType: "standard", Level: "Info", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "The application is using spring config server and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "config server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "config server",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\s)spring\\.config\\.import", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\s)spring\\.cloud\\.config\\.uri", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-properties-eureka-client-configs", RuleType: "standard", Level: "Info", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "The application is using eureka and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "eureka", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "eureka",}, { Value: "spring",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\s)eureka\\.client\\.(service-url|serviceUrl)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-properties-logging-console-appender", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "regex", DefaultPattern: "", Advice: "Please enable logging to console", Effort: 5, Readiness: 0, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(?i)(appender.*=.*ConsoleAppender|appender.*\\.type.*=.*Console|(org/springframework/boot/logging/logback/base.xml)|(org/springframework/boot/logging/log4j2/log4j2.xml))", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-redis-connection-string-user-password-properties", RuleType: "standard", Level: "Info", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Redis connection string, username and password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "redis", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "redis",}, { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-redis-connection-string-user-password-yaml", RuleType: "standard", Level: "Info", FileType: "yaml$|yml$", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "Redis connection string, username and password are detected", Effort: 0, Readiness: 0, Impact: "", Category: "redis", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "redis",}, { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-scheduled-job-annotation", RuleType: "standard", Level: "Warn", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "The application has scheduled jobs such as Quartz Scheduler tasks or cron jobs. Please be aware that after migrating to the cloud and scaling out, scheduled jobs in applications may run more than once per scheduled period and lead to unintended consequences.", Effort: 6, Readiness: 0, Impact: "", Category: "scheduled job", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "scheduled job",}, { Value: "faas",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Scheduled", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-scheduled-job-import", RuleType: "standard", Level: "Warn", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import\\s*%s.*", Advice: "The application has scheduled jobs such as Quartz Scheduler tasks or cron jobs. Please be aware that after migrating to the cloud and scaling out, scheduled jobs in applications may run more than once per scheduled period and lead to unintended consequences.", Effort: 6, Readiness: 0, Impact: "file", Category: "scheduled job", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "scheduled job",}, { Value: "faas",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-spring-boot-oss-support-version-gradle", RuleType: "standard", Level: "Warn", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Spring boot version is out of spring boot open source support support(https://spring.io/projects/spring-boot#support). If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org\\.springframework\\.boot(.*)([\\r\\n]*)version( *[:=]? *)[''\"]?(2\\.[4-7])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i)(springBootVersion)(.*)(2\\.[4-7])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.springframework\\.boot[''\"]?( *[:=]? *)[''\"]?(2\\.[4-7])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-oss-support-version-maven", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[groupId=\"org.springframework.boot\" and starts-with(version,\"%s\")]/version", Advice: "Spring boot version is out of spring boot open source support support(https://spring.io/projects/spring-boot#support). If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "2.4", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.6", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "2.7", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-starter", RuleType: "standard", Level: "Info", FileType: "gradle$|kts$|xml$", Target: "contents", Type: "regex", DefaultPattern: "", Advice: "The application has spring boot/spring cloud starter dependencies", Effort: 0, Readiness: 0, Impact: "", Category: "spring boot", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-spring-boot-support-version-gradle", RuleType: "standard", Level: "Warn", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Spring boot version is out of any spring boot support(https://spring.io/projects/spring-boot#support). Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org\\.springframework\\.boot(.*)([\\r\\n]*)version( *[:=]? *)[''\"]?(2\\.[0-3])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i)(springBootVersion)(.*)(2\\.[0-3])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.springframework\\.boot[''\"]?( *[:=]? *)[''\"]?(2\\.[0-3])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-support-version-maven", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[groupId=\"org.springframework.boot\" and starts-with(version,\"%s\")]/version", Advice: "Spring boot version is out of any spring boot support(https://spring.io/projects/spring-boot#support). Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-spring-boot-version-gradle", RuleType: "standard", Level: "Critical", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Spring boot version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org\\.springframework\\.boot(.*)([\\r\\n]*)version( *[:=]? *)[''\"]?(1\\.)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i)(springBootVersion)(.*)(1\\.)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org\\.springframework\\.boot[''\"]?( *[:=]? *)[''\"]?(1\\.)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-boot-version-maven", RuleType: "standard", Level: "Critical", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[groupId=\"org.springframework.boot\" and starts-with(version,\"%s\")]/version", Advice: "Spring boot version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring-boot-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-oss-support-version-gradle", RuleType: "standard", Level: "Warn", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(?i)(springCloudVersion|org\\.springframework\\.cloud)(.*)(%s)", Advice: "Spring cloud version is out of open source support support. If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring cloud",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "2020\\.|ilford", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-oss-support-version-maven", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "Spring cloud version is out of open source support support. If you don't have commercial support, please update to newer version", Effort: 4, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring cloud",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-spring-cloud-support-version-gradle", RuleType: "standard", Level: "Warn", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(?i)(springCloudVersion|org\\.springframework\\.cloud)(.*)(%s)", Advice: "Spring cloud version is out of any support. Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring cloud",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Hoxton|Greenwich|Finchley", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-support-version-maven", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "Spring cloud version is out of any support. Please update to newer version", Effort: 6, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring cloud",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-spring-cloud-version-gradle", RuleType: "standard", Level: "Critical", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(?i)(springCloudVersion|org\\.springframework\\.cloud)(.*)(%s)", Advice: "Spring cloud version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring cloud",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Edgware|Dalston", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-cloud-version-maven", RuleType: "standard", Level: "Critical", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "Spring cloud version is too low", Effort: 50, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring cloud",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-version-gradle", RuleType: "standard", Level: "Critical", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "(^|\\.|\\s)targetCompatibility(\\s|=)[^\\d]+(%s)", Advice: "JDK version is found to be lower than JAVA_8", Effort: 50, Readiness: 0, Impact: "", Category: "version", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "1\\.[0-7]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-version-maven", RuleType: "standard", Level: "Critical", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "//*[starts-with(java.version,\"%s\")]/java.version", Advice: "JDK version is found to be lower than JAVA_8", Effort: 50, Readiness: 0, Impact: "", Category: "java", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-xml-logging-console-appender", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "regex", DefaultPattern: "", Advice: "Please enable logging to console", Effort: 5, Readiness: 0, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(?i)(ConsoleAppender|type(\\s)*=(\\s)*\"console\"|</console>|org/springframework/boot/logging/logback/base.xml|org/springframework/boot/logging/log4j2/log4j2.xml)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yaml-config-client-configs", RuleType: "standard", Level: "Info", FileType: "(yaml$|yml$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "The application is using spring config server and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "config server", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "config server",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.spring.config.import", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "$.spring.cloud.config.uri", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yaml-eureka-client-configs", RuleType: "standard", Level: "Info", FileType: "(yaml$|yml$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "The application is using eureka and setting the connection string", Effort: 0, Readiness: 0, Impact: "", Category: "eureka", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "eureka",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.eureka.client[\"service-url\",\"serviceUrl\"]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-yaml-json-log4j2-fileAppender", RuleType: "standard", Level: "Warn", FileType: "yaml$|yml$|json$|jsn$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "Replace file appender with console appender", Effort: 3, Readiness: 0, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "log4j2",}, { Value: "logging",}, { Value: "log2file",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-yamljson-logging-console-appender", RuleType: "standard", Level: "Warn", FileType: "yaml$|yml$|json$|jsn$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "Please enable logging to console", Effort: 5, Readiness: 0, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$[*][\"appenders\",\"Appenders\"][?(@.console || @.Console || @.appender[?(@.type=~/(?i)console/)])]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-zipkin-maven", RuleType: "standard", Level: "Warn", FileType: "xml$", Target: "contents", Type: "xpath", DefaultPattern: "", Advice: "The application uses Zipkin. Update the application to use Application Insights(https://learn.microsoft.com/en-us/azure/spring-apps/how-to-application-insights) instead if migrating to Azure", Effort: 5, Readiness: 0, Impact: "", Category: "zipkin", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "azure",}, { Value: "zipkin",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "//*[contains(groupId, \"zipkin\") or contains(artifactId, \"zipkin\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "properties-port", RuleType: "standard", Level: "Info", FileType: "properties$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "The application is setting the server port. Please be aware of potential port reliance issues during the migration process", Effort: 0, Readiness: 0, Impact: "", Category: "port", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "port",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(^|\\s)server\\.port", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "windows-dynamic-link-library", RuleType: "standard", Level: "Critical", FileType: "dll$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "This Dynamic-Link Library is Microsoft Windows platform dependent. It needs to be replaced with a Linux-style shared library", Effort: 10, Readiness: 0, Impact: "", Category: "os", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "os",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "dll$", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "windows-file-path", RuleType: "standard", Level: "Critical", FileType: "java$|properties$|xml$|yaml$|yml$|json$|jsn$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "This file system path is Microsoft Windows platform dependent. It needs to be replaced with a Linux-style path", Effort: 4, Readiness: 0, Impact: "", Category: "os", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "os",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(\"|''|`|\\s|^)(?:[a-zA-Z]\\:|\\\\\\\\[\\w\\s\\.]+\\\\[\\w\\s\\.$]+)([\\\\\\/][^\\n\\t]+)+", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "yaml-nonstandard-port", RuleType: "standard", Level: "Info", FileType: "(yaml$|yml$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "The application is setting the server port. Please be aware of potential port reliance issues during the migration process", Effort: 0, Readiness: 0, Impact: "", Category: "port", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "port",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "$.server.port", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spark", RuleType: "", Level: "", FileType: "(java$|kt$|kts$|scala$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Spark", Effort: 0, Readiness: 9, Impact: "", Category: "web", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "spark",}, { Value: "batch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*import\\s*spark.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.apache.spark.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring", RuleType: "", Level: "", FileType: "(xml$|gradle$|java$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Spring Batch", Effort: 0, Readiness: 9, Impact: "", Category: "batch", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "spring-batch",}, { Value: "batch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*spring-batch-core.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.springframework.batch.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "build-ant-maven", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Align with standard build system", Effort: 0, Readiness: 0, Impact: "", Category: "buildSystem", Criticality: "",Container: 0, Cloud: 0,
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
        
            { Name: "build-gradle", RuleType: "standard", Level: "", FileType: "gradle$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Gradle is being used", Effort: 0, Readiness: 0, Impact: "", Category: "buildSystem", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "build-system",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "build.gradle", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-package-gradle", RuleType: "standard", Level: "", FileType: "gradle$|kts$", Target: "line", Type: "regex", DefaultPattern: "apply\\s*plugin:\\s*['']%s['']", Advice: "Application Server coupling detected.  Consider repackaging artifact as either war or executable jar", Effort: 100, Readiness: 0, Impact: "", Category: "packaging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "gradle",}, { Value: "build-system",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ear", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jar", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "war", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-package-maven-ant", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application Server coupling detected", Effort: 100, Readiness: 0, Impact: "", Category: "packaging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "maven",}, { Value: "build-system",}, },
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
        
            { Name: "SNAP-java-ver-Maven-Ant", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Legacy Java detected.  Consider upgrading.", Effort: 100, Readiness: 0, Impact: "", Category: "java-ver", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "java-version",}, { Value: "build-system",}, },
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
        
            { Name: "dotnet-cache", RuleType: "standard", Level: "", FileType: "(cs$|vb$|config$|csproj$|json$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Migrate to VMware GemFire", Effort: 100, Readiness: 5, Impact: "", Category: "cache", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Memcache", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "memcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "Apache.Geode", Advice: "Direct usage of Apache Geode .NET Client APIs in .NET code for data management and caching.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: ".CreateRegionFactory", Advice: "Direct usage of Apache Geode .NET Client APIs in .NET code for data management and caching.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: ".GetQueryService", Advice: "Direct usage of Apache Geode .NET Client APIs in .NET code for data management and caching.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: "HazelcastClient", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hazelcast", Recipe: "", },
             { Type: "", Pattern: "", Value: "HazelcastOptions", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hazelcast", Recipe: "", },
             { Type: "", Pattern: "", Value: "Hazelcast", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hazelcast", Recipe: "", },
             }, },
        
            { Name: "dotnet-FileCacheModule", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "", Effort: 1000, Readiness: 5, Impact: "", Category: "caching", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/location/system.webServer/serverRuntime[@frequentHitThreshold or @frequentHitTimePeriod][1]", Advice: "Caching should be managed externally. Relying on IIS Modules won't work in Linux environment. Alternatives are distributed cache solutions like Redis, SQL Server distributed cache, NCache distributed cache.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "/configuration/system.webServer/staticContent/clientCache[@cacheControlMode][1]", Advice: "Output cache profiles etc. via configuration.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cache-control-mode", Recipe: "", },
             }, },
        
            { Name: "java-cache", RuleType: "standard", Level: "", FileType: "(java$|xml$|gradle$|properties$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Distributed caches must be remediated to function in K8S", Effort: 50, Readiness: 10, Impact: "", Category: "cache", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "cache",}, },
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
             { Type: "", Pattern: "", Value: "com.tangosol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "tangosol", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.commons-jcs-jcache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcs-jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.websphere.cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websphere-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.jboss.ha.cachemanager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jboss-ha-cachemanager", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "javax-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.spy.memcached", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "net-spy-memcached", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.opensymphony.oscache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "opensymphony-oscache", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.shiftone-cache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "shiftone-casche", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.websphere.objectgrid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "websphere-objectgrid", Recipe: "", },
             { Type: "", Pattern: "", Value: "getConfigurableCacheFactory", Advice: "Make sure its externalized and using Cloud Friendly cache implementation.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "NamedCache", Advice: "Make sure its externalized and using Cloud Friendly cache implementation.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "ReflectionExtractor", Advice: "Make sure its externalized and using Cloud Friendly cache implementation.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "ChainedExtractor", Advice: "Make sure its externalized and using Cloud Friendly cache implementation.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.cache", Advice: "Make sure its externalized and using Cloud Friendly cache implementation.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jcache", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*spring-boot-starter-cache.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-boot-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*spring.cache.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-boot-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.springframework.cache.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-boot-cache", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.jackey", Advice: "Migrate to ValKey (Redis) tile or Tanzu Data Hub", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "redis", Recipe: "", },
             { Type: "", Pattern: "", Value: "redis", Advice: "Migrate to ValKey (Redis) tile or Tanzu Data Hub", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "redis", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.lettuce", Advice: "Migrate to ValKey (Redis) tile or Tanzu Data Hub", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "redis", Recipe: "", },
             { Type: "", Pattern: "", Value: "Redis", Advice: "Migrate to ValKey (Redis) tile or Tanzu Data Hub", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "redis", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.geode.cache.*", Advice: "Direct usage of Apache Geode APIs in Java, focusing on region operations, client configurations, and data management.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: "new ClientCacheFactory", Advice: "Direct usage of Apache Geode APIs in Java, focusing on region operations, client configurations, and data management.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: "createClientRegionFactory", Advice: "Direct usage of Apache Geode APIs in Java, focusing on region operations, client configurations, and data management.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: "clientRegionFactory.create", Advice: "Direct usage of Apache Geode APIs in Java, focusing on region operations, client configurations, and data management.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.geode", Advice: "Direct usage of Apache Geode APIs in Java, focusing on region operations, client configurations, and data management.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             { Type: "", Pattern: "", Value: "geode-core", Advice: "Direct usage of Apache Geode APIs in Java, focusing on region operations, client configurations, and data management.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-geode", Recipe: "", },
             }, },
        
            { Name: "python-db-redis", RuleType: "standard", Level: "", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Redis is being used", Effort: 1, Readiness: 10, Impact: "", Category: "cache", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "redis",}, { Value: "cache",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "redis", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-cfenv-gradle", RuleType: "standard", Level: "", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "Application appears to be using cf env library to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-cloud-cloudfoundry-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-cfenv-maven", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(>.*%s<)", Advice: "Application appears to be using cf env library to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java-cfenv-boot", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-connector-gradle", RuleType: "standard", Level: "", FileType: "gradle$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "Application appears to be using cloud connectors library to inject environment variables provided by TAS. Follow the guidance listed here https://docs.cloudfoundry.org/buildpacks/java/configuring-service-connections.html", Effort: 100, Readiness: 7, Impact: "", Category: "tas", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-boot-starter-cloud-connectors", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-connectors-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-cloudfoundry-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-spring-service-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-connector-maven", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(>.*%s<)", Advice: "Application appears to be using cloud connectors library to inject environment variables provided by TAS. Follow the guidance listed here https://docs.cloudfoundry.org/buildpacks/java/configuring-service-connections.html", Effort: 100, Readiness: 7, Impact: "", Category: "tas", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "spring-boot-starter-cloud-connectors", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-connectors-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-cloudfoundry-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-cloud-spring-service-connector", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-environment-connector-python", RuleType: "standard", Level: "", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "Application appears to be using Cloudfoundry python client to inject environment variables provided by TAS.", Effort: 1, Readiness: 7, Impact: "", Category: "tas", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "CloudFoundryClient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "tas-spring-code-changes", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(>.*%s<)", Advice: "Remove any of these annotations from Spring Java configuration classes (provided by Spring Cloud Connectors). Replace them with the Spring SPeL or Spring Boot configuration (https://docs.cloudfoundry.org/buildpacks/java/configuring-service-connections.html#spring) options listed above.", Effort: 100, Readiness: 7, Impact: "", Category: "tas", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tas",}, { Value: "env-config",}, { Value: "dependencies",}, { Value: "spring",}, },
            Profiles:
            []Profile{  { Value: "tas",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ServiceScan", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CloudScan", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "apache-zookeeper", RuleType: "", Level: "", FileType: "(java$|py$|js$|ts$|yml$|yaml$|go$|rb$|scala$|xml$|properties$)", Target: "line", Type: "regex", DefaultPattern: "(?i)%s", Advice: "Usage of apache zookeeper", Effort: 0, Readiness: 10, Impact: "", Category: "other", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "apache-zookeeper",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*apache.*zookeeper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.zookeeper.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "hardcode-uri", RuleType: "standard", Level: "", FileType: "(java$|^vb$|py$|go$|aspx$|^c$|h$|cs$|csx$|cpp$|cob$|cfm$|cfml$|dockerfile$|jsp$|php$|^r$|^rb$|^ts$|yaml$|yml$|json$)", Target: "line", Type: "regex", DefaultPattern: "(?:(?:^|[^\"])|(?:^|[^ =])\"|(?:^|[^:]) \"|(?:^|[^s])=\"|(?:^|[^\"]): \"|(?:^|[^n])s=\"|(?:^|[^d])\": \"|(?:^|[^l])ns=\"|(?:^|[^e])d\": \"|(?:^|[^m])lns=\"|(?:^|[^v])ed\": \"|(?:^|[^x])mlns=\"|(?:^|[^l])ved\": \"|(?:^|[^o])lved\": \"|(?:^|[^s])olved\": \"|(?:^|[^e])solved\": \"|(?:^|[^r])esolved\": \")(%s)\\:{1}\\/{2}[a-zA-Z0-9]", Advice: "Found hard-coded URI. Make configurable, put into environment or config map", Effort: 3, Readiness: 8, Impact: "", Category: "env-config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "hardcoded-uri",}, { Value: "configuration",}, },
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
        
            { Name: "java-hardIP", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Hardcoded IP addresses are problematic in K8S", Effort: 1, Readiness: 8, Impact: "", Category: "hard-ip", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "hard-ip",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-php-url-not-hardcoded", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s://).*", Advice: "Url Hardcoding in code is anti-pattern", Effort: 200, Readiness: 1000, Impact: "", Category: "anti-pattern", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "code_smell",}, { Value: "configuration",}, },
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
        
            { Name: "plaintext-creds", RuleType: "standard", Level: "", FileType: "", Target: "line", Type: "regex", DefaultPattern: "(%s)\\s*=+\\s*[\"''][^${[(\\s][^\"${[(\\s].*", Advice: "Found plain text crdentials. Never save passwords or login information in files. Make configurable, put into environment variables or config maps in a kubernetes like environment. Credentials could also be set and retrieved from cloud friendly crdential vaults.", Effort: 0, Readiness: 9, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "security",}, { Value: "configuration",}, },
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
        
            { Name: "ruby-hardIP", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Hardcoded IP addresses are problematic in cloud/container env", Effort: 2, Readiness: 8, Impact: "", Category: "hard-ip", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "hard-ip",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "docker-non-root-user", RuleType: "standard", Level: "", FileType: "(Dockerfile$|Containerfile$)", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Shows evidence of avoiding root privledges", Effort: -100, Readiness: 1000, Impact: "", Category: "dockerSecurity", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "non-root-user",}, { Value: "docker",}, { Value: "container",}, },
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
        
            { Name: "docker-sudo", RuleType: "standard", Level: "", FileType: "(Dockerfile$|Containerfile$)", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Using root inside a container is a serious vulnerability.", Effort: 100, Readiness: 10, Impact: "", Category: "docker", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "docker",}, { Value: "sudo",}, { Value: "container",}, },
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
        
            { Name: "faas-meta", RuleType: "fire-once", Level: "", FileType: "meta$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "App should be started in the shortest time possible", Effort: 200, Readiness: 2, Impact: "", Category: "boottime", Criticality: "",Container: 0, Cloud: 0,
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
        
            { Name: "kube-config", RuleType: "", Level: "", FileType: "(yaml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Kubernetes Config", Effort: 0, Readiness: 9, Impact: "", Category: "platform", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "kubernetes",}, { Value: "platform",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*kind:\\s.*Deployment.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*kind:\\s.*Service.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*kind:\\s.*Pod.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*kind:\\s.*Policy.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*kind:\\s.*Ingress.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "kube-splunk-config", RuleType: "", Level: "", FileType: "(yml$|yaml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Splunk", Effort: 0, Readiness: 9, Impact: "", Category: "logging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "Splunk",}, { Value: "logging",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*splunk.com.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "splunk.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "openshift-config", RuleType: "", Level: "", FileType: "(yaml$|json$|yml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Openshift Config", Effort: 0, Readiness: 9, Impact: "", Category: "platform", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "openshift",}, { Value: "platform",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*openshift.io.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-nosql-database", RuleType: "standard", Level: "", FileType: "(cs$|vb$|config$|csproj$|json$)", Target: "line", Type: "regex", DefaultPattern: ".*(%s)", Advice: "", Effort: 100, Readiness: 5, Impact: "", Category: "database", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "database",}, { Value: "nosql",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Apache.Ignite", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-ignite", Recipe: "", },
             { Type: "", Pattern: "", Value: "Cassandra", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cassandra", Recipe: "", },
             { Type: "", Pattern: "", Value: "GraphQL", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "graphql", Recipe: "", },
             { Type: "", Pattern: "", Value: "MongoDB", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mongodb", Recipe: "", },
             { Type: "", Pattern: "", Value: "mongodb", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nosql", Recipe: "", },
             }, },
        
            { Name: "dotnet-rdbms-database", RuleType: "standard", Level: "", FileType: "(cs$|vb$|config$|csproj$|json$)", Target: "line", Type: "regex", DefaultPattern: ".*(%s)", Advice: "", Effort: 100, Readiness: 5, Impact: "", Category: "database", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "database",}, { Value: "rdbms",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ConnectionStrings", Advice: "Remove connection strings from files, use environment variables (or mount configmap into pod in a docker like environment)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "configuration", Recipe: "", },
             { Type: "", Pattern: "", Value: "connectionStrings", Advice: "Remove connection strings from files, use environment variables (or mount configmap into pod in a docker like environment)", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "configuration", Recipe: "", },
             { Type: "", Pattern: "", Value: "IBM.Data.DB2", Advice: "IBM.Data.DB2 can require a special procedure so that the driver's native components are deployed with the application.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-db2", Recipe: "", },
             { Type: "", Pattern: "", Value: "SqlServer", Advice: "Leverage MySQL Tile or k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mssql", Recipe: "", },
             { Type: "", Pattern: "", Value: "SqlClient", Advice: "Leverage MySQL Tile or k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mssql", Recipe: "", },
             { Type: "", Pattern: "", Value: "MySql", Advice: "Leverage MySQL Tile or k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mssql", Recipe: "", },
             { Type: "", Pattern: "", Value: "Oracle.DataAccess", Advice: "Oracle unmanaged driver requires including binaries with app -- can use buildpack.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "Npgsql", Advice: "Direct usage of Npgsql for PostgreSQL database interactions in .NET code.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "postgresql", Recipe: "", },
             { Type: "", Pattern: "", Value: "Elasticsearch", Advice: "Make sure to have reachable ELK stack from deployed app", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "elasticsearch", Recipe: "", },
             }, },
        
            { Name: "java-in-memory-database", RuleType: "standard", Level: "", FileType: "(java$|gradle$|yml$|properties$|xml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "", Effort: 0, Readiness: 5, Impact: "", Category: "persistence", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "in-memory",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*hazelcast", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hazelcast", Recipe: "", },
             }, },
        
            { Name: "java-nosql-database", RuleType: "standard", Level: "", FileType: "(java$|gradle$|yml$|properties$|xml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Migrate to VMware GemFire", Effort: 10, Readiness: 5, Impact: "", Category: "database", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "database",}, { Value: "nosql",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*apache.ignite.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-ignite", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*cassandra.*", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "cassandra", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*import\\s*graphql.*", Advice: "Usage of Graphql", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "graphql", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*com.graphql.*", Advice: "Usage of Graphql", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "graphql", Recipe: "", },
             { Type: "", Pattern: "", Value: "mongodb", Advice: "Migrate to VMware GemFire, Spring Data makes it easy to migrate to GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mongodb", Recipe: "", },
             }, },
        
            { Name: "java-rdbms-database", RuleType: "standard", Level: "", FileType: "(java$|gradle$|yml$|properties$|xml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "", Effort: 0, Readiness: 5, Impact: "", Category: "persistence", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "rdbms",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.springframework.jdbc", Advice: "Consult 3rd party documentation", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-jdbc", Recipe: "", },
             { Type: "", Pattern: "", Value: "couchbase", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "couchbase", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdbc:db2", Advice: "Detection of db2 Usage, use Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-db2", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.db2.jcc.DB2Driver", Advice: "Detection of db2 Usage, use Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-db2", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.db2", Advice: "Detection of db2 Usage, use Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-db2", Recipe: "", },
             { Type: "", Pattern: "", Value: "elasticsearch", Advice: "Migrate to VMware GemFire", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "elasticsearch", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.apache.hadoop.*", Advice: "Usage of Hadoop", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hadoop", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*hadoop-*", Advice: "Usage of Hadoop", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "hadoop", Recipe: "", },
             { Type: "", Pattern: "", Value: "mariadb", Advice: "Migrate to Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mariadb", Recipe: "", },
             { Type: "", Pattern: "", Value: "jdbc:sqlserver", Advice: "Detection of mssql Usage", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mssql", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.microsoft.sqlserver.*", Advice: "Detection of mssql Usage", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mssql", Recipe: "", },
             { Type: "", Pattern: "", Value: "sqlserver", Advice: "Migrate to Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sqlserver", Recipe: "", },
             { Type: "", Pattern: "", Value: "mysql", Advice: "Detection of MySQL Usage", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "mysql", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle", Advice: "Migrate to Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle", Recipe: "", },
             { Type: "", Pattern: "", Value: "postgresql", Advice: "Leverage Postgres, Postgres Tile or k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "postgres", Recipe: "", },
             { Type: "", Pattern: "", Value: "sqlite", Advice: "Detection of SQLite Usage", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sqlite", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.talend", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "talend", Recipe: "", },
             { Type: "", Pattern: "", Value: "oracle.odi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "oracle-odi", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.ibm.is", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-is", Recipe: "", },
             { Type: "", Pattern: "", Value: "net.sf.jasper", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jasper", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.pentaho", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "pentaho", Recipe: "", },
             { Type: "", Pattern: "", Value: "StoredProcedureQuery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             { Type: "", Pattern: "", Value: "CallableStatement", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jdbc", Recipe: "", },
             }, },
        
            { Name: "java-rds-connection-string", RuleType: "standard", Level: "Info", FileType: "(yaml$|yml$|json$|jsn$|properties$)", Target: "contents", Type: "yamlpath", DefaultPattern: "", Advice: "relational database service connection string, username or password are detected. Migrate to Tanzu Postgres or MySQL", Effort: 0, Readiness: 0, Impact: "", Category: "relational database service", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "relational database service",}, { Value: "database",}, { Value: "rdbms",}, },
            Profiles:
            []Profile{  { Value: "Azure Spring Apps",}, },
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
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)datasource\\.(.*\\.)?(url|jdbc-url|u\\-r\\-l)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)datasource\\.(.*\\.)?(username|user|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)r2dbc\\.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)r2dbc\\.(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)jdbc\\.url", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(^|\\.|\\s)jdbc\\.(username|password)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "=\\s*jdbc:.+:.+", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-database", RuleType: "standard", Level: "", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "", Effort: 1, Readiness: 10, Impact: "", Category: "database", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "peewee",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "peewee", Advice: "Check target platform has support for this library", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "peewee", Recipe: "", },
             { Type: "", Pattern: "", Value: "sqlite", Advice: "Consider migration to an external database.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sqlite", Recipe: "", },
             { Type: "", Pattern: "", Value: "SqliteDatabase", Advice: "Consider migration to an external database.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sqlite", Recipe: "", },
             }, },
        
            { Name: "sqlserver-ssis", RuleType: "fire-once", Level: "", FileType: "(dtsx$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "SSIS is not supported on CloudFoundry.", Effort: 100, Readiness: 0, Impact: "", Category: "Unsupported modules", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "ssis",}, { Value: "etl",}, { Value: "sql",}, { Value: "modernize",}, { Value: "container",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DTS", Advice: "SSIS is not supported on CloudFoundry. Consider leaving the packages in an external SQL Server deployment or rewrite them in a cloud native ETL Framework like Spring Cloud Data Flow.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-spring-swagger", RuleType: "", Level: "", FileType: "(java$|yml$|yaml$|scala$|xml$|properties$)", Target: "line", Type: "regex", DefaultPattern: "(?i)%s", Advice: "Usage of spring swagger", Effort: 0, Readiness: 10, Impact: "", Category: "api", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "api",}, { Value: "spring-swagger",}, { Value: "documentation",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*io.*swagger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-io-swagger-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.*swagger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-io-swagger-detect", Recipe: "", },
             { Type: "", Pattern: "", Value: "import.*spring.*swagger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-swagger-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring.*swagger", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-swagger-detect", Recipe: "", },
             }, },
        
            { Name: "config-dotnet-webConfig", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Upgrade to .Net Core", Effort: 0, Readiness: 0, Impact: "", Category: "Config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-config",}, { Value: "modernize-to-net-core",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-RequestFilteringModule", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Refer to platform documentation", Effort: 100, Readiness: 1, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "request-filtering",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/requestFiltering/requestLimits[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-StaticCompressionModule", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuration", Effort: 100, Readiness: 5, Impact: "", Category: "unsupported-iis-module", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "static-compression",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpCompression[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-StaticFileModule", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Ensure compatible configuation", Effort: 400, Readiness: 5, Impact: "", Category: "Unsupported modules", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "static-file-module",}, { Value: "web-server",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-iis_module-TokenCacheModule", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Caches windows security tokens for password based authentication schemes (anonymous authentication; basic authentication; IIS client certificate authentication). Ensure compabible configuration.", Effort: 1000, Readiness: 5, Impact: "", Category: "unsupported-iis-module", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "cache",}, { Value: "token-cache-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/globalModules/add[@name=\"TokenCacheModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-file-based-config", RuleType: "standard", Level: "", FileType: "cs", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Externalize configuration to environment or use ConfigMap as file-mount into a K8S pod", Effort: 10, Readiness: 3, Impact: "", Category: "env-config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "env-config",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".AddJsonFile\\(\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-fileIO-read", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s*\\(.*", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform.", Effort: 3, Readiness: 8, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "File.Open", Advice: "Calling File.Open", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-fileIO", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*\\(.*", Advice: "Relying on the local filesystem to store data is unreliable in a cloud platform. Containers are ephemerals and data could be lost. Data could be stored in database or the application should leverage File Sharing Systems like NFS Persistent Volume (PV) in Kubernetes environment.", Effort: 100, Readiness: 8, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-filepath", RuleType: "standard", Level: "", FileType: "(cs|yaml|yml|json)", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform. Containers are ephemerals and data could be lost. Data could be stored in database or the application should leverage File Sharing Systems like NFS Persistent Volume (PV) in Kubernetes environment.", Effort: 8, Readiness: 3, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-ipv4-addresses", RuleType: "standard", Level: "", FileType: "(yaml|yml|cs$|json)", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Found hard-coded IPv4s. Make configurable, put into environment variables. Leverage config maps in a kubernetes like environment.", Effort: 3, Readiness: 8, Impact: "", Category: "hard-ip", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "hard-ip",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  { Value: ".GeneratedCodeAttribute\\(.*[,]",}, { Value: "AssemblyFileVersion\\(",}, { Value: "AssemblyVersion\\(",}, { Value: "Version=",}, { Value: "\"assemblyVersion\".*[:].\".*\"",}, { Value: "\"fileVersion\".*[:].\".*\"",}, { Value: "PackageReference.*Version=\"",}, { Value: "GeneratedCode\\(",}, { Value: "([12]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01]))",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "[^_](([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-launchProcess", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Launching additional processes within a container like environment is not recommended. In a targeted docker like environment it is generally recommended that you separate areas of concern by using one service per container.", Effort: 300, Readiness: 7, Impact: "", Category: "process-launch", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "process-launch",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "dotnet-serilog", RuleType: "standard", Level: "", FileType: "cs$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Logging with Serilog. Make sure not to log to file. Remote logging sinks need to be reachable on network. Logging to the Event Log is not recommended for cloud native apps. Write to or manage logfiles. Instead, each running process should write its event stream, unbuffered, to stdout.", Effort: 3, Readiness: 8, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "dotnet-sharepoint", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "SharePoint is not supported on CloudFoundry.", Effort: 100, Readiness: 0, Impact: "", Category: "unsupported-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "sharepoint",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.SharePoint", Advice: "SharePoint is not supported on CloudFoundry.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-transactions", RuleType: "fire-once", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Potential use of distributed transactions which are unsupported", Effort: 100, Readiness: 10, Impact: "", Category: "transaction", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "distributed-transaction",}, { Value: "modernize",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TransactionScope", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-bindings", RuleType: "standard", Level: "", FileType: "config$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Refers to documentation when using WCF, some protocols might not be supported", Effort: 0, Readiness: 0, Impact: "", Category: "wcf", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-wcf-protocols-ws", RuleType: "fire-once", Level: "", FileType: "config$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Unsupported protocols", Effort: 1000, Readiness: 0, Impact: "", Category: "wcf", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<ws.+HttpBinding>", Advice: "Many features of WS* protocols are problematic in the cloud like distributed transactions and reliable sessions", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-protocols", RuleType: "fire-once", Level: "", FileType: "config$", Target: "line", Type: "regex", DefaultPattern: "", Advice: "Unsupported protocols", Effort: 500, Readiness: 0, Impact: "", Category: "wcf", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-wcf-ssl", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "When using HTTPS, terminate SSL at load balancer", Effort: 1, Readiness: 10, Impact: "", Category: "wcf", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "ssl",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-windows-code-access-security", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "unsupported", Effort: 10, Readiness: 9, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "dotnet-windows-components", RuleType: "fire-once", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported on TAS", Effort: 1000, Readiness: 9, Impact: "", Category: "", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "com+",}, { Value: "api",}, { Value: "distributed-transaction",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-windows-presentation-foundation", RuleType: "fire-once", Level: "", FileType: "xaml", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Desktop technology only available on Windows platform", Effort: 500, Readiness: 1, Impact: "", Category: "windows-desktop", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-desktop",}, { Value: "windows-wpf",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<Page\\sxmlns=\".*xaml\\/presentation\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsForms", RuleType: "fire-once", Level: "", FileType: "(cs$|vb$|csproj$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Forms module is not supported in the cloud. Refactor to a web application.", Effort: 1000, Readiness: 5, Impact: "", Category: "Unsupported modules", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-forms",}, { Value: "windows-desktop",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Windows.Forms", Advice: "Windows Forms module is not supported. Refactor to a web application.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsRegistry", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "^.*%s.*", Advice: "External configurations should be made available by environment variables or some another external service (Property Management Service).", Effort: 100, Readiness: 2, Impact: "", Category: "windows-registry", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-registry",}, { Value: "configuration",}, },
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
        
            { Name: "dotnet-windowsServices", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Traditional Windows services will not run on linux environments. Windows services should be turned into their own deployable console application.", Effort: 0, Readiness: 4, Impact: "", Category: "windows-services", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-services",}, { Value: "process",}, { Value: "modernize",}, },
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
        
            { Name: "java-apache-camel", RuleType: "", Level: "", FileType: "(java$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Apache Camel", Effort: 0, Readiness: 9, Impact: "", Category: "integration-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "integration-framework",}, { Value: "apache-camel",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*org.apache.camel.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.camel.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "mulesoft-config-raml", RuleType: "", Level: "", FileType: "(raml$|xml$|json$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Mulesoft", Effort: 0, Readiness: 9, Impact: "", Category: "api", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "api",}, { Value: "mulesoft",}, { Value: "api-gateway",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*#%RAML.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*minMuleVersion.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*http[s]*://.*mulesoft.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*<mule\\s*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "bootCDI", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "cdi", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cdi",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "bootEJB", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "bootJAXWS", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "jaxws", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jax-ws",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "bootJSF", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "jsf", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "bootMDB", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "mdb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "bootSTRUTS", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "struts", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "frontend",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts2-core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "struts-2", Recipe: "", },
             }, },
        
            { Name: "bootTXN", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 2, Readiness: 2, Impact: "", Category: "txn", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "txn",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "bootWEBSOCKET", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Automatic remediation with Bootifier", Effort: 0, Readiness: 2, Impact: "", Category: "WebServlet", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-socket",}, { Value: "migrate-off-legacy-server",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "config-sessionState", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "App should be executed as a stateless process. User session information could be store in the database (Ex https://www.c-sharpcorner.com/article/configure-sql-server-session-state-in-asp-net-core/). A better approach is to use a Redis or SQL Server distributed cache, which doesn't require sticky sessions. For more information, see Distributed caching in ASP.NET Core (https://learn.microsoft.com/en-us/aspnet/core/performance/caching/distributed?view=aspnetcore-7.0).", Effort: 200, Readiness: 2, Impact: "", Category: "stateful", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "session",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/sessionState[@mode=\"InProc\" or @mode=\"StateServer\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-3rdPartyImports", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 0, Readiness: 0, Impact: "", Category: "third-party", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "third-party",}, { Value: "dependencies",}, },
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
        
            { Name: "java-3rdPartySecurity-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 7, Readiness: 10, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "java-MBeans", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\:{1}", Advice: "MBean is application server specific, change to refactor using cloud friendly technology", Effort: 100, Readiness: 10, Impact: "", Category: "websphere", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "websphere",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WebSphere:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-processexit", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[.]%s[ (].*$", Advice: "Refer to IBM documentation", Effort: 9, Readiness: 10, Impact: "", Category: "process-exit", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "process-exit",}, { Value: "container",}, },
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
        
            { Name: "java-alarmD-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Integrate with new alarmD service or migrate to Spring Boot scheduler", Effort: 5, Readiness: 0, Impact: "", Category: "ThirdParty", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "alarm-d",}, { Value: "scheduler",}, { Value: "faas",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.tfcci.ucs.manager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-apache-poi", RuleType: "", Level: "", FileType: "(java$|py$|js$|ts$|yml$|yaml$|go$|rb$|scala$|xml$|properties$)", Target: "line", Type: "regex", DefaultPattern: "(?i)%s", Advice: "Usage of apache poi", Effort: 0, Readiness: 10, Impact: "", Category: "other", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "document",}, { Value: "apache-poi",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*apache.*poi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-poi-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "apache.*poi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-poi-detect", Recipe: "", },
             }, },
        
            { Name: "java-apache-wink", RuleType: "", Level: "", FileType: "(java$|py$|js$|ts$|yml$|yaml$|go$|rb$|scala$|xml$|properties$)", Target: "line", Type: "regex", DefaultPattern: "(?i)%s", Advice: "Usage of apache wink", Effort: 0, Readiness: 10, Impact: "", Category: "api", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "api",}, { Value: "apache-wink",}, { Value: "rest-api",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*apache.*wink", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-wink-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "apache.*wink", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-wink-detect", Recipe: "", },
             }, },
        
            { Name: "java-apacheFop-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "[ .]%s[ .({]", Advice: "Usage requires configuration remdiation", Effort: 5, Readiness: 10, Impact: "", Category: "thirdParty", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "apache-fop",}, { Value: "third-party",}, { Value: "dependencies",}, },
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
        
            { Name: "java-batch", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Remove jndi provider or move to TKG", Effort: 10, Readiness: 7, Impact: "", Category: "jndi", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "batch",}, { Value: "web-container",}, { Value: "container",}, },
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
        
            { Name: "java-batchAnnotations", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Batch processing can include long running processes consider using k8s job scheduler", Effort: 10, Readiness: 7, Impact: "", Category: "batch", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "batch",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "BatchProperty", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-corba", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: ".*[ .]%s[ (].*", Advice: "Replace with cloud-friendly framework or move to TKG", Effort: 10, Readiness: 6, Impact: "", Category: "corba", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "rpc",}, { Value: "corba",}, { Value: "non-standard-protocol",}, { Value: "container",}, },
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
        
            { Name: "java-ejb-invocation", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 10, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, { Value: "ejb",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJB\\(lookup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-mdb", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult MDB documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "mdb",}, { Value: "messaging",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-ejb-rmi", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Removing RMI calls from client applications.", Effort: 10, Readiness: 4, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateless",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-ejb-stateful-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateful",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.ejb.Stateful", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-ejb-stateful", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Refer EJB stateful/stateless documentation", Effort: 100, Readiness: 100, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "ejb",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-ejb-stateless", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Removing RMI calls from client applications.", Effort: 10, Readiness: 4, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateless",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-faces-flow-Annotations", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Batch processing can include long running processes", Effort: 10, Readiness: 7, Impact: "", Category: "batch", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, { Value: "batch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FlowScoped", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Review usage to determine how the customized state of JSF flow is being used and determined if it can be externalized.", Effort: 100, Readiness: 100, Impact: "", Category: "jsf-flow", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, { Value: "batch",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.faces.flow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-faces-flow", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Review usage to determine how the customized state of JSF flow is being used and determined if it can be externalized.", Effort: 10, Readiness: 10, Impact: "", Category: "jsf-flow", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf-flow",}, { Value: "jsf",}, { Value: "batch",}, },
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
        
            { Name: "java-fileIO", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Move to cloud friendly alternative storage service", Effort: 8, Readiness: 8, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "file storage",}, { Value: "stateful",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-file-system", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\:{1}\\/{2}", Advice: "Use backing storage service", Effort: 100, Readiness: 10, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "stateful",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "file", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-handles-term", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*%s", Advice: "For containerization, the TERM signal must be handled, this pattern is a positive finding.", Effort: 1, Readiness: 6, Impact: "", Category: "term-signal", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "term",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Runtime.getRuntime().addShutdownHook", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-iop", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (.].*", Advice: "Remote Method Invocations create coupling between componets. Move to cloud friendly alternatives such as REST endpoints.", Effort: 3, Readiness: 6, Impact: "", Category: "iop", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "non-standard-protocol",}, { Value: "corba",}, { Value: "legacy-protocol",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-java-fx-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Java-fx is not cloud compatible and requires the JRE on the remote device.", Effort: 1000, Readiness: 100, Impact: "", Category: "java-fx", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "java-fx",}, { Value: "desktop-app",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javafx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jaxrs-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Refer to platform documentation", Effort: 1, Readiness: 8, Impact: "", Category: "jax-rs", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jax-rs",}, { Value: "rest",}, { Value: "tas",}, },
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
        
            { Name: "java-jcaAnnotations", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s$", Advice: "Convert to a cloud friendly backing service.", Effort: 100, Readiness: 7, Impact: "", Category: "jca", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "jca",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-jersey-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Refer to 3rd party organization for cloud affinity of library", Effort: 5, Readiness: 8, Impact: "", Category: "jersey", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jersey",}, { Value: "rest",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.sun.jersey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "jersey-.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-boot-starter-jersey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jks", RuleType: "standard", Level: "", FileType: "jks$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "Make sure to externalize jks store", Effort: 7, Readiness: 7, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "security",}, { Value: "jks",}, { Value: "mutual-auth",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*\\.jks", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-jndi", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Refactor jndi calls", Effort: 250, Readiness: 7, Impact: "", Category: "jndi", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "web-container",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-jni", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*%s\\s*native\\s*", Advice: "A few conditions have to be met to make JNI calls", Effort: 1000, Readiness: 7, Impact: "", Category: "jni", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "native",}, { Value: "jni",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-jpa", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ @.]%s[ (].*", Advice: "JPA will work inside of Cloud Native applications, make sure to use best practices to externalize connection parameters.", Effort: 2, Readiness: 6, Impact: "", Category: "jpa", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jpa",}, { Value: "web-profile",}, { Value: "entity-manager",}, { Value: "database",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-jsf", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*$", Advice: "Consider migrating to modern frameworks that have better support for the cloud.", Effort: 5, Readiness: 4, Impact: "", Category: "jsf", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "ui",}, { Value: "frontend",}, { Value: "modernize",}, },
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
        
            { Name: "java-jsp", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider migrating to modern UI frameworks that have better support for the cloud.", Effort: 2, Readiness: 5, Impact: "", Category: "jsp", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ui",}, { Value: "jsp",}, { Value: "web-container",}, { Value: "frontend",}, { Value: "modernize",}, },
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
        
            { Name: "java-jta", RuleType: "standard", Level: "", FileType: "(java$|yml$|yaml$|scala$|xml$|properties$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Distributed transactions are problematic and should be remediated.  Consider Eventual Consistency pattern.", Effort: 500, Readiness: 6, Impact: "file", Category: "jta", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "transaction",}, { Value: "jta",}, { Value: "web-profile",}, { Value: "modernize",}, { Value: "database",}, { Value: "spring-transaction",}, },
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
             { Type: "", Pattern: "", Value: "import.*org.springframework.transaction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-transaction-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.springframework.transaction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".PlatformTransactionManager", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import.*spring.*tx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-tx-import", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring.*tx", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "spring-tx-detect", Recipe: "", },
             }, },
        
            { Name: "java-jvm-runtimeConfigProps", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Do not change these properties at runtime in application code", Effort: 50, Readiness: 0, Impact: "", Category: "config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "config",}, { Value: "configuration",}, },
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
        
            { Name: "java-logging-file-appender", RuleType: "standard", Level: "", FileType: "xml$|properties$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Replace file appender with console appender", Effort: 3, Readiness: 5, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "logging",}, { Value: "log2file",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-logging-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Change to an implementation of SLF4J i.e. Logback", Effort: 3, Readiness: 5, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-message-driven-annotations", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "To convert a message driven bean to spring cloud stream with rabbitmq", Effort: 10, Readiness: 0, Impact: "", Category: "annotations", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-messageDrivenBeans", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*@%s.*$", Advice: "Refer to platform documentation", Effort: 7, Readiness: 10, Impact: "", Category: "mdb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "mdb",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "EJB\\(lookup", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-metrics", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Indicates use of a metrics collection library, which supports containerization", Effort: -10, Readiness: 7, Impact: "", Category: "metrics", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "metrics",}, { Value: "monitoring",}, },
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
        
            { Name: "java-mulesoft-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "There are several changes required to move a Mule Project to PCF", Effort: 100, Readiness: 10, Impact: "", Category: "etl", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mulesoft",}, { Value: "etl",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.mule", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mulesoft-intf", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ {].*", Advice: "Refer to platform documentation", Effort: 10, Readiness: 10, Impact: "", Category: "etl", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mulesoft",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-netflix-healthcheck", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Indicates existance of healthcheck endpoint, which is positive finding", Effort: -100, Readiness: 7, Impact: "", Category: "health-check", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "health-check",}, { Value: "monitoring",}, },
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
        
            { Name: "java-nio", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Make sure to use Cloud Friendly storage provider.", Effort: 8, Readiness: 8, Impact: "", Category: "nio", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "nio",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-nonstandard-protocol", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Ensure cloud friendly communication protocols", Effort: 50, Readiness: 10, Impact: "", Category: "protocol", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "non-standard-protocol",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-persistence", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consult 3rd party documentation", Effort: 5, Readiness: 5, Impact: "", Category: "persistence", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "javax-persistence",}, { Value: "io",}, { Value: "database",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.persistence", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-portUsage", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^*.(%s?:.*):(\\d*)\\/?(.*)", Advice: "Ensure port usage is cloud-friendly or use TKG", Effort: 5, Readiness: 6, Impact: "", Category: "port-usage", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "port-usage",}, { Value: "tas",}, { Value: "container",}, },
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
        
            { Name: "java-remoteEJB", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 10, Readiness: 10, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "annotations",}, { Value: "ejb",}, { Value: "fullprofile",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-remoteWebService-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 10, Readiness: 10, Impact: "", Category: "web-service", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-service",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-resource-cci", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 70, Readiness: 7, Impact: "", Category: "jca", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cci",}, { Value: "resource-adapter",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-resource-spi", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider rearchitecting the application to use Cloud friendly remote communications - http or messaging", Effort: 10, Readiness: 70, Impact: "", Category: "jca", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "resource-adapter",}, { Value: "spi",}, { Value: "jca",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-restlet-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "The Restlet library appears to be dead at this point in time. Consider upgrading to a Cloud friendly UI framework.", Effort: 10, Readiness: 7, Impact: "", Category: "restlet", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "rest",}, { Value: "restlet",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.restlet", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-rpc-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Adapt cloud friendly protocol", Effort: 2, Readiness: 10, Impact: "", Category: "web-service", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-service",}, { Value: "rpc",}, { Value: "soap",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-security-annotations", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Java EE security appears to be used", Effort: 5, Readiness: 0, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "java-security", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Application is using a java keystore", Effort: 10, Readiness: 0, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "java-servlet-session", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "HTTP Session Java API Import. Make sure that externalized data store is used for session", Effort: 10, Readiness: 0, Impact: "", Category: "servlet", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "session",}, { Value: "web-container",}, { Value: "tas",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.servlet.http.HttpSession", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-servlet", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Servlet Java API Import", Effort: 1, Readiness: 0, Impact: "", Category: "servlet", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, { Value: "tas",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "javax.servlet.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-slf4j-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Review logging configuration and remove file appenders.", Effort: 1, Readiness: 0, Impact: "", Category: "slf4j", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "slf4j",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.slf4j", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-soap", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.* %s.*[(].*", Advice: "Consider upgrading to modern cloud native messaging", Effort: 4, Readiness: 3, Impact: "", Category: "soap", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-service",}, { Value: "soap",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "java-springboot-annotations", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Spring Boot is a positive score", Effort: -100, Readiness: 0, Impact: "", Category: "spring-boot", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "java-spring-integration", RuleType: "", Level: "", FileType: "(xml$|gradle$|java$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Spring Integration", Effort: 0, Readiness: 9, Impact: "", Category: "messaging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "spring-integration",}, { Value: "messaging",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*spring-integration-core.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*spring-integration-file.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*@EnableIntegration.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.springframework.integration.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-springframework", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "contains", DefaultPattern: "", Advice: "Presence of spring framework may indicate the app should target TAS", Effort: -100, Readiness: 0, Impact: "", Category: "springFramework", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "spring",}, { Value: "database",}, },
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
        
            { Name: "java-stateful-annotations", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Relies on application server for shared state and will require a rewrite to Stateless or externalize state storage", Effort: 10, Readiness: 0, Impact: "", Category: "stateful", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "modernize",}, },
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
        
            { Name: "java-stateless-annotations", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Consider rearchitecting if decisions is made to not use application server", Effort: 5, Readiness: 0, Impact: "", Category: "stateless", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "stateless",}, { Value: "ejb",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stateless", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ejb", Recipe: "", },
             }, },
        
            { Name: "java-struts-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 6, Readiness: 7, Impact: "", Category: "struts", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "web-app",}, { Value: "web-container",}, { Value: "frontend",}, { Value: "container",}, },
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
        
            { Name: "java-struts", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 6, Readiness: 7, Impact: "", Category: "struts", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "web-app",}, { Value: "web-container",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "java-swing", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 500, Readiness: 100, Impact: "", Category: "Swing", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "desktop-app",}, { Value: "ui",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-system-config", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(System.)%s\\(", Advice: "Review usage of environment variables and system properties and externalize.", Effort: 3, Readiness: 10, Impact: "", Category: "config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "config",}, { Value: "configuration",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
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
        
            { Name: "java-systemLoad", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(System.)%s", Advice: "Remediate to cloud friendly implentation", Effort: 1000, Readiness: 10, Impact: "", Category: "process-launch", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jni",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "Azure Spring Apps",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "loadLibrary", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "load", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-tangosol", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ ({.<].*", Advice: "Convert to cloud friendly cache implementation", Effort: 10, Readiness: 6, Impact: "", Category: "cache", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "tangosol",}, { Value: "modernize",}, },
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
        
            { Name: "java-threadUsage-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Use Managed Executor", Effort: 100, Readiness: 10, Impact: "", Category: "threading", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "threading",}, { Value: "tas",}, { Value: "container",}, },
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
        
            { Name: "java-tibco-jms", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Integrating with TIBCO BusinessWorks JMS queues from a Spring application requires vendor-specific implementation", Effort: 7, Readiness: 6, Impact: "", Category: "tibco", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "tibco",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TibjmsConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-transaction-annotations", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^\\s*@%s", Advice: "Review if distributed transcations are used and consider rearchitecting using eventual consistency", Effort: 5, Readiness: 0, Impact: "", Category: "transaction", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "transaction",}, { Value: "modernize",}, },
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
        
            { Name: "java-transportSecurity", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Servlet data protection is used, make sure that it is supported by target platform runtime", Effort: 10, Readiness: 10, Impact: "", Category: "Security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "java-websockets-import", RuleType: "standard", Level: "", FileType: "java$", Target: "line", Type: "regex", DefaultPattern: "^import\\s*%s.*", Advice: "Make sure that target platform supports websocket api", Effort: 100, Readiness: 100, Impact: "", Category: "websockets", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-socket",}, { Value: "io",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "wsdl-soap", RuleType: "standard", Level: "", FileType: "wsdl$", Target: "file", Type: "regex", DefaultPattern: "", Advice: "Consider upgrading to cloud friendly communication protocol", Effort: 4, Readiness: 3, Impact: "", Category: "soap", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "wsdl",}, { Value: "web-service",}, { Value: "soap",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*\\.wsdl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-2-1", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 10, Readiness: 10, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "javaee",}, { Value: "full-profile",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_2_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-3-0", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "ejb", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "javaee",}, { Value: "ejb-lite",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ejb-jar_3_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ejb-jar_3_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ejb-jar_3_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-mdb", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling", Effort: 7, Readiness: 7, Impact: "", Category: "MDB", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mdb",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "fullprofile",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<ejb-ref>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MessageDriven", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-ejb-remote", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application server coupling.  Consider rearchitecting to use cloud friendly communication protocol", Effort: 100, Readiness: 100, Impact: "", Category: "ejb", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "remote-ejb",}, { Value: "ejb",}, { Value: "javaee",}, { Value: "full-profile",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<remote>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-facelets", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 2, Readiness: 2, Impact: "", Category: "Facelets", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "facelets",}, { Value: "jsf",}, { Value: "web-app",}, { Value: "modernize",}, { Value: "container",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "jsf/facelets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jee", RuleType: "standard", Level: "", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Convert to Spring based application configuration or use importResource", Effort: 20, Readiness: 8, Impact: "", Category: "Config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "full-profile",}, { Value: "javaee",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "xml-jsf-1-x", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 1000, Readiness: 0, Impact: "", Category: "jsf", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, { Value: "container",}, { Value: "modernize",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_1_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-facesconfig_1_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-facesconfig_1_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-jsf-2-x", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 700, Readiness: 700, Impact: "", Category: "jsf", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-app",}, { Value: "container",}, { Value: "modernize",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-facesconfig_2_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-facesconfig_2_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-facesconfig_2_2.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-facesconfig_2_3.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-localJNDI", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "WebSphere does not allow local JNDI, refer to documentation", Effort: 10, Readiness: 10, Impact: "", Category: "jndi", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jndi",}, { Value: "local",}, { Value: "web-container",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<local-jndi-name>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-myfaces", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 5, Readiness: 2, Impact: "", Category: "myfaces", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "container",}, { Value: "modernize",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "org.apache.myfaces", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-portlet-1-0", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 500, Readiness: 700, Impact: "", Category: "version-portlet", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "portlet",}, { Value: "web-container",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "portlet-app_1_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-2-x", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 700, Readiness: 700, Impact: "", Category: "servlet", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_2_3.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-app_2_4.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-app_2_5.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-servlet-3-x", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Web Server coupling. Consider upgrading to modern framework.", Effort: 8, Readiness: 10, Impact: "", Category: "servlet", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "servlet",}, { Value: "web-container",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "web-app_3_0.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "web-app_3_1.xsd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-session-scoped-beans", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Relies on application server for shared state and will require a rewrite to Stateless or externalize state storage", Effort: 100, Readiness: 100, Impact: "", Category: "session", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "stateful",}, { Value: "session",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "scope=\"session\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-statefulEJB", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Relies on application server for shared state and will require a rewrite to Stateless or externalize state storage", Effort: 1000, Readiness: 1000, Impact: "", Category: "ejb", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "ejb",}, { Value: "stateful",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<session-type>Stateful</session-type>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-struts", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern cloud native UI framework", Effort: 10, Readiness: 10, Impact: "", Category: "struts", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "struts",}, { Value: "ui",}, { Value: "frontend",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "struts-config_1_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "struts-2.2.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "struts-2.3.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "struts-2.4.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "struts-2.5.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "struts.apache.org/tags-tiles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_2_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_2_1.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "tiles-config_3_0.dtd", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-tomahawk", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 600, Readiness: 200, Impact: "", Category: "tomahawk", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-profile",}, { Value: "tomahawk",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "myfaces.apache.org/tomahawk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-transportSecurity", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Application Server coupling", Effort: 2, Readiness: 10, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "xml-trinidad", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider upgrading to modern UI framework", Effort: 10, Readiness: 7, Impact: "", Category: "trinidad", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "jsf",}, { Value: "web-profile",}, { Value: "trinidad",}, { Value: "frontend",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "myfaces.apache.org/trinidad/config", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-webprofile", RuleType: "standard", Level: "", FileType: "xml$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 50, Readiness: 1000, Impact: "", Category: "config", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "web-profile",}, { Value: "container",}, { Value: "modernize",}, },
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
        
            { Name: "xml-xa-dataSource", RuleType: "fire-once", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Consider rearchitecting if decision is made to move off application server", Effort: 1000, Readiness: 10, Impact: "", Category: "datasource", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "transaction",}, { Value: "jta",}, { Value: "web-profile",}, { Value: "container",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<xa-datasource", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-angular-html", RuleType: "", Level: "", FileType: "(html$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Angular", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "angular",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "[*]ngFor=", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "[*]ngIf=", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-angular-package", RuleType: "", Level: "", FileType: "(json$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Angular", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "angular",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "@angular\\/.*:", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-angular-typescript", RuleType: "", Level: "", FileType: "(ts$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Angular", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "angular",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*''@angular.*'';", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-jquery-code", RuleType: "", Level: "", FileType: "(js$|html$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Jquery", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "jquery",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "[$]\\(document\\).ready\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "<script.*src=\".*jquery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "require\\(''jquery", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-jquery-package", RuleType: "", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Jquery", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "jquery-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "PackageReference.*\"jQuery\"", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-knockout-code", RuleType: "", Level: "", FileType: "(js$|html$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Knockout JS", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "knockout-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<script.*src=.*knockout", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ko.applyBindings\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ko.observable\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-knockout-package", RuleType: "", Level: "", FileType: "(json$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Knockout JS", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "knockout-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "\"knockout\".*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-react-code", RuleType: "", Level: "", FileType: "(js$|ts$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of React JS", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "react-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "React.Component", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "import\\sReact\\sfrom\\s''react''", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-react-package", RuleType: "", Level: "", FileType: "(json$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of React JS", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "react-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "\"react\".*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-vue-code", RuleType: "", Level: "", FileType: "(js$|vue$|html$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Vue JS", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "vue-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*from\\s''vue''", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "vue.js", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "new\\sVue\\s*[(]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "javascript-vue-package", RuleType: "", Level: "", FileType: "(json$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Vue JS", Effort: 0, Readiness: 9, Impact: "", Category: "javascript-framework", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "javascript-framework",}, { Value: "vue-js",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "\"vue\".*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-bool-comparison", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(!|=)=\\\\s+%s", Advice: "Boolean literals should be avoided in comparison expressions == and != to improve code readability.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "bool-compare",}, { Value: "readability",}, { Value: "container",}, { Value: "modernize",}, { Value: "frontend",}, },
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
        
            { Name: "js-cache", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Using caching in memory will not scale when used in a multi nodes environment. Use distributed caching like redis instead.", Effort: 20, Readiness: 1000, Impact: "", Category: "session_management", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "frontend",}, },
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
        
            { Name: "js-cipher-algorithms-should-be-robust", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s+.+,+)", Advice: "Weak cipher algorithms are used. A general recommendation is to only use cipher algorithms intensively tested and promoted by the cryptographic community.", Effort: 10, Readiness: 1000, Impact: "", Category: "vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "frontend",}, },
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
             { Type: "", Pattern: "", Value: "RC5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "MD5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "SHA-1", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "WEP", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "CAST-128", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Blowfish", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-continue", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(\\\\s+)?%s(\\\\s+)?;", Advice: "continue is an unstructured control flow statement. It makes code less testable, less readable and less maintainable. Structured control flow statements such as if should be used instead.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "testability",}, { Value: "modernize",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "continue", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-document-write.yaml", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "The use of document.write where native DOM alternatives such as document.createElement are more appropriate. document.write has been grossly misused over the years and has quite a few disadvantages, including that if it’s executed after the page has been loaded, it can actually overwrite the page. Opting for more DOM-friendly methods such as document.createElement is more favourable", Effort: 5, Readiness: 1000, Impact: "", Category: "code_smell", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "modernize",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "document.write", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-fileIO", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".%s", Advice: "Cloud-native applications should not rely on a local filesystem to store information, configuration or logs. Some alternatives could be cloud storage (buckets), storing data in a shared database or leveraging shared mounted volumes in a kubernetes like environment.", Effort: 30, Readiness: 8, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "modernize",}, { Value: "container",}, { Value: "frontend",}, },
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
        
            { Name: "js-function", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\\\s+[A-Z]", Advice: "Shared naming conventions allow teams to collaborate efficiently. This rule checks that all function names match a provided regular expression.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "function",}, { Value: "faas",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "function", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-jwt-signed-verify-with-strong-cipher-algorithms", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".*\\s*{\\s(%s\\s*:+\\s*.none.\\s*}|.*\\s*(%ss:+\\s*\\[\\s*.*,+\\s*.*none.*\\s*]+))", Advice: "If a JSON Web Token (JWT) is not signed with a strong cipher algorithm (or not signed at all) an attacker can forge it and impersonate user identities. Do not use none algorithm to sign or verify the validity of a token. Do not use a token without verifying its signature before.", Effort: 10, Readiness: 1000, Impact: "", Category: "vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "algorithm", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-md5-sha1", RuleType: "", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".*(%s\\(.*\\))", Advice: "", Effort: 10, Readiness: 1000, Impact: "", Category: "security", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  },
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
        
            { Name: "js-stdout", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\s*\\(.*\\)", Advice: "Debug statements are always useful during development. But include them in production code - particularly in code that runs client-side - and you run the risk of inadvertently exposing sensitive information, slowing down the browser, or even erroring-out the site for some users.", Effort: 2, Readiness: 4, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "modernize",}, { Value: "security",}, { Value: "frontend",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "console.log", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-symbol", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "new\\\\s+%s", Advice: "Symbol is a primitive type introduced in ECMAScript2015. Its instances are mainly used as unique property keys. An instance can only be created by using Symbol as a function. Using Symbol with the new operator will raise a TypeError.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "type-error",}, { Value: "frontend",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Symbol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-throw-literal", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(%s)\\\\s+(\\\\d|''.+''|\\\".+\\\")", Advice: "It is a bad practice to throw something that's not derived at some level from Error. If you can't find an existing Error type that suitably conveys what you need to convey, then you should extend Error to create one.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "exception",}, { Value: "frontend",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "throw", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "js-md5-sha1-noncompliant", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*(1.1)+", Advice: "TLS1.1 is not secure", Effort: 5, Readiness: 1000, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "frontend",}, },
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
        
            { Name: "js-var", RuleType: "standard", Level: "", FileType: "js$", Target: "line", Type: "regex", DefaultPattern: "(\\\\s+)?%s\\\\s+", Advice: "ECMAScript 2015 introduced the let and const keywords for block-scope variable declaration. Using const creates a read-only (constant) variable.", Effort: 1, Readiness: 10, Impact: "", Category: "design", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "read-only",}, { Value: "frontend",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "var", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-logging", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Logging to the Event Log is not recommended for cloud native apps. Write to or manage logfiles. Instead, each running process should write its event stream, unbuffered, to stdout. https://docs.lacunasoftware.com/en-us/articles/amplia/on-premises/windows/enable-stdout-log.html", Effort: 100, Readiness: 3, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "dotnet-logging-config", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Logging to the Event Log is not recommended for cloud native apps. write to or manage logfiles. Instead, each running process should write its event stream, unbuffered, to stdout. https://docs.lacunasoftware.com/en-us/articles/amplia/on-premises/windows/enable-stdout-log.html", Effort: 100, Readiness: 5, Impact: "", Category: "logging", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "dotnet-splunk", RuleType: "", Level: "", FileType: "(cs$|vb$|config$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Splunk", Effort: 0, Readiness: 9, Impact: "", Category: "logging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "Splunk",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*using\\sSplunk.Client", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*http[s]*://.*splunk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Splunk.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-splunk", RuleType: "", Level: "", FileType: "(java$|properties$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Splunk", Effort: 0, Readiness: 9, Impact: "", Category: "logging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "splunk",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.splunk.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*Authorization.*Splunk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i).*Splunk.*Log.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i).*Log.*Splunk.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "(?i).*splunk_.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*http[s]*://.*splunk", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*Splunk.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*com.splunk.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "log2file-import", RuleType: "standard", Level: "", FileType: "(jsp$|java$)", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Logging should be to console", Effort: 1, Readiness: 8, Impact: "", Category: "log2file", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "log2file",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "java.util.logging.FileHandler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "java.util.logging", Recipe: "", },
             }, },
        
            { Name: "log4j-properties", RuleType: "standard", Level: "", FileType: "log4j2.properties$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Refer to platform documentation", Effort: 1, Readiness: 0, Impact: "", Category: "log2file", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "log2file",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "property.filename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "log4j-xml", RuleType: "standard", Level: "", FileType: "log4j2.xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Refer to platform documentation", Effort: 1, Readiness: 0, Impact: "", Category: "log2file", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "log2file",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "property.filename", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "mainframe-code", RuleType: "", Level: "", FileType: "(cpy$|cob$|cbl$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Mainframe Code", Effort: 1000, Readiness: 9, Impact: "", Category: "mainframe", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "mainframe",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*DIVISION.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*DISPLAY.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*INSPECT.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*STOP RUN.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*PROGRAM-ID.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*[0-9][0-9]\\s.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-azure-servicebus", RuleType: "", Level: "", FileType: "(cs$|vb$|csproj$|vbproj$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Azure Service Bus API usage in .NET code, including client creation and message operations.", Effort: 0, Readiness: 9, Impact: "", Category: "messaging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "messaging",}, { Value: "azure-servicebus",}, { Value: ".NET",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "new ServiceBusClient", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".CreateSender", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "new ServiceBusMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".SendMessageAsync", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".ReceiveMessageAsync\\(\\)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Azure.Messaging.ServiceBus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-messaging", RuleType: "", Level: "", FileType: "(cs$|vb$|config$|csproj$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "", Effort: 5, Readiness: 1, Impact: "", Category: "messaging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "messaging",}, { Value: ".NET",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Apache.NMS.*", Advice: "Usage of ActiveMQ library in .NET project files and XML configuration settings. Migrate to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*Flume.*", Advice: "Direct usage of Apache Flume .NET Client APIs in .NET code, indicating logging and data forwarding to Flume.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flume", Recipe: "", },
             { Type: "", Pattern: "", Value: "Storm.*", Advice: "Integration with Apache Storm using the net-storm-multilang-adapter in .NET applications.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-storm", Recipe: "", },
             { Type: "", Pattern: "", Value: "IBMMQ", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibm-mq", Recipe: "", },
             { Type: "", Pattern: "", Value: "Msmq", Advice: "MSMQ is a proprietary Windows server-based messaging solution that is not supported by Linux based containers. Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "msmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "Kafka", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kafka", Recipe: "", },
             { Type: "", Pattern: "", Value: "Solace", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "solace", Recipe: "", },
             { Type: "", Pattern: "", Value: "MassTransit", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator. Out of box support for RabbitMQ", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "masstransit", Recipe: "", },
             { Type: "", Pattern: "", Value: "NServiceBus", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator. Out of box support for RabbitMQ", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nservicebus", Recipe: "", },
             { Type: "", Pattern: "", Value: "DotPulsar", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-pulsar", Recipe: "", },
             { Type: "", Pattern: "", Value: "rocketmq", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rocketmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "zeromq", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "zeromq", Recipe: "", },
             { Type: "", Pattern: "", Value: "NATS.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nats", Recipe: "", },
             { Type: "", Pattern: "", Value: "RabbitMQ", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rabbitmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "amqmdnet", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "webspheremq", Recipe: "", },
             }, },
        
            { Name: "dotnet-mulesoft", RuleType: "", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Usage of Mulesoft", Effort: 0, Readiness: 9, Impact: "", Category: "api", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "api",}, { Value: "mulesoft",}, { Value: "api-gateway",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RAML.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "Raml.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "AMF.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-azure-servicebus", RuleType: "", Level: "", FileType: "(java$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Implementation of Azure Service Bus API in Java code.", Effort: 0, Readiness: 1, Impact: "", Category: "messaging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "messaging",}, { Value: "azure-servicebus",}, { Value: "java",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "new ServiceBusClientBuilder\\(\\)", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "new ServiceBusMessage", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceBusMessageBatch", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "ServiceBusReceivedMessageContext", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "azure-messaging-servicebus", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-messaging", RuleType: "standard", Level: "", FileType: "(java$|gradle$|yml$|properties$|xml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Remediate any persistence issues", Effort: 7, Readiness: 7, Impact: "", Category: "messaging", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "messaging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "amq:broker", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "createActiveMQConnection", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "import.*org.apache.activemq", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "new NMSConnectionFactory", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "NMSCorrelationID", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "new Uri\\(\"activemq:", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "ActiveMQConnectionFactory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.flink.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flink", Recipe: "", },
             { Type: "", Pattern: "", Value: ".getExecutionEnvironment\\(\\)", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flink", Recipe: "", },
             { Type: "", Pattern: "", Value: ".setStreamTimeCharacteristic\\(", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flink", Recipe: "", },
             { Type: "", Pattern: "", Value: "flink*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flink", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.flume.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flume", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*Flume*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flume", Recipe: "", },
             { Type: "", Pattern: "", Value: "flume-*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-flume", Recipe: "", },
             { Type: "", Pattern: "", Value: ".storm.", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-storm", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.apache.storm", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-storm", Recipe: "", },
             { Type: "", Pattern: "", Value: "storm-*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-storm", Recipe: "", },
             { Type: "", Pattern: "", Value: "javax.jms", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jms", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*QueueConnectionFactory.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jms", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*TopicConnectionFactory.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jms", Recipe: "", },
             { Type: "", Pattern: "", Value: "kafka", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kafka", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-kafka", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kafka", Recipe: "", },
             { Type: "", Pattern: "", Value: "io.nats", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "nats", Recipe: "", },
             { Type: "", Pattern: "", Value: "pulsar", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "apache-pulsar", Recipe: "", },
             { Type: "", Pattern: "", Value: "import.*com.ibm.mq.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibmmq", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*MQ.*", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "ibmmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "msmq.", Advice: "MSMQ is a proprietary Windows server-based messaging solution that is not supported by Linux based containers. Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "msmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "com.rabbitmq", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rabbitmq", Recipe: "", },
             }, },
        
            { Name: "java-mulesoft", RuleType: "", Level: "", FileType: "(java$|cls$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Detection of Mulesoft Usage", Effort: 0, Readiness: 1, Impact: "", Category: "api", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "api",}, { Value: "mulesoft",}, { Value: "messaging",}, { Value: "streaming",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "com.mulesoft", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".mulesoft.com", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "www.mulesoft.org", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*mule-.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.mule.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-messaging", RuleType: "", Level: "", FileType: "(py$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "", Effort: 5, Readiness: 1, Impact: "", Category: "messaging", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "messaging",}, { Value: "python",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "pyactivemq", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "stomp", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "activemq", Recipe: "", },
             { Type: "", Pattern: "", Value: "kafka", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "kafka", Recipe: "", },
             { Type: "", Pattern: "", Value: "pika", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rabbitmq", Recipe: "", },
             { Type: "", Pattern: "", Value: "aio_pika", Advice: "Move to RabbitMQ tile or RabbitMQ k8s operator", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rabbitmq", Recipe: "", },
             }, },
        
            { Name: "dotnet-ISAPI-Filters-config", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 1000, Readiness: 10, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/isapiFilters[1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-ISAPI-Filters-vbcs", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 1000, Readiness: 0, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "isapi-filter",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-ip-address", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported in .netCore, Use IHttpContextAccessor instead", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "debt",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Request.UserHostAddress", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-MSMQ-vbcs", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "MSMQ is a proprietary Windows server-based messaging solution that is not supported by Linux based containers. Leverage bridging solutions like https://kubemq.io/helping-net-msmq-based-services-to-migrate-to-kubernetes/ on kubernetes", Effort: 10, Readiness: 1, Impact: "", Category: "unsupported-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "msmq",}, { Value: "messaging",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-iis_module-UriCacheModule", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux. Implements a generic cache for URL-specific server state; such as configuration. With this module; the server only reads configuration for the first request for a particular URL. And reuse it on subsequent requests until it changes.", Effort: 1000, Readiness: 5, Impact: "", Category: "Unsupported IIS modules", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "uri-cache",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/globalModules/add[@name=\"UriCacheModule\"][1]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-child-action", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*\\[%s\\].*", Advice: "Not supported, replace with a new View Component feature https://learn.microsoft.com/en-us/aspnet/core/mvc/views/view-components?view=aspnetcore-7.0", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "mvc",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ChildActionOnly", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-classic-2-0", RuleType: "standard", Level: "", FileType: "(asp$)", Target: "line", Type: "regex", DefaultPattern: "\\s%s.*", Advice: "Not supported by many PAS. ASP .net pages will have to be replaced by ASP .NetCore", Effort: 10, Readiness: 9, Impact: "", Category: "Unsupported", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "asp-classic",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-asp-machine-key", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Not supported in .NetCore. Will have to be replaced https://learn.microsoft.com/en-us/aspnet/core/security/data-protection/compatibility/replacing-machinekey?view=aspnetcore-7.0", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "encryption",}, { Value: "config",}, { Value: "asp-classic",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/machineKey", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-membership", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Not supported, replaced by ASP.NET Core Identity. Manage users in Database.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "authentication",}, { Value: "asp",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/membership", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-mvc-form-collection", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: "%s", Advice: "Not supported, Replace with IFormCollection", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mvc",}, { Value: "asp",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*Create.*\\(FormCollection.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-mvc-model-update", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*\\s%s<.*", Advice: "Not supported in .NetCore, Replace with TryUpdateModelAsync", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mvc",}, { Value: "asp",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "TryUpdateModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-mvc", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported on .netCore. Migrate to ASP .NetCore and replace with Microsoft.AspNetCore.Mvc", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "mvc",}, { Value: "asp",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Web.Mvc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-session-context", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported, Replace with IHttpContextAccessor to manage session data", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "session",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "HttpContextBase", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-asp-web-form", RuleType: "standard", Level: "", FileType: "(aspx$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Unsupported-netcore. Not supported by many PAS. ASP .net pages will have to be replaced by ASP .NetCore. See https://learn.microsoft.com/en-us/dotnet/architecture/porting-existing-aspnet-apps/migrate-web-forms", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore.", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "asp-web-form",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-classic-asp-net", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported-netcore. Not supported by many PAS. ASP .net pages will have to be replaced by ASP .NetCore. See https://learn.microsoft.com/en-us/dotnet/architecture/porting-existing-aspnet-apps/migrate-web-forms", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "asp-classic",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Web.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-delegate-begininvoke", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "BeginInvoke and EndInvoke methods on delegates are supported on .NET Core. See https://devblogs.microsoft.com/dotnet/migrating-delegate-begininvoke-calls-for-net-core/", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "netcore",}, { Value: "mvc",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-identity-foundation", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Identity Foundation (WIF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "WIF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.IdentityModel.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-AnonymousAuthentication", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/authentication/anonymousAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Authentication", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux. Use Microsoft.AspNetCore.Authentication.Negotiate instead.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  { Value: ".authentication.*mode=\"None\"",}, },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/authentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Authorization", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/authorization", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-CertificateMappingAuthentication", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/authentication/clientCertificateMappingAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-DefaultDocument", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/defaultDocument", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-DigestAuthentication", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/authentication/digestAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-DirectoryListing", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/directoryBrowse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-HttpErrors", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpErrors", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-HttpLogging", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpLogging", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-HttpRedirection", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpRedirect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-IisClientCertificateMappingAuthentication", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/iisClientCertificateMappingAuthentication", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-IpSecurity", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/security/ipSecurity", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-IsapiCgiRestriction", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/isapiCgiRestriction", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-OutputCache", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/caching/outputCache", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-ProtocolSupport", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/httpProtocol", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-ServerSideInclude", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/serverSideInclude", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-iis_module-Tracing", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
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
        
            { Name: "dotnet-iis_module-Validation", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows and IIS server dependency. Not supported in .netCore targeting linux.", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "iis-module",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.webServer/validation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-javascript-serializer", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Use JSON.NET instead.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "JavaScriptSerializer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-managed-addin", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Managed Addin Framework (MAF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.AddIn", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-app-metrics-extensions-hosting", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "App.Metrics.Extensions.Hosting", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-mvc5", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.Mvc5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-owin", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.Owin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-webapi2-owin", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.WebApi2.Owin", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-autofac-webapi2", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Autofac.WebApi2", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-dotnetopenauth-core", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "DotNetOpenAuth.Core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-fluentvalidation-mvc5", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FluentValidation.MVC5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-fluentvalidation-webapi", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "FluentValidation.WebAPI", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-glimpse", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Glimpse", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-imageprocessor-web", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ImageProcessor.Web", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-imageprocessor", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ImageProcessor", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-imageresizer", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ImageResizer", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-friendlyurls-core", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.FriendlyUrls.Core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-identity-entityframework", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.Identity.EntityFramework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-mvc", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.Mvc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-signalr", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.SignalR", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-aspnet-webapi-versioning", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.AspNet.WebApi.Versioning", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-azure-common-dependencies", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Azure.Common.Dependencies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-build-locator", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Build.Locator", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-build-tasks-git", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Build.Tasks.Git", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-codeanalysis-workspaces-msbuild", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CodeAnalysis.Workspaces.MSBuild", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-coreassemblies", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.CoreAssemblies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-deployment", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.Deployment", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-workflow", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.Workflow", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-crmsdk-xrmtooling-coreassembly", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.CrmSdk.XrmTooling.CoreAssembly", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-data-sqlclient-sni", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Data.SqlClient.SNI", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-diagnostics-tracing-eventsource-redist", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Diagnostics.Tracing.EventSource.Redist", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net45", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net45", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net461", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net461", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net462", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net462", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net472", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net472", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies-net48", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies.net48", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-netframework-referenceassemblies", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.NETFramework.ReferenceAssemblies", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-cors", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Cors", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-diagnostics", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Diagnostics", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-filesystems", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.FileSystems", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-activedirectory", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.ActiveDirectory", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-oauth", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.OAuth", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-openidconnect", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.OpenIdConnect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-security-wsfederation", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Security.WsFederation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-selfhost", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.SelfHost", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-staticfiles", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.StaticFiles", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-owin-testing", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Owin.Testing", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-servicefabric-aspnetcore-abstractions", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.ServiceFabric.AspNetCore.Abstractions", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-servicefabric-aspnetcore-kestrel", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.ServiceFabric.AspNetCore.Kestrel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-visualstudio-azure-containers-tools-targets", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.VisualStudio.Azure.Containers.Tools.Targets", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-visualstudio-azure-fabric-msbuild", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.VisualStudio.Azure.Fabric.MSBuild", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-visualstudio-slowcheetah", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.VisualStudio.SlowCheetah", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-microsoft-web-redissessionstateprovider", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Microsoft.Web.RedisSessionStateProvider", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-miniprofiler", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "MiniProfiler", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-ninject-mvc5", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Ninject.MVC5", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-ninject-web-common-webhost", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Ninject.Web.Common.WebHost", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-ninject-web-webapi", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Ninject.Web.WebApi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-nlog-extended", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "NLog.Extended", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-nlog-web", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "NLog.Web", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-portable-dataannotations", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Portable.DataAnnotations", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-razorengine", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "RazorEngine", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-runtime-aot-system-reflection-primitives", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "runtime.aot.System.Reflection.Primitives", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-sqlitepclraw-provider-e_sqlite3-net45", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "SQLitePCLRaw.provider.e_sqlite3.net45", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-stub-system-data-sqlite-core-netframework", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Stub.System.Data.SQLite.Core.NetFramework", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-system-data-sqlite-linq", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Data.SQLite.Linq", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-system-net-http-formatting-extension", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Net.Http.Formatting.Extension", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-system-reactive-windows-threading", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Reactive.Windows.Threading", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-telerik-dataaccess-core", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Telerik.DataAccess.Core", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-telerik-dataaccess-web", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Telerik.DataAccess.Web", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-unity-aspnet-webapi", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Unity.AspNet.WebApi", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-package-unity-mvc", RuleType: "standard", Level: "", FileType: "(config$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported in .NetCore", Effort: 10, Readiness: 10, Impact: "", Category: "package", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Unity.Mvc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-reflection-assembly", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Use AssemblyLoadContext. Evidence is not supported in .NET Core.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Assembly.Load\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-script-blocks", RuleType: "standard", Level: "", FileType: "(xsl)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Script blocks are supported only in .NET Framework. They are not supported on .NET Core or .NET 5 or later.", Effort: 10, Readiness: 10, Impact: "", Category: "WCF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "msxsl:script", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-serialization", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not supported in .netCore", Effort: 10, Readiness: 1, Impact: "", Category: "unsupported-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "serialization",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Runtime.Serialization", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-system-speech", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "System.Speech won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Speech.Recognition", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-system-web-services", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Upgrade to CoreWCF https://devblogs.microsoft.com/dotnet/upgrading-a-wcf-service-to-dotnet-6/ or more modern rest api", Effort: 10, Readiness: 10, Impact: "", Category: "WCF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-web-services",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Web.Services", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-vb-razor-view-engine", RuleType: "standard", Level: "", FileType: "(vbhtml$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Not supported in .netCore https://aka.ms/vb-angular-and-web-api", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore.", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "asp",}, { Value: "asp-web-form",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "@Code", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-service-model-config", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Upgrade to CoreWCF https://devblogs.microsoft.com/dotnet/upgrading-a-wcf-service-to-dotnet-6/", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-iis-module", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.serviceModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wcf-service-model", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Upgrade to CoreWCF https://devblogs.microsoft.com/dotnet/upgrading-a-wcf-service-to-dotnet-6/", Effort: 10, Readiness: 10, Impact: "", Category: "WCF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "communication",}, { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.ServiceModel", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wf-activity", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "WF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Activities.Presentation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wf-integration", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "WF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "XamlIntegration", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-wf-presentation", RuleType: "standard", Level: "", FileType: "(xaml$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 500, Readiness: 1, Impact: "", Category: "WF", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wf",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Activities.Core.Presentation", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-enterprise-services", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Not Supported on .Net Core", Effort: 10, Readiness: 9, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "com+",}, { Value: "api",}, { Value: "distributed-transaction",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.EnterpriseService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-remoting-code", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported, consider inter-process communication (IPC) System.IO.Pipes class or the MemoryMappedFile class.Also StreamJsonRpc or ASP.NET Core (either using gRPC or RESTful Web API services).", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "remoting",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Runtime.Remoting.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-remoting", RuleType: "standard", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Unsupported, consider inter-process communication (IPC) System.IO.Pipes class or the MemoryMappedFile class.Also StreamJsonRpc or ASP.NET Core (either using gRPC or RESTful Web API services).", Effort: 500, Readiness: 5, Impact: "", Category: "Unsupported", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "remoting",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.runtime.remoting", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windows-workflow-foundation", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Use alternatives, see CoreWF and CoreWCF, runs on windows only and would need to be ran by a console application", Effort: 10, Readiness: 1, Impact: "", Category: "unsupported-netcore", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-wcf",}, { Value: "windows-workflow-service",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "WorkflowService", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-workflow", RuleType: "standard", Level: "", FileType: "(cs$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Workflow Foundation (WF) won't be ported to .NET Core. See https://aka.ms/unsupported-netfx-api.", Effort: 10, Readiness: 10, Impact: "", Category: "API", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "api",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "netcore",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Workflow.", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-allow-url-in-config", RuleType: "standard", Level: "", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=1$", Advice: "allow_url_fopen and allow_url_include allow code to be read into a script from URL’s. The ability to suck in executable code from outside your site, coupled with imperfect input cleansing could lay your site bare to attackers explicitly disable allow_url_fopen and allow_url_include'", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-cache", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\[.*\\]", Advice: "However, the default configuration is to store session data in a temporary file on the local disk. Again, this will not work if you’re using multiple nodes", Effort: 200, Readiness: 1000, Impact: "", Category: "session_management", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-cgi-force-redirect-enabled", RuleType: "standard", Level: "", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+0$", Advice: "The cgi.force_redirect php.ini configuration is on by default, and it prevents unauthenticated access to scripts when PHP is running as a CGI.", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "cgi.force_redirect", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-deprecated-feature-parse_str", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "([^_]%s)\\s*\\(.*[^,]\\)", Advice: "parse_str() without second argument is deprecated", Effort: 1, Readiness: 1000, Impact: "", Category: "deprecation", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "deprecated",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "parse_str", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-deprecated-feature-should-not-be-used", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "([^_]%s)\\s*\\(+.*\\)", Advice: "Deprecated language features are those that have been retained temporarily for backward compatibility, but which will eventually be removed from the language. In effect, deprecation announces a grace period to allow the smooth transition from the old features to the new ones. In that period, no use of the deprecated features should be added to the code. Refactor or upgrade to use Php 7+", Effort: 1, Readiness: 1000, Impact: "", Category: "deprecation", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "deprecated",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-disabled-enable-dl", RuleType: "standard", Level: "", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+1$", Advice: "enable_dl is on by default and allows open_basedir restrictions, which limit the files a script can access, to be ignored enable_dl should be explicitly turned off in php.ini", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "enable_dl", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-disabled-file-uploads", RuleType: "standard", Level: "", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+1$", Advice: "file_uploads is an on-by-default PHP configuration that allows files to be uploaded to your site.", Effort: 5, Readiness: 1000, Impact: "", Category: "Vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "file_uploads", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-drupal", RuleType: "", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "Drupal application", Effort: 100, Readiness: 0, Impact: "", Category: "drupal", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "drupal",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "drupal_route", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "modernize", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "container", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-enable-session-use-trans-sid", RuleType: "standard", Level: "", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s)=+1$", Advice: "PHP’s session.use_trans_sid automatically appends the user’s session id to urls when cookies are disabled.", Effort: 100, Readiness: 1000, Impact: "", Category: "Vulnerability", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "vulnerability",}, { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "session.use_trans_sid", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-file-system-manipulation", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Filesystem manipulation is not encouraged in cloud-native applications. Keep your content off the filesystem", Effort: 500, Readiness: 1000, Impact: "", Category: "filesystem", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-function-method-naming-convention", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: "^(%s)\\s+[a-z]+[a-zA-Z0-9]*\\(.*\\)", Advice: "Shared naming conventions allow teams to collaborate efficiently. All function names match a provided regular expression.", Effort: 1, Readiness: 1000, Impact: "", Category: "naming_convention", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "code-smell",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "function", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-global-variable", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(\\$%s)\\s*\\[.*\\]", Advice: "The default configuration is to store session data in a temporary file on the local disk. Again, this will not work if you’re using multiple nodes", Effort: 200, Readiness: 1000, Impact: "", Category: "session_management", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "cache",}, { Value: "session",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-goto-stmt-should-not-be-used", RuleType: "standard", Level: "", FileType: "ini$", Target: "line", Type: "regex", DefaultPattern: "\\s*(%s\\s*.+)", Advice: "goto is an unstructured control flow statement. It makes code less readable and maintainable. Structured control flow statements such as if, for, while, continue or break should be used instead.", Effort: 1, Readiness: 1000, Impact: "", Category: "code_smell", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "unstructured",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "goto", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-md5-sha1-noncompliant", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s\\(.*\\))", Advice: "Cryptographic hash algorithms such as MD2, MD4, MD5, MD6, HAVAL-128, HMAC-MD5, DSA (which uses SHA-1), RIPEMD, RIPEMD-128, RIPEMD-160, HMACRIPEMD160 and SHA-1 are no longer considered secure, because it is possible to have collisions (little computational effort is enough to find two or more different inputs that produce the same hash). Safer alternatives, such as SHA-256, SHA-512, SHA-3 are recommended, and for password hashing, it is even better to use algorithms that do not compute too quickly, like bcrypt, scrypt, argon2 or pbkdf2 because it slows down brute force attacks.", Effort: 10, Readiness: 1000, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-reading-std-Input-security-sensitive", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Reading Standard Input is security-sensitive. It has led in the past to the following vulnerabilities:CVE-2005-2337,CVE-2017-11449", Effort: 50, Readiness: 1000, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "security",}, { Value: "filesystem",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "php-references-should-not-be-passed-to-func-call", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*\\s*\\(%s+.*\\)", Advice: "Passing a reference to a function parameter means that any modifications the method makes to the parameter will be made to the original value as well, since references have the effect of pointing two variables at the same memory space. This feature can be difficult to use correctly, particularly if the callee is not expecting a reference, and the improper use of references in function calls can make code less efficient rather than more efficient.", Effort: 1, Readiness: 1000, Impact: "", Category: "code_smell", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "code_smell",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "&", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-removing-accent-marks", RuleType: "", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Most web guides will suggest using PHP’s iconv() function to remove diacritics. However iconv() often has trouble with UTF-8 input and will sometimes produce surprising errors.", Effort: 5, Readiness: 1000, Impact: "", Category: "code-smell", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "code-smell",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "iconv", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "php-socket-security-sensitive", RuleType: "standard", Level: "", FileType: "php$", Target: "line", Type: "regex", DefaultPattern: ".*(%s)\\s*\\(.*\\)", Advice: "Using sockets is security-sensitive. It has led in the past to the following vulnerabilities:CVE-2011-178,CVE-2017-5645,CVE-2018-6597", Effort: 200, Readiness: 1000, Impact: "", Category: "anti-pattern", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "anti-pattern",}, { Value: "security",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "python-cf", RuleType: "standard", Level: "", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*import(\\s*|=\")%s.*$", Advice: "Check for cloud foundry support.", Effort: -10, Readiness: 10, Impact: "", Category: "cloudfoundry", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "container",}, { Value: "tas",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "load_from_vcap_services", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "python-fileIO", RuleType: "standard", Level: "", FileType: "py$", Target: "line", Type: "regex", DefaultPattern: "^.*[ .]%s[ (].*", Advice: "Relying on the local filesystem to store state is unreliable in a cloud platform. Since containers are immutable, restarts will lose any written changes. Refactor this logic to use an external service to store state.", Effort: 8, Readiness: 8, Impact: "", Category: "io", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "io",}, { Value: "modernize",}, { Value: "container",}, },
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
        
            { Name: "ruby-file-io", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "File i/o in environments that are emphemeral is a bad practice", Effort: 2, Readiness: 0, Impact: "", Category: "io", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "io",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "IO", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-gemfile", RuleType: "", Level: "", FileType: "Gemfile$", Target: "file", Type: "simple-text", DefaultPattern: "", Advice: "Ruby on Rails", Effort: 5, Readiness: 1000, Impact: "", Category: "docker", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "ruby-gemfile",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Gemfile", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-kernel", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "Kernel operations my indicate actions that are not cloud/container friendly", Effort: 2, Readiness: 0, Impact: "", Category: "kernel", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "kernel",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Kernel\\.open\\(", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-ldap", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "Use of ldap in a cloud/container env is not a best practice", Effort: 1, Readiness: 0, Impact: "", Category: "windows-reg", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-reg",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "LDAP", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-log2file", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "Logging to a file", Effort: 2, Readiness: 0, Impact: "", Category: "servlet", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "log2file",}, { Value: "logging",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Logger.new\\(''", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-requires", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\brequire\\s+[''\"]%s[''\"].*$", Advice: "Consult 3rd party documentation", Effort: 300, Readiness: 10, Impact: "", Category: "third-party", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "third-party",}, { Value: "dependencies",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "listen", Advice: "This gem is used for file system monitoring and can cause issues when used in a container environment. It can lead to high CPU usage and memory leaks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "file system monitoring", Recipe: "", },
             { Type: "", Pattern: "", Value: "rb-inotify", Advice: "This gem is similar to `listen` and can also cause issues in a container environment. It can lead to high CPU usage and memory leaks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "system monitoring", Recipe: "", },
             { Type: "", Pattern: "", Value: "rb-fsevent", Advice: "This gem is used for file system monitoring on macOS and can cause issues when used in a container environment. It can lead to high CPU usage and memory leaks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "macOS file monitoring", Recipe: "", },
             { Type: "", Pattern: "", Value: "rb-kqueue", Advice: "This gem is used for file system monitoring on BSD-based systems and can cause issues when used in a container environment. It can lead to high CPU usage and memory leaks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "BSD file monitoring", Recipe: "", },
             { Type: "", Pattern: "", Value: "rb-fchange", Advice: "This gem is used for file system monitoring and can cause issues when used in a container environment. It can lead to high CPU usage and memory leaks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "file system monitoring", Recipe: "", },
             { Type: "", Pattern: "", Value: "docker-sync", Advice: "docker-sync usage indicates this application has been preppred to run in a container", Effort: -200, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "docker-compose", Advice: "docker-compose usage indicates this application has been preppred to run in a container", Effort: -200, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "sinatra", Advice: "Sinatra applications require special consideration migrating to the cloud/container", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "sinatra", Recipe: "", },
             { Type: "", Pattern: "", Value: "socket", Advice: "Socket traffic may pose a problem in some containership/cloud settings", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "socket", Recipe: "", },
             { Type: "", Pattern: "", Value: "gtk3", Advice: "Desktop applications cannot run in containers/clouds", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "desktop-app", Recipe: "", },
             { Type: "", Pattern: "", Value: "sinatra/activerecord", Advice: "While ActiveRecord is a powerful ORM, its tight coupling to SQL databases and reliance on a centralized connection pool can make it difficult to horizontally scale in cloud environments.", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "ORM", Recipe: "", },
             { Type: "", Pattern: "", Value: "activerecord", Advice: "While ActiveRecord is a powerful ORM, its tight coupling to SQL databases and reliance on a centralized connection pool can make it difficult to horizontally scale in cloud environments.", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "ORM", Recipe: "", },
             { Type: "", Pattern: "", Value: "rmagick", Advice: "While RMagick is a powerful image processing library, its dependence on ImageMagick can make it difficult to install and configure in certain cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "image-processing", Recipe: "", },
             { Type: "", Pattern: "", Value: "nokogiri", Advice: "While Nokogiri is a popular XML parsing library, its reliance on libxml2 can make it difficult to install and configure in certain cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "xml-parsing", Recipe: "", },
             { Type: "", Pattern: "", Value: "unicorn", Advice: "While Unicorn is a popular web server, its master-worker architecture can make it difficult to horizontally scale in cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-server", Recipe: "", },
             { Type: "", Pattern: "", Value: "puma", Advice: "While Puma is another popular web server, its multi-process architecture can make it difficult to vertically scale in cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "web-server", Recipe: "", },
             { Type: "", Pattern: "", Value: "resque", Advice: "While Resque is a popular background job processor, its reliance on Redis can make it difficult to horizontally scale in cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "job-processor", Recipe: "", },
             { Type: "", Pattern: "", Value: "sidekiq", Advice: "While Sidekiq is another popular background job processor, its reliance on Redis can make it difficult to horizontally scale in cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "job-processor", Recipe: "", },
             { Type: "", Pattern: "", Value: "delayedjob", Advice: "While DelayedJob is a popular background job processor, its reliance on a centralized queue can make it difficult to horizontally scale in cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "job-processor", Recipe: "", },
             { Type: "", Pattern: "", Value: "activestorage", Advice: "While ActiveStorage is a convenient feature for handling files in Rails applications, its reliance on a centralized storage solution can make it difficult to horizontally scale in cloud environments.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "rails", Recipe: "", },
             { Type: "", Pattern: "", Value: "win32ole", Advice: "This gem provides an interface to the Windows OLE (Object Linking and Embedding) system, which allows Ruby programs to interact with Windows applications and controls.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "windows", Recipe: "", },
             { Type: "", Pattern: "", Value: "win32api", Advice: "This gem provides an interface to the Windows API (Application Programming Interface), which allows Ruby programs to interact with Windows system functions and services.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "windows", Recipe: "", },
             { Type: "", Pattern: "", Value: "win32gui", Advice: "This gem provides an interface to the Windows GUI (Graphical User Interface), which allows Ruby programs to create and manipulate Windows GUI elements such as windows, buttons, and menus.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "windows", Recipe: "", },
             { Type: "", Pattern: "", Value: "win32process", Advice: "This gem provides an interface to the Windows process management system, which allows Ruby programs to create, manage, and communicate with Windows processes.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "windows", Recipe: "", },
             { Type: "", Pattern: "", Value: "win32service", Advice: "This gem provides an interface to the Windows service management system, which allows Ruby programs to create, manage, and communicate with Windows services.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "windows", Recipe: "", },
             { Type: "", Pattern: "", Value: "osx_app", Advice: "This gem provides an interface to the OSX application programming interface (API), which allows Ruby programs to interact with OSX applications and services.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osx", Recipe: "", },
             { Type: "", Pattern: "", Value: "osx_core_foundation", Advice: "This gem provides an interface to the OSX Core Foundation framework, which allows Ruby programs to interact with OSX system services and frameworks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osx", Recipe: "", },
             { Type: "", Pattern: "", Value: "osx_foundation", Advice: "This gem provides an interface to the OSX Foundation framework, which allows Ruby programs to interact with OSX system services and frameworks.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osx", Recipe: "", },
             { Type: "", Pattern: "", Value: "osx_gui", Advice: "This gem provides an interface to the OSX GUI (Graphical User Interface), which allows Ruby programs to create and manipulate OSX GUI elements such as windows, buttons, and menus.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osx", Recipe: "", },
             { Type: "", Pattern: "", Value: "osx_process", Advice: "This gem provides an interface to the OSX process management system, which allows Ruby programs to create, manage, and communicate with OSX processes.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osx", Recipe: "", },
             { Type: "", Pattern: "", Value: "osx_service", Advice: "This gem provides an interface to the OSX service management system, which allows Ruby programs to create, manage, and communicate with OSX services.", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "osx", Recipe: "", },
             { Type: "", Pattern: "", Value: "ruby-linux", Advice: "This gem provides a set of classes and methods for interacting with the Linux operating system, including file system manipulation, process management, and network communication.", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "linux", Recipe: "", },
             { Type: "", Pattern: "", Value: "linux-sys", Advice: "This gem provides a set of classes and methods for interacting with the Linux kernel, including system calls, device drivers, and kernel modules", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "linux", Recipe: "", },
             { Type: "", Pattern: "", Value: "linux-users", Advice: "This gem provides a set of classes and methods for managing Linux users and groups, including user authentication, password management, and group membership.", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "linux", Recipe: "", },
             { Type: "", Pattern: "", Value: "linux-network", Advice: "This gem provides a set of classes and methods for managing Linux network interfaces, including IP addresses, routing tables, and network protocols.", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "linux", Recipe: "", },
             { Type: "", Pattern: "", Value: "linux-storage", Advice: "This gem provides a set of classes and methods for managing Linux storage devices, including disk partitions, file systems, and storage drivers.", Effort: 1, Readiness: 0, Criticality: "", Category: "", Tag: "linux", Recipe: "", },
             }, },
        
            { Name: "ruby-routing", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "Ruby on rails application", Effort: 1, Readiness: 0, Impact: "", Category: "ruby-on-rails", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "ruby-on-rails",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "ActionController", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-rpc", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "^%s", Advice: "Remote procedure calls can be problematic in cloud/container envs.", Effort: 5, Readiness: 0, Impact: "", Category: "rpc", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "rpc",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "rpc", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "ruby-windows-reg", RuleType: "", Level: "", FileType: "rb$", Target: "line", Type: "regex", DefaultPattern: "\\b%s", Advice: "Windows Registry is not available in cloud/container envs", Effort: 2, Readiness: 0, Impact: "", Category: "windows-reg", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-reg",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "Registry", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-quartz", RuleType: "", Level: "", FileType: "(java$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of Quartz Scheduling", Effort: 0, Readiness: 9, Impact: "", Category: "scheduling", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "quartz",}, { Value: "scheduler",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: ".*org.springframework.scheduling.quartz.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*org.quartz.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: ".*spring-boot-starter-quartz*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "config-encryption", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Use of encrypted sections is problematic in cloud environments. Configurations should be externalized as environment variables instead of encrypting them in configuration file.", Effort: 300, Readiness: 0, Impact: "", Category: "security", Criticality: "",Container: 0, Cloud: 0,
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
        
            { Name: "config-security", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Use of generated machine keys is problematic", Effort: 100, Readiness: 100, Impact: "", Category: "security", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "auto-gen-key",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/system.web/machineKey[contains(@validationKey, \"AutoGenerate\") or contains(@decryptionKey, \"AutoGenerate\")]", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-WindowsAuth-config", RuleType: "fire-once", Level: "", FileType: "config$", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Windows Integrated Authentication will be problematic in the cloud as containers won't be joining Windows Domains. Use another authentication mechanism more cloud friendly like SSO, LDAP Remote Calls, Local...", Effort: 500, Readiness: 1, Impact: "", Category: "Unsupported modules", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-auth",}, { Value: "security",}, },
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
        
            { Name: "dotnet-WindowsAuth-csvb", RuleType: "fire-once", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Windows Integrated Authentication will be problematic in the cloud as containers won't be joining Windows Domains. Use another authentication mechanism more cloud friendly like SSO, LDAP Remote Calls, Local...", Effort: 500, Readiness: 1, Impact: "", Category: "Unsupported modules", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "windows-auth",}, { Value: "security",}, },
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
        
            { Name: "dotnet-dotnetopenauth-config", RuleType: "", Level: "", FileType: "(config$)", Target: "file", Type: "xpath", DefaultPattern: "", Advice: "Usage of Dotnet Open Auth", Effort: 0, Readiness: 9, Impact: "", Category: "authentication", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "oauth-authentication",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "/configuration/dotNetOpenAuth", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "DotNetOpenAuth", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-security", RuleType: "standard", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Relying on Windows certificate stores is problematic in a cloud environment. Certificate stores could be shared via shared volume mounts. Ideally you should decouple certificates and SSL from your application. SSL could also be handled at the load balancer level.", Effort: 3, Readiness: 10, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "dotnet-windows-application-domain", RuleType: "fire-once", Level: "", FileType: "(cs$|vb$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Unsupported, For .NET Core, there is exactly one AppDomain. Isolation and unloading are provided through AssemblyLoadContext. Security boundaries should be provided by process boundaries and appropriate remoting techniques", Effort: 1000, Readiness: 9, Impact: "", Category: "windows-domain", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "app-domain",}, { Value: "security",}, { Value: "unsupported-netcore",}, { Value: "modernize",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "AppDomain.CreateDomain", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "dotnet-windowsPrincipal", RuleType: "standard", Level: "", FileType: "(cs$|vb$|csproj$)", Target: "line", Type: "regex", DefaultPattern: ".*%s.*", Advice: "Operations requiring a Windows domain are not supported in a container like environment. Features relying on Windows authentication will be problematic if targeting linux environment.", Effort: 100, Readiness: 8, Impact: "", Category: "Unsupported modules", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "windows-principal",}, { Value: "windows-domain",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "System.Security.Principal", Advice: "Operations requiring a Windows domain are not supported", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-oauth2", RuleType: "", Level: "", FileType: "(java$|xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of OAuth2", Effort: 0, Readiness: 9, Impact: "", Category: "authentication", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "authentication",}, { Value: "oauth",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "import.*org.springframework.security.oauth2*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "org.springframework.security.oauth2.jwt.Jwt.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "jwt-authentication", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-boot-starter-oauth2.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             { Type: "", Pattern: "", Value: "spring-security-oauth2-.*", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "java-saml", RuleType: "", Level: "", FileType: "(xml$|gradle$)", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Usage of SAML", Effort: 0, Readiness: 9, Impact: "", Category: "authentication", Criticality: "",Container: 0, Cloud: 0,
            Tags:
            []Tag{  { Value: "authentication",}, { Value: "saml",}, { Value: "security",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, { Value: "compass-discovery",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "opensaml", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "xml-clientCert", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Avoid reliance on SSL", Effort: 5, Readiness: 10, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
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
        
            { Name: "xml-ejb-resource-mgr-aut", RuleType: "standard", Level: "", FileType: "xml$", Target: "line", Type: "regex", DefaultPattern: "(%s)", Advice: "Security implementation might need to be refactored", Effort: 10, Readiness: 10, Impact: "", Category: "security", Criticality: "medium",Container: 100, Cloud: 100,
            Tags:
            []Tag{  { Value: "security",}, { Value: "web-app",}, { Value: "modernize",}, { Value: "container",}, },
            Profiles:
            []Profile{  { Value: "cloud-suitability",}, },
            Excludepatterns:
            []ExcludePattern{  },
            Recipes:
            []Recipe{  },
            Patterns:
            []Pattern{  { Type: "", Pattern: "", Value: "<res-auth>Container</res-auth>", Advice: "", Effort: 0, Readiness: 0, Criticality: "", Category: "", Tag: "", Recipe: "", },
             }, },
        
            { Name: "yaml-test", RuleType: "standard", Level: "", FileType: "(yaml|yml)$", Target: "file", Type: "yamlpath", DefaultPattern: "", Advice: "", Effort: 1000, Readiness: 10, Impact: "", Category: "", Criticality: "medium",Container: 100, Cloud: 100,
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
