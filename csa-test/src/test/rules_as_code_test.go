package main

import (
	"fmt"
	"test/csa-app/csa"
	"test/csa-app/model"
	"test/csa-app/util"
	matt "test/mat"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Rules
var rulesCloudBlockerDir = "/Users/scarbonell/Workspace/bankofamerica/csa-test/rules/cloud_blockers"
var cloudBlockerRuleSet = csa.Setup(rulesCloudBlockerDir)
var rulesCorePortabilityDir = "/Users/scarbonell/Workspace/bankofamerica/csa-test/rules/core_portability"
var corePortabilityRuleSet = csa.Setup(rulesCorePortabilityDir)
var rulesCloudSuitabilityDir = "/Users/scarbonell/Workspace/bankofamerica/csa-test/rules/cloud_suitability"
var cloudSuitabilityRuleSet = csa.Setup(rulesCloudSuitabilityDir)

var ruleCount = len(cloudBlockerRuleSet) + len(corePortabilityRuleSet) + len(cloudSuitabilityRuleSet)
var rulesCovered = make(map[string]struct{})
var rulesTested [100]string
var testCount = 0

// Code Samples
var sampleBaseDir = "/Users/scarbonell/Workspace/bankofamerica/csa-test/src/test/test_samples"

/*************************************
// Test Rule is Valid / Compiles
// Test Rule targets the right file types
// Test Match
// Test Finding text
*************************************/

func TestCloudBlockerSuite(t *testing.T) {

	// File IO
	testRuleByName(t, "File IO Create/Update/Delete", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-fileIO"), "file_io.cs", true, 9, "null")
	// Session in Memory
	testRuleByName(t, "Session in Memory", csa.RuleByName(t, cloudBlockerRuleSet, "config-sessionState"), "server_session.config", true, 1, "null")
	// Windows Forms
	testRuleByName(t, "Windows Form", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-windowsForms"), "windows_form.cs", true, 2, "using System.Windows.Forms;public class Form : System.Windows.Forms.ContainerControl")
	// Domain
	testRuleByName(t, "Application Domain", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-windows-application-domain"), "test.cs", true, 1, "AppDomain newDomain = AppDomain.CreateDomain(\"newDomain\", evidence, setup);")
	// File Caching
	testRuleByName(t, "File Caching Config", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-FileCacheModule"), "file_caching_module.config", true, 1, "null")
	// Launch Process
	testRuleByName(t, "Launch Process", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-launchProcess"), "launch_process.cs", true, 3, " Process() Process () Process.")
	// Windows Registry
	testRuleByName(t, "Windows Registry", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-windowsRegistry"), "windows_registry.cs", true, 3, "null")
	// Dotnet Logging
	testRuleByName(t, "Dotnet Logging", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-logging"), "dotnet_logging.vb", true, 2, "null")
	testRuleByName(t, "Dotnet Logging Config", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-logging-config"), "dotnet_logging_config.config", true, 1, "null")
	// File Path
	testRuleByName(t, "File Path", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-filepath"), "file_path.cs", true, 1, "null")
	// Dotnet Transactions
	//testRuleByName(t, "Dotnet Transaction", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-transactions"), "dotnet_transactions.cs", true, 1, "null")
	// File Based Config
	testRuleByName(t, "File Based Config", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-file-based-config"), "dotnet-file-based-config.cs", true, 1, "null")
	// Security
	testRuleByName(t, "Security", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-security"), "dotnet-security.cs", true, 1, "null")
	// Serilog
	testRuleByName(t, "Serilog", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-serilog"), "dotnet-serilog.cs", true, 2, "null")
	// WPF
	testRuleByName(t, "Windows Presentation Foundation", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-windows-presentation-foundation"), "dotnet-windows-presentation-foundation.xaml", true, 1, "null")
	// Windows Principal
	testRuleByName(t, "Windows Principal", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-windowsPrincipal"), "dotnet-windowsPrincipal.cs", true, 1, "null")
	// Windows Services
	testRuleByName(t, "Windows Services", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-windowsServices"), "dotnet-windowsServices.cs", true, 3, "null")
	// Encryption Config
	testRuleByName(t, "Encryption Config", csa.RuleByName(t, cloudBlockerRuleSet, "config-encryption"), "config-encryption.config", true, 1, "null")
	// Connection String
	testRuleByName(t, "Connection String", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-connectionstrings"), "dotnet-connectionstrings.config", true, 1, "null")
	// Connection String
	testRuleByName(t, "Connection String", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-database-access"), "dotnet-database-access.config", true, 1, "null")
	// WCF WS Protocol
	//testRuleByName(t, "WCF WS", csa.RuleByName(t, cloudBlockerRuleSet, "dotnet-wcf-protocols-ws"), "dotnet-wcf-protocols-ws.config", true, 1, "null")
	// Javascript Anti Cloud Patterns
	testRuleByName(t, "JS FileIO", csa.RuleByName(t, cloudBlockerRuleSet, "js-fileIO"), "js-fileIO.js", true, 1, "null")
	testRuleByName(t, "JS Cache", csa.RuleByName(t, cloudBlockerRuleSet, "js-cache"), "js-cache.js", true, 3, "null")

}

func TestCorePortabilitySuite(t *testing.T) {
	// ISAPI Filters
	testRuleByName(t, "ISAPI Filters Config", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-ISAPI-Filters-config"), "dotnet-ISAPI-Filters-config.config", true, 1, "null")
	testRuleByName(t, "ISAPI Filters VBCS", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-ISAPI-Filters-vbcs"), "dotnet-ISAPI-Filters-vbcs.cs", true, 2, "null")

	// IIS Modules
	testRuleByName(t, "Anonymous Auth IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-AnonymousAuthentication"), "dotnet-iis_module-AnonymousAuthentication.config", true, 1, "null")
	testRuleByName(t, "Authentication IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-Authentication"), "dotnet-iis_module-Authentication.config", true, 1, "null")
	testRuleByName(t, "Authorization IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-Authorization"), "dotnet-iis_module-Authorization.config", true, 1, "null")
	testRuleByName(t, "Cert Mapping Auth IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-CertificateMappingAuthentication"), "dotnet-iis_module-CertificateMappingAuthentication.config", true, 1, "null")
	testRuleByName(t, "Document IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-DefaultDocument"), "dotnet-iis_module-DefaultDocument.config", true, 1, "null")
	testRuleByName(t, "Digest Authentication IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-DigestAuthentication"), "dotnet-iis_module-DigestAuthentication.config", true, 1, "null")
	testRuleByName(t, "Directory Listing IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-DirectoryListing"), "dotnet-iis_module-DirectoryListing.config", true, 1, "null")
	testRuleByName(t, "HTTP Errors IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-HttpErrors"), "dotnet-iis_module-HttpErrors.config", true, 1, "null")
	testRuleByName(t, "HTTP Logging IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-HttpLogging"), "dotnet-iis_module-HttpLogging.config", true, 1, "null")
	testRuleByName(t, "HTTP Redirection IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-HttpRedirection"), "dotnet-iis_module-HttpRedirection.config", true, 1, "null")
	testRuleByName(t, "Client Certificate Mapping IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-IisClientCertificateMappingAuthentication"), "dotnet-iis_module-IisClientCertificateMappingAuthentication.config", true, 1, "null")
	testRuleByName(t, "IP Security IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-IpSecurity"), "dotnet-iis_module-IpSecurity.config", true, 1, "null")
	testRuleByName(t, "ISAPI CGI Registration IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-IsapiCgiRestriction"), "dotnet-iis_module-IsapiCgiRestriction.config", true, 1, "null")
	testRuleByName(t, "Output Cache IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-OutputCache"), "dotnet-iis_module-OutputCache.config", true, 1, "null")
	testRuleByName(t, "Protocol Support IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-ProtocolSupport"), "dotnet-iis_module-ProtocolSupport.config", true, 1, "null")
	testRuleByName(t, "Server Side Include IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-ServerSideInclude"), "dotnet-iis_module-ServerSideInclude.config", true, 1, "null")
	testRuleByName(t, "Tracing IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-Tracing"), "dotnet-iis_module-Tracing.config", true, 2, "null")
	testRuleByName(t, "Validation IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-Validation"), "dotnet-iis_module-Validation.config", true, 1, "null")
	testRuleByName(t, "URI Cache IIS Module", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-iis_module-UriCacheModule"), "dotnet-iis_module-UriCacheModule.config", true, 1, "null")

	// WF
	testRuleByName(t, "Windows Workflow Foundation", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-windows-workflow-foundation"), "dotnet-windows-workflow-foundation.cs", true, 1, "null")
	// Enterprise Services
	testRuleByName(t, "Enterprise Services", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-windows-enterprise-services"), "dotnet-windows-enterprise-services.cs", true, 1, "null")
	//Dotnet IP Address
	testRuleByName(t, "Dotnet IP Address", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-ip-address"), "dotnet-ip-address.cs", true, 1, "null")
	// ASP Child Action
	testRuleByName(t, "ASP Child Action", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-child-action"), "dotnet-asp-child-action.cs", true, 1, "null")
	// ASP Classic 2.0
	testRuleByName(t, "ASP Classic 2.0", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-classic-2-0"), "dotnet-asp-classic-2-0.asp", true, 4, "null")
	// ASP Machine Key
	testRuleByName(t, "ASP Machine Key", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-machine-key"), "dotnet-asp-machin-key.config", true, 1, "null")
	// ASP Machine Key
	testRuleByName(t, "ASP Membership", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-membership"), "dotnet-asp-membership.config", true, 1, "null")
	// ASP Form Collection
	testRuleByName(t, "ASP Form Collection", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-mvc-form-collection"), "dotnet-asp-mvc-form-collection.cs", true, 1, "null")
	// ASP Model Update
	testRuleByName(t, "ASP Model Update", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-mvc-model-update"), "dotnet-asp-mvc-model-update.cs", true, 1, "null")
	// ASP MVC
	testRuleByName(t, "ASP MVC", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-mvc"), "dotnet-asp-mvc.cs", true, 1, "null")
	// ASP Session Context
	testRuleByName(t, "ASP Session Context", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-session-context"), "dotnet-asp-session-context.cs", true, 1, "null")
	// ASP Web Form
	testRuleByName(t, "ASP Web Form", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-asp-web-form"), "dotnet-asp-web-form.aspx", true, 2, "null")
	// Windows Remoting
	testRuleByName(t, "Windows Remoting", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-windows-remoting"), "dotnet-windows-remoting.config", true, 1, "null")
	// Windows Remoting
	testRuleByName(t, "Windows Remoting Code", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-windows-remoting-code"), "dotnet-windows-remoting-code.cs", true, 3, "null")
	// Delegate
	testRuleByName(t, "Begin Invoke", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-delegate-begininvoke"), "dotnet-delegate-begininvoke.cs", true, 2, "null")
	// WCF
	testRuleByName(t, "WCF Config", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-wcf-service-model-config"), "dotnet-wcf-service-model-config.config", true, 1, "null")
	// WCF
	testRuleByName(t, "WCF Service", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-wcf-service-model"), "dotnet-wcf-service-model.cs", true, 1, "null")
	// Razor Web
	testRuleByName(t, "VB Razor Web", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-vb-razor-view-engine"), "dotnet-vb-razor-view-engine.vbhtml", true, 1, "null")
	// Script Blocks
	testRuleByName(t, "Script Blocks", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-script-blocks"), "dotnet-script-blocks.xsl", true, 2, "null")
	// Reflection Assembly
	testRuleByName(t, "Reflection Assembly Load", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-reflection-assembly"), "dotnet-reflection-assembly.cs", true, 1, "null")
	// WF
	testRuleByName(t, "WF ActivityXaml Services", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-wf-integration"), "dotnet-wf-integration.cs", true, 1, "null")
	// WF
	testRuleByName(t, "WF Core Presentation", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-wf-presentation"), "dotnet-wf-presentation.xaml", true, 3, "null")
	// WF
	testRuleByName(t, "WF Core Presentation", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-wf-activity"), "dotnet-wf-activity.cs", true, 1, "null")
	// Manage Addin
	testRuleByName(t, "Managed Addin Framework", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-managed-addin"), "dotnet-managed-addin.cs", true, 1, "null")
	// WIF
	testRuleByName(t, "Windows Identity Foundation", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-identity-foundation"), "dotnet-identity-foundation.cs", true, 1, "null")
	// System Speech
	testRuleByName(t, "System Speech", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-system-speech"), "dotnet-system-speech.cs", true, 1, "null")
	// ASP Net
	testRuleByName(t, "Classic ASP Net", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-classic-asp-net"), "dotnet-classic-asp-net.cs", true, 5, "null")
	// Workflow Foundation
	testRuleByName(t, "Worflow Foundation", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-workflow"), "dotnet-workflow.cs", true, 7, "null")
	// Javascript Serializer
	testRuleByName(t, "Javascript Serializer", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-javascript-serializer"), "dotnet-javascript-serializer.cs", true, 1, "null")
	// Windows MSMQ
	testRuleByName(t, "Windows MSMQ", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-MSMQ-vbcs"), "windows_msmq.cs", true, 15, "null")
	// Serialization dotnet-serialization
	testRuleByName(t, "Serialization", csa.RuleByName(t, corePortabilityRuleSet, "dotnet-serialization"), "dotnet-serialization.cs", true, 2, "null")

}

func TestCloudSuitabilitySuite(t *testing.T) {
	// Plan text Credentials
	testRuleByName(t, "Plain Text Credentials", csa.RuleByName(t, cloudSuitabilityRuleSet, "plaintext-creds"), "plain_text_credentials.cs", true, 2, "null")
	testRuleByName(t, "Plain Text Credentials False positive", csa.RuleByName(t, cloudSuitabilityRuleSet, "plaintext-creds"), "plain_text_credentials_false_positive.cs", false, 0, "null")

	// Hard Coded Rules
	testRuleByName(t, "Hard Coded IP False Positive", csa.RuleByName(t, cloudSuitabilityRuleSet, "dotnet-ipv4-addresses"), "hard_code_ip_false_positive.cs", false, 0, "null")
	testRuleByName(t, "Hard Coded IP", csa.RuleByName(t, cloudSuitabilityRuleSet, "dotnet-ipv4-addresses"), "hard_code_ip.cs", true, 1, "conn.connect(\"http://10.180.142.31/\")")

	// Hard Coded URI
	testRuleByName(t, "Hard Coded URI False Positive", csa.RuleByName(t, cloudSuitabilityRuleSet, "hardcode-uri"), "hard_code_uri_false_positive.cs", false, 0, "null")
	testRuleByName(t, "Hard Coded URI", csa.RuleByName(t, cloudSuitabilityRuleSet, "hardcode-uri"), "hard_code_uri.cs", true, 1, "https://some-uri.com")
}

func TestCoverage(t *testing.T) {
	t.Log("\033[43m\033[30m  #Test:", testCount, "\033[0m", "\033[0m")
	t.Log("\033[43m\033[30m  Rules Covered with Test:", len(rulesCovered), "/", ruleCount, "\033[0m", "\033[0m")
}

func TestExportForMat(t *testing.T) {
	ruleList := append(cloudSuitabilityRuleSet, corePortabilityRuleSet...)
	ruleList = append(ruleList, cloudBlockerRuleSet...)
	matt.Export(ruleList)

	for _, r := range ruleList {
		if rulesCovered[r.Name] != struct{}{} {
			fmt.Println("// Untested Rule: ")
			fmt.Println(r.Name)
		}
	}

}

// Easily test one rule against targeted file

func testRuleByName(t *testing.T, testName string, rule model.Rule, targetFilePath string, expectedMatch bool, expectedMatchCount int, expectedValue string) {
	var exists = struct{}{}
	rulesCovered[rule.Name] = exists
	testCount++
	t.Log("\033[94m[Test: " + testName + "] Rule: \033[33m" + rule.Name + "\033[0m || File: " + targetFilePath + "\033[0m")
	isValid, err := rule.IsValid()
	assert.Equal(t, err, nil, "The rule should not fail during validation")
	assert.Equal(t, isValid, true, "The rule should be valid")

	// Loading target file
	f := util.GetFile(sampleBaseDir + "/" + targetFilePath)

	// Does the rule apply ?
	applies := rule.Applies(f.GetCleanedExt(), f.Name)
	assert.Equal(t, applies, true, "\033[91mThe rule does not apply to this file type\033[0m")

	//Is there a match
	fileFindings, value := csa.AnalyzeFile(rule, f)

	assert.Equal(t, expectedMatch, (fileFindings > 0), "\033[91mThe match result was different than expected\033[0m")
	assert.Equal(t, expectedMatchCount, fileFindings, "\033[91mThe number of matches was different than expected\033[0m")

	if expectedValue != "null" && expectedValue != "" {
		assert.Equal(t, expectedValue, value, "\033[91mThe collected finding text did not match expectation\033[0m")
	}
}