#!/bin/sh
rm -r -f test-dist
mkdir test-dist
mkdir test-dist/rules
cp -r -f rules/ test-dist/rules
cp -r -f test-cases/ test-dist/test-cases
cp -r -f src/test/test-samples/ test-dist/test-samples

cd src/test/
go test -c -o unit-test
GOOS=windows CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOARCH=amd64 go test -c -o unit-test-w
GOOS=linux GOARCH=amd64 go test -c -o unit-test-l
cd ../../
mv src/test/unit-test test-dist/unit-test
mv src/test/unit-test-l test-dist/unit-test-l
mv src/test/unit-test-w test-dist/unit-test-w
chmod +x test-dist/unit-test

