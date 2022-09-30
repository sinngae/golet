package json_serial

import (
	"encoding/json"
	"fmt"
)

func ExampleBytes() {
	data := struct {
		Data []byte
	}{
		[]byte("abcd"),
	}
	body, _ := json.Marshal(data)
	println(body)
	data2 := &struct {
		Data []byte
	}{}
	_ = json.Unmarshal(body, data2)
	str := fmt.Sprintf("%s", data2)
	println(str)
	// output:
	//{"Data":"YWJjZA=="}
	//&{abcd}
}

func ExampleJson() {
	var rawData = map[string]interface{}{}
	data := []byte(`{"jwt":"abc","list":[{"country":"CO","tn""}],"timestamp":1645757412}`)
	e := json.Unmarshal(data, &json.RawMessage{})
	if e == nil { // must be a json
		// reflect data
		if e := json.Unmarshal(data, &rawData); e != nil {
			return
		}
	}
	println("end")
	// output:

}
