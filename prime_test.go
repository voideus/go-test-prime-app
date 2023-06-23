package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	input := 3

	result := IsPrime(input)

	assert.True(t, result)
}

func Test_IsPrimeTableTest(t *testing.T) {
	testCases := []struct {
		number   int
		expected bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{11, true},
		{1, false},
		{4, false},
		{9, false},
		{15, false},
	}

	for _, tc := range testCases {
		result := IsPrime(tc.number)
		assert.Equal(t, tc.expected, result)
	}
}
