package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name: "example 1",
			s:    "leetcode",
			wordDict: []string{
				"leet", "code",
			},
			expected: true,
		},
		{
			name: "example 2",
			s:    "applepenapple",
			wordDict: []string{
				"apple",
				"pen",
			},
			expected: true,
		},
		{
			name: "example 3",
			s:    "catsanddog",
			wordDict: []string{
				"cats",
				"dog",
				"sand",
				"and",
				"cat",
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := wordBreak(test.s, test.wordDict)

			if !assert.Equal(t, output, test.expected) {
				t.Fatalf(
					"test case %v - expected: %v, got %v", test.name, test.expected, output,
				)
			}
		})
	}
}
