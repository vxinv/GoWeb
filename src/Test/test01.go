package Test

import (
	"fmt"
)
type Base struct{
	a int
	c int
	d []string 
}
type Top struct{
	base Base
	aaa int
	bbb string
}
type one interface{
	say()
}
type two interface{
	sing()
}
func (t Top) sing(){
	fmt.Println("hello,sing");
}
func (t Top) say(){
	fmt.Println("hello,base")
}
func  ymain(){
  d:=new(Top);
  fmt.Println(d)
  fmt.Println(append(d.base.d,"hello,world"))
  var t two
  t=*d
  yes,ok:=t.(two)
  if ok{
	  comma,ok:=yes.(Top)
	  if ok{
		  fmt.Println(comma.base.c)
	  }
	  fmt.Println(yes)
  }

}