package main

import (
	"math"
)

// Point represents a geometric point with x & y values
type Point struct {
	x float64
	y float64
}

// Line represents a pair of points & a geometric line
type Line struct {
	p1 Point
	p2 Point
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
		} else if k >= halfLen {
			c[i] = a[j]
			j++
		} else if b[k].y < a[j].y {
			c[i] = b[k]
			k++
		} else {
			c[i] = a[j]
			j++
		}
	}
	return c
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

func sortPairOfPointsByX(p1 Point, p2 Point) (Point, Point) {
	if p1.x < p2.x {
		return p1, p2
	}
	return p2, p1
}

func computeEucledianDistanceSquare(p1 Point, p2 Point) float64 {
	return math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)
}

func closestPairBaseCase(points []Point) (Point, Point) {
	length := len(points)
	min := math.Inf(1)
	var p1, p2 Point
	var dist float64
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			dist = computeEucledianDistanceSquare(points[i], points[j])
			if dist < min {
				min = dist
				p1, p2 = points[i], points[j]
			}
		}
	}
	return sortPairOfPointsByX(p1, p2)
}

func filterPointsWithinMinDistance(medianPoint Point, py []Point, minDist float64) []Point {
	medianXCoordinate := medianPoint.x
	var filteredPoints []Point
	for i := 0; i < len(py); i++ {
		p1 := py[i]
		p2 := Point{medianXCoordinate, py[i].y}
		dist := computeEucledianDistanceSquare(p1, p2)
		if dist <= minDist {
			filteredPoints = append(filteredPoints, py[i])
		}
	}
	return filteredPoints
}

func closestPairSplitPoints(px []Point, py []Point, minDist float64) (Point, Point) {
	medianPoint := px[len(px)/2-1]
	filteredPoints := filterPointsWithinMinDistance(medianPoint, py, minDist)
	return closestPairBaseCase(filteredPoints)
}

func shortestPairOfPoints(lines ...Line) (Point, Point, float64) {
	var m1, m2 Point
	min := math.MaxFloat64

	for _, line := range lines {
		dist := computeEucledianDistanceSquare(line.p1, line.p2)
		if dist < min {
			m1, m2 = line.p1, line.p2
			min = dist
		}
	}
	return m1, m2, min
}

func closestPairOfPoints(px []Point, py []Point) (Point, Point) {
	length := len(px)
	halfLen := length / 2
	if length <= 3 {
		return closestPairBaseCase(px)
	}
	pxSortedByX := mergeSortPoints(px[:halfLen], mergeListOfPointsSortedByXCoordinate)
	pxSortedByY := mergeSortPoints(px[:halfLen], mergeListOfPointsSortedByYCoordinate)
	l1, l2 := closestPairOfPoints(pxSortedByX, pxSortedByY)
	minDistLeftPair := computeEucledianDistanceSquare(l1, l2)
	pySortedByX := mergeSortPoints(px[halfLen:], mergeListOfPointsSortedByXCoordinate)
	pySortedByY := mergeSortPoints(px[halfLen:], mergeListOfPointsSortedByYCoordinate)
	r1, r2 := closestPairOfPoints(pySortedByX, pySortedByY)
	minDistRightPair := computeEucledianDistanceSquare(r1, r2)
	var minDist float64
	if minDistLeftPair < minDistRightPair {
		minDist = minDistLeftPair
	} else {
		minDist = minDistRightPair
	}
	s1, s2 := closestPairSplitPoints(px, py, minDist)
	lines := []Line{{l1, l2}, {r1, r2}, {s1, s2}}
	m1, m2, min := shortestPairOfPoints(lines...)
	if min == 0.0 {
		m1, m2, _ = shortestPairOfPoints([]Line{{l1, l2}, {r1, r2}}...)
	}
	return m1, m2
}
