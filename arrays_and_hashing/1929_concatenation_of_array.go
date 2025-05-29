package arraysandhashing

func GetConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)

	for i, num := range nums {
		res[i] = num
		res[i+len(nums)] = num
	}

	return res
}
