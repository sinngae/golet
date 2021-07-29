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
		bin, err := json.MarshalIndent(&data, "", "\t")
		assert.NoError(t, err)
		fmt.Println(bin)
	})
}

func TestRawJsonUnmarshal(t *testing.T) {
	t.Run("test unmarshal raw json", func(t *testing.T) {
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

		var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)
		var colors []Color
		err := json.Unmarshal(j, &colors)
		assert.NoError(t, err)

		for _, c := range colors {
			var dst interface{}
			switch c.Space {
			case "RGB":
				dst = new(RGB)
			case "YCbCr":
				dst = new(YCbCr)
			}
			err := json.Unmarshal(c.Point, dst)
			assert.NoError(t, err)
			fmt.Println(c.Space, dst)
		}
	})
}
