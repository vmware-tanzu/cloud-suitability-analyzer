#!/bin/bash
set -e

echo "~~~> Requested OS builds..: $1"
echo "~~~> Version..............: $2"
export OS=$1
export VERSION=$2

rm -rf csa-app/exe/* csa-app/frontend/build

#--- building web app
./buildWebUI.sh

#--- binding UI and other resources
./build-CSA-Bind-UI.sh
