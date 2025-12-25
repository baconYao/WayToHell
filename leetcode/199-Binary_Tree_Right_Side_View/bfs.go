package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// q := []*TreeNode{root}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	// output := []int{}
	output := make([]int, 0)

	for len(q) > 0 {
		currentLen := len(q)
		output = append(output, q[currentLen-1].Val)
		for i := 0; i < currentLen; i++ {
			// pop out
			node := q[0]
			q = q[1:]
			// insert children
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return output
}
