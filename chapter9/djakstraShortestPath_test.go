package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstraShortestPath(t *testing.T) {
	frankfurt := Vertex{0, false}
	stuttgart := Vertex{0, false}
	cologne := Vertex{0, false}
	berlin := Vertex{0, false}
	vertices := Vertices{
		&frankfurt,
		&stuttgart,
		&cologne,
		&berlin,
	}
	edges := Edges{
		Edge{100, &frankfurt, &stuttgart},
		Edge{200, &frankfurt, &cologne},
		Edge{300, &stuttgart, &berlin},
		Edge{400, &cologne, &berlin},
	}
	graph := Graph{vertices, edges}
	dijkstraShortestPath(graph, &frankfurt)
	assert.Equal(t, 0, frankfurt.totalDistance)
	assert.Equal(t, 100, stuttgart.totalDistance)
	assert.Equal(t, 200, cologne.totalDistance)
	assert.Equal(t, 400, stuttgart.totalDistance)
}

func TestElementExists(t *testing.T) {
	frankfurt := Vertex{0, false}
	stuttgart := Vertex{0, false}
	cologne := Vertex{0, false}
	berlin := Vertex{0, false}
	brussels := Vertex{0, false}
	vertices := Vertices{
		&frankfurt,
		&stuttgart,
		&cologne,
		&berlin,
	}
	got := vertices.elementExists(&frankfurt)
	wanted := true
	assert.Equal(t, wanted, got)

	got = vertices.elementExists(&brussels)
	wanted = false
	assert.Equal(t, wanted, got)
}
