package main

import "fmt"

func one(ch chan int,i int){
	ch <-i
}
func two(ch chan int)  {
	ch<- 1
}
func main()  {
	ch1:=make(chan int)
	ch2:=make(chan int)
	for i:=0;i<10;i++{
		i=i*i
		go one(ch1,i);
		if(i==9){
			go two(ch2)
		}
	}
	for i:=0;i<10 ;i++  {

	}
	fmt.Println("hello")
}
