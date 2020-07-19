package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirectedConnectedComponents(t *testing.T) {
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
		{&vertex4, &vertex5},
	}
	undirectedGraph := Graph{vertices, edges}
	// [][]*Vertex := undirectedConnectedComponents(Graph)
	undirectedConnectedComponents(undirectedGraph)
	firstVertexCCValue := undirectedGraph.Vertices[0].Value
	assert.Equal(t, firstVertexCCValue, undirectedGraph.Vertices[1].Value)
	assert.Equal(t, firstVertexCCValue, undirectedGraph.Vertices[2].Value)
	assert.Equal(t, undirectedGraph.Vertices[4].Value, undirectedGraph.Vertices[3].Value)
}

func TestBfs(t *testing.T) {
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
		{&vertex4, &vertex5},
	}
	graph := Graph{vertices, edges}
	startingVertex := &vertex1
	ccNum := 1
	bfs(graph, startingVertex, ccNum)
	assert.Equal(t, true, graph.Vertices[0].Explored)
	assert.Equal(t, true, graph.Vertices[1].Explored)
	assert.Equal(t, true, graph.Vertices[2].Explored)
	assert.Equal(t, false, graph.Vertices[3].Explored)
	assert.Equal(t, false, graph.Vertices[4].Explored)
	assert.Equal(t, ccNum, graph.Vertices[0].Value)
	assert.Equal(t, ccNum, graph.Vertices[1].Value)
	assert.Equal(t, ccNum, graph.Vertices[2].Value)
}
