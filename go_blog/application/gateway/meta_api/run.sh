#!/usr/bin/env bash

export ENV=local
pwd=$(pwd)

cd ./output || exit

exec ./manager-api
