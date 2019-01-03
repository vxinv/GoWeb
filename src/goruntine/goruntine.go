package main

import (
	"fmt"
)
// 像map一样，通道是一个使用make创建的数据结构的引用。当复制或者作为参数传递到一个函数时，复制的是引用
func mains(){
natura:=make(chan int);
squares:=make(chan int);
//counter
go func(){
	for x:=0;x<100;x++{
		natura<-x;
	}
	close(natura)
}()
//squares
//ok模式比较笨拙
//使range循环可以可以在接收最后一个值后关闭循环
go func(){
	for x:=range natura{
		squares<-x*x;
	}
	close(squares)
}()
//print
for x:=range squares {
	fmt.Println(x)
 }
}
