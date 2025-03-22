package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	// Arrange

	testTable := []struct {
		numbers  []int
		expected int
	}{
		{numbers: []int{1, 2, 4, 4, 0, -120, -0, 2}, expected: 4},
		{numbers: []int{0, 0, 0, 0, 0, 0, 0}, expected: 0},
		{numbers: []int{120, 120, 120, 120, 999999999999}, expected: 999999999999},
		{numbers: []int{-1054, -12030, -1220, -0, -1, 1}, expected: 1},
	}

	// Act
	for _, testCase := range testTable {
		result := Max(testCase.numbers)

		t.Logf("Calling Max(%v), result %d\n", testCase.numbers, result)

		// Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result.  Expect %d, got %d", testCase.expected, result),
		)
	}
}

func TestMaxIndex(t *testing.T) {
	// Arrange
	numbers := []int{1, 2, 4, 5, 6, 0, -120, -1, 40, -0}
	expected := 8

	// Act
	result := MaxIndex(numbers)

	// Assert
	if result != expected {
		t.Errorf("incorrect result. Expect %d, got %d", expected, result)
	}
}
