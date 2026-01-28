package trees

import "slices"

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}

		leftTree := dfs(root.Left)
		rightTree := dfs(root.Right)

		withRoot := root.Val + leftTree[1] + rightTree[1]
		withoutRoot := slices.Max(leftTree) + slices.Max(rightTree)

		return []int{withRoot, withoutRoot}
	}

	return slices.Max(dfs(root))
}
