package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	a "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(a.MajorityElementII([]int{1, 1, 2, 2, 3}))
}
