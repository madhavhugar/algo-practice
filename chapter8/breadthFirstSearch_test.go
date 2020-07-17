package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreathFirstSearch(t *testing.T) {
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
	breathFirstSearch(graph, &vertex1)
	assert.Equal(t, true, graph.Vertices[0].Explored)
	assert.Equal(t, true, graph.Vertices[1].Explored)
	assert.Equal(t, true, graph.Vertices[2].Explored)
	assert.Equal(t, false, graph.Vertices[3].Explored)
	assert.Equal(t, false, graph.Vertices[4].Explored)
}

func TestAddToSmallQueue(t *testing.T) {
	var q queue
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: false}
	q.add(&vertex1)
	assert.Equal(t, &vertex1, q[0])
	q.add(&vertex2)
	assert.Equal(t, &vertex2, q[1])
}

func TestRemoveFromSmallQueue(t *testing.T) {
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: false}
	var q queue
	q.add(&vertex1)
	q.add(&vertex2)
	removed := q.remove()
	assert.Equal(t, &vertex1, removed)
	assert.Equal(t, 1, len(q))
	removed = q.remove()
	assert.Equal(t, &vertex2, removed)
	assert.Equal(t, 0, len(q))
}
