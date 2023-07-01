#!/bin/sh

set -ex

docker build -t fuxu/gohello:build . -f Dockerfile.build
docker create --name extract fuxu/gohello:build
docker cp extract:/hello ./hello
docker rm -f extract
docker build -t fuxu/gohello .
