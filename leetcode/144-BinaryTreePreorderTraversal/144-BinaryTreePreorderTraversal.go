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

func preorderTraversal(root *TreeNode) []int {
	var valueList []int
	traverse(root, &valueList)
	return valueList
}

func traverse(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}

	*values = append(*values, root.Val)
	traverse(root.Left, values)
	traverse(root.Right, values)
}
