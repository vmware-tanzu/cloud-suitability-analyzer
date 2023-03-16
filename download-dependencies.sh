#!/bin/bash

MAC_CLIS=(nodejs npm go go-bindata musl mingw-w64 goreleaser)
LINUX_CLIS=(nodejs npm go go-bindata musl-tools gcc-mingw-w64 build-essential)
OS=$(uname)

install_mac_clis() {
    for cli in "${MAC_CLIS[@]}"; do
        if [[ -z $(brew list | grep $cli) ]]; then
            brew install $cli
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

