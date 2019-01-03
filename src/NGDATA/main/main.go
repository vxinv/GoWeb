package main

import (
	"net/http"
	"fmt"
)
func sayHi(w http.ResponseWriter, r  *http.Request){
	//do nothing
	fmt.Println(r.Body)
}
func main() {
	//需要实现ServerHttp接口
	http.HandleFunc("/", sayHi)
	err:=http.ListenAndServeTLS(":9090", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("hello,world")
	}
}


