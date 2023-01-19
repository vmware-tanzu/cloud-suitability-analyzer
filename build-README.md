# Build Release Candidate Cross Platforms

## MACOS

./build-Local.sh -O

Executable will be generated here => csa-app/exe/csa

## LINUX & WINDOWS

Release builds containing all required GO dependencies can be generated using a docker build

1. Update Release Version in build-Dockerfile

ARG releaseVersion=v3.2.10-rc1

2. Run Docker Built

docker build -f build-Dockerfile -t csa-release:3.2.10-rc1 .

3. Extract Generated Artifacts

docker run -v /Users/[somewhere on your machine]:/dist csa-release:3.2.10-rc1
