package trees

func GoodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 1

	var dfs func(node *TreeNode, prev int)
	dfs = func(node *TreeNode, prev int) {
		if node == nil {
			return
		}

		if node.Val >= prev {
			res++
		}
		dfs(node.Left, max(node.Val, prev))
		dfs(node.Right, max(node.Val, prev))
	}

	dfs(root.Left, root.Val)
	dfs(root.Right, root.Val)

	return res
}
