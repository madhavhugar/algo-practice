package main

// Problem 3.3
// In a sorted array. Decide whether or not there is an index i such that A[i] = i

func doesIndexEqualElementExist(nums []int, startIndex int, endIndex int) (bool, int) {
	length := len(nums)
	halfLen := length / 2

	if length == 1 {
		if startIndex == nums[0] {
			return true, nums[0]
		}
		return false, nums[0]
	}

	midIndex := startIndex + (endIndex-startIndex)/2

	if nums[halfLen] > midIndex {
		return doesIndexEqualElementExist(nums[0:halfLen], startIndex, startIndex+halfLen)
	} else if nums[halfLen] < midIndex {
		return doesIndexEqualElementExist(nums[halfLen:length], startIndex+halfLen, endIndex)
	}
	return true, nums[halfLen]
}
