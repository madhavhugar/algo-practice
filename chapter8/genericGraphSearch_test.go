package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToQueue(t *testing.T) {
	vertices := []Vertex{
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
	}
	var q Queue
	q.add(vertices[0])
	assert.Equal(t, q[0], vertices[0])
	q.add(vertices[1])
	assert.Equal(t, q[1], vertices[1])
}

func TestRemoveFromQueue(t *testing.T) {
	vertex1 := Vertex{1, false}
	vertex2 := Vertex{2, false}
	var q Queue
	q = []Vertex{
		vertex1,
		vertex2,
	}
	element := q.remove()
	assert.Equal(t, vertex1, element)
	element = q.remove()
	assert.Equal(t, vertex2, element)
}

func TestGenericGraphSearch(t *testing.T) {
	vertex1 := Vertex{1, false}
	vertex2 := Vertex{2, false}
	vertex3 := Vertex{3, false}
	vertex4 := Vertex{4, false}
	vertex5 := Vertex{5, false}
	vertices := []*Vertex{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
	}
	edges := []Edge{
		{vertices[0], vertices[1]},
		{vertices[0], vertices[2]},
		{vertices[1], vertices[2]},
	}
	graph := Graph{
		Edges:    edges,
		Vertices: vertices,
	}
	startingVertex := graph.Vertices[0]

	genericGraphSearch(graph, startingVertex)

	assert.Equal(t, graph.Vertices[0].Explored, true)
	assert.Equal(t, graph.Vertices[1].Explored, true)
	assert.Equal(t, graph.Vertices[2].Explored, true)
	assert.Equal(t, graph.Vertices[3].Explored, false)
	assert.Equal(t, graph.Vertices[4].Explored, false)
}
