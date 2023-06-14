#!/bin/sh
rm -r -f csa-app/test-exe
mkdir csa-app/test-exe
mkdir csa-app/test-exe/rules
cp -r -f rules/ csa-app/test-exe/rules
cp -r -f csa-app/tests/test-cases/ csa-app/test-exe/test-cases
cp -r -f csa-app/tests/test-samples/ csa-app/test-exe/test-samples

cd csa-app/tests
go test -c -o unit-test
GOOS=windows CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOARCH=amd64 go test -c -o unit-test-w.exe
GOOS=linux GOARCH=amd64 go test -c -o unit-test-l
cd ../../
mv csa-app/tests/unit-test csa-app/test-exe/unit-test
mv csa-app/tests/unit-test-l csa-app/test-exe/unit-test-l
mv csa-app/tests/unit-test-w.exe csa-app/test-exe/unit-test-w.exe
mv csa-app/tests/ReadMe-public.md csa-app/test-exe/ReadMe.md

cd csa-app/test-exe/
zip -r rule-test.zip *
cd ../../

rm -r -f csa-app/test-exe/rules
rm -r -f csa-app/test-exe/test-cases
rm -r -f csa-app/test-exe/test-samples
rm -r -f csa-app/test-exe/unit*


