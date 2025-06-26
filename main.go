package main

import (
	"fmt"

	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/2d_dynamic_programming"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	g "leetcode/graphs"
	_ "leetcode/greedy"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	testCases := []struct {
		n      int
		edges  [][]int
		expect int
	}{
		{
			n:      3,
			edges:  [][]int{{0, 1}, {0, 2}},
			expect: 2,
		},
		{
			n:      5,
			edges:  [][]int{{0, 1}, {1, 2}, {3, 4}},
			expect: 2,
		},
		{
			n:      5,
			edges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			expect: 1,
		},
		{
			n:      4,
			edges:  [][]int{{0, 1}, {2, 3}},
			expect: 2,
		},
		{
			n:      4,
			edges:  [][]int{},
			expect: 4,
		},
		{
			n:      1,
			edges:  [][]int{},
			expect: 1,
		},
		{
			n:      2,
			edges:  [][]int{{1, 0}},
			expect: 1,
		},
	}

	for i, tc := range testCases {
		result := g.CountComponents(tc.n, tc.edges)
		fmt.Printf("Test case %d: n=%d, edges=%v\n", i+1, tc.n, tc.edges)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expect, result)
		if result == tc.expect {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
		fmt.Println()
	}
}
