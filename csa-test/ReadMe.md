## CSA Rule Unit Testings ##

### Compile Unit Test Executable ###
- cd csa-test/ 
- sh build-unit-test-executable // It will generate the unit test executable for linux/mac/windows
- Artifacts are created in /test-dist/

### Run Unit Test Executable ###
cd csa-test/test-dist/
- WORK_DIR=[base dir] ./unit-test -test.v

### Sample Test ###

1. Choose the code sample to test against
- Place the file with the right extension under "/test-dist/test-samples" directory
2. Write a YAML test case for it in "/test-dist/test-cases"
```
  - name: "[Name of the Test]"
    rule-name: [unique-name-for-rule] //ex: plaintext-creds
    test-filename: [file-with-sample-of-code] //ex: plain_text_credentials.cs
    assert: [true or false] //Is a match expected? ex: true
    assert-count:  [number] //how many matches are expected? ex: 3
    assert-value: "[text captured by the rule]" //Text expected to be captured by the rule ex: password=$453556 or "null" to skip that check
```