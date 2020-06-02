#!/bin/bash
set -ex

export GOPATH=/gopath

# IN=$PWD
# OUT=/gen
# protoc -I=$IN/serving -I $IN/tensorflow \
#     --go_out=plugins=grpc:$OUT \
#     $IN/serving/tensorflow_serving/apis/*.proto

# IN=$PWD/tensorflow
# protoc -I=$IN --go_out=plugins=grpc:$OUT \
#     $IN/tensorflow/core/framework/*.proto
    
# protoc -I=$IN --go_out=plugins=grpc:$OUT \
#     $IN/tensorflow/core/protobuf/{saver,meta_graph}.proto
    
# protoc -I=$IN --go_out=plugins=grpc:$OUT \
#     $IN/tensorflow/core/example/*.proto

IN=$PWD
OUT=/gen
protoc -I=$IN/serving -I $IN/tensorflow \
    --go_out=plugins=grpc:$OUT \
    $IN/serving/tensorflow_serving/apis/*.proto

# tensorflow/tensorflow/core/framework/*.proto

IN=$PWD/tensorflow
protoc -I=$IN --go_out=plugins=grpc:$OUT \
    $IN/tensorflow/core/framework/*.proto

protoc -I=$IN --go_out=plugins=grpc:$OUT \
    $IN/tensorflow/core/example/*.proto
    
protoc -I=$IN --go_out=plugins=grpc:$OUT \
    $IN/tensorflow/core/protobuf/saver.proto

protoc -I=$IN --go_out=plugins=grpc:$OUT \
    $IN/tensorflow/core/protobuf/meta_graph.proto
