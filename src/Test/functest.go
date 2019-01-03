package main

import (
	"fmt"
)

func hello(){
	fmt.Println("hello lixin")
}

type FFF func()

func (f FFF)world(){
	fmt.Println("hello world")
	f()
}

func main(){
	FFF(hello).world()
}