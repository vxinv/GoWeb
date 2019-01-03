package scheduler

import "crawler/engine"

type SimpleSheduler struct {
    workerChan chan engine.Request
 }

func (s *SimpleSheduler) ConfigWorkerchan(ch chan engine.Request) {
	s.workerChan=ch
}

func (s *SimpleSheduler ) Submit(r engine.Request) {
	//send  request down to worker chan
	go func() {
		s.workerChan<-r
	}()
}


