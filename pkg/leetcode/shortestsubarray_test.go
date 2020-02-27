package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestsubarray(t *testing.T) {
	testCases := [][]int{
		{1, 1, 1},
	}

	for i := range testCases {
		assert.Equal(t, testCases[i][0],
			shortestSubArray(testCases[i][2:], testCases[i][1]))
	}
}
