package trees

func KthSmallest(root *TreeNode, k int) int {
	numbers := make([]int, 0)

	var inOrderTraversal func(root *TreeNode)
	inOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraversal(root.Left)
		numbers = append(numbers, root.Val)
		inOrderTraversal(root.Right)
	}

	inOrderTraversal(root)

	return numbers[k-1]
}
