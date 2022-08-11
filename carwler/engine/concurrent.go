package engine

import (
	"log"
)

// ConcurrentEngine 并发
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		creatWorker(in, out)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item is : %+v", item) // %v表示不转义
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

func creatWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			result, err := worker(<-in)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
