package main

type matrixDimension struct {
	x0 int;
	x1 int;
	y0 int; 
	y1 int;
}

func initialize2DMatrix(x int, y int) [][]int {
	s := make([][]int, x)
	for l := 0; l < x; l++ {
		s[l] = make([]int, y)
	}
	return s
}

func createSubMatrix(original [][]int, new *[][]int, dim *matrixDimension) {
	for i, k := dim.x0, 0; i < dim.x1; i, k = i+1, k+1 {
		for j, l := dim.y0, 0; j < dim.y1; j, l = j+1, l+1 {
			(*new)[k][l] = original[i][j]
		}
	}
}

func splitIntoSubMatrices(z [][]int, n int) ([][]int, [][]int, [][]int, [][]int) {
	a := initialize2DMatrix(n/2, n/2)
	b := initialize2DMatrix(n/2, n/2)
	c := initialize2DMatrix(n/2, n/2)
	d := initialize2DMatrix(n/2, n/2)
	createSubMatrix(z, &a, &matrixDimension{x0: 0, x1: n/2, y0: 0, y1: n/2})
	createSubMatrix(z, &b, &matrixDimension{x0: 0, x1: n/2, y0: n/2, y1: n})
	createSubMatrix(z, &c, &matrixDimension{x0: n/2, x1: n, y0: 0, y1: n/2})
	createSubMatrix(z, &d, &matrixDimension{x0: n/2, x1: n, y0: n/2, y1: n})
	return a, b, c, d
}

func addMatrices(a [][]int, b [][]int, n int) [][]int {
	c := initialize2DMatrix(n, n)
	for i := 0; i < n; i ++ {
		for j := 0; j < n; j ++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func copySubMatrix(combined *[][]int, sub [][]int, dim *matrixDimension) {
	for i, k := dim.x0, 0; i < dim.x1; i, k = i+1, k+1 {
		for j, l := dim.y0, 0; j < dim.y1; j, l = j+1, l+1 {
			(*combined)[i][j] = sub[k][l]
		}
	}
}

func combineSubMatrices(a [][]int, b [][]int, c [][]int, d [][]int, n int) [][]int {
	z := initialize2DMatrix(n, n)
	copySubMatrix(&z, a, &matrixDimension{x0: 0, x1: n/2, y0: 0, y1: n/2})
	copySubMatrix(&z, b, &matrixDimension{x0: 0, x1: n/2, y0: n/2, y1: n})
	copySubMatrix(&z, c, &matrixDimension{x0: n/2, x1: n, y0: 0, y1: n/2})
	copySubMatrix(&z, d, &matrixDimension{x0: n/2, x1: n, y0: n/2, y1: n})
	return z
}

func recursiveMatrixMultiplication(x [][]int, y [][]int, n int) [][]int {
	if n == 1 {
		z := initialize2DMatrix(n, n)
		z[0][0] = x[0][0] * y[0][0]
		return z
	}
	a, b, c, d := splitIntoSubMatrices(x, n)
	e, f, g, h := splitIntoSubMatrices(y, n)
	ae := recursiveMatrixMultiplication(a, e, n/2)
	bg := recursiveMatrixMultiplication(b, g, n/2)
	af := recursiveMatrixMultiplication(a, f, n/2)
	bh := recursiveMatrixMultiplication(b, h, n/2)
	ce := recursiveMatrixMultiplication(c, e, n/2)
	dg := recursiveMatrixMultiplication(d, g, n/2)
	cf := recursiveMatrixMultiplication(c, f, n/2)
	dh := recursiveMatrixMultiplication(d, h, n/2)

	aebg := addMatrices(ae, bg, n/2)
	afbh := addMatrices(af, bh, n/2)
	cedg := addMatrices(ce, dg, n/2)
	cfdh := addMatrices(cf, dh, n/2)
	return combineSubMatrices(aebg, afbh, cedg, cfdh, n)
}
