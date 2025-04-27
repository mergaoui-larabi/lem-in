package main

import (
	"fmt"
	"time"

	"lem-in/graph"
	extract "lem-in/parse"
)

func main() {
	start := time.Now()
	var coords []graph.Room

	newGraph := graph.Graph{Colony: make(map[string][]string)}
	err := extract.Parse("./tests/maps/audit/badexample00", &newGraph, &coords)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, r := range newGraph.Colony {

		fmt.Println(i, r)
	}

	// paths := solver.FindPaths(newGraph.Colony, newGraph.Start.Name, newGraph.End.Name)
	// fmt.Println(paths)

	// solver.Solver(&newGraph)

	fmt.Println("THIS EXEC TOOK", time.Since(start))
}
