#!/bin/sh
rm -r -f custom-rules
rm -r -f output/csa.db
mkdir custom-rules
cp -r -f ../rules/cloud_blockers custom-rules
cp -r -f ../rules/cloud_suitability custom-rules
cp -r -f ../rules/core_portability custom-rules
cp -r -f ../rules/package_portability custom-rules
cp -r -f ../rules/third_party_packages custom-rules
# Trigger creation of the DB required to import rules
./csa --database-dir output -p app-empty

./csa --database-dir output rules delete SNAP-java-package-Gradle
./csa --database-dir output rules delete dotnet-windowsRegistry
./csa --database-dir output rules delete java-fileIO
./csa --database-dir output rules delete js-function
./csa --database-dir output rules delete php-references-should-not-be-passed-to-func-call
./csa --database-dir output rules delete php-comments-should-not-be-at-eol
./csa --database-dir output rules delete dotnet-transactions
./csa --database-dir output rules delete java-activemq
./csa --database-dir output rules delete dotnet-iis_module-Validation
./csa --database-dir output rules delete SNAP-build-Gradle
./csa --database-dir output rules delete dotnet-asp-classic-2-0
./csa --database-dir output rules delete yaml-test
./csa --database-dir output rules delete xml-webLogic-1-5
./csa --database-dir output rules delete dotnet-database-access
./csa --database-dir output rules delete php-global-variable
./csa --database-dir output rules delete dotnet-logging
./csa --database-dir output rules delete js-fileIO
./csa --database-dir output rules delete java-struts
./csa --database-dir output rules delete java-faces-flow-Annotations
./csa --database-dir output rules delete php-file-system-manipulation
./csa --database-dir output rules delete dotnet-ISAPI-Filters-vbcs
./csa --database-dir output rules delete python-redis
./csa --database-dir output rules delete java-metrics
./csa --database-dir output rules delete dotnet-iis_module-DigestAuthentication
./csa --database-dir output rules delete java-apacheFop-import
./csa --database-dir output rules delete dotnet-serilog-elasticsearch
./csa --database-dir output rules delete java-nio
./csa --database-dir output rules delete java-servlet
./csa --database-dir output rules delete dotnet-wcf-bindings
./csa --database-dir output rules delete dotnet-windows-workflow-foundation
./csa --database-dir output rules delete java-mdb
./csa --database-dir output rules delete java-ejb-mdb
./csa --database-dir output rules delete xml-webLogic-1-9
./csa --database-dir output rules delete xml-ejb-3-0
./csa --database-dir output rules delete log2file-import
./csa --database-dir output rules delete websphere-cluster-jacl
./csa --database-dir output rules delete java-corba
./csa --database-dir output rules delete dotnet-windows-application-domain
./csa --database-dir output rules delete java-security
./csa --database-dir output rules delete java-mongo-cassandra
./csa --database-dir output rules delete donet-windows-remoting
./csa --database-dir output rules delete java-ehcache
./csa --database-dir output rules delete java-stateless-annotations
./csa --database-dir output rules delete dotnet-iis_module-CertificateMappingAuthentication
./csa --database-dir output rules delete java-nonstandard-protocol
./csa --database-dir output rules delete java-weblogic-import
./csa --database-dir output rules delete java-batchAnnotations
./csa --database-dir output rules delete java-jboss
./csa --database-dir output rules delete php-enable-session-use-trans-sid
./csa --database-dir output rules delete hardcode-uri
./csa --database-dir output rules delete xml-trinidad
./csa --database-dir output rules delete xml-webLogic-1-8
./csa --database-dir output rules delete java-batch
./csa --database-dir output rules delete xml-ejb-3-1
./csa --database-dir output rules delete dotnet-sharepoint
./csa --database-dir output rules delete xml-tomahawk
./csa --database-dir output rules delete xml-myfaces
./csa --database-dir output rules delete dotnet-windowsServices
./csa --database-dir output rules delete java-3rdPartyImports
./csa --database-dir output rules delete SNAP-SQL
./csa --database-dir output rules delete xml-portlet-1-0
./csa --database-dir output rules delete java-security-annotations
./csa --database-dir output rules delete dotnet-fileIO-read
./csa --database-dir output rules delete java-systemLoad
./csa --database-dir output rules delete js-document-write.yaml
./csa --database-dir output rules delete bootTXN
./csa --database-dir output rules delete java-jaxrs-import
./csa --database-dir output rules delete dotnet-iis_module-HttpErrors
./csa --database-dir output rules delete js-jwt-signed-verify-with-strong-cipher-algorithms
./csa --database-dir output rules delete xml-xa-dataSource
./csa --database-dir output rules delete java-tangosol
./csa --database-dir output rules delete java-swing
./csa --database-dir output rules delete java-ejb-rmi
./csa --database-dir output rules delete dotnet-StaticFileModule
./csa --database-dir output rules delete xml-webLogic-1-4
./csa --database-dir output rules delete xml-jsf-1-2
./csa --database-dir output rules delete dotnet-asp-mvc
./csa --database-dir output rules delete java-websockets-import
./csa --database-dir output rules delete java-persistence
./csa --database-dir output rules delete dotnet-windowsPrincipal
./csa --database-dir output rules delete docker-dockerFile
./csa --database-dir output rules delete xml-ejb-resource-mgr-aut
./csa --database-dir output rules delete java-logging-import
./csa --database-dir output rules delete dotnet-launchProcess
./csa --database-dir output rules delete faas-meta
./csa --database-dir output rules delete java-mqseries
./csa --database-dir output rules delete dotnet-windowsForms
./csa --database-dir output rules delete java-jersey-import
./csa --database-dir output rules delete dotnet-ISAPI-Filters-config
./csa --database-dir output rules delete dotnet-iis_module-DefaultDocument
./csa --database-dir output rules delete java-java-fx-import
./csa --database-dir output rules delete dotnet-iis_module-UriCacheModule
./csa --database-dir output rules delete dotnet-iis_module-Tracing
./csa --database-dir output rules delete xml-websphere
./csa --database-dir output rules delete java-alarmD-import
./csa --database-dir output rules delete bootEJB
./csa --database-dir output rules delete docker-sudo
./csa --database-dir output rules delete dotnet-asp-membership
./csa --database-dir output rules delete js-stdout
./csa --database-dir output rules delete dotnet-iis_module-IpSecurity
./csa --database-dir output rules delete dotnet-iis_module-HttpRedirection
./csa --database-dir output rules delete config-sessionState
./csa --database-dir output rules delete dotnet-db2-unmanaged
./csa --database-dir output rules delete js-var
./csa --database-dir output rules delete java-weblogic
./csa --database-dir output rules delete xml-webLogic-1-3
./csa --database-dir output rules delete xml-struts-tiles
./csa --database-dir output rules delete js-throw-literal
./csa --database-dir output rules delete SNAP-build-Ant-Maven
./csa --database-dir output rules delete java-jni
./csa --database-dir output rules delete dotnet-iis_module-Authorization
./csa --database-dir output rules delete xml-jsf-2-0
./csa --database-dir output rules delete bootJSF
./csa --database-dir output rules delete dotnet-windows-presentation-foundation
./csa --database-dir output rules delete xml-jsf-2-1
./csa --database-dir output rules delete java-mulesoft-intf
./csa --database-dir output rules delete java-hazelcast
./csa --database-dir output rules delete dotnet-asp-session-context
./csa --database-dir output rules delete xml-statefulEJB
./csa --database-dir output rules delete dotnet-iis_module-ProtocolSupport
./csa --database-dir output rules delete java-jsp
./csa --database-dir output rules delete xml-webLogic-1-2
./csa --database-dir output rules delete java-handles-term
./csa --database-dir output rules delete xml-messageDrivenBeans
./csa --database-dir output rules delete java-resource-cci
./csa --database-dir output rules delete wsdl-soap
./csa --database-dir output rules delete dotnet-asp-mvc-form-collection
./csa --database-dir output rules delete java-rpc-import
./csa --database-dir output rules delete php-socket-security-sensitive
./csa --database-dir output rules delete plaintext-creds
./csa --database-dir output rules delete dotnet-iis_module-AnonymousAuthentication
./csa --database-dir output rules delete java-servlet-session
./csa --database-dir output rules delete java-system-config
./csa --database-dir output rules delete java-struts-import
./csa --database-dir output rules delete dotnet-FileCacheModule
./csa --database-dir output rules delete dotnet-iis_module-IsapiCgiRestriction
./csa --database-dir output rules delete dotnet-ip-address
./csa --database-dir output rules delete java-jcaAnnotations
./csa --database-dir output rules delete php-goto-stmt-should-not-be-used
./csa --database-dir output rules delete SNAP-java-package-GradleJar
./csa --database-dir output rules delete php-deprecated-feature-parse_str
./csa --database-dir output rules delete java-restlet-import
./csa --database-dir output rules delete xml-jboss
./csa --database-dir output rules delete dotnet-iis_module-OutputCache
./csa --database-dir output rules delete python-cf
./csa --database-dir output rules delete java-file-system
./csa --database-dir output rules delete java-ejb-stateful
./csa --database-dir output rules delete SNAP-SQL-properties
./csa --database-dir output rules delete dotnet-ipv4-addresses
./csa --database-dir output rules delete js-cipher-algorithms-should-be-robust
./csa --database-dir output rules delete config-dotnet-webConfig
./csa --database-dir output rules delete java-jetty-import
./csa --database-dir output rules delete xml-facelets
./csa --database-dir output rules delete xml-session-scoped-beans
./csa --database-dir output rules delete bootJAXWS
./csa --database-dir output rules delete dotnet-MSMQ-vbcs
./csa --database-dir output rules delete dotnet-asp-web-form
./csa --database-dir output rules delete java-jta
./csa --database-dir output rules delete xml-jee
./csa --database-dir output rules delete dotnet-WindowsAuth-csvb
./csa --database-dir output rules delete php-function-method-naming-convention
./csa --database-dir output rules delete xml-struts-2-5
./csa --database-dir output rules delete dotnet-iis_module-DirectoryListing
./csa --database-dir output rules delete docker-non-root-user
./csa --database-dir output rules delete config-security
./csa --database-dir output rules delete java-faces-flow
./csa --database-dir output rules delete bootCDI
./csa --database-dir output rules delete java-transportSecurity
./csa --database-dir output rules delete xml-jsf-2-2
./csa --database-dir output rules delete java-remoteEJB
./csa --database-dir output rules delete java-MBeans
./csa --database-dir output rules delete java-jsf
./csa --database-dir output rules delete xml-webprofile
./csa --database-dir output rules delete python-rabbitmq
./csa --database-dir output rules delete php-removing-accent-marks
./csa --database-dir output rules delete dotnet-fileIO
./csa --database-dir output rules delete python-sqlite
./csa --database-dir output rules delete bootJDBC
./csa --database-dir output rules delete xml-servlet-2-3
./csa --database-dir output rules delete java-rabbitmq-import
./csa --database-dir output rules delete java-portUsage
./csa --database-dir output rules delete php-cache
./csa --database-dir output rules delete java-cache-dist-import
./csa --database-dir output rules delete xml-clientCert
./csa --database-dir output rules delete php-php-url-not-hardcoded
./csa --database-dir output rules delete xml-activeMQ
./csa --database-dir output rules delete php-cgi-force-redirect-enabled
./csa --database-dir output rules delete java-remoteWebService-import
./csa --database-dir output rules delete php-allow-url-in-config
./csa --database-dir output rules delete log4j-xml
./csa --database-dir output rules delete js-bool-comparison
./csa --database-dir output rules delete java-soap
./csa --database-dir output rules delete java-jndi
./csa --database-dir output rules delete SNAP-java-package-Maven-Ant
./csa --database-dir output rules delete php-reading-std-Input-security-sensitive
./csa --database-dir output rules delete xml-tomcat
./csa --database-dir output rules delete python-db-peewee
./csa --database-dir output rules delete xml-localJNDI
./csa --database-dir output rules delete dotnet-wcf-protocols
./csa --database-dir output rules delete python-fileIO
./csa --database-dir output rules delete xml-jsf-2-3
./csa --database-dir output rules delete dotnet-asp-mvc-model-update
./csa --database-dir output rules delete dotnet-wcf-protocols-ws
./csa --database-dir output rules delete xml-ejb-mdb-import
./csa --database-dir output rules delete dotnet-wcf-service-model
./csa --database-dir output rules delete dotnet-asp-machin-key
./csa --database-dir output rules delete xml-ejb-2-1
./csa --database-dir output rules delete php-disabled-file-uploads
./csa --database-dir output rules delete java-springboot-annotations
./csa --database-dir output rules delete xml-struts-2-4
./csa --database-dir output rules delete js-md5-sha1
./csa --database-dir output rules delete SNAP-java-ver-Maven-Ant
./csa --database-dir output rules delete dotnet-wcf-ssl
./csa --database-dir output rules delete dotnet-iis_module-IisClientCertificateMappingAuthentication
./csa --database-dir output rules delete java-iop
./csa --database-dir output rules delete java-ejb-stateful-import
./csa --database-dir output rules delete xml-ejb-remote
./csa --database-dir output rules delete xml-struts-1-1
./csa --database-dir output rules delete php-deprecated-feature-should-not-be-used
./csa --database-dir output rules delete config-encryption
./csa --database-dir output rules delete java-jpa
./csa --database-dir output rules delete java-ejb-stateless
./csa --database-dir output rules delete weblogic-cluster-config
./csa --database-dir output rules delete bootSTRUTS
./csa --database-dir output rules delete java-ws2liberty-methods
./csa --database-dir output rules delete bootMDB
./csa --database-dir output rules delete dotnet-windows-code-access-security
./csa --database-dir output rules delete dotnet-asp-child-action
./csa --database-dir output rules delete python-db-redis
./csa --database-dir output rules delete java-ejb-invocation
./csa --database-dir output rules delete dotnet-file-based-config
./csa --database-dir output rules delete dotnet-windows-components
./csa --database-dir output rules delete dotnet-asp-session-state
./csa --database-dir output rules delete xml-servlet-3-1
./csa --database-dir output rules delete xml-webLogic-1-7
./csa --database-dir output rules delete log4j-properties
./csa --database-dir output rules delete java-mulesoft-import
./csa --database-dir output rules delete java-stateful-annotations
./csa --database-dir output rules delete xml-jsf-1-1
./csa --database-dir output rules delete xml-servlet-2-5
./csa --database-dir output rules delete java-resource-spi
./csa --database-dir output rules delete dotnet-oracle-umanaged
./csa --database-dir output rules delete dotnet-security
./csa --database-dir output rules delete SNAP-ETL-import
./csa --database-dir output rules delete xml-transportSecurity
./csa --database-dir output rules delete dotnet-HttpCacheModule
./csa --database-dir output rules delete java-jcache
./csa --database-dir output rules delete java-messageDrivenBeans
./csa --database-dir output rules delete java-springframework
./csa --database-dir output rules delete dotnet-RequestFilteringModule
./csa --database-dir output rules delete dotnet-serilog
./csa --database-dir output rules delete xml-ehcache-Config
./csa --database-dir output rules delete java-hardIP
./csa --database-dir output rules delete xml-ejb-3-2
./csa --database-dir output rules delete xml-struts-2-3
./csa --database-dir output rules delete java-3rdPartySecurity-import
./csa --database-dir output rules delete dotnet-connectionstrings
./csa --database-dir output rules delete js-symbol
./csa --database-dir output rules delete java-netflix-healthcheck
./csa --database-dir output rules delete php-disabled-enable-dl
./csa --database-dir output rules delete java-jks
./csa --database-dir output rules delete java-logging-file-appender
./csa --database-dir output rules delete sqlserver-ssis
./csa --database-dir output rules delete dotnet-windows-enterprise-services
./csa --database-dir output rules delete java-jvm-runtimeConfigProps
./csa --database-dir output rules delete dotnet-filepath
./csa --database-dir output rules delete java-tibco-jms
./csa --database-dir output rules delete xml-struts-2-2
./csa --database-dir output rules delete dotnet-iis_module-HttpLogging
./csa --database-dir output rules delete bootWEBSOCKET
./csa --database-dir output rules delete php-md5-sha1-noncompliant
./csa --database-dir output rules delete java-transaction-annotations
./csa --database-dir output rules delete dotnet-iis_module-ServerSideInclude
./csa --database-dir output rules delete dotnet-WindowsAuth-config
./csa --database-dir output rules delete js-cache
./csa --database-dir output rules delete dotnet-iis_module-Authentication
./csa --database-dir output rules delete java-message-driven-annotations
./csa --database-dir output rules delete xml-weblogic
./csa --database-dir output rules delete xml-servlet-2-4
./csa --database-dir output rules delete java-slf4j-import
./csa --database-dir output rules delete xml-jsf-1-0
./csa --database-dir output rules delete java-threadUsage-import
./csa --database-dir output rules delete java-cache-import
./csa --database-dir output rules delete java-ws2liberty-import
./csa --database-dir output rules delete xml-servlet-3-0
./csa --database-dir output rules delete xml-weblogic-1-6
./csa --database-dir output rules delete js-md5-sha1-noncompliant
./csa --database-dir output rules delete java-glassfish-import
./csa --database-dir output rules delete java-jms
./csa --database-dir output rules delete java-processexit
./csa --database-dir output rules delete java-faces-flow-import
./csa --database-dir output rules delete dotnet-StaticCompressionModule
./csa --database-dir output rules delete js-continue
./csa --database-dir output rules delete dotnet-iis_module-TokenCacheModule

./csa --database-dir output rules import --rules-dir=custom-rules
./csa --database-dir output -p dotnetApps --cicd-dir="/Users/scarbonell/Workspace/boa-csa/cloud-suitability-analyzer/csa-test/exe/CICD-outputs" --cicd-file-name="findings.csv"
#./csa --database-dir output ui