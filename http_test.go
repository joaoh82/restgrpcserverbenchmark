package main

import (
	"net/http"
	"testing"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"
	"golang.org/x/net/http2"
)

const noWorkers = 2

func BenchmarkHTTP11GetWithWorkers(b *testing.B) {
	client.Transport = &http.Transport{
		TLSClientConfig: createTLSConfigWithCustomCertificate(),
	}
	requestQueue := make(chan Request)
	defer startWorkers(&requestQueue, noWorkers)()
	b.ResetTimer() // don't count worker initialization time
	for i := 0; i < b.N; i++ {
		requestQueue <- Request{Path: "https://localhost:9191/get", Random: &pb.Random{}}
	}
}

// This results in:
// // Get https://localhost:9191: dial tcp 127.0.1.1:9191: socket: too many open files
// func BenchmarkHTTP1Get(b *testing.B) {
// 	client.Transport = &http.Transport{
// 		TLSClientConfig: createTLSConfigWithCustomCertificate(),
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(b.N)
// 	for i := 0; i < b.N; i++ {
// 		go func() {
// 			get("https://localhost:9191", &pb.Random{})
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

// HTTP2 test without workers
// func BenchmarkHTTP2Get(b *testing.B) {
// 	client.Transport = &http2.Transport{
// 		TLSClientConfig: createTLSConfigWithCustomCertificate(),
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(b.N)
// 	for i := 0; i < b.N; i++ {
// 		go func() {
// 			get("https://localhost:9191", &pb.Random{})
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

func BenchmarkHTTP2GetWithWorkers(b *testing.B) {
	client.Transport = &http2.Transport{
		TLSClientConfig: createTLSConfigWithCustomCertificate(),
	}
	requestQueue := make(chan Request)
	defer startWorkers(&requestQueue, noWorkers)()
	b.ResetTimer() // don't count worker initialization time
	for i := 0; i < b.N; i++ {
		requestQueue <- Request{Path: "https://localhost:9191/get", Random: &pb.Random{}}
	}
}
