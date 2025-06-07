package trees

import "math"

func IsValidBSTOptimised(root *TreeNode) bool {
	var valid func(root *TreeNode, left, right int) bool
	valid = func(root *TreeNode, left, right int) bool {
		if root == nil {
			return true
		}

		if root.Val >= right || root.Val <= left {
			return false
		}

		return valid(root.Left, left, root.Val) && valid(root.Right, root.Val, right)
	}

	return valid(root, math.MinInt, math.MaxInt)
}

func IsValidBST(root *TreeNode) bool {
	res := make([]int, 0)

	var inOrderTraversal func(root *TreeNode)
	inOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrderTraversal(root.Left)
		res = append(res, root.Val)
		inOrderTraversal(root.Right)
	}

	inOrderTraversal(root)

	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}

	return true
}
