package main

import (
	"fmt"

	_ "leetcode/1d_dynamic_programming"
	dp "leetcode/2d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Print(dp.UniquePaths(3, 7))
}
