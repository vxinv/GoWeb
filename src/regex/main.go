package main

import (
	"fmt"
	"regexp"
)
const text=`My Eamil address monneyloveme8866@Gmail.com
  emali is iii@org.com
  emali is ggg@qq.com
`
func main(){
	//  . + 代表任意字符
	/*  re,err:= regexp.Compile(`[a-zA-Z0-9]+@.+\..+`)
	if err!=nil{
		panic(err)
	}*/
	//match:=re.FindString(text)
	//匹配多个字符串  -1代表所有匹配的
	//match:=re.FindAllString(text,-1)
	//提取的功能
	re,err:= regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	if err!=nil{
		panic(err)
	}
	match:=re.FindAllStringSubmatch(text,-1)
	fmt.Println(match)
}