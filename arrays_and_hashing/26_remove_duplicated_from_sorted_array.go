package arraysandhashing

func RemoveDuplicates(nums []int) int {
	left, right := 0, 0
	for right < len(nums) {
		nums[left] = nums[right]
		for right < len(nums) && nums[right] == nums[left] {
			right++
		}
		left++
	}

	return left
}
