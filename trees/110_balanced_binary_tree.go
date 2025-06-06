package trees

import "math"

func IsBalanced(root *TreeNode) bool {
	var balanced func(root *TreeNode) []int
	balanced = func(root *TreeNode) []int {
		if root == nil {
			return []int{1, 0}
		}

		nodeLeft := balanced(root.Left)
		nodeRight := balanced(root.Right)

		res := 0
		if math.Abs(float64(nodeLeft[1]-nodeRight[1])) <= 1 &&
			nodeLeft[0] == nodeRight[0] &&
			nodeLeft[0] == 1 {
			res = 1
		}

		return []int{res, max(nodeLeft[1], nodeRight[1]) + 1}
	}

	return balanced(root)[0] == 1
}
