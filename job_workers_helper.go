package main

import (
	"sync"

	"github.com/joaoh82/restgrpcserverbenchmark/pb"
)

type Request struct {
	Path   string
	Random *pb.Random
}

const stopRequestPath = "STOP"

func startWorkers(requestQueue *chan Request, noWorkers int) func() {
	var wg sync.WaitGroup
	for i := 0; i < noWorkers; i++ {
		startWorker(requestQueue, &wg)
	}
	// Returns a function that stops as many workers as were just started
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
