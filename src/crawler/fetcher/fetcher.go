package fetcher
import (
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)
func Fetch(url string )([]byte,error){
	resp,err:=http.Get(url)  
	if err!=nil{
     return nil,err
	}
	defer resp.Body.Close();
	all:=[]byte{}
	if resp.StatusCode==http.StatusOK{
		utf8reader:=transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
		all,err=ioutil.ReadAll(utf8reader)
		if err!=nil{
			return  nil,errors.New("error")
		}
	}
		// fmt.Printf("%s\n",all)
		return all,nil
}