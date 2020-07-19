package main

func bfs(graph Graph, startingVertex *Vertex, ccNum int) {
	// mark starting vertex as explored
	startingVertex.Explored = true
	// add starting vertex to the queue
	var q queue
	q.add(startingVertex)
	startingVertex.Value = ccNum
	// loop until all connected neighbours (of startingVertex) are explored
	for len(q) > 0 {
		v := q.remove()
		for _, edge := range graph.Edges {
			// when edge meets the criteria (v, w) where w is unexplored,
			// add it to the queue
			if edge.Head == v && !edge.Tail.Explored {
				q.add(edge.Tail)
				edge.Tail.Explored = true
				edge.Tail.Value = ccNum
			} else if edge.Tail == v && !edge.Head.Explored {
				q.add(edge.Head)
				edge.Head.Explored = true
				edge.Head.Value = ccNum
			}
		}
	}
}

func undirectedConnectedComponents(undirectedGraph Graph) {
	for i, vertex := range undirectedGraph.Vertices {
		if !vertex.Explored {
			bfs(undirectedGraph, vertex, i)
		}
	}
}
