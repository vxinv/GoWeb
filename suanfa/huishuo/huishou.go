package main
import (
	"fmt"
)
type Node struct{
	Layer int
	IsRight bool
	IsLeft bool
	UpNode *Node
	LeftNode *Node
	RightNode *Node
}
var cw int =0
var cv int =0
var mw [6]int=[6]int{0,2,5,4,2,0}
var mv [6]int=[6]int{0,10,3,5,4,0}
var bestp int=0
func main(){
		root:=&Node{Layer:0,IsRight:true,IsLeft:true}
		Look(root)
	}
 func  Look(currentNode *Node){
       if(currentNode.Layer==5){
		   bestp=cv
		   fmt.Println(bestp)
		   return 
	    }
	   if(cw+mw[currentNode.Layer]<10){
		   node:=Node{Layer:currentNode.Layer+1,IsLeft:true,IsRight:true}
		   currentNode.LeftNode=&node
		   currentNode.IsLeft=false
		   node.UpNode=currentNode
		   cw+=mw[currentNode.Layer]
		   cv+=mv[currentNode.Layer]
		   Look(&node)
		   cv-=mv[currentNode.Layer]
		   cw-=mw[currentNode.Layer]
		}else{
           currentNode.IsLeft=false
		}
	    if(cv+Computer(currentNode.Layer+1)>bestp){
			 node:=Node{Layer:currentNode.Layer+1,IsLeft:true,IsRight:true}
			 currentNode.RightNode=&node
			 node.UpNode=currentNode
			 currentNode.IsRight=false
			 Look(&node)
        }else{
             currentNode.IsRight=false
		}
}
func Computer(i int) int  {
	num:=0
    for g:=i;g<6;g++{
        num+=mv[i]
	}
	return num
}	