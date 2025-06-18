package trees

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1Leaves := make([]int, 0)
	root2Leaves := make([]int, 0)

	var dfs func(root *TreeNode, slice *[]int)
	dfs = func(root *TreeNode, slice *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*slice = append(*slice, root.Val)
			return
		}
		dfs(root.Left, slice)
		dfs(root.Right, slice)
	}

	dfs(root1, &root1Leaves)
	dfs(root2, &root2Leaves)

	ptr1 := 0
	ptr2 := 0

	for ptr1 < len(root1Leaves) && ptr2 < len(root2Leaves) {
		if root1Leaves[ptr1] != root2Leaves[ptr2] {
			return false
		}
		ptr1++
		ptr2++
	}

	return len(root1Leaves) == len(root2Leaves)
}
