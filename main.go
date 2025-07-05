package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/2d_dynamic_programming"
	_ "leetcode/advanced_graphs"
	_ "leetcode/arrays_and_hashing"
	b "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/greedy"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func main() {
	testCases := []struct {
		digits string
		expect []string
	}{
		{
			digits: "23",
			expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "2",
			expect: []string{"a", "b", "c"},
		},
		{
			digits: "",
			expect: []string{},
		},
		{
			digits: "9",
			expect: []string{"w", "x", "y", "z"},
		},
		{
			digits: "234",
			expect: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
	}

	for i, tc := range testCases {
		result := b.LetterCombinations(tc.digits)
		fmt.Printf("Test case %d: digits=%s\n", i+1, tc.digits)
		fmt.Printf("Expected: %v\n", tc.expect)
		fmt.Printf("Got: %v\n", result)
		if fmt.Sprintf("%v", result) == fmt.Sprintf("%v", tc.expect) {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
		fmt.Println()
	}
}
