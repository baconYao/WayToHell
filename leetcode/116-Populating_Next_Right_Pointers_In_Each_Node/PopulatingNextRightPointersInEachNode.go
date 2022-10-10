package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//  method 1
// https://labuladong.github.io/algo/2/21/37/
// 概念: 從第二層開始，兩個點視為一個點
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	// 傳入兩個點
	traverse(root.Left, root.Right)
	return root
}

func traverse(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	// 左點連接右點
	node1.Next = node2

	// 點1 的左連到右
	traverse(node1.Left, node1.Right)
	// 點2 的左連到右
	traverse(node2.Left, node2.Right)
	// 點1的左連到點2的右
	traverse(node1.Right, node2.Left)
}

// method 2
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}

	connect2(root.Left)
	connect2(root.Right)
	return root
}
