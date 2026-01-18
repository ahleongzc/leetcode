package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:   "example 1",
			digits: "23",
			expected: []string{
				"ad",
				"ae",
				"af",
				"bd",
				"be",
				"bf",
				"cd",
				"ce",
				"cf",
			},
		},
		{
			name:   "example 2",
			digits: "2",
			expected: []string{
				"a",
				"b",
				"c",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := letterCombinations(test.digits)
			if !assert.ElementsMatch(t, output, test.expected) {
				t.Fatalf(
					"test case: %v - expected %v, got %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
