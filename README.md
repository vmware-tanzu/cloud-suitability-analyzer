# cloud-suitability-analyzer



## Overview
Note: The web user interface has been added to this OSS version.


## Try it out

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

**Run the build script**

Make sure your in the root directory of your project and run the build script
```
$> ./buildDocker.sh [executable type] or buildDocker.ps1
```

**Executable Type Options**
```
   O     creates a OSX executable
   L     creates a Linux executable
   W     creates a Windows executable
   OWL   builds all three
```

**Verify docker created directory has correct ownership**
```
$> sudo chown -R $USER:`id -g -n $USER` $WORKING_DIR/go/exe
```

**Check that the exe runs**

The executable(s) can be found in  `<project root dir>/go/exe` directory

## Documentation

## Contributing

The cloud-suitability-analyzer project team welcomes contributions from the community. Before you start working with cloud-suitability-analyzer, please
read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be
signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on
as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Cloud Suitability Analyzer is released under the BSD-2 license. Please see [LICENSE.txt](https://github.com/vmware-samples/cloud-suitability-analyzer/blob/master/LICENSE.txt)
