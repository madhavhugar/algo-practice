package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSelectMergeSort(t *testing.T) {
	nums := []int{6, 8, 9, 2}
	wanted := 6
	nThSmallest := 2
	got, err := rselectMergeSort(nums, nThSmallest)
	assert.Equal(t, wanted, got)
	assert.Equal(t, nil, err)

	nums = []int{6, 8, 9, 2}
	wantedErr := errors.New("Array smaller than requested nTh smallest number")
	nThSmallest = 8
	_, err = rselectMergeSort(nums, nThSmallest)
	assert.Equal(t, wantedErr, err)
}

func TestMergeSortForRSelect(t *testing.T) {
	nums := []int{4, 2, 7, 3, 8, 1, 5, 6}
	wanted := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := mergeSortForRSelect(nums)
	assert.Equal(t, wanted, got)
}

func TestMerger(t *testing.T) {
	a := []int{1, 3, 6}
	b := []int{2, 4, 5}
	wanted := []int{1, 2, 3, 4, 5, 6}
	got := merger(a, b)
	assert.Equal(t, wanted, got)
}
