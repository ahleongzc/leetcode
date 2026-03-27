package trees

import "slices"

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	var dfs func(root *TreeNode, value int)
	dfs = func(root *TreeNode, value int) {
		if root == nil {
			return
		}
		root.Val = value
		if root.Left != nil {
			dfs(root.Left, 2*value+1)
		}
		if root.Left != nil {
			dfs(root.Right, 2*value+2)
		}
	}

	dfs(root, 0)

	return FindElements{
		root: root,
	}
}

func (this *FindElements) Find(target int) bool {
	// 0 == even, 1 == odd
	steps := make([]int, 0)
	for target != 0 {
		if target%2 == 0 {
			target = (target - 2) / 2
			steps = append(steps, 0)
		} else {
			target = (target - 1) / 2
			steps = append(steps, 1)
		}
	}

	slices.Reverse(steps)
	ptr := this.root
	for _, step := range steps {
		if ptr == nil {
			return false
		}
		if step == 0 {
			ptr = ptr.Right
		} else {
			ptr = ptr.Left
		}
	}

	return ptr != nil
}
