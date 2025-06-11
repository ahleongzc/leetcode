package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	bt "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(bt.SubsetXORSum([]int{5, 1, 6}))
}
