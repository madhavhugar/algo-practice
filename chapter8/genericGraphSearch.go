package main

type Vertex struct {
	Value    interface{}
	Explored bool
}

type Edge struct {
	Head *Vertex
	Tail *Vertex
}

type Graph struct {
	Vertices []*Vertex
	Edges    []Edge
}

type Queue []Vertex

func (q *Queue) add(v Vertex) {
	*q = append(*q, v)
}

func (q *Queue) remove() Vertex {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}

func genericGraphSearch(graph Graph, startingVertex *Vertex) {
	// mark startingVertex as explored
	startingVertex.Explored = true
	// go through all the edges until at least there is
	// one pair which has a combination
	// (explored, unexplored) or (unexplored, explored) in undirected graph
	// (explored, unexplored) in directed graph
	done := true
	var matchFound bool
	for done {
		matchFound = false
		for _, edge := range graph.Edges {
			if edge.Head.Explored && !edge.Tail.Explored {
				edge.Tail.Explored = true
				matchFound = true
			} else if !edge.Head.Explored && edge.Tail.Explored {
				edge.Head.Explored = true
				matchFound = true
			}
		}
		if !matchFound {
			return
		}
	}
}
