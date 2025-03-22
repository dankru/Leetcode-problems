package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// Arrange
	testTable := []struct {
		number   int32
		expected bool
	}{
		{
			number:   1,
			expected: true,
		},
		{
			number:   121,
			expected: true,
		},
		{
			number:   -121,
			expected: false,
		},
		{
			number:   10,
			expected: false,
		},
		{
			number:   1001,
			expected: true,
		},
		{
			number:   1000,
			expected: false,
		},
		{
			number:   10000,
			expected: false,
		},
	}
	// Act
	for _, testCase := range testTable {
		result := isPalindrome(testCase.number)

		t.Logf("Calling is Palindrome(%v) result: %v", testCase.number, result)

		// Assert
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Unexpected result: %v, expected: %v", result, testCase.expected))
	}
}
