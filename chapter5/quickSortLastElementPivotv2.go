package main

func chooseLastElementPivotV2(nums []int, beg int, end int) (int, int) {
	return nums[end-1], end - 1
}

func quickSortLastElementPivotV2(nums []int, beg int, end int) {
	if beg >= end {
		return
	}
	pivot, currPivotIndex := chooseLastElementPivotV2(nums, beg, end)
	i := beg
	for j := beg; j < end-1; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	newPivotIndex := i
	nums[currPivotIndex], nums[newPivotIndex] = nums[newPivotIndex], nums[currPivotIndex]
	quickSortLastElementPivotV2(nums, beg, newPivotIndex)
	quickSortLastElementPivotV2(nums, newPivotIndex+1, end)
}
