package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"MyLargstProgram/MgoInit"
	"MyLargstProgram/Mytype"
)

type person struct {
    Xuehao string
}
func main(){
	counter()
}
func  counter(){
  session:=MgoInit.Session()
  dbuser:=session.DB("NGDATA").C("user")
  var users []Mytype.User
  err:=dbuser.Find(bson.M{"usernewnotify.flag":bson.M{"$lte":-1}}).All(&users)
  if(err!=nil){
    fmt.Println(err)
  }
  fmt.Println(len(users[0].Usernewnotify))
}

func test(){
	session, err := mgo.Dial("211.159.175.75:27017")
    if err!=nil{
    	panic(err)
	}
	defer session.Close()
	//Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("NGDATA").C("user")
    err=c.Insert(&person{"1504815041"})
	if(err!=nil){
		panic(err)
	}
	/*M是map[String]接口{}map的方便别名，用于以本机方式处理BSON。
	例如：bson.m{“a”：1，“b”：true}除了对等效映射
	类型所做的操作之外，没有对此类型进行特殊处理。
	映射中的元素将被转储到一个未定义的有序中。还请参阅bson.d类型以获得有
	序的替代方案。*/
	persons:=[]person{}
	err=c.Find(bson.M{"xuehao":"1504815041"}).All(&persons)
	if(err!=nil){
		panic(err)
		fmt.Println("hello world")
	}
    fmt.Println(persons)
}



















