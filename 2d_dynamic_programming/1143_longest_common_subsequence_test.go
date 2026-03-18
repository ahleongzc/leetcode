package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name, text1, text2 string
		expected           int
	}{
		{
			name:     "example 1",
			text1:    "cat",
			text2:    "crabt",
			expected: 3,
		},
		{
			name:     "example 2",
			text1:    "abcd",
			text2:    "abcd",
			expected: 4,
		},
		{
			name:     "example 3",
			text1:    "abcd",
			text2:    "efgh",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := longestCommonSubsequence(test.text1, test.text2)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
