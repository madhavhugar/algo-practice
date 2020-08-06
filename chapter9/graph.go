package main

type Vertex struct {
	totalDistance int
	explored      bool
}

type Edge struct {
	distance int
	head     *Vertex
	tail     *Vertex
}

type Vertices []*Vertex

type Edges []Edge

type Graph struct {
	vertices Vertices
	edges    Edges
}

func (v *Vertices) elementExists(element *Vertex) bool {
	exists := false
	for _, vertex := range *v {
		if vertex == element {
			exists = true
		}
	}
	return exists
}
