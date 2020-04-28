package main

import (
	"testing"
	
	assert "github.com/stretchr/testify/assert"
)

func TestRecursiveMatrixMultiplication(t *testing.T) {
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
	got := recursiveMatrixMultiplication(a, b, 2)
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
	got = recursiveMatrixMultiplication(a, b, 2)
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
	got = recursiveMatrixMultiplication(a, b, 4)
	assert.Equal(t, got, wanted)
}
