package main

import (
	"fmt"
)

func bbmain(){
    w:= [6]int{0,2,5,4,2,3}
	v:=[6]int{0,6,3,5,4,6}
    mv:= [6][11]int{}
	for i:=1;i<6;i++{
       for j:=1;j<11;j++{
		   if j<w[i]{
			   mv[i][j]=mv[i-1][j]
		   }else {
			   mv[i][j]=max(mv[i-1][j],mv[i-1][j-w[i]]+v[i])
		   }
	   }
	}
	fmt.Println(mv)
}
func max( a int, b int) int {
      if a>b{
		  return a
	  }
	  return b
}