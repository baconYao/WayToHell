package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method一
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

// helper 用於判斷 root 的左邊全都要小於 root 且 root 的右邊全都要大於 root
func helper(root, min, max *TreeNode) bool {
	// 等於 nil 相當於走完了 node
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return helper(root.Left, min, root) && helper(root.Right, root, max)
}

// 與 Method一 不同寫法
func isValidBST2(root *TreeNode) bool {
	max := math.MaxInt64
	min := -math.MaxInt64

	var dfs func(*TreeNode, int, int) bool

	dfs = func(tn *TreeNode, i1, i2 int) bool {
		if tn == nil {
			return true
		}

		if tn.Val > i1 && tn.Val < i2 {
			return dfs(tn.Left, i1, tn.Val) && dfs(tn.Right, tn.Val, i2)
		}

		return false
	}

	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}
