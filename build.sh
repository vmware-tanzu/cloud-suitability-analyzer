#!/bin/bash

echo "Build started at $(date -u +%s)"

#---   this corrects and issue with go mod tidy
export GOSUMDB=off

#--- set module mode
export GO111MODULE=on

cleanup() {
  rm -rf ${PWD}/csa-app/dist/* ${PWD}/csa-app/frontend/build
}

compilePackageFrontEnd() {
  echo "~~~> Compile/Minify UI"
  pushd ${PWD}/csa-app/frontend

    echo "~~~> Running npm ci"
    export NODE_OPTIONS="--max_old_space_size=4096"

    #npm ci -s --no-optional #> /dev/null 2>&1
    npm ci -s --no-optional --force --legacy-peer-deps

#    if [ $? -eq 0 ]
#    then
#      echo "~~~> npm ci succeeded"
#    else
#      echo "~~~> npm ci failed" >&2
#      exit 1
#    fi


    echo "~~~> Running npm production-build"
    npm run -s production-build

    if [ $? -eq 0 ]
    then
      echo "~~~> production-build succeeded"
    else
      echo "~~~> production-build failed " >&2
      exit 1
    fi

    bindDataFile=${PWD}/csa-app/frontend/resources/web-site.go

    if [ -f "$bindDataFile" ]; then
      echo "~~~> web-site.go found, removing"
      rm $bindDataFile
    fi

    echo "~~~> Binding in web assets"
    mkdir -p ${PWD}/csa-app/frontend/resources
    go-bindata -o resources/web-site.go -pkg resources build/...

    if [ $? -eq 0 ]
    then
      echo "~~~> go-bindata succeeded!"
    else
      echo "~~~> go-bindata failed " >&2
      exit 1
    fi

  popd
}

stashFiles() {
  git stash
}

#--- Run go generate
runGoGenerate() {
  pushd ${PWD}/csa-app

    echo "~~~> Binding rules and bins"
    go generate #>&2

    if [ $? -eq 0 ]
    then
      echo "~~~> go generate succeeded"
    else
      echo "~~~> go generate failed " >&2
      exit 1
    fi

    echo "~~~> Updating dependencies"
    go mod tidy
    if [ $? -eq 0 ]
    then
      echo "~~~> go mod tidy succeeded"
    else
      echo "~~~> go mod tidy failed " >&2
      exit 1
    fi
  popd
}

generateExecutables() {
  pushd ${PWD}/csa-app
    OS=$(uname)
    if [[ "$OS" == "Linux" ]]; then
      echo "Building executables for linux and windows"
      goreleaser build --skip-validate --snapshot --id='linux' --id='windows' --clean
    elif [[ "$OS" == "Darwin" && "$RELEASE" == true ]]; then
      echo "Building executables for darwin, linux and windows"
        goreleaser build --clean
    elif [[ "$OS" == "Darwin" ]]; then
      echo "Building executables for darwin, linux and windows"
      goreleaser build --skip-validate --snapshot --clean
    fi
  popd
}

helpText() {
  echo "./build.sh - Generate binaries"
  echo "  -h|--help       Help!!"
  echo "  -r|--release    Specify this flag if you want to generate the release builds"
}

while [[ $# -gt 0 ]]; do
  case $1 in
    -r|--release)
      RELEASE=true
      shift
      shift
      ;;
    -h|--help)
      helpText
      exit 1
      ;;
    *)
      echo "Invalid option"
      helpText
      exit 1
      ;;
  esac  
done

./download-dependencies.sh
cleanup
compilePackageFrontEnd
runGoGenerate



if [[ ! -z "$RELEASE" ]]; then
  stashFiles
fi

generateExecutables

echo "Build ended at $(date -u +%s)"
