package binarytreerightsideview

// Definition of a binary tree node
type TreeNode2[T any] struct {
	Data  T
	Left  *TreeNode2[T]
	Right *TreeNode2[T]
}

// rightSideView returns the right side view of a binary tree
func rightSideView2(root *TreeNode2[int]) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}

	dfs(root, 0, &output)

	return output
}

func dfs(node *TreeNode2[int], level int, rside *[]int) {
	if level == len(*rside) {
		*rside = append(*rside, node.Data)
	}

	for _, child := range []*TreeNode2[int]{node.Right, node.Left} {
		if child != nil {
			dfs(child, level+1, rside)
		}
	}
}
