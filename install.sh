#!/bin/bash

BIN_NAME="happy-app-server.fcgi"
INSTALL_DIR="${GOPATH}/bin"

if [ -z ${GOPATH} ]; then
    echo "Error: GOPATH not set"
    exit
fi

if [ ! -d "${INSTALL_DIR}" ]; then
    echo "Error: ${INSTALL_DIR} doesn't exist"
    exit
fi

go build -o "${BIN_NAME}"
mv "${BIN_NAME}" "${INSTALL_DIR}"
