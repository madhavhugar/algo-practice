package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestQuickSortNaivePivot(t *testing.T) {
	nums := []int{}
	quickSortNaivePivot(nums, 0, 0)
	wanted := []int{}
	assert.Equal(t, wanted, nums)

	nums = []int{3}
	quickSortNaivePivot(nums, 0, 1)
	wanted = []int{3}
	assert.Equal(t, wanted, nums)

	nums = []int{3, 8, 2, 5, 1, 4, 7, 6}
	quickSortNaivePivot(nums, 0, len(nums))
	wanted = []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, wanted, nums)
}
