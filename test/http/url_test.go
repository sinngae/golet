package http

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func ExampleUrl() {
	path := url.Values{}
	path.Set("abc", "==def")
	path.Set("ab", "def")
	str := path.Encode()
	fmt.Println(str)
	data := `"` + str + `"`
	body, err := json.MarshalIndent(struct {
		Val  int             `json:"val"`
		Data json.RawMessage `json:"data"`
	}{
		Val:  1,
		Data: []byte(data),
	}, "", "")
	fmt.Println(err)
	fmt.Println(body)
	// output:

}
