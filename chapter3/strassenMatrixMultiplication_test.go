package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestStrassenMatrixMultiplication(t *testing.T) {
	a := [][]int{
		{1,2},
		{3,4},
	}
	b := [][]int{
		{2,0},
		{1,2},
	}
	wanted := [][]int{
		{4,4},
		{10,8},
	}
	got := strassenMatrixMultiplication(a, b, 2)
	assert.Equal(t, got, wanted)

	a = [][]int{
		{2,0},
		{1,2},
	}
	b = [][]int{
		{1,2},
		{3,4},
	}
	wanted = [][]int{
		{2,4},
		{7,10},
	}
	got = strassenMatrixMultiplication(a, b, 2)
	assert.Equal(t, got, wanted)

	a = [][]int{
		{1,2,3,4},
		{3,4,5,6},
		{5,6,7,8},
		{7,8,9,10},
	}
	b = [][]int{
		{1,2,3,4},
		{3,4,5,6},
		{5,6,7,8},
		{7,8,9,10},
	}
	wanted = [][]int{
		{50,60,70,80},
		{82,100,118,136},
		{114,140,166,192},
		{146,180,214,248},
	}
	got = strassenMatrixMultiplication(a, b, 4)
	assert.Equal(t, got, wanted)
}

func TestAddMatricesStrassen(t *testing.T) {
	a := [][]int{
		{1,2},
		{3,4},
	}
	b := [][]int{
		{2,0},
		{1,2},
	}
	wanted := [][]int{
		{3,2},
		{4,6},
	}
	got := addMatricesStrassen(a, b, 2)
	assert.Equal(t, got, wanted)
}

func TestSubtractMatricesStrassen(t *testing.T) {
	a := [][]int{
		{1,2},
		{3,4},
	}
	b := [][]int{
		{2,0},
		{1,2},
	}
	wanted := [][]int{
		{-1,2},
		{2,2},
	}
	got := substractMatricesStrassen(a, b, 2)
	assert.Equal(t, got, wanted)
}

func TestSplitIntoSubMatricesStrassen(t *testing.T) {
	a := [][]int{
		{1,2},
		{3,4},
	}
	e,f,g,h := splitIntoSubMatricesStrassen(a, 2)
	assert.Equal(t, e, [][]int{{1}})
	assert.Equal(t, f, [][]int{{2}})
	assert.Equal(t, g, [][]int{{3}})
	assert.Equal(t, h, [][]int{{4}})
}

func TestCombineSubMatricesIntoMatrixStrassen(t *testing.T) {
	e := [][]int{{1}}
	f := [][]int{{2}}
	g := [][]int{{3}}
	h := [][]int{{4}}
	wanted := [][]int{
		{1,2},
		{3,4},
	}
	got := combineSubMatricesIntoMatrixStrassen(e, f, g, h, 2)
	assert.Equal(t, got, wanted)
}