package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	b "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(b.SubsetsWithDup([]int{2, 1, 2}))
}
