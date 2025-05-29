package twopointers

func MergeAlternately(word1 string, word2 string) string {
	mergedString := make([]rune, 0)
	startWordOne := 0
	startWordTwo := 0

	for startWordOne < len(word1) && startWordTwo < len(word2) {
		mergedString = append(mergedString, rune(word1[startWordOne]))
		mergedString = append(mergedString, rune(word2[startWordTwo]))
		startWordOne++
		startWordTwo++
	}

	for startWordOne < len(word1) {
		mergedString = append(mergedString, rune(word1[startWordOne]))
		startWordOne++
	}

	for startWordTwo < len(word2) {
		mergedString = append(mergedString, rune(word2[startWordTwo]))
		startWordTwo++
	}

	return string(mergedString)
}