package main
 import(
	 "net/http"
	 "fmt"
	 "io/ioutil"
	 "golang.org/x/text/encoding/simplifiedchinese"
     "golang.org/x/text/transform"
 )
 func main(){
	 req,_:=http.NewRequest("GET","http://222.139.215.217/jwweb/",nil);
	 req.Header.Set("Referer","http://www.nyist.edu.cn/")
	 client:=&http.Client{}
	 resp,err:=client.Do(req)
	 if err!=nil{
		fmt.Println(err)
	 }else {
		    utf8reader:=transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
		  	body, _ := ioutil.ReadAll(utf8reader)
    		fmt.Println("response Body:", string(body))
	 }
 }