package main

import (
	"fmt"
)
//如果一个结构体潜入了含有两个相同字段的结构体 必须指定名字
type A struct{
	Name string
}
type B struct{
	Name string
}
type C struct{
	A 
	B
}
func(a A)show(){
	fmt.Println(a.Name)
}
func(b B)show(){
	fmt.Println(b.Name)
}

func main(){
	c:=C{A{"小敏"},B{"小海"}}
    c.B.show()
}
