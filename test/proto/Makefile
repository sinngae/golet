.PHONY: all release

PROTO_FILES = $(wildcard ./proto/*.proto)
#PB_GO_FILES = $(patsubst ./proto/%.proto, ./go/%/%.pb.go, $(PROTO_FILES))

all:target

pbf:
	#echo ${PROTO_FILES}
	#echo ${PB_GO_FILES}

target: ${PB_GO_FILES}

#$(PB_GO_FILES):./go/%/%.pb.go:./proto/%.proto
#	protoc -I=/usr/local/include/ -I=${GOPATH}/src/ -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf/ -I=./proto/ \
#	--go_out=./go/ --go-grpc_out=./go/ --govalidators_out=./go/ $^