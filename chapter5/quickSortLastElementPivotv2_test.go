package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortLastElementPivotV2(t *testing.T) {
	input := []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted := []int{1, 2, 3, 4, 5, 6, 7, 8}
	quickSortLastElementPivotV2(input, 0, len(input))
	assert.Equal(t, wanted, input)
}

func TestChooseLastElementPivotV2(t *testing.T) {
	input := []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted, index := 6, 7
	got, pivotIndex := chooseLastElementPivotV2(input, 0, len(input))
	assert.Equal(t, wanted, got)
	assert.Equal(t, pivotIndex, index)

	input = []int{3, 8, 2, 5, 1, 4, 7, 6}
	wanted, index = 5, 3
	got, pivotIndex = chooseLastElementPivotV2(input, 2, 4)
	assert.Equal(t, wanted, got)
	assert.Equal(t, pivotIndex, index)
}
