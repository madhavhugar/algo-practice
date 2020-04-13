package main

/*
	Assumptions:
	- array of length is a power of 2
*/

func mergeSort(numbers []int, beg int, end int) []int {
	len := end - beg
	// Special case where input slice is empty
	if len == 0 {
		return numbers
	}
	if len == 1 {
		return numbers[beg:end]
	}
	halfLen := len / 2
	c := mergeSort(numbers, beg, beg+halfLen)
	d := mergeSort(numbers, beg+halfLen, end)
	return merge(c, d, len)
}

func merge(a []int, b []int, len int) []int {
	var i, j int = 0, 0
	var merged []int
	for k := 0; k < len; k++ {
		if (j >= len/2) || ((i < len/2) && (a[i] < b[j])) {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}
	return merged
}
