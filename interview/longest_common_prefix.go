package interview

func longestCommonPrefix(input []string) string {
	// O(k) where k is the smallest length of the words
	output := ""

	iterations := len(input[0])
	// O(n)
	for _, s := range input {
		if len(s) < iterations {
			iterations = len(s)
		}
	}
	
	// O(k) where k is the smallest length
	for index := range iterations {
		curr := input[0][index]
		same := true
		
		// O(n)
		for i := 1; i < len(input); i++ {
			currWord := input[i][index]
			if currWord != curr {
				return output
			}
		}

		if same {
			output += string(curr)
		}
	}
	
	// O(kn)
	return output
}
