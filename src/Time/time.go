package main
import (
	"fmt"
	"io/ioutil"
)
func main(){
 Counter()
 writecounter()
}
func writecounter() {
    d1 := []byte("3")
    err := ioutil.WriteFile("counter.txt", d1, 0644)
    check(err)
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func Counter(){
	counter,err:=ioutil.ReadFile("counter.txt")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(counter)
	scounter:=string(counter)
	fmt.Println(scounter)
}