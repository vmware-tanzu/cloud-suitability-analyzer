#!/bin/sh
rm -r -f test-dist
mkdir test-dist
mkdir test-dist/rules
cp -r -f rules/ test-dist/rules
cp -r -f test-cases/ test-dist/test-cases
cp -r -f src/test/test-samples/ test-dist/test-samples

cd src/test/
go test -c -o unit-test
cd ../../
mv src/test/unit-test test-dist/unit-test
chmod +x test-dist/unit-test

