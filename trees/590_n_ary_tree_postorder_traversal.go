package trees

type Node struct {
	Val      int
	Children []*Node
}

func Postorder(root *Node) []int {
	res := make([]int, 0)

	var postOrderTraversal func(root *Node)
	postOrderTraversal = func(root *Node) {
		if root == nil {
			return
		}

		for _, child := range root.Children {
			postOrderTraversal(child)
		}
		res = append(res, root.Val)
	}

	postOrderTraversal(root)

	return res
}
