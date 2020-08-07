package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDjakstraShortestPathDirected(t *testing.T) {
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
	djakstraShortestPathDirected(graph, &frankfurt)
	assert.Equal(t, 0, frankfurt.totalDistance)
	assert.Equal(t, 100, stuttgart.totalDistance)
	assert.Equal(t, 200, cologne.totalDistance)
	assert.Equal(t, 400, berlin.totalDistance)

	home := Vertex{0, false}
	a := Vertex{0, false}
	b := Vertex{0, false}
	school := Vertex{0, false}
	vertices = Vertices{
		&home,
		&a,
		&b,
		&school,
	}
	edges = Edges{
		Edge{1, &home, &a},
		Edge{4, &home, &b},
		Edge{2, &a, &b},
		Edge{6, &a, &school},
		Edge{3, &b, &school},
	}
	graph = Graph{vertices, edges}
	djakstraShortestPathDirected(graph, &home)
	assert.Equal(t, 0, home.totalDistance)
	assert.Equal(t, 6, school.totalDistance)
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
