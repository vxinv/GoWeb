package main

import (
	"fmt"
)
type tpmrature int
func ma(x,y int)(int){
	return x+y
}
type aaa int 
type bbb int	


func m(f func(int,int,)(int),x int,y int) int { 
	z:=f(x,y)
	return z
}
func main(){
	L:=m(ma,1,2)
	fmt.Println(L)
	var x aaa=1
	var y bbb=2
	z:=x+aaa(y)
	fmt.Println(z)
}


