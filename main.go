package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/2d_dynamic_programming"
	ag "leetcode/advanced_graphs"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/greedy"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {

	fmt.Println(ag.MinCostConnectPoints([][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}))

}
