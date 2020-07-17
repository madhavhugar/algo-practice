package main

type queue []*Vertex

func (q *queue) add(element *Vertex) {
	*q = append(*q, element)
}

func (q *queue) remove() *Vertex {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}

func breathFirstSearch(graph Graph, startingVertex *Vertex) {
	// mark starting vertex as explored
	startingVertex.Explored = true
	// add first vertex to queue
	var q queue
	q.add(startingVertex)
	done := false
	for !done {
		// select an each edge and mark it's neighbors as explored
		// and add the neighbors to the queue
		v := q.remove()
		for _, edge := range graph.Edges {
			if edge.Head == v && !edge.Tail.Explored {
				q.add(edge.Tail)
				edge.Tail.Explored = true
			} else if edge.Tail == v && !edge.Head.Explored {
				q.add(edge.Head)
				edge.Head.Explored = true
			}
		}
		if len(q) == 0 {
			done = true
		}
	}
}
