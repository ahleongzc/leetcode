package main

import (
	"fmt"

	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/2d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	g "leetcode/greedy"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(g.FindBuildings([]int{4, 2, 3, 2, 1}))
}
