package main

import (
	"sort"
)

func quickSortMedianPivot(nums []int, beg int, end int) {
	if beg >= end {
		return
	}
	pivot := getMedianPivot(nums, beg, end)
	i := beg + 1

	for j := beg + 1; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[beg], nums[i-1] = nums[i-1], nums[beg]

	quickSortMedianPivot(nums, beg, i-1)
	quickSortMedianPivot(nums, i, end)
}

func getMedianPivot(nums []int, beg int, end int) int {
	mid := (end - beg) / 2
	values := []int{nums[beg], nums[mid], nums[end-1]}
	sort.Ints(values)
	median := values[1]

	if median == nums[mid] {
		nums[beg], nums[mid] = nums[mid], nums[beg]
	} else if median == nums[end-1] {
		nums[beg], nums[end-1] = nums[end-1], nums[beg]
	}
	return nums[beg]
}
