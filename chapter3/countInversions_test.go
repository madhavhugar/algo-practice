package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestCountingInversions(t *testing.T) {
	given, len := []int{1, 2, 3, 4, 5, 6, 7, 8}, 8
	_, gotCount := countInversions(given, len)
	wanted := 0
	assert.Equal(t, wanted, gotCount)

	given, len = []int{1, 3, 5, 2, 4, 6, 7, 8}, 8
	_, gotCount = countInversions(given, len)
	wanted = 3
	assert.Equal(t, wanted, gotCount)

	given, len = []int{8}, 1
	gotList, gotCount := countInversions(given, len)
	wantedList, wantedCount := []int{8}, 0
	assert.Equal(t, wantedList, gotList)
	assert.Equal(t, wantedCount, gotCount)

}

func TestSliceListIntoTwo(t *testing.T) {
	given := []int{1, 2, 3, 4, 5, 6, 7, 8}
	gotA, gotB := sliceListIntoTwo(given, 8)
	wantedA, wantedB := []int{1, 2, 3, 4}, []int{5, 6, 7, 8}
	assert.Equal(t, gotA, wantedA)
	assert.Equal(t, gotB, wantedB)

	given = []int{1, 2}
	gotA, gotB = sliceListIntoTwo(given, 2)
	wantedA, wantedB = []int{1}, []int{2}
	assert.Equal(t, gotA, wantedA)
	assert.Equal(t, gotB, wantedB)
}

func TestMergeLists(t *testing.T) {
	givenA, givenB := []int{1, 2, 3, 4}, []int{5, 6, 7, 8}
	gotList, gotCount := mergeLists(givenA, givenB, 8)
	wantedList, wantedCount := []int{1, 2, 3, 4, 5, 6, 7, 8}, 0
	assert.Equal(t, gotList, wantedList)
	assert.Equal(t, gotCount, wantedCount)

	givenA, givenB = []int{1, 2, 4}, []int{0, 3, 8}
	gotList, gotCount = mergeLists(givenA, givenB, 6)
	wantedList, wantedCount = []int{0, 1, 2, 3, 4, 8}, 4
	assert.Equal(t, gotList, wantedList)
	assert.Equal(t, gotCount, wantedCount)
}
