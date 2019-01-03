package main
import "fmt"
type node struct {
	Data   int
	Lchild *node
	Rchild *node
}
type Trees struct {
	root *node
	size int
}
func NewTree(i int) *Tree {
	n := node{
		Data: i,
	}
	t := Tree{
		root: &n,
		size: 1,
	}
	return &t
}
func (t *Tree) AddNode(i int) {
	if t.root == nil {
		panic("must init tree first")
	}
	t.addNode(t.root, i)
}
func (t *Tree) addNode(parentNode *node, i int) {
	if parentNode.Data < i {
		if  parentNode.Lchild == nil {
			parentNode.Lchild = &node{
				Data: i,
			}
			t.size++
			return
		} else {
			t.addNode(parentNode.Lchild, i)
		}
	} else {
		if  parentNode.Rchild == nil {
			parentNode.Rchild = &node{
				Data: i,
			}
			t.size++
			return
		} else {
			t.addNode(parentNode.Rchild, i)
		}
	}
}

func main() {
	t := NewTree(0)
	t.AddNode(1)
	t.AddNode(2)
	t.AddNode(-1)
	t.AddNode(-5)
	fmt.Println(t)
	fmt.Println(t.root.Lchild.Data)
}
