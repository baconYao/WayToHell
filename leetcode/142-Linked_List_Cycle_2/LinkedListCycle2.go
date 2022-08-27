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
func DetectCycle(head *ListNode) *ListNode {
	s, f := head, head
	hasCycle := false
	for f != nil && f.Next != nil {
		// slow 一次移動一個
		s = s.Next
		// fast 一次移動兩個
		f = f.Next.Next
		if s == f {
			hasCycle = true
			// 相遇後，將 s 指回 head
			s = head
			break
		}

	}
	// 沒有發現
	if !hasCycle {
		return nil
	}
	// 找 cycle 的起始點
	for {
		if s == f {
			return s
		}
		// 只要一起同時移動相同距離即可
		s = s.Next
		f = f.Next
	}
}

func DetectCycle2(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		// slow 一次移動一個
		s = s.Next
		// fast 一次移動兩個
		f = f.Next.Next
		if s == f {
			// 相遇後，將 s 指回 head
			s = head
			// 找 cycle 的起始點
			for {
				if s == f {
					return s
				}
				// 只要一起同時移動相同距離即可
				s = s.Next
				f = f.Next
			}
		}

	}
	// 沒有發現
	return nil
}
