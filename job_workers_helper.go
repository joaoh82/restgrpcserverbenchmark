package main

import (
	"context"
	"sync"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"
)

type Request struct {
	Path   string
	Random *pb.Random
}

const stopRequestPath = "STOP"

func startWorkers(requestQueue *chan Request, noWorkers int, startWorker func(*chan Request, *sync.WaitGroup)) func() {
	var wg sync.WaitGroup
	for i := 0; i < noWorkers; i++ {
		startWorker(requestQueue, &wg)
	}
	return func() {
		wg.Add(noWorkers)
		stopRequest := Request{Path: stopRequestPath}
		for i := 0; i < noWorkers; i++ {
			*requestQueue <- stopRequest
		}
		wg.Wait()
	}
}

func startWorker(requestQueue *chan Request, wg *sync.WaitGroup) {
	go func() {
		for {
			request := <-*requestQueue
			if request.Path == stopRequestPath {
				wg.Done()
				return
			}
			get(request.Path, request.Random)
		}
	}()
}

func startPostWorker(requestQueue *chan Request, wg *sync.WaitGroup) {
	go func() {
		for {
			request := <-*requestQueue
			if request.Path == stopRequestPath {
				wg.Done()
				return
			}
			post(request.Path, request.Random, request.Random)
		}
	}()
}

func getStartGRPCWorkerFunction(client pb.RandomServiceClient) func(*chan Request, *sync.WaitGroup) {
	return func(requestQueue *chan Request, wg *sync.WaitGroup) {
		go func() {
			for {
				request := <-*requestQueue
				if request.Path == stopRequestPath {
					wg.Done()
					return
				}
				client.DoSomething(context.TODO(), request.Random)
			}
		}()
	}
}
