package main

type matrixDimensions struct {
	x0 int
	x1 int
	y0 int
	y1 int
}

func initialize2DMatrixStrassen(x int, y int) [][]int {
	z := make([][]int, x)
	for i := 0; i < x; i++ {
		z[i] = make([]int, y)
	}
	return z
}

func addMatricesStrassen(a [][]int, b [][]int, n int) [][]int {
	c := initialize2DMatrixStrassen(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func substractMatricesStrassen(a [][]int, b [][]int, n int) [][]int {
	c := initialize2DMatrixStrassen(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

func copyValuesIntoSubMatrix(actual *[][]int, sub *[][]int, dim *matrixDimensions) {
	for i, m := 0, dim.x0; m < dim.x1; i, m = i+1, m+1 {
		for j, n := 0, dim.y0; n < dim.y1; j, n = j+1, n+1 {
			(*sub)[i][j] = (*actual)[m][n]
		}
	}
}

func splitIntoSubMatricesStrassen(z [][]int, n int) ([][]int, [][]int, [][]int, [][]int) {
	a := initialize2DMatrixStrassen(n/2, n/2)
	b := initialize2DMatrixStrassen(n/2, n/2)
	c := initialize2DMatrixStrassen(n/2, n/2)
	d := initialize2DMatrixStrassen(n/2, n/2)
	copyValuesIntoSubMatrix(&z, &a, &matrixDimensions{x0: 0, x1: n / 2, y0: 0, y1: n / 2})
	copyValuesIntoSubMatrix(&z, &b, &matrixDimensions{x0: 0, x1: n / 2, y0: n / 2, y1: n})
	copyValuesIntoSubMatrix(&z, &c, &matrixDimensions{x0: n / 2, x1: n, y0: 0, y1: n / 2})
	copyValuesIntoSubMatrix(&z, &d, &matrixDimensions{x0: n / 2, x1: n, y0: n / 2, y1: n})
	return a, b, c, d
}

func copyValuesIntoParentMatrixStrassen(parent *[][]int, sub *[][]int, dim *matrixDimensions) {
	for i, m := 0, dim.x0; m < dim.x1; i, m = i+1, m+1 {
		for j, n := 0, dim.y0; n < dim.y1; j, n = j+1, n+1 {
			(*parent)[m][n] = (*sub)[i][j]
		}
	}
}

func combineSubMatricesIntoMatrixStrassen(a [][]int, b [][]int, c [][]int, d [][]int, n int) [][]int {
	z := initialize2DMatrixStrassen(n, n)
	copyValuesIntoParentMatrixStrassen(&z, &a, &matrixDimensions{x0: 0, x1: n / 2, y0: 0, y1: n / 2})
	copyValuesIntoParentMatrixStrassen(&z, &b, &matrixDimensions{x0: 0, x1: n / 2, y0: n / 2, y1: n})
	copyValuesIntoParentMatrixStrassen(&z, &c, &matrixDimensions{x0: n / 2, x1: n, y0: 0, y1: n / 2})
	copyValuesIntoParentMatrixStrassen(&z, &d, &matrixDimensions{x0: n / 2, x1: n, y0: n / 2, y1: n})
	return z
}

func strassenMatrixMultiplication(x [][]int, y [][]int, n int) [][]int {
	// Base condition
	if n == 1 {
		z := initialize2DMatrixStrassen(1, 1)
		z[0][0] = x[0][0] * y[0][0]
		return z
	}
	// split x into sub-matrices
	a, b, c, d := splitIntoSubMatricesStrassen(x, n)
	// split y into sub-matrices
	e, f, g, h := splitIntoSubMatricesStrassen(y, n)

	// selectively multiply according to strassen's formula
	//	P1 = A · (F - H)
	fh := substractMatricesStrassen(f, h, n/2)
	p1 := strassenMatrixMultiplication(a, fh, n/2)
	// P2 = (A + B) · H
	ab := addMatricesStrassen(a, b, n/2)
	p2 := strassenMatrixMultiplication(ab, h, n/2)
	// P3 = (C + D) · E
	cd := addMatricesStrassen(c, d, n/2)
	p3 := strassenMatrixMultiplication(cd, e, n/2)
	// P4 = D · (G - E)
	ge := substractMatricesStrassen(g, e, n/2)
	p4 := strassenMatrixMultiplication(d, ge, n/2)
	// P5 = (A + D) · (E + H)
	ad := addMatricesStrassen(a, d, n/2)
	eh := addMatricesStrassen(e, h, n/2)
	p5 := strassenMatrixMultiplication(ad, eh, n/2)
	// P6 = (B - D) . (G + H)
	bd := substractMatricesStrassen(b, d, n/2)
	gh := addMatricesStrassen(g, h, n/2)
	p6 := strassenMatrixMultiplication(bd, gh, n/2)
	// P7 = (A - C) · (E + F)
	ac := substractMatricesStrassen(a, c, n/2)
	ef := addMatricesStrassen(e, f, n/2)
	p7 := strassenMatrixMultiplication(ac, ef, n/2)

	// Merge the sub-matrices
	/*
		P5 + P4 - P2 + P6 	| 			P1 + P2
		----------------------------------------------
					P3 + P4 			| P1 + P5 - P3 - P7
	*/
	firstQuadrant := substractMatricesStrassen(addMatricesStrassen(addMatricesStrassen(p5, p4, n/2), p6, n/2), p2, n/2)
	secondQuadrant := addMatricesStrassen(p1, p2, n/2)
	thirdQuadrant := addMatricesStrassen(p3, p4, n/2)
	fourthQuadrant := substractMatricesStrassen(addMatricesStrassen(p1, p5, n/2), addMatricesStrassen(p3, p7, n/2), n/2)
	res := combineSubMatricesIntoMatrixStrassen(firstQuadrant, secondQuadrant, thirdQuadrant, fourthQuadrant, n)
	return res
}
