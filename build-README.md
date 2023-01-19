# Build Executables for each platform

## MACOS

./build-Local.sh -O

Executable will be generated here => csa-app/exe/csa

## LINUX & WINDOWS

Release builds containing all required GO dependencies can be generated using a docker build

docker build -f build-Dockerfile -t csa-build:0.0.1 .
