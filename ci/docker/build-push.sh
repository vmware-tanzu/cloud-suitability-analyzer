#!/bin/bash
NAME=inokes/node-go-cross-build
TAG="$(git log -1 --pretty=%h)"
IMG=${NAME}:${TAG}
LATEST="${NAME}:latest"
cmd1="docker build -t ${IMG} ."
cmd2="docker tag ${IMG} ${LATEST}"
cmd3="docker push ${NAME}"
eval $cmd1 && eval $cmd2 && eval $cmd3