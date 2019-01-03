package main
import (
	"fmt"
)
type SetName interface{
   setname(s string);
}
type USB interface{
  print() ;
}
type Phone struct{
	name string
}
type Computer struct{
	name string
}

func (c *Computer)play( u USB){
  u.print();
}
func (c *Computer)show(s SetName, n string){
    s.setname(n);
	}

func (p *Phone) print (){
	fmt.Println(p.name);
}

type MP3 struct{
	name string
}

func (m MP3)print(){
	fmt.Println(m.name);
}
func (p *Phone)setname(n string){
	p.name=n;
	p.print();
}
func (m MP3) setname(n string){
	m.name=n;
	m.print();
}
func main(){
	//简单接口
	computer:=new(Computer);
	computer.name="lianxiang"
	phone:=new(Phone);
	mpa3:=MP3{"BUBUGAO"};
	phone.name="huawei"
	computer.play(phone);
	computer.play(mpa3);
	//接口的类型判断
	var us USB
	us=phone
	switch us.(type){
	case *Phone:
        fmt.Println("phone")
	case MP3:
        fmt.Println("mp3")
	}
	//接口的类型断言
	computer.show(mpa3, "SONY");
	//这个地方要注意的是传给computer的是一个mpa3的副本 所以要改值要用*
	fmt.Println(mpa3.name);
	computer.show(phone,"IPONE");
	fmt.Println(phone.name);
}

