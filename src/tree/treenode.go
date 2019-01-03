package main
import "fmt"
type Tree struct {
	Left *Tree
	Right *Tree
	Value int
}
func Insert (t *Tree,v int) {
	if t==nil{
		t=&Tree{nil,nil,v}
		fmt.Println(1)
		return
	}
	if  v<t.Value{
		if(t.Left==nil){
			t.Left=&Tree{nil,nil,v}
		}else {
			Insert(t.Left,v)
		}
		return
	}else {
		if(t.Right==nil){
			t.Right=&Tree{nil,nil,v}
		}else {
			Insert(t.Right,v)
		}
		return
	}

}
func main()  {
	t:=&Tree{nil,nil,2}
	Insert(t,1)
	Insert(t,2)
	Insert(t,4)
	Insert(t,3)
	Insert(t,6)
	Insert(t,-1)
	Insert(t,11)
	Insert(t,21)
	Insert(t,41)
	Insert(t,31)
	Insert(t,61)
	Insert(t,-11)
	Insert(t,6)
	Insert(t,4)
	Insert(t,45)
	Insert(t,31)
	Insert(t,26)
	Insert(t,-11)
	fmt.Println(t.Right.Right.Right.Right.Value)
}