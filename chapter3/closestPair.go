package main

type Point struct {
	x float64
	y float64
}

func mergeSortPoints(points []Point, mergeFunc func([]Point, []Point) []Point) []Point {
	if len(points) == 1 {
		return points
	}
	a, b := splitPointsList(points)
	c := mergeSortPoints(a, mergeFunc)
	d := mergeSortPoints(b, mergeFunc)

	return mergeFunc(c, d)
}

func splitPointsList(points []Point) ([]Point, []Point) {
	halfLen := len(points) / 2
	return points[:halfLen], points[halfLen:]
}

func mergeListOfPointsSortedByXCoordinate(a []Point, b []Point) []Point {
	halfLen := len(a)
	length := halfLen * 2
	c := make([]Point, length)
	j := 0
	k := 0

	for i := 0; i < length; i++ {
		if j >= halfLen {
			c[i] = b[k]
			k++
		} else if k >= halfLen {
			c[i] = a[j]
			j++
		} else if a[j].x < b[k].x {
			c[i] = a[j]
			j++
		} else {
			c[i] = b[k]
			k++
		}
	}
	return c
}

func mergeListOfPointsSortedByYCoordinate(a []Point, b []Point) []Point {
	halfLen := len(a)
	length := halfLen * 2
	c := make([]Point, length)
	j, k := 0, 0

	for i := 0; i < length; i++ {
		if j >= halfLen {
			c[i] = b[k]
			k++
		} else if b[k].y < a[j].y {
			c[i] = b[k]
			k++
		} else if k >= halfLen {
			c[i] = a[j]
			j++
		} else {
			c[i] = a[j]
			j++
		}
	}
	return c
}

// func closestPairBaseCase(points []Point, len int) {
// 	for i := 0; i < len; i++ {

// 	}
// }

// func closestPairOfPoints(pX []Point, pY []Point, len int) {
// 	if len <= 3 {
// 		return closestPairBaseCase(pX, len)
// 	}
// }
