package lists

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	dummy := &ListNode{-1, nil}
	p := dummy

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
