package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	binarysearch "leetcode/binary_search"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(binarysearch.ShipWithinDays([]int{3, 2, 2, 4, 1, 4}, 4))
}
