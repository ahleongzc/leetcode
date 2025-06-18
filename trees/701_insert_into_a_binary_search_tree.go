package trees

func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val > root.Val {
		root.Right = InsertIntoBST(root.Right, val)
	} else {
		root.Left = InsertIntoBST(root.Left, val)
	}

	return root
}
