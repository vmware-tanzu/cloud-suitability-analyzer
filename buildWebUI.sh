#!/bin/bash

start_time="$(date -u +%s)"


#---   this corrects and issue with go mod tidy
export GOSUMDB=off

#--- set module mode
export GO111MODULE=on

export WORKING_DIR=$GOPATH/src/github.com/vmware-samples/cloud-suitability-analyzer

echo "~~~> Compile/Minify UI"
pushd ${WORKING_DIR}/go/frontend > /dev/null
  
  echo "~~~> Running npm ci"

  
  npm ci -s --no-optional > /dev/null 2>&1

  if [ $? -eq 0 ]
  then
    echo "~~~> npm ci succeeded"
  else
    echo "~~~> npm ci failed" >&2
    exit 1
  fi
  

  echo "~~~> Running npm production-build"
  npm run -s production-build

  if [ $? -eq 0 ]
  then
    echo "~~~> production-build succeeded"
  else
    echo "~~~> production-build failed " >&2
    exit 1
  fi

  bindDataFile=${WORKING_DIR}/go/frontend/resources/web-site.go

  if [ -f "$bindDataFile" ]; then
    echo "~~~> web-site.go found, removing"
    rm $bindDataFile
  fi

  echo "~~~> Binding in web assets"
  mkdir -p ${WORKING_DIR}/go/frontend/resources
  go-bindata -o resources/web-site.go -pkg resources build/...
    
  if [ $? -eq 0 ]
  then
    echo "~~~> go-bindata succeeded!"
  else
    echo "~~~> go-bindata failed " >&2
    exit 1
  fi

popd > /dev/null

cd ${WORKING_DIR}

runGenerate=0

echo "~~~> Checking for rule changes"
git diff --exit-code rules
rulesChanged=$?

echo "~~~> Checking for bin changes"
git diff --exit-code bins
binsChanged=$?

echo "~~~> Checking for scoring model changes"
git diff --exit-code score-models
modelsChanged=$?


if [ $rulesChanged -eq 1 ]
then
  echo "~~~> Rules changed"
  runGenerate=1
fi

if [ $binsChanged -eq 1 ]
then
  echo "~~~> Bins changed"
  runGenerate=1
fi

if [ $modelsChanged -eq 1 ]
then
  echo "~~~> Scoring models changed"
  runGenerate=1
fi

#--- these go files have to exist to make the compiler happy, they may be overwritten by go generate

BootstrapFile="go/model/Bootstrap.go"
BinBootstrapFile="go/model/BinBootstrap.go"
ScoringModelBootstrapFile="go/model/ScoringModelBootstrap.go"
if [ ! -f "$BootstrapFile" ]; then
    echo "~~~> $BootstrapFile is missing"
    cp go/model-cache/Bootstrap.go go/model/Bootstrap.go
    runGenerate=1
fi

if [ ! -f "$BinBootstrapFile" ]; then
    echo "~~~> $BinBootstrapFile is missing"
    cp go/model-cache/BinBootstrap.go go/model/BinBootstrap.go
    runGenerate=1
fi

if [ ! -f "$ScoringModelBootstrapFile" ]; then
    echo "~~~> $ScoringModelBootstrapFile is missing"
    cp go/model-cache/ScoringModelBootstrap.go go/model/ScoringModelBootstrap.go
    runGenerate=1
fi


if [ $runGenerate -eq 1 ]
then

   pushd ${WORKING_DIR}/go > /dev/null

   echo "~~~> Binding rules and bins"
   go generate

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
popd > /dev/null
fi


