package main

import "fmt"

type adjacencyMatrix struct {
	vertices [][]int
}

func main() {
	// We represent a simple undirected graph
	// =1=-----=2=
	fmt.Println("two vertex graph")
	fmt.Printf("=1=-----=2=\n")
	am := adjacencyMatrix{[][]int{
		{0, 1},
		{1, 0},
	}}
	for i := 0; i < len(am.vertices); i++ {
		for j := 0; j < len(am.vertices); j++ {
			if am.vertices[i][j] == 1 {
				fmt.Printf("(%d => %v)", i+1, j+1)
			}
			fmt.Println()
		}
	}

	// We represent a simple undirected graph
	// =1=-----=2=------=3=
	//          |
	//         =4=
	fmt.Println("\nfour vertex graph")
	fmt.Printf("=1=-----=2=------=3=\n")
	fmt.Printf("         |          \n")
	fmt.Printf("        =4=         \n\n")
	am = adjacencyMatrix{[][]int{
		{0, 1, 0, 0},
		{1, 0, 1, 1},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
	}}
	for i := 0; i < len(am.vertices); i++ {
		for j := 0; j < len(am.vertices); j++ {
			if am.vertices[i][j] == 1 {
				fmt.Printf("(%d => %v)", i+1, j+1)
			}
		}
		fmt.Println()
	}
}
