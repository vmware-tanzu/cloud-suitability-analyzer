## CSA Rule Unit Testings ##

This is a Unit Test Framework that can easily be used to curate and test a rule set.

It makes it easy to test a rule set and ensure that rules will fire when expected. 

Very often rules might never trigger or might trigger and pick up a false positive.

Now rules can be tested against specific source code samples. The framework checks the following:
- The rule syntax is correct and compiles
- The rule fires against the right set of file extensions
- The rule fires properly given a source code sample
- The rule fires the right amount of times given a source code sample
- The rule finding captured text is what should be expected given a source code sample

It is completely independant of CSA which does not have to be in the environment to test the rules. 

It is also platform independant and can run on Linux, Mac, Windows.

Made that way, it could easily be embedded into a CI pipeline for users who need to source controle and manage changes to their curated rule set.

### Compile Unit Test Executable ###
- cd csa-test/ 
- sh build-RuleTest.sh // It will generate the unit test executable for linux/mac/windows
- Artifacts are created in /test-dist/

### Run Unit Test Executable ###
- unzip rule-test.zip
- cd rule-test
- WORK_DIR=[base dir] ./rule-test -test.v

Mac: ./rule-test
Linux: ./rule-test-l
Windows: ./rule-test-w

### Sample Test ###

1. Choose the code sample to test against
- Place the file with the right extension under "/rule-test/test-samples" directory
2. Write a YAML test case for it in "/rule-test/test-cases"
```
  - name: "[Name of the Test]"
    rule-name: [unique-name-for-rule] //ex: plaintext-creds
    test-filename: [file-with-sample-of-code] //ex: plain_text_credentials.cs
    assert: [true or false] //Is a match expected? ex: true
    assert-count:  [number] //how many matches are expected? ex: 3
    assert-value: "[text captured by the rule]" //Text expected to be captured by the rule ex: password=$453556 or "null" to skip that check
```