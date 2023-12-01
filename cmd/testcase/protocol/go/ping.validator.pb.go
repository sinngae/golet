// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ping.proto

package protocol

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PingBody) Validate() error {
	if this.Message != nil {
		if !(len(*(this.Message)) > 2) {
			return github_com_mwitkow_go_proto_validators.FieldError("Message", fmt.Errorf(`value '%v' must have a length greater than '2'`, *(this.Message)))
		}
		if !(len(*(this.Message)) < 10) {
			return github_com_mwitkow_go_proto_validators.FieldError("Message", fmt.Errorf(`value '%v' must have a length smaller than '10'`, *(this.Message)))
		}
	}
	return nil
}
