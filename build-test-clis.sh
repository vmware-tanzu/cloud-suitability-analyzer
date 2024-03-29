#!/bin/bash

OUTPUT_DIR=$PWD/csa-app/test-exe

if [[ -d $OUTPUT_DIR ]]; then
    rm -rf $OUTPUT_DIR
fi

mkdir -p $OUTPUT_DIR

cp -r $PWD/rules $OUTPUT_DIR/

pushd $PWD/csa-app/tests
    cp -r $PWD/test-cases $OUTPUT_DIR/
    cp -r $PWD/test-samples $OUTPUT_DIR/

    go test -c -o $OUTPUT_DIR/rule-test
    GOOS=windows CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOARCH=amd64 go test -c -o $OUTPUT_DIR/rule-test-w.exe
    GOOS=linux GOARCH=amd64 go test -c -o $OUTPUT_DIR/rule-test-l
popd

cp $PWD/csa-app/tests/README-public.md $OUTPUT_DIR/README.md

pushd $OUTPUT_DIR
    zip -r test.zip *
    rm -rf rule* test-* README.md
popd