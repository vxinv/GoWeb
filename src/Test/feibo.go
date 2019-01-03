package main

import (
	"fmt"
)

func main(){
	ch:=make( chan int)
	quit:=make( chan bool)
	go func(){
       for i:=0; i<8;i++{
			num:=<-ch
			fmt.Println(num)
	   }
	   quit<-true
	}()
	//从chan读取内容
	fibo(ch ,quit)
}
func fibo(ch chan<- int,quit <-chan bool ){
	  x ,y:=1,1 
	  //不停的监听chanel 数据的流动
    for{
		select {
		case ch<-x:
            x,y=y,x+y
		case <-quit:
			fmt.Println("end")
			return 
		}
	}
	  
}
