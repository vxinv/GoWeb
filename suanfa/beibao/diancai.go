package main
import (
	"fmt"
)
type Node struct{
	Queue [6]int
	LeftNode *Node
	RightNode *Node
	Value int
	Layer int
}
func main(){
	LayerArray:=map[int][]*Node{}
	ValueArry:=[5]int{18,19,17,6,7}
	root:= &Node{Layer:0}
    LayerArray[0]=append(LayerArray[0],root);
	for i , n:= range ValueArry{
		  for _,v:= range LayerArray[i]{
			lq:=v.Queue
			lq[i+1]=0
			rq:=v.Queue
			rq[i+1]=1 
			ln:=&Node{Layer:i+1,Value:n,Queue:lq}
		    rn:=&Node{Layer:i+1,Value:n,Queue:rq}
			  v.RightNode=rn
			  v.LeftNode=ln
			  LayerArray[i+1]=append(LayerArray[i+1],rn,ln)
		  }
	  }
    Order(root)
	}
	func Order(currentNode *Node){
        if(currentNode!=nil){
			Order(currentNode.RightNode)
			Computer(currentNode.Queue)
			
			Order(currentNode.LeftNode)
		}
	}
	func Computer(que [6]int){
		 ValueArry:=[5]int{18,19,17,6,7}
		 num:=0
		 for i ,v:=range ValueArry{
             num+=v*que[i+1]
		 }
		 fmt.Println(num)
	}