package arraysandhashing

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num
}

func ScoreOfString(s string) int {
	score := 0

	for i := 1; i < len(s); i++ {
		score += abs(int(s[i]) - int(s[i-1]))
	}

	return score
}
