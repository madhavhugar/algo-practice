package main

type Stack []*Vertex

func (s *Stack) add(element *Vertex) {
	*s = append(*s, element)
}

func (s *Stack) remove() *Vertex {
	lastIndex := len(*s) - 1
	removed := (*s)[lastIndex]
	*s = (*s)[0:lastIndex]
	return removed
}

func depthFirstSearch(graph Graph, startingVertex *Vertex) {
	// mark all vertices as unexplored
	// add the starting vertex to the stack
	var s Stack
	s.add(startingVertex)
	// loop through until the stack is not empty
	for len(s) > 0 {
		// remove element from stack
		removed := s.remove()
		// if removed vertex is not explored
		if !removed.Explored {
			// 	- mark it as explored
			removed.Explored = true
			//	- add all it's neigbors to the stack
			for _, edge := range graph.Edges {
				if edge.Head == removed {
					s.add(edge.Tail)
				} else if edge.Tail == removed {
					s.add(edge.Head)
				}
			}
		}
	}
}
