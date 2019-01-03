package main

import (
	"fmt"
	"time"
)
func one(ch chan int ){
  for i:=0;i<10;i++{
	  ch <-i
	  time.Sleep(time.Second)
  }
  close(ch)
}
func two(ch chan int,ch2 chan int){
	for x:=range ch{   //如果不使用for循环,函数只调用一次就结束了
	    ch2<-x*x
	}
	close(ch2)
}

func main(){
  ch1:=make(chan int)
  ch2:=make(chan int)
  go one(ch1)
  go two(ch1,ch2)
  for x:=range ch2{
	  fmt.Println(x)
  }
}