package twoddp

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)

	var dfs func(iS1, iS2 int) bool
	dfs = func(iS1, iS2 int) bool {
		key := [2]int{iS1, iS2}
		if val, ok := dp[key]; ok {
			return val
		}

		if iS1 == len(s1) && iS2 == len(s2) {
			return true
		}

		res := false
		iS3 := iS1 + iS2

		if iS1 < len(s1) && s1[iS1] == s3[iS3] {
			res = res || dfs(iS1+1, iS2)
		}

		if iS2 < len(s2) && s2[iS2] == s3[iS3] {
			res = res || dfs(iS1, iS2+1)
		}

		dp[key] = res
		return dp[key]
	}

	return dfs(0, 0)
}
