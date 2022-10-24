package maximumbinarytree

import "math"

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

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	idx := -1
	maxValue := math.MinInt
	for i := low; i <= high; i++ {
		if nums[i] > maxValue {
			idx = i
			maxValue = nums[i]
		}
	}

	root := &TreeNode{}
	root.Val = maxValue
	root.Left = build(nums, low, idx-1)
	root.Right = build(nums, idx+1, high)

	return root
}
