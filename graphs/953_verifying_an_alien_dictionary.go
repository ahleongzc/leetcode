package graphs

func isAlienSorted(words []string, order string) bool {
	orderList := make(map[byte]int)
	for i := range len(order) {
		orderList[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		curr := words[i-1]
		next := words[i]

		for index := range len(curr) {
			if index == len(next) {
				return false
			}

			if curr[index] != next[index] {
				if orderList[curr[index]] > orderList[next[index]] {
					return false
				}
				break
			}
		}
	}

	return true
}
