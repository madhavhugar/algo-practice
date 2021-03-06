package main

func resetGraphAsUnexplored(graph Graph) {
	for _, vertex := range graph.Vertices {
		vertex.Explored = false
	}
}

var sccNum int

func dfsPlus(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	startingVertex.Value = sccNum
	for _, edge := range graph.Edges {
		if edge.Head == startingVertex && !edge.Tail.Explored {
			edge.Tail.Value = sccNum
			dfsPlus(graph, edge.Tail)
		}
	}
}

func kosaraju(graph Graph) {
	order := topoSortReverse(graph)
	resetGraphAsUnexplored(graph)
	sccNum = len(graph.Vertices)
	for _, vertex := range order {
		dfsPlus(graph, vertex)
		sccNum--
	}
}

var order []*Vertex
var orderNum int

func dfsReverse(graph Graph, startingVertex *Vertex) {
	if startingVertex.Explored {
		return
	}
	startingVertex.Explored = true
	for _, edge := range graph.Edges {
		if edge.Tail == startingVertex && !edge.Head.Explored {
			dfsReverse(graph, edge.Head)
			order = append(order, startingVertex)
		}
	}
	startingVertex.Value = orderNum
	orderNum--
	order = append([]*Vertex{startingVertex}, order...)
}

func topoSortReverse(graph Graph) []*Vertex {
	orderNum = len(graph.Vertices)
	for _, vertex := range graph.Vertices {
		if !vertex.Explored {
			dfsReverse(graph, vertex)
		}
	}
	return order
}
