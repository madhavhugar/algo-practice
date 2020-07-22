package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsTopologicalOrder(t *testing.T) {
	vertex1 := Vertex{Value: 0, Explored: false}
	vertex2 := Vertex{Value: 0, Explored: false}
	vertex3 := Vertex{Value: 0, Explored: false}
	vertex4 := Vertex{Value: 0, Explored: false}
	vertices := []*Vertex{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
	}
	edges := []Edge{
		{&vertex1, &vertex2},
		{&vertex1, &vertex3},
		{&vertex2, &vertex4},
		{&vertex3, &vertex4},
	}
	graph := Graph{vertices, edges}
	dfsTopologicalOrder(graph)
	assert.Equal(t, 1, vertex1.Value)
	assert.Equal(t, true, (2 <= vertex2.Value.(int)) && (vertex2.Value.(int) <= 3))
	assert.Equal(t, true, (2 <= vertex3.Value.(int)) && (vertex3.Value.(int) <= 3))
	assert.Equal(t, 4, vertex4.Value)
}

func TestDfs(t *testing.T) {
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: false}
	vertex3 := Vertex{Value: 3, Explored: false}
	vertex4 := Vertex{Value: 4, Explored: false}
	vertex5 := Vertex{Value: 5, Explored: false}
	vertices := []*Vertex{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
	}
	edges := []Edge{
		{&vertex1, &vertex2},
		{&vertex1, &vertex3},
		{&vertex2, &vertex3},
	}
	graph := Graph{vertices, edges}
	startingVertex := &vertex1
	dfsDirected(graph, startingVertex)
	assert.Equal(t, true, graph.Vertices[0].Explored)
	assert.Equal(t, true, graph.Vertices[1].Explored)
	assert.Equal(t, true, graph.Vertices[2].Explored)
	assert.Equal(t, false, graph.Vertices[3].Explored)
	assert.Equal(t, false, graph.Vertices[4].Explored)
}
