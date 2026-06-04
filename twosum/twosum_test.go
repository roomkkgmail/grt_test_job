package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Case 1: [2, 7, 11, 15], target 9",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Case 2: [3, 2, 4], target 6",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Case 3: [3, 3], target 6",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
