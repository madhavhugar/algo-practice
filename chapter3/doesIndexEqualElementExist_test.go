package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestDoesIndexEqualElementExist(t *testing.T) {
	nums := []int{-1,0,2,5,8,9}
	gotExists, gotNum := doesIndexEqualElementExist(nums, 0, len(nums))
	wantedExists, wantedNum := true, 2
	assert.Equal(t, wantedExists, gotExists)
	assert.Equal(t, wantedNum, gotNum)

	nums = []int{-1,0,1,2,3,5}
	gotExists, gotNum = doesIndexEqualElementExist(nums, 0, len(nums))
	wantedExists, wantedNum = true, 5
	assert.Equal(t, wantedExists, gotExists)
	assert.Equal(t, wantedNum, gotNum)

	nums = []int{-1,0,1,2,3,6}
	gotExists, gotNum = doesIndexEqualElementExist(nums, 0, len(nums))
	wantedExists = false
	assert.Equal(t, wantedExists, gotExists)
}
