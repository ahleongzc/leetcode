package main

import (
	"fmt"
	_ "leetcode/1d_dynamic_programming"
	_ "leetcode/2d_dynamic_programming"
	_ "leetcode/advanced_graphs"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/binary_search"
	_ "leetcode/graphs"
	_ "leetcode/greedy"
	_ "leetcode/linked_list"
	s "leetcode/sliding_window"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

func generateAlternatingString(n int, a, b rune) string {
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			runes[i] = a
		} else {
			runes[i] = b
		}
	}
	return string(runes)
}

func main() {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"All same chars", "AAAA", 2, 4},
		{"Single replacement needed", "ABAB", 2, 4},
		{"Mixed chars, small k", "AABABBA", 1, 4},
		{"No replacements allowed", "ABAB", 0, 1},
		{"Change all if needed", "ABCDE", 5, 5},
		{"Empty string", "", 3, 0},
		{"Single char string", "Z", 1, 1},
		{"Long same with small break", "AAAAABAAAAA", 1, 11},
		{"Large alternating string", generateAlternatingString(20, 'A', 'B'), 10, 20},
	}

	for _, tt := range tests {
		result := s.CharacterReplacement(tt.s, tt.k)
		pass := result == tt.expected
		fmt.Printf("%-30s s=%q k=%d → got %d (expected %d) ✅=%v\n",
			tt.name, tt.s, tt.k, result, tt.expected, pass)
	}
}
