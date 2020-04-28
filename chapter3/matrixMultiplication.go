package main

func matrixMultiplication(a [][]int, b [][]int, n int) [][]int {
	var sum int
	c := make([][]int, n)
	for l := 0; l < n; l++ {
		c[l] = make([]int, n)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			sum = 0
			for j := 0; j < n; j++ {
				sum = sum + (a[k][j] * b[j][i])
			}
			c[k][i] = sum
		}
	}
	return c
}
