package main
type tree struct {
	value int
	left,right *tree
}
func Sort(values []int)  {
	var root *tree
	for _,v :=range values{
		root=add(root,v)
	}
}

type trees struct {
	root *tree
	Size int
}
func add(root *tree, i int) *tree {
	if root==nil{
		root=new(tree)
		root.value=i
		return root
	}
	if(i<root.value){
		root.left=add(root.left,i)
	}else {
		root.right = add(root.right, i)
	}
	return root
}
func main()  {

}
