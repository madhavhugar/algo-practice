package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestKosaRaju(t *testing.T) {
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
	kosaraju(graph)
	assert.Equal(t, graph.Vertices[0].Value.(int), 1)
	assert.Equal(t, graph.Vertices[3].Value.(int), 4)

	vertex1 = Vertex{Value: 0, Explored: false}
	vertex2 = Vertex{Value: 0, Explored: false}
	vertex3 = Vertex{Value: 0, Explored: false}
	vertex4 = Vertex{Value: 0, Explored: false}
	vertex5 := Vertex{Value: 0, Explored: false}
	vertex6 := Vertex{Value: 0, Explored: false}
	vertex7 := Vertex{Value: 0, Explored: false}
	vertex8 := Vertex{Value: 0, Explored: false}
	vertex9 := Vertex{Value: 0, Explored: false}
	vertex10 := Vertex{Value: 0, Explored: false}
	vertex11 := Vertex{Value: 0, Explored: false}
	vertices = []*Vertex{
		&vertex1,
		&vertex2,
		&vertex3,
		&vertex4,
		&vertex5,
		&vertex6,
		&vertex7,
		&vertex8,
		&vertex9,
		&vertex10,
		&vertex11,
	}
	edges = []Edge{
		{&vertex1, &vertex5},
		{&vertex2, &vertex9},
		{&vertex3, &vertex1},
		{&vertex4, &vertex2},
		{&vertex4, &vertex9},
		{&vertex5, &vertex3},
		{&vertex6, &vertex8},
		{&vertex6, &vertex11},
		{&vertex7, &vertex4},
		{&vertex7, &vertex5},
		{&vertex8, &vertex9},
		{&vertex8, &vertex10},
		{&vertex8, &vertex11},
		{&vertex9, &vertex5},
		{&vertex9, &vertex7},
		{&vertex10, &vertex2},
		{&vertex10, &vertex6},
		{&vertex11, &vertex3},
	}
	graph = Graph{vertices, edges}
	kosaraju(graph)
	sccOne := graph.Vertices[0].Value.(int)
	assert.Equal(t, sccOne, graph.Vertices[0].Value.(int))
	assert.Equal(t, sccOne, graph.Vertices[2].Value.(int))
	assert.Equal(t, sccOne, graph.Vertices[4].Value.(int))
	sccTwo := graph.Vertices[5].Value.(int)
	assert.Equal(t, sccTwo, graph.Vertices[5].Value.(int))
	assert.Equal(t, sccTwo, graph.Vertices[7].Value.(int))
	assert.Equal(t, sccTwo, graph.Vertices[9].Value.(int))
	sccThree := graph.Vertices[1].Value.(int)
	assert.Equal(t, sccThree, graph.Vertices[1].Value.(int))
	assert.Equal(t, sccThree, graph.Vertices[3].Value.(int))
	assert.Equal(t, sccThree, graph.Vertices[8].Value.(int))
	assert.Equal(t, sccThree, graph.Vertices[6].Value.(int))
	sccFour := graph.Vertices[10].Value.(int)
	assert.Equal(t, sccFour, graph.Vertices[10].Value.(int))

	assert.NotEqual(t, sccOne, sccTwo)
	assert.NotEqual(t, sccOne, sccThree)
	assert.NotEqual(t, sccOne, sccFour)
	assert.NotEqual(t, sccTwo, sccThree)
	assert.NotEqual(t, sccTwo, sccFour)
	assert.NotEqual(t, sccThree, sccFour)
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
