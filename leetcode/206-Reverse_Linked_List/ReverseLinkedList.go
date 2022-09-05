package reverselinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Method 1 - iterate
// https://leetcode.com/problems/reverse-linked-list/discuss/2337797/Go-Easy-Solution-oror-LinkedLists-oror-Iterative-oror-Recursion-oror-Fastest-Submission
func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode = nil
	var current *ListNode = head
	var next *ListNode = nil

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

// Method 2 - recursive
func ReverseList2(head *ListNode) *ListNode {
	// reverse 要用來做頭尾交換
	// 第一個參數要指向第二個參數
	// 亦即 head node 要指向 nil
	return reverse(head, nil)
}

func reverse(current, previous *ListNode) *ListNode {
	if current == nil {
		return previous
	}
	var next *ListNode = current.Next
	current.Next = previous
	return reverse(next, current)
}
