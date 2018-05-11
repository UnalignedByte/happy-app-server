#!/bin/bash

set -e
./install.sh

BIN="${GOPATH}/bin/happy-app-server.fcgi"
if [ ! -f "${BIN}" ]; then
    echo "${BIN} doesn't exist"
    exit 1
fi

"${BIN}" $1
