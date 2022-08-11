package scheduler

import "u2pppw/carwler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 发送request to work chan
	s.workerChan <- request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
