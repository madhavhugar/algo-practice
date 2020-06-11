package main

func quickSortNaivePivot(nums []int, beg int, end int) {
	if beg >= end {
		return
	}
	pivot := nums[beg]
	i := beg + 1
	for j := beg + 1; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i-1], nums[beg] = nums[beg], nums[i-1]
	quickSortNaivePivot(nums, beg, i-1)
	quickSortNaivePivot(nums, i, end)
}
