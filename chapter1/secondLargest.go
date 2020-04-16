package main

import "errors"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLargestNumber(numbers []int, n int) int {
	if n == 1 {
		return numbers[0]
	}
	halfLen := n / 2
	b := findLargestNumber(numbers[0:halfLen], halfLen)
	c := findLargestNumber(numbers[halfLen:n], halfLen)

	return max(b, c)
}

func findSecondLargestNumber(numbers []int, n int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("Cannot pass empty numbers array")
	}
	largestNumber := findLargestNumber(numbers, n)
	nextLargest := numbers[0]

	for i := 1; i < n; i++ {
		if (numbers[i] > nextLargest) && (numbers[i] < largestNumber) {
			nextLargest = numbers[i]
		}
	}
	return nextLargest, nil
}

