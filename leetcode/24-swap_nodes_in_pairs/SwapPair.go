package swappair

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{Val: 0, Next: head}
	curr := head
	next := curr.Next
	head = next

	for {
		prev.Next = next
		curr.Next = next.Next
		next.Next = curr
		prev = curr
		curr = curr.Next
		if curr == nil {
			break
		}
		next = curr.Next
		if next == nil {
			break
		}

	}
	return head
}

// // Recursive Way
// func swapPairs(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }

//     first := head
//     second := head.Next

//     first.Next = swapPairs(second.Next)
//     second.Next = first

//     return second
// }
