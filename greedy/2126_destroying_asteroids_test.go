package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidsDestroyed(t *testing.T) {
	tests := []struct {
		name     string
		mass     int
		asteroid []int
		expected bool
	}{
		{
			name: "ex 1",
			mass: 10,
			asteroid: []int{
				3, 9, 19, 5, 21,
			},
			expected: true,
		},
		{
			name: "ex 2",
			mass: 5,
			asteroid: []int{
				4, 9, 23, 4,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := asteroidsDestroyed(test.mass, test.asteroid)
		if !assert.Equal(t, test.expected, actual) {
			t.Fatal(test.name)
		}
	}
}
