package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/2d_dynamic_programming"
	_ "leetcode/advanced_graphs"
	ah "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/greedy"
	_ "leetcode/linked_list"
	_ "leetcode/sliding_window"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	fmt.Println(ah.TopKFrequent(nums, 2))
}
