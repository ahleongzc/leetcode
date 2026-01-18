package helper

import (
	"sort"
	"unicode"
)

func IsRuneAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func ReverseString(s string) string {
	runes := make([]rune, 0)

	for i := len(s) - 1; i >= 0; i-- {
		runes = append(runes, rune(s[i]))
	}

	return string(runes)
}

func RemoveRuneFromString(s string, i int) string {
	return s[:i] + s[i+1:]
}

func Normalize2DIntSlices(combos [][]int) {
	for _, c := range combos {
		sort.Ints(c)
	}

	sort.Slice(combos, func(i, j int) bool {
		a, b := combos[i], combos[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}
