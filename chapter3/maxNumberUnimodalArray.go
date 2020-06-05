package main

// Problem 3.2
// How a unimodal array looks like
// [ascending, max, descending] or [max, descending] or [ascending, max]
// Compute the maximum element of a unimodal array

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxNumInUniModalList(nums []int) int {
	length := len(nums)
	halfLen := length / 2
	if length == 2 {
		return maxInt(nums[0], nums[1])
	}

	left := nums[halfLen-1]
	mid := nums[halfLen]
	right := nums[halfLen+1]

	if left <= mid && mid < right {
		return maxNumInUniModalList(nums[halfLen:length])
	} else if left < mid && mid >= right {
		return mid
	}
	return maxNumInUniModalList(nums[0:halfLen])
}
