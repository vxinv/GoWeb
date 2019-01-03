package main

import (
	"net/http/httputil"
	"fmt"
	"net/http"
)
func main(){
  resp,err:=http.Get("http://www.imooc.com")
  if err!=nil{
	  panic(err)
  }
  defer resp.Body.Close()
  response, err:=httputil.DumpResponse(resp,true)
  if err!=nil{
	  panic(err)
  }
  fmt.Println(string(response))
}