package main

import (
	"fmt"

	"lem-in/graph"
	extract "lem-in/parse"
	"lem-in/solver"
)

func main() {
	var coords []graph.Room

	newGraph := graph.Graph{Colony: make(map[string][]string)}
	err := extract.Parse("./tests/example00.txt", &newGraph, &coords)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newGraph.Colony)

	paths := solver.FindPaths(newGraph.Colony, newGraph.Start.Name, newGraph.End.Name)
	fmt.Println(paths)
	
	solver.Solver(&newGraph)
}
