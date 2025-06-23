package trees

func RemoveLeafNodes(root *TreeNode, target int) *TreeNode {
	isLeaf := func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Left == nil && root.Right == nil {
			return true
		}
		return false
	}

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode{
		if root == nil {
			return nil
		}

		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)

		if isLeaf(root) && root.Val == target {
			return nil
		}

		return root
	}

	return dfs(root)
}
