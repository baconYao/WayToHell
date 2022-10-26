package constructbinarytreefrompreorderandinordertraversal

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
// 參考 https://labuladong.github.io/algo/2/21/38/ 詳解
// 利用 前序和中序 的排列特性組成 tree
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = preorder[0]
	idx := -1
	for k, v := range inorder {
		if v == root.Val {
			idx = k
			break
		}
	}
	root.Left = BuildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = BuildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}
