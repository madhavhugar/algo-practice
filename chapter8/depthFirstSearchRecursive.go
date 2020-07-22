package main

func depthFirstSearchRecursive(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Head == startingVertex {
			edge.Tail.Explored = true
			depthFirstSearchRecursive(graph, edge.Tail)
		} else if edge.Tail == startingVertex {
			edge.Head.Explored = true
			depthFirstSearchRecursive(graph, edge.Head)
		}
	}
}
