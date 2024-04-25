package main

import (
	"context"
	"fmt"
	"log"
	"net"

	protocol "github.com/sinngae/golet/sample/server/internal/grpc/protocol/go"

	"google.golang.org/grpc"
)

const (
	port = ":6606"
)

type server struct {
	protocol.protocol
}

func (s *server) Ping(ctx context.Context, in *protocol.PingBody) (*protocol.PingBody, error) {
	fmt.Printf("ping: %s", in.GetMessage())
	return &protocol.PingBody{
		Message: in.Message,
	}, nil
}

func main() {
	lst, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protocol.RegisterHeartBeatServer(s, &server{})
	_ = s.Serve(lst)
}
