package main

import "math/rand"

func rselect(nums []int, nThSmallest int, beg int, end int) int {
	if beg >= end {
		return 0
	}
	index := beg + rand.Intn(end-beg)

	nums[beg], nums[index] = nums[index], nums[beg]
	pivot := nums[beg]
	i := beg + 1

	for j := beg + 1; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[beg], nums[i-1] = nums[i-1], nums[beg]
	position := i

	if position == nThSmallest {
		return nums[position-1]
	} else if nThSmallest < position {
		return rselect(nums, nThSmallest, beg, i-1)
	}
	return rselect(nums, nThSmallest, i, end)
}
