package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortLastElementPivot(t *testing.T) {
	input := []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted := []int{1, 2, 3, 4, 5, 6, 7, 8}
	quickSortLastElementPivot(input, 0, len(input))
	assert.Equal(t, wanted, input)
}

func TestChooseLastElementPivot(t *testing.T) {
	input := []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted := 6
	got := chooseLastElementPivot(input, 0, len(input))
	assert.Equal(t, wanted, got)

	input = []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted = 5
	got = chooseLastElementPivot(input, 2, 4)
	assert.Equal(t, wanted, got)
}
