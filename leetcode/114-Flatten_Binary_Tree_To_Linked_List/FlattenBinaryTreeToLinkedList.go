package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1
// 東哥解法: https://labuladong.github.io/algo/2/21/37/
// 概念就是將每棵左子樹並到其右子樹
/*		1   			   1
/*	   / \				  / \
/*    2   5       -->    2   5		-->    Answer
/*   / \   \              \   \
/*  3   4   6			   3   6
/*							\
/*							 4
*/
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	Flatten(root.Left)
	Flatten(root.Right)
	// 在後序位實作

	/* 紀錄該樹原本的左右節點 */
	var leftNode *TreeNode = root.Left
	var rightNode *TreeNode = root.Right

	/* 將樹的右節點指標指左節點，使得原本是左邊的節點成為新的右節點 */
	root.Left = nil
	root.Right = leftNode

	/* 將原本的右節點接在新的右節點後面 */
	var p *TreeNode = root
	// 移動指標到新的右子樹的最後一個節點
	for p.Right != nil {
		p = p.Right
	}
	// 合併新舊右節點
	p.Right = rightNode
}

// ===========================================

// Method 2
func Flatten2(root *TreeNode) {
	helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	/* 取得尾部節點 */
	endLeft := helper(root.Left)
	endRight := helper(root.Right)

	if endLeft == nil && endRight == nil {
		return root
	} else if endLeft == nil {
		return endRight
	} else if endRight == nil {
		root.Right = root.Left
		root.Left = nil
		return endLeft
	}

	/* 右子樹接到左子樹尾部，左子樹移到變成右子樹 */
	endLeft.Right = root.Right
	root.Right = root.Left
	root.Left = nil

	return endRight
}

// ===========================================

func Flatten3(root *TreeNode) {
	if root == nil {
		return
	}

	Flatten3(root.Left)
	Flatten3(root.Right)

	// assume that left and right pointer is already a linked list
	var rightMostLeft **TreeNode = &root.Left
	/* 和東哥的做法雷同，這裡是移動左子樹的指標 */
	for *rightMostLeft != nil {
		// 這裏會指向左子樹最後一個節點的右子，此時應為 nil。
		// 此做法是為了下一行方便直接將原本右子樹併過來
		rightMostLeft = &((*rightMostLeft).Right)
	}
	*rightMostLeft = root.Right
	root.Right = root.Left
	root.Left = nil
}

// ===========================================
// Morris 遍歷, O(1) 空間解法：尋找前驅節點 (Predecessor)
func Flatten4(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			predecessor := curr.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			predecessor.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}

// ===========================================
// 時間複雜度：O(n)，其中 n 是二元樹的節點總數，因為每個節點都會被訪問一次。
// 空間複雜度：O(h)，其中 h 是樹的高度，主要消耗在遞迴的系統棧空間。在最壞情況下（樹呈線性），空間複雜度為 O(n)。
func Flatten5(root *TreeNode) {
	var prev *TreeNode

	// 定義一個內部遞迴函數
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 1. 先處理右子樹
		helper(node.Right)
		// 2. 再處理左子樹
		helper(node.Left)

		// 3. 核心操作：將當前節點轉換為鏈表節點
		node.Right = prev
		node.Left = nil

		// 4. 更新 prev，讓下一個節點可以接到當前節點上
		prev = node
	}

	helper(root)
}
