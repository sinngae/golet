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
	data := []byte(`{"jwt":"eyJhY2NvdW50Ijoic2VsbGVyLmZ1bGxmaWxsbWVudCIsImFsZyI6IkhTMjU2IiwidHlwIjoiSldUIn0.eyJkYXRhIjp7InNsc190bl9saXN0IjpbeyJjb3VudHJ5IjoiQ08iLCJzbHNfdG4iOiJDTzIyMjMwNzA1OTk1NTEwS0gifV0sInRpbWVzdGFtcCI6MTY0NTc1NzQxMn0sInRpbWVzdGFtcCI6MTY0NTc1NzQxMn0.wfBl7_7p0kOBG9YhvgWBq4WU8Bu6yV3MW_jCkdgGyCo","sls_tn_list":[{"country":"CO","sls_tn":"CO22230705995510KH"}],"timestamp":1645757412}`)
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
