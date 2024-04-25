package main

import (
	"context"
	"log"
	"os"

	protocol2 "github.com/sinngae/golet/sample/server/internal/grpc/protocol/go"

	"google.golang.org/grpc"
)

const (
	address    = "localhost:6100"
	defaultMsg = "hi,work"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer conn.Close()
	c := protocol2.protocol.NewHeartBeatClient(conn)

	msg := defaultMsg
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	r, err := c.Ping(context.Background(), &protocol2.PingBody{Message: &msg})
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	log.Printf("ping: %s", r.GetMessage())
}
