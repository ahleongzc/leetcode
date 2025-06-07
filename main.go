package main

import (
	"fmt"
	oneddp "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(oneddp.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
