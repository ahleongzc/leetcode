package twopointers

func ReverseString(s []byte) {
	start := 0
	end := len(s) - 1

	for start < end {
		startCopy := s[start]
		s[start] = s[end]
		s[end] = startCopy

		start++
		end--
	}
}
