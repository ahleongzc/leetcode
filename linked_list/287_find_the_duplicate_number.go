package linkedlist

func findDuplicate(nums []int) int {
	// 1. find the intersection between fast and slow
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	// 2. find the p = x portion
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}
}
