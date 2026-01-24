package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		order    string
		expected bool
	}{
		{
			name: "example 1",
			words: []string{
				"hello",
				"leetcode",
			},
			order:    "hlabcdefgijkmnopqrstuvwxyz",
			expected: true,
		},
		{
			name: "example 2",
			words: []string{
				"word",
				"world",
				"row",
			},
			order:    "worldabcefghijkmnpqstuvxyz",
			expected: false,
		},
		{
			name: "example 3",
			words: []string{
				"apple",
				"app",
			},
			order:    "abcdefghijklmnopqrstuvwxyz",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isAlienSorted(test.words, test.order)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf("test case: %v - expected: %v, got: %v", test.name, test.expected, output)
			}
		})
	}
}
