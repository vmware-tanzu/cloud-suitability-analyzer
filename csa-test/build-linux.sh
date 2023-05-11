#!/bin/sh
rm -r -f test-dist
mkdir test-dist
mkdir test-dist/rules
cp -r -f rules/ test-dist/rules
cp -r -f test-cases/ test-dist/test-cases
cp -r -f src/test/test-samples/ test-dist/test-samples

cd src/test/
GOOS=linux GOARCH=amd64 go test -c -o unit-test-l
cd ../../
mv src/test/unit-test-l test-dist/unit-test-l
chmod +x test-dist/unit-test-l

