package lists

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	dummy := &ListNode{-1, nil}
	p := dummy
	carry := 0

	for p1 != nil || p2 != nil || carry > 0 {
		value := carry
		if p1 != nil {
			value += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			value += p2.Val
			p2 = p2.Next
		}
		carry = value / 10
		value = value % 10
		p.Next = &ListNode{value, nil}
		p = p.Next
	}
	return dummy.Next
}
