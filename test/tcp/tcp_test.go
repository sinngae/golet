package tcp

import (
	"fmt"
	"net/http"
)

func ExampleTcpServer() {
	//lst, err := net.Listen("tcp", "127.0.0.1")
	err := http.ListenAndServe("127.0.0.1", nil)
	fmt.Printf("%s", err)
	//output:

}
