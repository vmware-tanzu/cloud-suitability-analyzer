#!/bin/bash

start_time="$(date -u +%s)"


#---   this corrects and issue with go mod tidy
export GOSUMDB=off

#--- set module mode
export GO111MODULE=on

export WORKING_DIR=$GOPATH/src/github.com/vmware-samples/cloud-suitability-analyzer


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
  echo "**** Rules changed ****"
  runGenerate=1
fi

if [ $binsChanged -eq 1 ]
then
  echo "**** Bins changed ****"
  runGenerate=1
fi

if [ $modelsChanged -eq 1 ]
then
  echo "**** Scoring models changed ****"
  runGenerate=1
fi
BootstrapFile="go/model/Bootstrap.go"
BinBootstrapFile="go/model/BinBootstrap.go"
ScoringModelBootstrapFile="go/model/ScoringModelBootstrap.go"


if [ ! -f "$BootstrapFile" ]; then
    echo "**** $BootstrapFile is missing ****"
    cp go/model-cache/Bootstrap.go go/model/Bootstrap.go
    runGenerate=1
fi

if [ ! -f "$BinBootstrapFile" ]; then
    echo "**** $BinBootstrapFile is missing ****"
    cp go/model-cache/BinBootstrap.go go/model/BinBootstrap.go
    runGenerate=1
fi

if [ ! -f "$ScoringModelBootstrapFile" ]; then
    echo "**** $ScoringModelBootstrapFile is missing ****"
    cp go/model-cache/ScoringModelBootstrap.go go/model/ScoringModelBootstrap.go
    runGenerate=1
fi


if [ $runGenerate -eq 1 ]
then

   pushd ${WORKING_DIR}/go

   echo "~~~> Binding rules and bins"
   go generate

  if [ $? -eq 0 ]
  then
    echo "go generate succeeded"
  else
    echo "go generate failed " >&2
    exit 1
  fi

   echo "~~~> Updating dependencies" 
   go mod tidy
  if [ $? -eq 0 ]
  then
    echo "go mod tidy succeeded"
  else
    echo "go mod tidy failed " >&2
    exit 1
  fi
popd
   popd
fi

end_time="$(date -u +%s)"

elapsed="$(($end_time-$start_time))"
echo "Native  execution time: $elapsed seconds"

echo "~~~> Mounting docker image"
time docker run -it -e OS=$1 -e VERSION=$2 -e NATIVETIME=$elapsed -v ${GOPATH}/src/github.com/vmware-samples/cloud-suitability-analyzer/:/go/src/github.com/vmware-samples/cloud-suitability-analyzer neur0manc3r/node-go-cross-build /go/src/github.com/vmware-samples/cloud-suitability-analyzer/buildlocalFull.sh 

