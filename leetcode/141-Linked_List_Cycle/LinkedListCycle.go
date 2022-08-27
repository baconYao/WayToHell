package linkedlistcycle

// 注意以下 case
// [] 1
// [1] 1
// [1,1,1,1] -1

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

// Method 1
// Fast Slow pointer
// 只要有一個迴圈，兩個 pointer 一定最終會遇到
func hasCycle(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		// slow 一次移動一個
		s = s.Next
		// fast 一次移動兩個
		f = f.Next.Next
		if s == f {
			return true
		}

	}
	return false
}
