package main

import (
	"math"
)

func djakstraShortestPathDirected(graph Graph, startingVertex *Vertex) {
	for _, vertex := range graph.vertices {
		vertex.totalDistance = math.MaxInt64
	}
	startingVertex.totalDistance = 0

	var iterated = make(Vertices, len(graph.vertices))
	iterated = Vertices{startingVertex}

	continueIteration := true
	for continueIteration {
		continueIteration = false
		for _, edge := range graph.edges {
			if iterated.elementExists(edge.head) &&
				(!iterated.elementExists(edge.tail) || edge.tail.totalDistance > edge.head.totalDistance+edge.distance) {
				continueIteration = true
				currentLen := edge.head.totalDistance + edge.distance
				if edge.tail.totalDistance > currentLen {
					edge.tail.totalDistance = currentLen
					iterated = append(iterated, edge.tail)
				}
			}
		}
	}
}
