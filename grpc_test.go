package main

import (
	"log"
	"testing"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"

	"google.golang.org/grpc"
)

func BenchmarkGRPCWithWorkers(b *testing.B) {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial. %v", err)
	}
	client := pb.NewRandomServiceClient(conn)
	requestQueue := make(chan Request)
	defer startWorkers(&requestQueue, noWorkers, getStartGRPCWorkerFunction(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- Request{
			Path: "http://localhost:9090",
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}
