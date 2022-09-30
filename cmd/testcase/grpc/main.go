package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/sinngae/golet/cmd/testcase/protocol/go"
)

const (
	port = ":6606"
)

type server struct {
	pb.UnimplementedHeartBeatServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingBody) (*pb.PingBody, error) {
	fmt.Printf("ping: %s", in.GetMessage())
	return &pb.PingBody{
		Message: in.Message,
	}, nil
}

func main() {
	lst, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHeartBeatServer(s, &server{})
	s.Serve(lst)
}
