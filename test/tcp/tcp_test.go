package tcp

import (
	"net"
	"net/http"
)

func ExampleTcpServer() {
	net.Listen("127.0.0.1")
	http.ListenAndServe()
	//output:

}
