package main

type Point struct {
	x int
	y int
}

func closestPairBaseCase(points []Point, len int) {
	for i := 0; i < len; i++ {

	}
}

func closestPairOfPoints(pX []Point, pY []Point, len int) {
	if len <= 3 {
		return closestPairBaseCase(pX, len)
	}
}
