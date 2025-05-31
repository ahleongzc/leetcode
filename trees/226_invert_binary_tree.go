package trees

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := InvertTree(root.Left)
	right := InvertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}
