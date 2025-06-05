package main

import (
	"fmt"
	_ "leetcode/arrays_and_hashing"
	_ "leetcode/backtracking"
	_ "leetcode/linked_list"
	_ "leetcode/stack"
	stack "leetcode/stack"
	_ "leetcode/two_pointers"
)

func main() {
	fmt.Println(stack.AsteroidCollision([]int{20, 50, -20, -20}))
}
