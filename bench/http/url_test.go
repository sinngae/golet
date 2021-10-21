package http

import (
	"fmt"
	"net/url"
)

func ExampleUrl() {
	path := url.Values{}
	path.Set("abc", "==def")
	str := path.Encode()
	fmt.Println(str)
	// output:

}
