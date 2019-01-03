package main
import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler/scheduler"
)
func main(){
	   e:=engine.ConcurrentEngnine{Scheduler:&scheduler.QueueScheduler{},
	   WorkerCount:100}
	   e.Run(engine.Request{Url:"http://city.zhenai.com/",
	   ParserFunc:parser.ProcessCity})
}