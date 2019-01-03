package main
import (
	"fmt"
	"io/ioutil"
)
func main(){
	data,err:=ioutil.ReadFile("./../main/tianmao.html")
	fmt.Println(string(data))
	fmt.Println(err)
}