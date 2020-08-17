$StartTime = $(get-date)
$Env:GOSUMDB="off"
#--- set module mode
$Env:GO111MODULE="on"
$Env:WORKING_DIR=$Env:GOPATH+'\src\github.com\vmware-samples\cloud-suitability-analyzer'

echo "~~~> Checking for rule changes"
git diff --exit-code rules
$rulesChanged=$LastExitCode

echo "~~~> Checking for bin changes"
git diff --exit-code bins
$binsChanged=$LastExitCode

echo "~~~> Checking for scoring model changes"
git diff --exit-code score-models
$modelsChanged=$LastExitCode

if($rulesChanged -or $binsChanged -or $modelsChanged) {
    $BootstrapFile = "go\model\Bootstrap.go"
    $BinBootstrapFile = "go\model\BinBootstrap.go"
    $ScoringModelBootstrapFile = "go\model\ScoringModelBootstrap.go"
    
    if(!(Test-Path $BootstrapFile)) {
        echo "**** $BootstrapFile is missing ****"
        Copy-Item go\model-cache\Bootstrap.go $BootstrapFile
    }
    
    if(!(Test-Path $BinBootstrapFile)) {
        echo "**** $BinBootstrapFile is missing ****"
        Copy-Item go\model-cache\BinBootstrap.go $BinBootstrapFile
    }
    
    if(!(Test-Path $ScoringModelBootstrapFile)) {
        echo "**** $ScoringModelBootstrapFile is missing ****"
        Copy-Item go\model-cache\ScoringModelBootstrap.go $ScoringModelBootstrapFile
    }

    echo "~~~> Binding rules and bins"
    pushd $Env:WORKING_DIR"\go"
    go generate
    go mod tidy
    popd
}

$elapsed = $(get-date) - $StartTime

docker run -it -e NATIVETIME=$elapsed -e OS=$args[0] -e VERSION=$args[1] -v $Env:GOPATH"/src/github.com/vmware-samples/cloud-suitability-analyzer/:/go/src/github.com/vmware-samples/cloud-suitability-analyzer" neur0manc3r/node-go-cross-build /bin/bash /go/src/github.com/vmware-samples/cloud-suitability-analyzer/buildlocalFull.sh 
