package main

import "testing"

func TestClosestPair(t *testing.T) {
	// Base case
	points := []Point{{x: 2, y: 2}, {x: 3, y: 3}, {x: 4, y: 4}}
	closestPairOfPoints(points, 3)
}
