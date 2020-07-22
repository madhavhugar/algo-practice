package main

func depthFirstSearchRecursive(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Head == startingVertex {
			depthFirstSearchRecursive(graph, edge.Tail)
		} else if edge.Tail == startingVertex {
			depthFirstSearchRecursive(graph, edge.Head)
		}
	}
}
