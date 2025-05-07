package main

import (
	"fmt"

	"lem-in/graph"
	extract "lem-in/parse"
	"lem-in/solver"
)

func main() {
	// start := time.Now()
	var coords []graph.Room

	newGraph := graph.Graph{Colony: make(map[string][]*graph.Room), Rooms: make(map[string]*graph.Room)}
	err := extract.Parse("./tests/maps/audit/example01", &newGraph, &coords)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("start", newGraph.Start.Name, "end", newGraph.End.Name)
	// for i, r := range newGraph.Colony {
	// 	fmt.Println(i, r)
	// }

	paths := solver.FindPaths(&newGraph, newGraph.Start.Name, newGraph.End.Name)
	fmt.Println(paths)

	// for _, s := range newGraph.Rooms {
	// 	fmt.Println(*s)
	// }

	// for _, s := range newGraph.Rooms {
	// 	fmt.Println(*s)
	// }

	// solver.Solver(&newGraph)

	// fmt.Println("THIS EXEC TOOK", time.Since(start))

}
