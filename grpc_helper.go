package main

import (
	"context"
	"log"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"
	"google.golang.org/grpc"
)

const serverAddr = "localhost:9090"

func random(c context.Context, input *pb.Random) (*pb.Random, error) {
	conn, err := grpc.Dial(serverAddr)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := pb.NewRandomServiceClient(conn)
	return client.DoSomething(c, &pb.Random{})
}
