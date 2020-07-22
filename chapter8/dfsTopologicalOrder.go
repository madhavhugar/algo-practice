package main

var label int

func dfsTopologicalOrder(graph Graph) {
	label = len(graph.Vertices)
	for _, vertex := range graph.Vertices {
		if !vertex.Explored {
			dfsDirected(graph, vertex)
		}
	}
}

func dfsDirected(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Head == startingVertex {
			dfsDirected(graph, edge.Tail)
		}
	}
	startingVertex.Value = label
	label--
}
