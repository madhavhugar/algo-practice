package main

import "errors"

func merger(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))
	i := 0
	j := 0
	for k := 0; k < len(c); k++ {
		if i >= len(a) {
			c[k] = b[j]
			j++
		} else if j >= len(b) {
			c[k] = a[i]
			i++
		} else if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else if b[j] < a[i] {
			c[k] = b[j]
			j++
		}
	}
	return c
}

func mergeSortForRSelect(nums []int) []int {
	// base case
	length := len(nums)
	if length == 1 {
		return nums
	}
	halfLen := length / 2
	a := mergeSortForRSelect(nums[:halfLen])
	b := mergeSortForRSelect(nums[halfLen:])

	return merger(a, b)
}

func rselectMergeSort(nums []int, nThSmallest int) (int, error) {
	if nThSmallest > len(nums) {
		return 0, errors.New("Array smaller than requested nTh smallest number")
	}
	sorted := mergeSortForRSelect(nums)
	return sorted[nThSmallest-1], nil
}
