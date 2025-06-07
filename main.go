package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	backtracking "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(backtracking.CombinationSum([]int{2, 3, 5}, 8))
}
