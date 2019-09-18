#!/usr/bin/env bash

if [[ -z ${GOPATH} ]]; then
    echo "Check: error: env: GOPATH not exists, aborted"
    exit 1
else
    echo "Check: GOPATH=${GOPATH}, OK"
fi

set -e
set -o pipefail

pushd  $(dirname $0) > /dev/null # GO TO PROJECT ROOT

# echo "Working dir: $(pwd -P)"

PLATFORM="$(uname -s | tr 'A-Z' 'a-z')"

export CGO_ENABLED=0
#GOOS=linux
export GOOS=${PLATFORM}


for dir in $(find ./ -type d  -maxdepth 1 -mindepth 1 -exec basename '{}' \; ); do
    if [[ ${dir} == "bin" ]]; then
        echo "skip bin ${dir}"
    elif [[ -f "${dir}/main.go" ]]; then
        go build -o "${GOPATH}/bin/${dir}" "${dir}/main.go"
        echo "${dir}/main.go => ${GOPATH}/bin/${dir}"
    fi
done

unset GOOS
unset CGO_ENABLED
unset PLATFORM

popd > /dev/null # EXIT FROM PROJECT ROOT
