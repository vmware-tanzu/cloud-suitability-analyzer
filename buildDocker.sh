#!/bin/bash
echo "~~~> Mounting docker image to build web ui"
docker run -it -e OS=$1 -e VERSION=$2 \
   -v ${GOPATH}/src/github.com/vmware-samples/cloud-suitability-analyzer/:/go/src/github.com/vmware-samples/cloud-suitability-analyzer \
   neur0manc3r/node-go-cross-build \
   /go/src/github.com/vmware-samples/cloud-suitability-analyzer/buildWebUI.sh 
echo "~~~> Mounting docker image to build go executable and bind in web"   
docker run -it -e OS=$1 -e VERSION=$2 \
   -v ${GOPATH}/src/github.com/vmware-samples/cloud-suitability-analyzer/:/go/src/github.com/vmware-samples/cloud-suitability-analyzer \
   neur0manc3r/node-go-cross-build \
   /go/src/github.com/vmware-samples/cloud-suitability-analyzer/build-CSA-Bind-UI.sh
