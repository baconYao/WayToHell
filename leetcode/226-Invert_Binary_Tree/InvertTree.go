package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1
// InvertTree inverts the binary tree
// Runtime: 0 ms, Memory Usage: 2.1 MB
func InvertTreeTraverse(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

// traverse inverts the binary tree
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	traverse(root.Left)
	traverse(root.Right)
}

// Method 2
// Runtime: 4 ms, Memory Usage: 2.2 MB
func InvertTreeDivde(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := InvertTreeDivde(root.Right)
	right := InvertTreeDivde(root.Left)

	root.Left = left
	root.Right = right

	return root
}

// Method 3
// Runtime: 0 ms, Memory Usage: 2.1 MB
func InvertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree3(root.Left)
	InvertTree3(root.Right)

	return root
}
