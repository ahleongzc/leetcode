package interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Write a function to find the longest common prefix among an array of strings.
// input：strs = ["flower","flow","flight"] output："fl"

func TestLongestCommonPre(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output string
	}{
		{
			name: "example 1",
			input: []string{
				"flower",
				"flow",
				"fl",
			},
			output: "fl",
		},
		{
			name: "example 2",
			input: []string{
				"going",
				"gabv",
				"golang",
			},
			output: "g",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			answer := longestCommonPrefix(test.input)
			if !assert.Equal(t, test.output, answer) {
				t.Fatalf("expected %v, got %v", test.output, answer)
			}
		})
	}
}
