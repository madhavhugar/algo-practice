package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestRSelect(t *testing.T) {
	nums := []int{6, 8, 9, 2}
	wanted := 6
	nThSmallest := 2
	got := rselect(nums, nThSmallest, 0, len(nums))
	assert.Equal(t, wanted, got)

	nums = []int{1, 2, 3, 4, 5, 6, 8, 9}
	wanted = 6
	nThSmallest = 6
	got = rselect(nums, nThSmallest, 0, len(nums))
	assert.Equal(t, wanted, got)

	nums = []int{1, 2, 6, 4, 5, 3, 8, 9}
	wanted = 6
	nThSmallest = 6
	got = rselect(nums, nThSmallest, 0, len(nums))
	assert.Equal(t, wanted, got)
}
