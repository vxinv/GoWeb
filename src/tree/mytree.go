package main

import "fmt"

type Mytree struct {
	Rchlid ,Lchlid *Mytree
	Data int
}
func addnode (my *Mytree,v int)  {
	if my==nil{
		my=&Mytree{nil,nil,v}
		fmt.Println("hello")
		return
	}
	if my.Data>v{
		if my.Lchlid==nil{
			my.Lchlid=&Mytree{nil,nil,v}
		}else {
			addnode(my.Lchlid,v)
		}
	}else {
		if my.Rchlid==nil{
			my.Rchlid=&Mytree{nil,nil,v}
		}else {
			addnode(my.Rchlid,v)
		}
	}
}
func main()  {
	mytree:=&Mytree{nil,nil,1}
	addnode(mytree,3)
	addnode(mytree,-5)
	addnode(mytree,6)
	addnode(mytree,2)
	addnode(mytree,-1)
	addnode(mytree,7)
	fmt.Println(mytree.Rchlid.Rchlid.Data)
}
