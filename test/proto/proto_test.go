package proto

import (
	"encoding/json"
	"testing"

	hiwork2 "github.com/sinngae/golet/test/proto/go/hiwork"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestAnyData1(t *testing.T) {
	//abc := "abc"
	//anyData, _ := anypb.New(&abc)
	myany := &hiwork2.AnyData{
		Omit: proto.String("test"),
		Params: &hiwork2.AnyData_Param{
			Param: []*anypb.Any{
				{TypeUrl: "abc"},
				{Value: []byte("test")},
			},
		},
	}
	str, err := json.Marshal(myany)
	assert.NoError(t, err)
	assert.NotNil(t, str)
}

func TestAnyData2(t *testing.T) {
	//abc := "abc"
	//anyData, _ := anypb.New(&abc)
	myany := &hiwork2.AnyData2{
		Params: &hiwork2.AnyData2_Param{
			Param: []string{
				"abc", "123",
			},
		},
	}
	str, err := json.Marshal(myany)
	assert.NoError(t, err)
	assert.NotNil(t, str)
}
