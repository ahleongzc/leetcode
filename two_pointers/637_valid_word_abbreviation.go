package twopointers

import (
	"strconv"
	"unicode"
)

func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}

func ValidWordAbbreviation(word string, abbr string) bool {
	ptrForWord := 0
	ptrForAbbr := 0

	for ptrForAbbr < len(abbr) && ptrForWord < len(word) {
		moveCount := 0
		endPtrForAbbr := ptrForAbbr
		for endPtrForAbbr < len(abbr) && isNumber(rune(abbr[endPtrForAbbr])) {
			if abbr[endPtrForAbbr] == byte('0') && endPtrForAbbr == ptrForAbbr {
				return false
			}
			endPtrForAbbr++
		}

		if endPtrForAbbr != ptrForAbbr {
			moveCount, _ = strconv.Atoi(abbr[ptrForAbbr:endPtrForAbbr])
			ptrForAbbr = endPtrForAbbr
		}

		if moveCount != 0 {
			ptrForWord += moveCount
			continue
		}

		if word[ptrForWord] != abbr[ptrForAbbr] {
			return false
		} else {
			ptrForWord++
			ptrForAbbr++
		}
	}

	return ptrForAbbr == len(abbr) && ptrForWord == len(word)
}
