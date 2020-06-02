#!/bin/bash
set -ex

TENSORFLOW_VERSION=v2.1.0

# for install tensorflow GO version.
wget https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.15.0.tar.gz -P /tmp/
tar -C /usr/local -xzf /tmp/libtensorflow-cpu-linux-x86_64-1.15.0.tar.gz
ldconfig

# already there, if you compile files on your computer, you will need these code.
# go get -u -v google.golang.org/grpc
# go get -u -v github.com/golang/protobuf/proto
# go get -u -v github.com/golang/protobuf/protoc-gen-go

go get -u -v golang.org/x/net/context
go get -u -v go.uber.org/ratelimit
go get -d -v github.com/tensorflow/tensorflow/tensorflow/go || true # ignore error.
# go get -u -v github.com/tensorflow/tensorflow/tensorflow/go # this is v2.2.0 is not stable

# switch branch to v2.1.0 because it can build protobuf files 
pushd $GOPATH/src/github.com/tensorflow/tensorflow
git checkout $TENSORFLOW_VERSION
popd

go get -v github.com/tensorflow/tensorflow/tensorflow/go