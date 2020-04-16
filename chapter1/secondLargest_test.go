package main

import "testing"

func TestSecondLargest(t *testing.T) {
	// Regular
	numbers := []int{8, 1, 3, 9, 2, 5, 4, 6}
	got, e := findSecondLargestNumber(numbers, 8)
	if got != 8 {
		t.Errorf("findSecondLargestNumber(numbers) == %d; wanted 8", got)
	}

	// Empty array
	numbers = []int{}
	got, e = findSecondLargestNumber(numbers, 0)
	if e == nil {
		t.Errorf("findSecondLargestNumber(numbers) == %d; wanted nil", got)
	}
	
	// Array containing duplicate elements
	numbers = []int{8, 1, 3, 9, 2, 9, 4, 8}
	got, e = findSecondLargestNumber(numbers, 8)
	if got != 8 {
		t.Errorf("findSecondLargestNumber(numbers) == %d; wanted 8", got)
	}
}
