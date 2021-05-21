#########################################
# File Name: protoc_gen.sh
# Author: Philo
# mail: 
# Created Time: 三 11/25 20:23:40 2020
########################################
#!/bin/bash
usage="Usage: $0 proto-filename"

protofile=$1

protoc \
-I=${GOPATH}/src \
-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
-I=./protocol/proto/ \
--go_out=protocol/go/  \
--govalidators_out=./protocol/go/ \
./protocol/proto/${protofile}.proto

protoc \
-I=${GOPATH}/src \
-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
-I=./protocol/proto/ \
--go-grpc_out=protocol/go/  \
--govalidators_out=./protocol/go/ \
./protocol/proto/${protofile}.proto
