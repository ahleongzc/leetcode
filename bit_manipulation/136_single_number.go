package bitmanipulation

func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	single := nums[0]
	for i := 1; i < len(nums); i++ {
		single ^= nums[i]
	}

	return single
}
