#!/bin/bash

set -ex

mkdir -p ./tmp/protoc
cd ./tmp/protoc
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v24.0-rc1/protoc-24.0-rc-1-linux-x86_64.zip
unzip protoc-24.0-rc-1-linux-x86_64.zip
cd ../../
mkdir -p ./bin
mv ./tmp/protoc/bin/protoc ./bin
chmod 755 ./bin/protoc
rm -rf ./tmp/protoc