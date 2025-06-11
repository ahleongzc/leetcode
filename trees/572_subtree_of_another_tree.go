package trees

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}
