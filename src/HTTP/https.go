package main

import (
	"net/http/httputil"
	"fmt"
	"net/http"
)
func main(){
  resp,err:=http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code")
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