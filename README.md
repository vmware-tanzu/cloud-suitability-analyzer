# Cloud Suitability Analyzer (Open Source)


# Latest Additions

<b>New Features:</b>

- [Exclude Patterns](https://github.com/vmware-tanzu/cloud-suitability-analyzer/blob/master/doc/CSA-UserManual.md#rules) - To help with fighting false positives
- [Profiles](https://github.com/vmware-tanzu/cloud-suitability-analyzer/blob/master/doc/CSA-UserManual.md#rules) - To help select what type of rules to be used during scanning
- [Rule Testing Framework](https://github.com/vmware-tanzu/cloud-suitability-analyzer/blob/master/doc/CSA-UserManual.md#rule-testing-framework) - To help with testing CSA rules against real code samples
- [Generate HTML and CSV Finding Reports](https://github.com/vmware-tanzu/cloud-suitability-analyzer/blob/master/doc/CSA-UserManual.md#generate-html-and-csv-reports) - Generate HTML and CSV Reports that can be leveraged into a CI pipeline

# Backlog/Feature requests

We are becoming more formal in our tracking of backlog with increased usage driving feature requests. You'll find our official backlog detailed in our [Roadmap](https://github.com/vmware-tanzu/cloud-suitability-analyzer/projects/1) 

To summarize, here's what is in progress:

1. Rules that look for the absence of a pattern
2. Allowing multiple instances of file patterns



# Quick Start

`csa` has pre-built binaries for Windows, OSX, and Linux. The quickest way to get starting using `csa` is to download the [latest release](https://github.com/vmware-tanzu/cloud-suitability-analyzer/releases/latest). Included with the release is the latest [user manual](doc/CSA-UserManual.md) in pdf format. A sample portfolio of applications is also included.

No install is required. Simply download the version that matches your operating system and move the executable to the directory of your choice. It's best to put the directory in your path to simplify running `csa`. Everything you need to use `csa` is inside this single executable.

To see all the options to run, simply type:

```bash
csa -h
```

For more detailed help instructions:

```bash
csa --help-long
```

To simplify usage, `csa` has a default mode of `analyze`, since that is it most frequently used. Typically, you point `csa` to a directory that contains sub-directories, each containing one application. `csa` considers this directory of sub-directories a portfilio. To tell `csa` to process the directory as a portfolio of applications, we use the`-p` flag on the command-line

```bash
csa -p <directory of application code>
```

`csa` will process all the subdirectories as applications and write the results to the console. Look in the directory where `csa` was copied to and you'll see a file, `csa.db`. This is a SQLite database full of it findings.

To see the results of the analysis, you'll need to call `csa` one more time and launch the browser-based user interface:

```bash
csa ui
```

`csa` will take you directly to your browser and activate the user interface.

# Building CSA 

If you want to contribute to `csa`, you'll need to build it.

### Prerequisites

* Golang - Installation instructions are [here](https://golang.org/doc/install)
* Docker - Instructions for [Get Docker](https://docs.docker.com/get-docker/)
* go-bindata - Install go-bindata [using go get](https://github.com/jteeuwen/go-bindata) 

  Alternatives if go get fails
  
  OSX  ```$> brew install go-bindata``` 
  
  Linux ```sudo apt install go-bindata```

### Clone the repo

Clone the repo into your the directory of your choice.


### Build & Run

## Generating all binaries from your MACOS

```
> ./build.sh
```

Executable will be generated here:
```
=> csa-app/dist/csa
=> csa-app/dist/csa-l
=> csa-app/dist/csa.exe
```

## Generating LINUX & WINDOWS binaries from your LINUX (DEBIAN) OS

### 1. Docker Option

Release builds containing all required GO dependencies can be generated using a docker build

1. Run Docker Build

```
$> docker build -f build-Dockerfile -t csa-release:latest .
```

2. Generate executables

```
$> docker run -v ${PWD}:/app -e VERSION=v3.2.10-rc1 csa-release:latest
```


### 2. Using goreleaser
```
> ./build.sh
```

```
Executable will be generated here 
=> csa-app/dist/csa-l
=> csa-app/dist/csa.exe
```

**Verify docker created directory has correct ownership**
```
$> sudo chown -R $USER:`id -g -n $USER` $WORKING_DIR/go/exe
```

**Check that the exe runs**

The executable(s) can be found in  `<project root dir>/csa-app/exe` directory

# Create release builds on Mac

To create the release build do the following:
* Tag release 

  ```
  git tag <tag-name>
  ```

* Execute 
  
  ```
  ./build.sh -r
  ```

* Push the tag out
  
  ```
  git push --tags
  ```

# Documentation

## Contributing

The cloud-suitability-analyzer project team welcomes contributions from the community. Before you start working with cloud-suitability-analyzer, please
read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be
signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on
as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Cloud Suitability Analyzer is released under the BSD-2 license. Please see [LICENSE.txt](https://github.com/vmware-samples/cloud-suitability-analyzer/blob/master/LICENSE.txt)
