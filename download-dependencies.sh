#!/bin/bash

MAC_CLIS=("nodejs" "npm" "go" "go-bindata" "musl-cross|filosottile/musl-cross/musl-cross" "mingw-w64" "goreleaser")
LINUX_CLIS=("nodejs" "npm" "go" "go-bindata" "musl-tools" "gcc-mingw-w64" "build-essential")
OS=$(uname)

install_mac_clis() {
    for cli in "${MAC_CLIS[@]}"; do
        package_name=$(echo ${cli} | awk '{split($0,a,"|"); print a[1]}')
        tap_name=$(echo ${cli} | awk '{split($0,a,"|"); print a[2]}')
        
        if [[ -z $(brew list | grep ${package_name}) ]]; then
            if [[ ! -z ${tap_name} ]]; then
                brew install ${tap_name}
            else
                brew install ${cli}
            fi
        fi
    done
}

install_linux_clis() {
    for cli in "${LINUX_CLIS[@]}"; do
        if [[ -z $(apt list | grep $cli) ]]; then
            apt-get install -y --no-cache $cli
        fi
    done
}

if [[ "$OS" == "Linux" ]]; then
    echo "installing required tools for $OS"
    install_linux_clis
elif [[ "$OS" == "Darwin" ]]; then
    echo "installing required tools for $OS"
    install_mac_clis
fi

