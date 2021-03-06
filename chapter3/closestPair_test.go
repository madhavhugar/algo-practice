package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestClosestPair(t *testing.T) {
	// Base case
	points := []Point{{2, 2}, {3, 3}, {5, 5}}
	gotP1, gotP2 := closestPairOfPoints(points, points)
	wantedP1, wantedP2 := Point{2, 2}, Point{3, 3}
	assert.Equal(t, wantedP1, gotP1)
	assert.Equal(t, wantedP2, gotP2)
	// When the closest pair is a left pair
	points = []Point{{1, 1}, {2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}}
	gotP1, gotP2 = closestPairOfPoints(points, points)
	wantedP1, wantedP2 = Point{1, 1}, Point{2, 2}
	assert.Equal(t, wantedP1, gotP1)
	assert.Equal(t, wantedP2, gotP2)
	// When the closest pair is a right pair
	points = []Point{{1, 1}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}, {15, 15}}
	px := mergeSortPoints(points, mergeListOfPointsSortedByXCoordinate)
	py := mergeSortPoints(points, mergeListOfPointsSortedByYCoordinate)
	gotP1, gotP2 = closestPairOfPoints(px, py)
	wantedP1, wantedP2 = Point{14, 14}, Point{15, 15}
	assert.Equal(t, wantedP1, gotP1)
	assert.Equal(t, wantedP2, gotP2)
	// When the closest pair is a split pair
	points = []Point{{-1, -1}, {2, 2}, {5, 5}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}}
	gotP1, gotP2 = closestPairOfPoints(points, points)
	wantedP1, wantedP2 = Point{5, 5}, Point{6, 6}
	assert.Equal(t, wantedP1, gotP1)
	assert.Equal(t, wantedP2, gotP2)
}

func TestClosestPairBaseCase(t *testing.T) {
	points := []Point{{1, 2}, {4, 5}, {2, 1}}
	cp1, cp2 := Point{1, 2}, Point{2, 1}
	gotP1, gotP2 := closestPairBaseCase(points)
	assert.Equal(t, cp1, gotP1)
	assert.Equal(t, cp2, gotP2)

	points = []Point{{1, 1}, {2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}}
	cp1, cp2 = Point{1, 1}, Point{2, 2}
	gotP1, gotP2 = closestPairBaseCase(points)
	assert.Equal(t, cp1, gotP1)
	assert.Equal(t, cp2, gotP2)
}

func TestClosestPairSplitPoints(t *testing.T) {
	points := []Point{{0, 0}, {2, 2}, {4, 4}, {6, 6}, {7, 7}, {10, 10}, {12, 12}, {14, 14}}
	minDist := 4.0
	gotP1, gotP2 := closestPairSplitPoints(points, points, minDist)
	wantedP1, wantedP2 := Point{6, 6}, Point{7, 7}
	assert.Equal(t, wantedP1, gotP1)
	assert.Equal(t, wantedP2, gotP2)
}

func TestFilterPointsWithinMinDistance(t *testing.T) {
	points := []Point{{1, 1}, {2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}}
	minDistSquared := 4.0
	got := filterPointsWithinMinDistance(Point{6, 6}, points, minDistSquared)
	wanted := []Point{{4, 4}, {6, 6}, {8, 8}}
	assert.Equal(t, wanted, got)
}

func TestComputeEucledianDistanceSquare(t *testing.T) {
	p1, p2 := Point{1, 2}, Point{2, 1}
	wanted := 2.0
	got := computeEucledianDistanceSquare(p1, p2)
	assert.Equal(t, wanted, got)
}

func TestSortPairOfPointsByX(t *testing.T) {
	p1, p2 := Point{4, 2}, Point{2, 1}

	gotP1, gotP2 := sortPairOfPointsByX(p1, p2)
	assert.Equal(t, gotP1, p2)
	assert.Equal(t, gotP2, p1)
}

func TestShortestPairOfPoints(t *testing.T) {
	line1 := Line{Point{4, 4}, Point{6, 6}}
	line2 := Line{Point{10, 10}, Point{12, 12}}
	line3 := Line{Point{1, 1}, Point{2, 2}}
	gotP1, gotP2, gotMin := shortestPairOfPoints(line1, line2, line3)
	wantedP1, wantedP2 := Point{1, 1}, Point{2, 2}
	assert.Equal(t, wantedP1, gotP1)
	assert.Equal(t, wantedP2, gotP2)
	assert.Greater(t, gotMin, -1.0)
}

func TestMergeSortPoints(t *testing.T) {
	points := []Point{{x: 2, y: 2}, {x: 1, y: 3}, {x: 6, y: 4}, {x: 4, y: 4}}
	gotSortByX := mergeSortPoints(points, mergeListOfPointsSortedByXCoordinate)
	wantedSortByX := []Point{{x: 1, y: 3}, {x: 2, y: 2}, {x: 4, y: 4}, {x: 6, y: 4}}
	assert.Equal(t, wantedSortByX, gotSortByX)

	points = []Point{{x: 1, y: 2}, {x: 4, y: 3}, {x: 2, y: 1}, {x: 3, y: 4}}
	gotSortByY := mergeSortPoints(points, mergeListOfPointsSortedByYCoordinate)
	wantedSortByY := []Point{{x: 2, y: 1}, {x: 1, y: 2}, {x: 4, y: 3}, {x: 3, y: 4}}
	assert.Equal(t, wantedSortByY, gotSortByY)

	points = []Point{{4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}, {1, 1}, {2, 2}}
	gotSortByY = mergeSortPoints(points, mergeListOfPointsSortedByYCoordinate)
	wantedSortByY = []Point{{1, 1}, {2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14}}
	assert.Equal(t, wantedSortByY, gotSortByY)
}

func TestSplitPointsList(t *testing.T) {
	points := []Point{{x: 2, y: 2}, {x: 1, y: 3}, {x: 6, y: 4}, {x: 4, y: 4}}
	gotA, gotB := splitPointsList(points)
	wantedA, wantedB := []Point{{x: 2, y: 2}, {x: 1, y: 3}}, []Point{{x: 6, y: 4}, {x: 4, y: 4}}
	assert.Equal(t, wantedA, gotA)
	assert.Equal(t, wantedB, gotB)
}

func TestMergeListOfPointsSortedByXCoordinate(t *testing.T) {
	inputA, inputB := []Point{{x: 1, y: 2}, {x: 4, y: 3}}, []Point{{x: 2, y: 4}, {x: 3, y: 4}}
	got := mergeListOfPointsSortedByXCoordinate(inputA, inputB)
	wanted := []Point{{x: 1, y: 2}, {x: 2, y: 4}, {x: 3, y: 4}, {x: 4, y: 3}}
	assert.Equal(t, wanted, got)
}

func TestMergeListOfPointsSortedByYCoordinate(t *testing.T) {
	inputA, inputB := []Point{{x: 1, y: 2}, {x: 4, y: 3}}, []Point{{x: 2, y: 1}, {x: 3, y: 4}}
	got := mergeListOfPointsSortedByYCoordinate(inputA, inputB)
	wanted := []Point{{x: 2, y: 1}, {x: 1, y: 2}, {x: 4, y: 3}, {x: 3, y: 4}}
	assert.Equal(t, wanted, got)
}
