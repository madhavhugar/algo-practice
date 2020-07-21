package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {
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
	depthFirstSearch(graph, startingVertex)
	assert.Equal(t, true, graph.Vertices[0].Explored)
	assert.Equal(t, true, graph.Vertices[1].Explored)
	assert.Equal(t, true, graph.Vertices[2].Explored)
	assert.Equal(t, false, graph.Vertices[3].Explored)
	assert.Equal(t, false, graph.Vertices[4].Explored)
}

func TestAddToStack(t *testing.T) {
	var stack Stack
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: false}
	stack.add(&vertex1)
	assert.Equal(t, &vertex1, stack[0])
	stack.add(&vertex2)
	assert.Equal(t, &vertex2, stack[1])
	assert.Equal(t, &vertex1, stack[0])
	assert.Equal(t, len(stack), 2)
}

func TestRemoveFromStack(t *testing.T) {
	var stack Stack
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: false}
	stack.add(&vertex1)
	stack.add(&vertex2)
	removed := stack.remove()
	assert.Equal(t, &vertex2, removed)
	removed = stack.remove()
	assert.Equal(t, &vertex1, removed)
	assert.Equal(t, len(stack), 0)
}
