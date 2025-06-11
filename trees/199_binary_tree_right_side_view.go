package trees

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	view := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		for range len(queue) - 1 {
			curr := queue[0]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			queue = queue[1:]
		}

		last := queue[0]
		view = append(view, last.Val)
		if last.Left != nil {
			queue = append(queue, last.Left)
		}
		if last.Right != nil {
			queue = append(queue, last.Right)
		}
		queue = queue[1:]
	}

	return view
}
