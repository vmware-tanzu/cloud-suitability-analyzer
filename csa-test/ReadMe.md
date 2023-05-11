## CSA Rule Unit Testings ##

### Run Tests ###
- cd into /src/test
- go test -v

### Compile Unit Test Executable ###
In csa-test/ root dir
- sh build.sh (current platform) or sh build-linux.sh for Linux Arch
- Artifacts are created in /dist

### Run Unit Test Executable ###
In dist/
- WORK_DIR=[base dir] ./unit-test-l -test.v


### Sample Test ###

- Sample code snippets should be placed in the "/test_sample" dir
- testRuleByName(t, "Title for the test", "csa.RuleByName(t, rules, "Name of the rule you are testing", "File or folder to test against", expected result => true or false, "Expected finding text")
Ex: testRuleByName(t, "Application Domain", csa.RuleByName(t, rules, "dotnet-windows-application-domain"), "test.cs", true, "AppDomain newDomain = AppDomain.CreateDomain(\"newDomain\", evidence, setup);")

### Run CSA against rule folder ###
- ./csa -p ./dotnetApps
- ./csa rules import --rules-dir=./DotNetRules/rules

### Run Lookbehind Translator ###
Change the values at the top of the script you want to translate
python3 lookbehind_translator.py