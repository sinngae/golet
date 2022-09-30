package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	listen, err := net.Listen("tcp", ":2333")
	if err != nil {
		log.Fatalf("net.Listen failed, err=%s\n", err)
	}
	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("server serve failed, err=%s\n", err)
	}
}
