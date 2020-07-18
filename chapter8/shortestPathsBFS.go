package main

import (
	"math"
)

type vertexWithDist struct {
	vertex   Vertex
	distance int
}

type edgeWithDistVertex struct {
	head *vertexWithDist
	tail *vertexWithDist
}

type graph struct {
	vertices []*vertexWithDist
	edges    []edgeWithDistVertex
}

type queueWithDistVertices []*vertexWithDist

func (q *queueWithDistVertices) add(element *vertexWithDist) {
	*q = append(*q, element)
}

func (q *queueWithDistVertices) remove() *vertexWithDist {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}

func shortestPathsBFS(g graph, startingVertex *vertexWithDist) {
	// mark the starting vertex as explored
	startingVertex.vertex.Explored = true
	// assign all vertices MAX length
	for _, v := range g.vertices {
		v.distance = math.MaxInt64
	}
	// assign distnce 0 to starting vertex
	startingVertex.distance = 0
	// add the starting vertex to the queue
	var q queueWithDistVertices
	q.add(startingVertex)
	// loop until all connected vertices (to the starting vertex) are explored
	for len(q) > 0 {
		// remove a vertex v from the queue
		v := q.remove()
		// loop through all the edges to find v's neighbors
		for _, edge := range g.edges {
			// find all the vertices w such that there is an edge (v,w)
			// and if w is unexplored, then
			// 		- assign it a distance
			// 		- mark it as explored
			if edge.head == v && !edge.tail.vertex.Explored {
				edge.tail.vertex.Explored = true
				edge.tail.distance = v.distance + 1
			} else if edge.tail == v && !edge.head.vertex.Explored {
				edge.head.vertex.Explored = true
				edge.head.distance = v.distance + 1
			}
		}
	}
}
