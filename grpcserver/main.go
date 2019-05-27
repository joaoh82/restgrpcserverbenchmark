package main

import (
	"context"
	"log"
	"net"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen. %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRandomServiceServer(s, &server{})
	log.Println("Starting gRPC server")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve. %v", err)
	}
}

func (s *server) DoSomething(_ context.Context, random *pb.Random) (*pb.Random, error) {
	random.RandomString = "[Updated] " + random.RandomString
	return random, nil
}
