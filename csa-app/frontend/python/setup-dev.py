#!/usr/bin/env python3

#
#          FILE:  setup-dev.py
#
#
#   DESCRIPTION:  Setup web development environment for CSA UI
#
#       OPTIONS:  ---
#  REQUIREMENTS:  ---
#          BUGS:  ---
#       COMPANY:  VMware by Broadcom
#       VERSION:  1.0
#       CREATED:  08/26/2020
#      REVISION:  1.0
# ===============================================================================

import argparse
import sys
import shutil
import os
import tarfile


parser = argparse.ArgumentParser(description='Setup CSA dev environment')

parser.add_argument('-z', '--zipFilePath', metavar='zipFilePath', action='store', required=True, dest='zipFilePath')

parser.add_argument('-m', '--mode', metavar='mode', action='store',
                    choices={'create'}, dest='mode',
                    help='modes: %(choices)s')

args = parser.parse_args()

# --- setup directories
goPath = os.getenv('GOPATH')
homePath = os.getenv('HOME')

# --- get absolute path of release file that was passed in
absoluteZipPath = os.path.abspath(args.zipFilePath)

# --- check that Go is installed
if not goPath:
    print("GOPATH not set, have you installed Go?")
    sys.exit(1)

# --- setup path to clone frontend repo
devPath = f"{goPath}/src/github.com/vmware-tanzu/"
if not os.path.isdir(devPath):
    print(f"Creating dev path: {devPath}")
    os.makedirs(devPath)

# --- unpack release
head_tail = os.path.split(absoluteZipPath)
zipFileReleaseName = head_tail[1]


head_tail = os.path.split(absoluteZipPath)


# --- untar or unzip
if absoluteZipPath.endswith("tar.gz"):
    print(f'Untarring {absoluteZipPath} release to {devPath}')
    tar = tarfile.open(absoluteZipPath, "r:gz")
    tar.extractall(devPath)
    tar.close()
elif absoluteZipPath.endswith("zip"):
    shutil.unpack_archive(absoluteZipPath, devPath)

# --- rename directory to standard development path
oldPath = os.path.splitext(zipFileReleaseName)[0]
oldPath = devPath + os.path.splitext(oldPath)[0]
newPath = devPath + "cloud-suitability-analyzer"
os.rename(oldPath, newPath)


# --- clone frontend repo
os.chdir(newPath+'/go')
csaUiGitUrl = "git@github.com:vmwarepivotallabs/csa-ui.git"
cloneCommand = 'git clone ' + csaUiGitUrl + " frontend"
os.system(cloneCommand)

# --- move required go code to frontend to re-enable UI
print(' Copying go files to CSA CLI to enable the web interface')
goFileDirectory = newPath + '/go/frontend/gomodules/'
goDestDirectory = newPath + '/go/'
factorySource = goFileDirectory + 'factory.txt'
csaSource = goFileDirectory + 'csa.txt'
sharedSource = goFileDirectory + 'Shared.txt'
factoryDest = goDestDirectory + 'backend/routes/factory.go'
csaDest = goDestDirectory + 'csa.go'
sharedDest = goDestDirectory + 'util/Shared.go'
shutil.copyfile(factorySource, factoryDest)
shutil.copyfile(csaSource, csaDest)
shutil.copyfile(sharedSource, sharedDest)

print('Done!')




