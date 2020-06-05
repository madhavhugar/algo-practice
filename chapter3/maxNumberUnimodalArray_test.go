package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestMaxNumInUniModalList(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2}
	got := maxNumInUniModalList(numbers)
	wanted := 6
	assert.Equal(t, wanted, got)

	numbers = []int{1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 2}
	got = maxNumInUniModalList(numbers)
	wanted = 6
	assert.Equal(t, wanted, got)

	numbers = []int{1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 6, 5, 4, 3, 2}
	got = maxNumInUniModalList(numbers)
	wanted = 6
	assert.Equal(t, wanted, got)

	numbers = []int{1, 2, 3, 4, 6, 5, 5, 5, 5, 5, 4, 3, 2}
	got = maxNumInUniModalList(numbers)
	wanted = 6
	assert.Equal(t, wanted, got)

	numbers = []int{6, 5, 5, 5, 5, 5, 4, 3, 2}
	got = maxNumInUniModalList(numbers)
	wanted = 6
	assert.Equal(t, wanted, got)

	numbers = []int{1, 2, 3, 4, 6}
	got = maxNumInUniModalList(numbers)
	wanted = 6
	assert.Equal(t, wanted, got)
}

func TestMaxInt(t *testing.T) {
	got := maxInt(3, 4)
	wanted := 4
	assert.Equal(t, wanted, got)

	got = maxInt(4, 4)
	wanted = 4
	assert.Equal(t, wanted, got)
}
