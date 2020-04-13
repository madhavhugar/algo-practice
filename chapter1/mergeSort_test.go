package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	// Regular
	numbers := []int{6, 3, 8, 1, 9, 15, 4, 2}
	wanted := []int{1, 2, 3, 4, 6, 8, 9, 15}
	sorted := mergeSort(numbers, 0, 8)
	assert.Equal(t, sorted, wanted)

	// Empty array
	numbers = []int{}
	wanted = []int{}
	sorted = mergeSort(numbers, 0, 0)
	assert.Equal(t, sorted, wanted)

	// Array with same elements
	numbers = []int{5, 5, 5, 5, 5, 5, 5, 5}
	wanted = []int{5, 5, 5, 5, 5, 5, 5, 5}
	sorted = mergeSort(numbers, 0, 8)
	assert.Equal(t, sorted, wanted)

	// Reverse sorted array
	numbers = []int{8, 7, 6, 5, 4, 3, 2, 1}
	wanted = []int{1, 2, 3, 4, 5, 6, 7, 8}
	sorted = mergeSort(numbers, 0, 8)
	assert.Equal(t, sorted, wanted)

	// Array containing duplicate elements
	numbers = []int{8, 7, 6, 5, 5, 3, 2, 3}
	wanted = []int{2, 3, 3, 5, 5, 6, 7, 8}
	sorted = mergeSort(numbers, 0, 8)
	assert.Equal(t, sorted, wanted)
}
