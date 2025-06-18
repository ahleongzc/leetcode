package trees

func DeleteNode(root *TreeNode, key int) *TreeNode {
	var findLargestInLeftSubtree func(node *TreeNode) *TreeNode
	findLargestInLeftSubtree = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Right == nil {
			return node
		}
		return findLargestInLeftSubtree(node.Right)
	}

	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// This will never throw a panic because if root.Left was nil, it would be handled in the previous case already
			largestNodeInLeftSubTree := findLargestInLeftSubtree(root.Left)
			root.Val = largestNodeInLeftSubTree.Val
			root.Left = DeleteNode(root.Left, largestNodeInLeftSubTree.Val)
		}
	}

	return root
}
