package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{numbers: []int{
			5, 2, 1, 3, 4},
			target:   9,
			expected: []int{0, 4}},
	}
	// Act
	for _, testCase := range testTable {
		result := twoSum(testCase.numbers, testCase.target)

		t.Logf("Calling twoSum(%v), result %d\n", testCase.numbers, result)

		// Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result.  Expect %d, got %d", testCase.expected, result),
		)
	}

}
