package main
import (
	"strconv"
	"fmt"
)
var a [3]int;
func main(){
	for i,_:= range a{
	a[i]=i
}
  slice1:=make([]string,100);
  for i,_:=range slice1{
	  slice1[i]=strconv.Itoa(i);
  }
 fmt.Println(slice1);
 fmt.Println(a);
}
