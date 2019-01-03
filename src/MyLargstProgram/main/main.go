package main
import (
	// "encoding/json"
	"fmt"
	"net/http"
	 "io/ioutil"
	// "gopkg.in/mgo.v2/bson"
	"MyLargstProgram/HandleFun"
	"MyLargstProgram/utils"
	"strconv"

)
var Counter int 
func Login(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello")
	 for k,v:=range r.Form{
		 fmt.Println(k,v)
	 }
}
func RCounter(){
	counter,err:=ioutil.ReadFile("counter.txt")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(strconv.Atoi(string(counter)))
	Counter,_=strconv.Atoi(string(counter))
}
func WCounter(Counter int) {
    lc := []byte(string(Counter))
    err := ioutil.WriteFile("counter.txt", lc, 0644)
    check(err)
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main(){
	// RCounter()
	// HandleFun.In(Counter)
	// defer func(){
	// 	Counter=HandleFun.Out()
	// 	WCounter(Counter)
	// 	fmt.Println("bye")
	// } ()
	fmt.Println("正在准备")
	http.HandleFunc("/test",HandleFun.Test)
	http.HandleFunc("/login",Login)
	http.HandleFunc("/getToken",utils.Gettoken)
	http.HandleFunc("/publish",HandleFun.Publish)
	http.HandleFunc("/loadmore",HandleFun.Loadmore)
	http.HandleFunc("/loadone",HandleFun.Loadone)
	http.HandleFunc("/uppinglun",HandleFun.Uppinglun)
	http.HandleFunc("/loginlocalserve",HandleFun.LoginLocalServer)
	http.HandleFunc("/writemessage",HandleFun.WriteMessage)
	http.HandleFunc("/readmessage",HandleFun.ReadMessage)
	http.HandleFunc("/readonemessage",HandleFun.ReadOneMessage)
	http.HandleFunc("/loaduserinfo",HandleFun.LoadUserinfo)
	http.HandleFunc("/unsetchat",HandleFun.UnsetChat)
	http.HandleFunc("/clearmessage",HandleFun.ClearMessage)
	http.HandleFunc("/zan",HandleFun.Zan)
	http.HandleFunc("/pullreplay",HandleFun.Pullreplay)
	http.Handle("/", http.FileServer(http.Dir("/root/build")))
	err:=http.ListenAndServe(":9090",nil)
	CHECHERRPR(err)
	// http.ListenAndServeTLS(":443","1_ijavascript.club_bundle.crt","2_ijavascript.club.key", nil)	
}   
func CHECHERRPR(e error){
	fmt.Println(e)
	panic(e)
}