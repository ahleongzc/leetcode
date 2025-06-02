package arraysandhashing

func TwoSumOnePass(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		t := target - num

		if val, ok := numMap[t]; ok && val != i {
			return []int{i, val}
		}
		numMap[num] = i
	}

	return []int{}
}

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}

	for i, num := range nums {
		t := target - num
		if val, ok := numMap[t]; ok && val != i {
			return []int{i, val}
		}
	}

	return []int{}
}
