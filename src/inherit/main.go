package main

import (
	"fmt"
)

type Student struct{
	Name string
	Age int
	Score int
}
//显示成绩
type Pupil struct{
	 Student
}
type Daxuesheng struct{
     Student
}
type Gaozhongsheng struct{
     Student
}
func (p *Student) show(){
	fmt.Println(p.Name+"is"+string(p.Age)+string(p.Score)+"")
}
type Brand struct{
	Name string
	Address string
}
type Goods struct{
	Name string
	Price int
}
type TV struct{
	Goods
    Brand
}

func main(){
	pupil :=Pupil{Student{"dachui",19,45}}
	pupil.show()
	tv:=TV{Goods{"电视机",11000},Brand{"创维","深圳市罗湖区"}}
	fmt.Println(tv)
}
