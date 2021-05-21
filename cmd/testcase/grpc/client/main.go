package main

import (
	"context"
	"log"
	"os"

	pb "github.com/sinngae/gland/cmd/testcase/protocol/go"
	"google.golang.org/grpc"
)

const (
	address    = "localhost:6606"
	defaultMsg = "hi,work"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewHeartBeatClient(conn)

	msg := defaultMsg
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	r, err := c.Ping(context.Background(), &pb.PingBody{Message: &msg})
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	log.Printf("ping: %s", r.GetMessage())
}
