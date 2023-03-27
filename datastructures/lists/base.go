package lists

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil}
}
