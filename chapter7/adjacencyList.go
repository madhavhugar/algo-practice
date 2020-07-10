package main

import "fmt"

type vertex struct {
	value int
	edges []*vertex
}

type adjacencyList struct {
	vertices []vertex
}

func main() {
	// We represent a simple undirected graph
	// =1=-----=2=
	fmt.Println("two vertex graph")
	fmt.Printf("=1=-----=2=\n")
	a := vertex{1, nil}
	b := vertex{2, nil}
	a.edges = []*vertex{&b}
	b.edges = []*vertex{&a}
	fmt.Printf("(%d => %v)\n", a.value, *&a.edges[0].value)
	fmt.Printf("(%d => %v)\n", b.value, *&b.edges[0].value)

	// We represent a simple undirected graph
	// =1=-----=2=------=3=
	//          |
	//         =4=
	fmt.Println("\nfour vertex graph")
	fmt.Printf("=1=-----=2=------=3=\n")
	fmt.Printf("         |          \n")
	fmt.Printf("        =4=         \n")
	c := vertex{1, nil}
	d := vertex{2, nil}
	e := vertex{3, nil}
	f := vertex{4, nil}
	c.edges = []*vertex{&d}
	d.edges = []*vertex{&c, &e, &f}
	e.edges = []*vertex{&d}
	f.edges = []*vertex{&d}
	fmt.Printf("(%d => %v)\n", c.value, *&c.edges[0].value)
	for i := 0; i < len(d.edges); i++ {
		fmt.Printf("(%d => %v)  ", d.value, *&d.edges[i].value)
	}
	fmt.Printf("\n(%d => %v)\n", e.value, *&e.edges[0].value)
	fmt.Printf("(%d => %v)\n", f.value, *&f.edges[0].value)
}
