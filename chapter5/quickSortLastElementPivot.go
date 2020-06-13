package main

func quickSortLastElementPivot(nums []int, beg int, end int) {
	if beg >= end {
		return
	}
	pivot := chooseLastElementPivot(nums, beg, end)
	i := beg + 1
	for j := beg + 1; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	pivotIndex := i - 1
	nums[beg], nums[pivotIndex] = nums[pivotIndex], nums[beg]
	quickSortLastElementPivot(nums, beg, pivotIndex)
	quickSortLastElementPivot(nums, pivotIndex+1, end)
}

func chooseLastElementPivot(nums []int, beg int, end int) int {
	nums[beg], nums[end-1] = nums[end-1], nums[beg]
	return nums[beg]
}
