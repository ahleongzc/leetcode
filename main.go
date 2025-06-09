package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	twopointers "leetcode/two_pointers"
)

func main() {
	fmt.Println(twopointers.NumRescueBoats([]int{2, 3, 2, 1}, 3))
}
