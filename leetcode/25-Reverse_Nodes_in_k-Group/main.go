package reversenodesinkgroup

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  Optimized solution for O(n) time complexity and O(1) space complexity
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	ptr := dummy

	for ptr != nil {
		tracker := ptr

		for i := 0; i < k; i++ {
			if tracker == nil {
				break
			}
			tracker = tracker.Next
		}
		if tracker == nil {
			break
		}

		previous, current := reverseLinkedList(ptr.Next, k)

		lastOfReversedNode := ptr.Next
		lastOfReversedNode.Next = current
		ptr.Next = previous
		ptr = lastOfReversedNode
	}

	return dummy.Next
}

func reverseLinkedList(head *ListNode, k int) (*ListNode, *ListNode) {
	var previous, current, next *ListNode = nil, head, nil
	for i := 0; i < k; i++ {
		// temporarily store the next node
		next = current.Next
		// reverse the current node
		current.Next = previous
		// before moving to the next node, point previous to the current node
		previous = current
		// move to the next node
		current = next
	}

	return previous, current
}
