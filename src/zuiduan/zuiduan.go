package  main
import (
	"time"
	"fmt"
)
type animal struct {
	age int;
	name string;
}
type Dog struct {
	animal
	color string;
}
type say interface {
	say()
}
func (an animal) Voice(){
	fmt.Println("动物"+an.name)
}
func (d Dog) say()  {
}
func (d Dog) Eat() {
	fmt.Println("我在吃东西")
}
func (a animal) say(){
	fmt.Println("吃东西")
}
type ob1 struct {
	s1 string
}
type ob2 struct {
	s2 string
}
type obj interface {
	obtest(int,string)(string);
}
func(one ob1)obtest(i int,s string)(string) {
	return one.s1;
}
func(one ob2)obtest(i int,s string)(string){
		return one.s2;
		}
type Employment struct{
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
}
func main()  {
	arr:=make([]obj,0);
	arr=append(arr,ob1{"lalal"});
	arr=append(arr,ob2{"lalala"});
	for i,v:=range arr{
		fmt.Println(i,v)
	}
	myslice:=make([]say,5);
	myslice=append(myslice,animal{age:12,name:"lalla"});
	var  xiaohua Dog;
	xiaohua=Dog{animal{age:13,name:"lalal"},"lalla"};
	fmt.Println(xiaohua);
	myslice=append(myslice,xiaohua);
	fmt.Println(myslice);
}



