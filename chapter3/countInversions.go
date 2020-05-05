package main

func countInversions(a []int, len int) ([]int, int) {
	if len == 1 {
		return []int{a[0]}, 0
	}
	x, cx := countInversions(a[0:len/2], len/2)
	y, cy := countInversions(a[len/2:len], len/2)

	m, c := mergeLists(x, y, len)
	return m, (c + cx + cy)
}

func sliceListIntoTwo(original []int, len int) ([]int, []int) {
	return original[0 : len/2], original[len/2 : len]
}

func mergeLists(a []int, b []int, len int) ([]int, int) {
	invCount := 0
	i, j := 0, 0
	c := make([]int, len)
	for k := 0; k < len; k++ {
		if (j >= len/2) || ((i < len/2) && (a[i] < b[j])) {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
			invCount = invCount + (len / 2) - i
		}
	}
	return c, invCount
}
