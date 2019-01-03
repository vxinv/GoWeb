
package engine
import (
	"fmt"
)
type ConcurrentEngnine struct {
   Scheduler Scheduler
   WorkerCount int
}

type  Scheduler interface {
	Submit(Request)
	ConfigWorkerchan(chan Request)
	WorkerReady(chan Request)
	Run()
}
func (ce *ConcurrentEngnine)Run( seeds ...Request) {
	out:=make(chan ParserRusult)
	ce.Scheduler.Run()
    for i:=0;i<ce.WorkerCount;i++{
         creatWorker(out,ce.Scheduler)
	}
	for _ , r :=range seeds{
		ce.Scheduler.Submit(r)
	}
	for {//这里会开启一个携程 ,这一句会一直等待out这个channal里面的值
		result:=<-out
		for _,item:=range result.Items{
			fmt.Println(item)
		}
		for _,Request:=range result.Requests{
			ce.Scheduler.Submit(Request)
		}
	}           
}
func creatWorker(out chan ParserRusult,s Scheduler)  {
	in:=make(chan Request)
	go func() {
		for{
			s.WorkerReady(in)
			request:=<-in
		    result ,err:=worker(request)
		    if(err!=nil){
		    	continue
			}
			out<-result
		}
	}()
}
