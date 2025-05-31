package trees

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		elements := queue[:]
		intermediaryResult := make([]int, 0)

		queue = make([]*TreeNode, 0)
		for _, element := range elements {
			intermediaryResult = append(intermediaryResult, element.Val)

			if element.Left != nil {
				queue = append(queue, element.Left)
			}

			if element.Right != nil {
				queue = append(queue, element.Right)
			}
		}

		result = append(result, intermediaryResult)
	}

	return result
} 
