package main

import (
	"fmt"
	"encoding/json"
)
type A struct {
    Num int
}
type B struct {
	Num int
}
//func(w http.Write,r *http.Response)
type Student struct {
	Name string
	Age int
}
//type HandlerFunc func(w http.Write,r *http.Response)

type Stu Student

//可以装换,结构体的字段要完全一样
//名字 个数 类型


/**
每个结构体上的字段可以写一个tag,可以通过反射机制获取 ,主要用于序列化和反序列化
*/
type Moster struct{
	Age int  `json:"age"`
	Name string `json:"name"`
	Skill string `json:"skill"`
}

func main()  {
	var a A
	var b B
	a = A(b)
	fmt.Println(a,b)
	var student Student
	var stu Stu
	stu=Stu(student)
	fmt.Println(stu,student)
	// 创建
	moster:=Moster{500,"牛魔王","芭蕉扇"}
	data,err:=json.Marshal(moster)
	if(err!=nil){

	}
	fmt.Println(string(data))
}
