package main
import "fmt"
type node struct{
	Value string
	prev *node
	next *node
}
type Object interface {} 
type LinkedList struct{
	First *node
	Last *node
    Size uint64
}
func (lk *LinkedList) Init() {
   lk.Last=nil;
   lk.First=nil
   lk.Size=0
}
func(lk *LinkedList)add(n *node){
	if (lk.Size==0) {
		lk.First=n
		lk.Last=n
		lk.Size+=1
	}else {
		lk.Last.next=n
		n.prev=lk.Last
		lk.Last=n
		lk.Size+=1
	}
}
func (lk *LinkedList)get(index int) string {
	tem:=lk.First;
	for i:=0;i<index;i++ {
		tem=tem.next

	}
	return tem.Value
}
func main(){
	lk:=LinkedList{};
	lk.Init();
	n:=&node{"李鑫",nil,nil}
	lk.add(n);
	n1:=&node{"总经理",nil,nil}
	lk.add(n1);
	n2:=&node{"总经理",nil,nil}
	lk.add(n2);
	n3:=&node{"总经理",nil,nil}
	lk.add(n3);
	n4:=&node{"总经理",nil,nil}
	lk.add(n4);
	fmt.Println(lk)
	fmt.Println(lk.get(2))
}
