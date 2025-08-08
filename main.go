package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	twoddp "leetcode/2d_dynamic_programming"
	_ "leetcode/advanced_graphs"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/greedy"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println("the answer is", twoddp.LongestCommonSubsequence("abc", "abc"))
}
