#!/usr/bin/env bash
set -e

VERSION=$(cat ./api/VERSION)
GITHASH=$(cat ./.githash 2> /dev/null || HASH="ref: HEAD"; while [[ $HASH == ref\:* ]]; do HASH="$(cat ".git/$(echo $HASH | cut -d \  -f 2)")"; done; echo ${HASH:0:7})

GOLDFLAGS="-X github.com/apisix/manager-api/internal/utils.version=${VERSION} -X github.com/apisix/manager-api/internal/utils.gitHash=${GITHASH}"

# Enter dry-run mode
if [ "$1" == "--dry-run" ]; then
    cd ./api && go run -ldflags "${GOLDFLAGS}" ./main.go
    exit 0
fi

set -x
export ENV=local
pwd=$(pwd)

rm -rf output && mkdir -p output/conf && mkdir -p output/dag-to-lua

# get dag-to-lua lib
if [[ ! -f "dag-to-lua-1.1/lib/dag-to-lua.lua" ]]; then
    wget https://github.com/api7/dag-to-lua/archive/v1.1.tar.gz -P /tmp
    tar -zxvf /tmp/v1.1.tar.gz -C /tmp
    cp -r /tmp/dag-to-lua-1.1/lib/* ./output/dag-to-lua
fi

# build
cd ./api && go build -o ../output/manager-api -ldflags "${GOLDFLAGS}" ./main.go && cd ..

cp ./api/conf/schema.json ./output/conf/schema.json
cp ./api/conf/customize_schema.json ./output/conf/customize_schema.json
cp ./api/conf/conf*.yaml ./output/conf/

echo "Build the Manager API successfully"
