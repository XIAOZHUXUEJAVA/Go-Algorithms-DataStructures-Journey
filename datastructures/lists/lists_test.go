package lists

import "testing"

func TestDisplayListNodes(t *testing.T) {
	list := NewSinglyLinkedList()
	list.AddAtBegin(5)
	list.AddAtBegin(4)
	list.AddAtBegin(3)
	list.AddAtBegin(2)
	list.AddAtBegin(1)
	list.DisplayListNodes()
}
