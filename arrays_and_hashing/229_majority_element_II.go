package arraysandhashing

func MajorityElementII(nums []int) []int {
	majorityElements := make([]int, 0)
	voting := make(map[int]int)

	for _, num := range nums {
		voting[num]++

		if len(voting) <= 2 {
			continue
		}

		// The reason for this is because you don't want to modify the hashmap while iterating over it
		newVoting := make(map[int]int)
		for candidate, voteCount := range voting {
			if voteCount > 1 {
				newVoting[candidate] = voteCount - 1
			}
		}

		voting = newVoting
	}

	// This would be O(2n) because voting has at most two elements
	for candidate := range voting {
		count := 0

		for _, num := range nums {
			if num == candidate {
				count++
			}
		}

		if count > len(nums)/3 {
			majorityElements = append(majorityElements, candidate)
		}

	}

	return majorityElements
}
