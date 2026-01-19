package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourseSchedu(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:       "example 1",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			expected: true,
		},
		{
			name:       "example 2",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			expected: false,
		},
		{
			name:       "example 3",
			numCourses: 5,
			prerequisites: [][]int{
				{1, 4}, {2, 4}, {3, 1}, {3, 2},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := canFinish(test.numCourses, test.prerequisites)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf(
					"test case: %v - expected: %v, got :%v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
