package main

import "math"

// PairOfPoints represents a set of two points
type PairOfPoints struct {
	p1 Point
	p2 Point
}

func calculateEucledianDistanceSquaredBF(p1 Point, p2 Point) float64 {
	return math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2)
}

func sortPointPair(points PairOfPoints) PairOfPoints {
	if points.p1.x < points.p2.x {
		return points
	}
	return PairOfPoints{points.p2, points.p1}
}

func closestPairBruteForce(points []Point, len int) PairOfPoints {
	var closestPair PairOfPoints
	var distance float64
	min := math.Inf(1)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			distance = calculateEucledianDistanceSquaredBF(points[i], points[j])
			if distance < min {
				min = distance
				closestPair = sortPointPair(PairOfPoints{points[i], points[j]})
			}
		}
	}
	return closestPair
}
