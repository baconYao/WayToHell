package main

import "container/list"

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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Method 1, DFS recursive
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// Method 2, DFS recursive
type Res struct {
	Depth int
}

func maxDepth2(root *TreeNode) int {
	r := Res{}
	r.dfs(root, 0)
	return r.Depth
}

func (r *Res) dfs(node *TreeNode, count int) {

	if node != nil {
		r.dfs(node.Left, count+1)
		r.dfs(node.Right, count+1)
	}
	if count > r.Depth {
		r.Depth = count
	}
}

// Method 3, BFS iteration
func maxDepth3(root *TreeNode) int {

	if root == nil {
		return 0
	}

	nodeList := []*TreeNode{root}
	depth := 0

	for len(nodeList) > 0 {

		for _, node := range nodeList {

			nodeList = nodeList[1:]

			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
		}

		depth++
	}

	return depth
}

// method 4, BFS iteration
func maxDepth4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	level := 0
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level++
	}
	return level
}
