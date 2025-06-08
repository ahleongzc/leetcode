package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	arrays_and_hashing "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
)

func main() {
	fmt.Println(arrays_and_hashing.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
}
