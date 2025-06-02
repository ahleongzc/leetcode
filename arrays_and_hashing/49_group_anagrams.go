package arraysandhashing

import (
	"sort"
)

// The time complexity is O(m * n) where m is the number of words and n is the average length of each word
func GroupAnagramsOptimised(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[[26]int][]string)

	for _, str := range strs {
		arr := [26]int{}
		for _, c := range str {
			arr[c%26]++
		}

		m[arr] = append(m[arr], str)
	}

	for _, strs := range m {
		res = append(res, strs)
	}

	return res
}

// The time complexity is O(m * nlogn) where m is the number of words and n is the average length of each word
func GroupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]int)

	for i, str := range strs {
		runeCopy := []rune(str)
		sort.Slice(runeCopy, func(i, j int) bool { return runeCopy[i] < runeCopy[j] })
		stringCopy := string(runeCopy)

		m[stringCopy] = append(m[stringCopy], i)
	}

	for _, indexes := range m {
		arr := make([]string, 0)
		for _, index := range indexes {
			arr = append(arr, strs[index])
		}
		res = append(res, arr)
	}

	return res
}
