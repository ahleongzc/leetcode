package main

import (
	"fmt"

	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	t "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	root := &t.TreeNode{Val: 1}

	// Level 1
	root.Left = &t.TreeNode{Val: 2}
	root.Right = &t.TreeNode{Val: 3}

	// Level 2
	root.Left.Left = &t.TreeNode{Val: 4}
	root.Left.Right = &t.TreeNode{Val: 5}
	root.Right.Right = &t.TreeNode{Val: 6}

	// Level 3
	root.Left.Left.Left = &t.TreeNode{Val: 7}
	root.Right.Right.Left = &t.TreeNode{Val: 8}
	root.Right.Right.Right = &t.TreeNode{Val: 9}

	// Level 4
	root.Left.Left.Left.Left = &t.TreeNode{Val: 10}
	root.Left.Left.Left.Right = &t.TreeNode{Val: 11}
	root.Right.Right.Right.Left = &t.TreeNode{Val: 12}

	// Level 5
	root.Right.Right.Right.Left.Right = &t.TreeNode{Val: 13}
	// Calculate diameter
	diameter := t.DiameterOfBinaryTree(root)

	fmt.Printf("Diameter of binary tree: %d\n", diameter)
}
