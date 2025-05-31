package graphs

func FindTheTownJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	judgeMap := make(map[int]int)
	peopleMap := make(map[int]bool)

	for _, val := range trust {
		truster := val[0]
		trustee := val[1]

		judgeMap[trustee]++
		peopleMap[truster] = true
	}

	for potentialJudge, trustCount := range judgeMap {
		if _, ok := peopleMap[potentialJudge]; !ok && trustCount == n-1 {
			return potentialJudge
		}
	}

	return -1
}
