package scheduler

import (
	"crawler/engine"
)
type  QueueScheduler struct {
	  requestChan chan engine.Request
	  workerChan  chan chan engine.Request
}

func (s *QueueScheduler) Submit(r  engine.Request) {
	s.requestChan<-r
}

func (s QueueScheduler) ConfigWorkerchan(r chan engine.Request) {

}

func (s *QueueScheduler) WorkerReady( r chan engine.Request) {
	s.workerChan<-r
}

func (s *QueueScheduler) Run() {
	s.workerChan =make(chan chan engine.Request)
	s.requestChan=make(chan engine.Request)
	go func() {
		var  requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0{
               activeRequest=requestQ[0]
               activeWorker=workerQ[0]
			}
			select {
			case r :=<-s.requestChan:
               requestQ=append(requestQ,r)
			case w :=<-s.workerChan:
				workerQ=append(workerQ,w)
			case activeWorker<-activeRequest:
				requestQ=requestQ[1:]
				workerQ=workerQ[1:]
			}
		}
	}()
}

