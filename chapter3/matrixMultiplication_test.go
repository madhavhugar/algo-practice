package main

import (
	"testing"
	
	assert "github.com/stretchr/testify/assert"
)

func TestMatrixMultiplication(t *testing.T) {
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
	got := matrixMultiplication(a, b, 2)
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
	got = matrixMultiplication(a, b, 2)
	assert.Equal(t, got, wanted)
}
