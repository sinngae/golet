package json_serial

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawJson(t *testing.T) {
	t.Run("test raw json", func(t *testing.T) {
		header := json.RawMessage(`{"request_id":"abc"}`)
		data := struct {
			Header *json.RawMessage `json:"header"`
			Body   string           `json:"body"`
		}{
			Header: &header,
			Body:   "hi, data",
		}
		bin, err := json.MarshalIndent(&data, "", "t")
		assert.NoError(t, err)
		fmt.Println(bin)
	})
}

func TestRawJsonUnmarshal(t *testing.T) {

	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var tests = []string{
		`{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}}`,
		`{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}`,
	}

	for _, tc := range tests {
		t.Run("test unmarshal raw json", func(t *testing.T) {
			c := &Color{}
			err := json.Unmarshal([]byte(tc), c)
			assert.NoError(t, err)

			var dst interface{}
			switch c.Space {
			case "RGB":
				dst = new(RGB)
			case "YCbCr":
				dst = new(YCbCr)
			}
			err = json.Unmarshal(c.Point, dst)
			assert.NoError(t, err)
			fmt.Println(c.Space, dst)
		})
	}
}

func TestRawNull(t *testing.T) {
	type Data struct {
		Typ string          `json:"type"`
		Val json.RawMessage `json:"val"` // delay parsing until we know Typ
	}

	tests := []string{
		`{"type":"a", "val": null}`,
		`{"type":"b"}`,
		`{"type":"c", "val":1}`,
		`{"type":"c", "val":"adf"}`,
	}

	for _, tc := range tests {
		data := &Data{}
		err := json.Unmarshal([]byte(tc), data)
		assert.NoError(t, err)
	}

}

func TestNumber(t *testing.T) {
	data := `{"val": "2-100"}`
	dest := &struct {
		Val json.Number `json:"val"`
	}{}

	err := json.Unmarshal([]byte(data), dest)
	if err != nil {
		println(err)
		return
	}
	val, _ := dest.Val.Int64()
	//val2 := map[int]int{1:2}[1]
	println(val)
}
