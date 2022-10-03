package diameterofbinarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxDiameter int = 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	maxDepth(root)
	return maxDiameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	// 記錄當前子樹的最長直徑
	currentMaxDiameter := leftMax + rightMax
	// 記錄最長直徑
	maxDiameter = Max(maxDiameter, currentMaxDiameter)
	// 回傳值給父節點
	// 1 表示離開此子節點，再挑選該子樹的最長路徑作為該父節點的其中一條最長路徑
	return 1 + Max(leftMax, rightMax)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// ======================================================

func DiameterOfBinaryTree2(root *TreeNode) int {
	diameter := 0
	// pass diameter by reference - acts as a global variable
	depth(root, &diameter)
	return diameter
}

func depth(root *TreeNode, diameter *int) int {

	var ld = 0
	var rd = 0

	if root == nil {
		return 0
	}

	if root.Left != nil {
		ld = 1 + depth(root.Left, diameter)
	}

	if root.Right != nil {
		rd = 1 + depth(root.Right, diameter)
	}

	currDiameter := ld + rd
	*diameter = Max(*diameter, currDiameter)

	return Max(ld, rd)
}
