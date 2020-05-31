package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// func TestClosestPair(t *testing.T) {
// 	// Base case
// 	points := []Point{{x: 2, y: 2}, {x: 3, y: 3}, {x: 4, y: 4}}
// 	closestPairOfPoints(points, 3)
// }

// func TestMergeSortPoints(t *testing.T) {
// 	points := []Point{{x: 2, y: 2}, {x: 1, y: 3}, {x: 6, y: 4}, {x: 4, y: 4}}
// 	got := mergeSortPointsByXCoordinate(points)
// 	wanted := []Point{{x: 1, y: 3}, {x: 2, y: 2}, {x: 4, y: 4}, {x: 6, y: 4}}
// 	assert.Equal(t, wanted, got)
// }

func TestMergeSortPoints(t *testing.T) {
	points := []Point{{x: 2, y: 2}, {x: 1, y: 3}, {x: 6, y: 4}, {x: 4, y: 4}}
	gotSortByX := mergeSortPoints(points, mergeListOfPointsSortedByXCoordinate)
	wantedSortByX := []Point{{x: 1, y: 3}, {x: 2, y: 2}, {x: 4, y: 4}, {x: 6, y: 4}}
	assert.Equal(t, wantedSortByX, gotSortByX)

	points = []Point{{x: 1, y: 2}, {x: 4, y: 3}, {x: 2, y: 1}, {x: 3, y: 4}}
	gotSortByY := mergeSortPoints(points, mergeListOfPointsSortedByYCoordinate)
	wantedSortByY := []Point{{x: 2, y: 1}, {x: 1, y: 2}, {x: 4, y: 3}, {x: 3, y: 4}}
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
