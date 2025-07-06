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
	ll "leetcode/linked_list"
	_ "leetcode/stack"
	_ "leetcode/trees"
	_ "leetcode/two_pointers"
)

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ll.ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ll.ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ll.ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice
func linkedListToSlice(head *ll.ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{1, 2, 3, 4},
			expect: []int{1, 4, 2, 3},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			expect: []int{1, 5, 2, 4, 3},
		},
		{
			input:  []int{1},
			expect: []int{1},
		},
		{
			input:  []int{},
			expect: []int{},
		},
		{
			input:  []int{1, 2},
			expect: []int{1, 2},
		},
	}

	for i, tc := range testCases {
		head := createLinkedList(tc.input)
		ll.ReorderList(head)
		result := linkedListToSlice(head)
		fmt.Printf("Test case %d: input=%v\n", i+1, tc.input)
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
