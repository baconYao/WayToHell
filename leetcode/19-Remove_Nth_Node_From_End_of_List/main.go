package removenthnodefromendoflist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	// Move P2 n steps forward
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	if p2 == nil {
		head = head.Next
		return head
	}

	for {
		if p2.Next == nil {
			tmp := p1.Next
			p1.Next = p1.Next.Next
			tmp.Next = nil
			break
		}
		p2 = p2.Next
		p1 = p1.Next
	}

	return head
}
