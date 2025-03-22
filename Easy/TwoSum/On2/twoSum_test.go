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
		expected [][]int
	}{
		{
			numbers:  []int{5, 2, 1, 3, 4},
			target:   9,
			expected: [][]int{{0, 4}, {4, 0}},
		},
		{
			numbers:  []int{5, 5, 1, 2},
			target:   10,
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			numbers:  []int{0, 1, 2, 4, 4},
			target:   4,
			expected: [][]int{{0, 3}, {3, 0}},
		},
		{
			numbers:  []int{0, 1, 2, 3, 4},
			target:   100,
			expected: [][]int{{}},
		},
	}
	// Act
	for _, testCase := range testTable {
		result := twoSum(testCase.numbers, testCase.target)

		t.Logf("Calling twoSum(%v), result %d\n", testCase.numbers, result)

		// Assert
		assert.Contains(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result.  Expect %d, got %d", testCase.expected, result),
		)
	}

}

func TestTwoSumAlternate(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		target   int
		expected [][]int
	}{
		{
			numbers:  []int{5, 2, 1, 3, 4},
			target:   9,
			expected: [][]int{{0, 4}, {4, 0}},
		},
		{
			numbers:  []int{5, 5, 1, 2},
			target:   10,
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			numbers:  []int{0, 1, 2, 4, 4},
			target:   4,
			expected: [][]int{{0, 3}, {3, 0}},
		},
		{
			numbers:  []int{0, 1, 2, 3, 4},
			target:   100,
			expected: [][]int{{}},
		},
	}
	// Act
	for _, testCase := range testTable {
		result := twoSumAlternate(testCase.numbers, testCase.target)

		t.Logf("Calling twoSumAlternate(%v), result %d\n", testCase.numbers, result)

		// Assert
		assert.Contains(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result.  Expect %d, got %d", testCase.expected, result),
		)
	}

}
