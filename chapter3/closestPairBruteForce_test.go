package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestClosestPairBruteForce(t *testing.T) {
	points := []Point{{1, 2}, {4, 5}, {2, 1}}
	cp1, cp2 := Point{1, 2}, Point{2, 1}
	closestPair := closestPairBruteForce(points, 3)
	assert.Equal(t, cp1, closestPair.p1)
	assert.Equal(t, cp2, closestPair.p2)

	points = []Point{{1, 1}, {2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}}
	cp1, cp2 = Point{1, 1}, Point{2, 2}
	closestPair = closestPairBruteForce(points, 8)
	assert.Equal(t, cp1, closestPair.p1)
	assert.Equal(t, cp2, closestPair.p2)
}

func TestCalculateEucledianDistanceSquared(t *testing.T) {
	p1, p2 := Point{1, 2}, Point{2, 1}
	wanted := 2.0
	got := calculateEucledianDistanceSquaredBF(p1, p2)
	assert.Equal(t, wanted, got)
}

func TestPointsPairSortBF(t *testing.T) {
	p1, p2 := Point{4, 2}, Point{2, 1}
	pointPair := PairOfPoints{p2, p1}

	got := sortPointPair(pointPair)
	assert.Equal(t, got.p1, p2)
	assert.Equal(t, got.p2, p1)
}
