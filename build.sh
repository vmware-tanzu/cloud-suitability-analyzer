export GOSUMDB=off

#--- set module mode
export GO111MODULE=on

export WORKING_DIR=${PWD}

echo "~~~> Compile/Minify UI"
pushd ${PWD}/csa-app/frontend

export NODE_OPTIONS="--max_old_space_size=4096"

npm ci -s --no-optional --force

npm run -s production-build

if [ $? -eq 0 ]
then
echo "~~~> production-build succeeded"
else
echo "~~~> production-build failed " >&2
exit 1
fi

  bindDataFile=resources/web-site.go

  if [ -f "$bindDataFile" ]; then
    echo "~~~> web-site.go found, removing"
    rm $bindDataFile
  fi

  echo "~~~> Binding in web assets"
  mkdir -p resources
  go-bindata -o resources/web-site.go -pkg resources build/...

popd

pushd ${PWD}/csa-app
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

    export GOSUMDB=off
    export G0111MODULE=onS
    export OUTPUT_DIR="exe"
    export LD_FLAGS="-X \"main.Version=$VERSION\"" 

    env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags "${LD_FLAGS}" -o ${OUTPUT_DIR}/csa csa.go
popd