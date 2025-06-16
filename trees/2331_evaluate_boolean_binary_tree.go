package trees

func EvaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	if root.Val == 2 {
		return EvaluateTree(root.Left) || EvaluateTree(root.Right)
	}

	return EvaluateTree(root.Left) && EvaluateTree(root.Right)
}
