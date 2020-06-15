package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestQuickSortMedianPivot(t *testing.T) {
	input := []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted := []int{1, 2, 3, 4, 5, 6, 7, 8}
	quickSortMedianPivot(input, 0, len(input))
	assert.Equal(t, wanted, input)

	input = []int{3, 8, 2, 5, 1, 5, 5, 6}
	wanted = []int{1, 2, 3, 5, 5, 5, 6, 8}
	quickSortMedianPivot(input, 0, len(input))
	assert.Equal(t, wanted, input)
}

func TestGetMedianPivot(t *testing.T) {
	input := []int{3, 1, 2}
	wanted := 2
	got := getMedianPivot(input, 0, len(input))
	assert.Equal(t, wanted, got)
}
