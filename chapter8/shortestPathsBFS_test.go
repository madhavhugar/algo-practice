package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathsBFS(t *testing.T) {
	vertex1 := vertexWithDist{Vertex{1, false}, 999}
	vertex2 := vertexWithDist{Vertex{2, false}, 999}
	vertex3 := vertexWithDist{Vertex{3, false}, 999}
	vertex4 := vertexWithDist{Vertex{4, false}, 999}
	vertex5 := vertexWithDist{Vertex{5, false}, 999}
	vertices := []*vertexWithDist{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
	}
	edges := []edgeWithDistVertex{
		{&vertex1, &vertex2},
		{&vertex1, &vertex3},
		{&vertex2, &vertex3},
	}
	graph := graph{vertices, edges}
	shortestPathsBFS(graph, &vertex1)
	assert.Equal(t, 0, vertex1.distance)
	assert.Equal(t, 1, vertex2.distance)
	assert.Equal(t, 1, vertex3.distance)
	assert.Equal(t, math.MaxInt64, vertex4.distance)
	assert.Equal(t, math.MaxInt64, vertex5.distance)
}

func TestAddToQueueWithDistVertex(t *testing.T) {
	var q queueWithDistVertices
	vertex1 := vertexWithDist{Vertex{1, false}, 999}
	vertex2 := vertexWithDist{Vertex{2, false}, 999}
	q.add(&vertex1)
	assert.Equal(t, &vertex1, q[0])
	q.add(&vertex2)
	assert.Equal(t, &vertex2, q[1])
}

func TestRemoveFromQueueWithDistVertex(t *testing.T) {
	vertex1 := vertexWithDist{Vertex{1, false}, 999}
	vertex2 := vertexWithDist{Vertex{2, false}, 999}
	var q queueWithDistVertices
	q.add(&vertex1)
	q.add(&vertex2)
	removed := q.remove()
	assert.Equal(t, &vertex1, removed)
	assert.Equal(t, 1, len(q))
	removed = q.remove()
	assert.Equal(t, &vertex2, removed)
	assert.Equal(t, 0, len(q))
}
