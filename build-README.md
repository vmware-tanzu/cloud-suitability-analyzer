# Build Release Candidate Cross Platforms

## MACOS

./build-Local.sh -O

Executable will be generated here => csa-app/exe/csa

## LINUX & WINDOWS

Release builds containing all required GO dependencies can be generated using a docker build

1. Run Docker Built

docker build -f build-Dockerfile -t csa-release:latest .

2. Extract Generated Artifacts

```
cd ~/.../cloud-suitability-analyzer
docker run -v ${PWD}:/app -e VERSION=v3.2.10-rc1 csa-release:latest
```
