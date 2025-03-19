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
	err := extract.Parse("./tests/example05.txt", &newGraph, &coords)
	if err != nil {
		fmt.Println(err)
		return
	}

	solver.Solver(&newGraph)
}
