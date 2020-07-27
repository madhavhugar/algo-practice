package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// To test SCC
func TestKosaRaju(t *testing.T) {

}

func TestTopoSortReverseUsingDfs(t *testing.T) {
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: false}
	vertex3 := Vertex{Value: 3, Explored: false}
	vertex4 := Vertex{Value: 4, Explored: false}
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
	topoSorted := topoSortReverse(graph)
	assert.Equal(t, graph.Vertices[0].Value.(int), 4)
	assert.Equal(t, graph.Vertices[3].Value.(int), 1)
	assert.Equal(t, &vertex4, topoSorted[0])
	assert.Equal(t, &vertex1, topoSorted[3])
}

// Test reverse DFS search => returns a slice with vertices in the order we need to invoke DFS
// DFS that assigns a SCC number
func TestDfsPlusGraphSearch(t *testing.T) {
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
	dfsPlus(graph, startingVertex)
	assert.Equal(t, graph.Vertices[0].Explored, true)
	assert.Equal(t, graph.Vertices[1].Explored, true)
	assert.Equal(t, graph.Vertices[2].Explored, true)
	assert.Equal(t, graph.Vertices[3].Explored, false)
	assert.Equal(t, graph.Vertices[4].Explored, false)

	assert.Equal(t, graph.Vertices[0].Value, 0)
	assert.Equal(t, graph.Vertices[1].Value, 0)
	assert.Equal(t, graph.Vertices[2].Value, 0)
	assert.Equal(t, graph.Vertices[3].Value, 4)
	assert.Equal(t, graph.Vertices[4].Value, 5)
}

// Reset the graph by marking all vertices unexplored
func TestResetGraphAsUnexplored(t *testing.T) {
	vertex1 := Vertex{Value: 1, Explored: false}
	vertex2 := Vertex{Value: 2, Explored: true}
	vertex3 := Vertex{Value: 3, Explored: false}
	vertex4 := Vertex{Value: 4, Explored: true}
	vertex5 := Vertex{Value: 5, Explored: true}
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
	resetGraphAsUnexplored(graph)
	assert.Equal(t, graph.Vertices[0].Explored, false)
	assert.Equal(t, graph.Vertices[1].Explored, false)
	assert.Equal(t, graph.Vertices[2].Explored, false)
	assert.Equal(t, graph.Vertices[3].Explored, false)
	assert.Equal(t, graph.Vertices[4].Explored, false)
}
