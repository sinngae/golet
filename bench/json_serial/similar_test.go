package json_serial

import (
	"encoding/json"
	"fmt"
)

func ExampleLineIds() {
	type (
		LineIds0 struct {
			Fm string
			fl string
		}
		LineIds struct {
			Fm string
			Fl string
		}
		LineIds2 struct {
			Fm string
			fl string
			Fl string
			FL string
		}
	)
	lineIds0 := LineIds0{}
	lineIds := LineIds{}
	lineIds2 := LineIds2{}
	err := json.Unmarshal([]byte(`{"FL":"LVN33", "LM":"LVN35"}`), &lineIds0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(`{"FL":"LVN33", "LM":"LVN35"}`), &lineIds)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(`{"FL":"LVN33", "LM":"LVN35"}`), &lineIds2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", lineIds)
	//output:

}
