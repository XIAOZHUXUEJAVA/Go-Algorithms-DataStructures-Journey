package lists

import "fmt"

type SinglyLinkedList struct {
	Size int
	Head *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{0, nil}
}

func (self *SinglyLinkedList) AddAtBegin(val int) {
	node := NewNode(val)
	node.Next = self.Head
	self.Head = node
	self.Size++
}

func (self *SinglyLinkedList) DisplayListNodes() {
	for p := self.Head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

func (self *SinglyLinkedList) Count() int {
	return self.Size
}
