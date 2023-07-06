#!/bin/bash



start_time="$(date -u +%s)"


#---   this corrects and issue with go mod tidy
export GOSUMDB=off

#--- set module mode
export GO111MODULE=on

export WORKING_DIR=${PWD}

echo "~~~> Compile/Minify UI"
pushd ${WORKING_DIR}/csa-app/frontend > /dev/null

  echo "~~~> Running npm ci"
  export NODE_OPTIONS="--max_old_space_size=4096"

  #npm ci -s --no-optional #> /dev/null 2>&1
  npm ci -s --no-optional --force --legacy-peer-deps

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

  bindDataFile=${WORKING_DIR}/csa-app/frontend/resources/web-site.go

  if [ -f "$bindDataFile" ]; then
    echo "~~~> web-site.go found, removing"
    rm $bindDataFile
  fi

  echo "~~~> Binding in web assets"
  mkdir -p ${WORKING_DIR}/csa-app/frontend/resources
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

#--- removing condidation go generate
runGenerate=1
if [ $runGenerate -eq 1 ]
then

   pushd ${WORKING_DIR}/csa-app > /dev/null

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
popd > /dev/null
fi


