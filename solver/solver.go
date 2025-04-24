package solver

import (
	"fmt"
	"strconv"

	"lem-in/graph"
)

func Solver(Graph *graph.Graph) {
	var _ [][]string
	distirbution, paths := AntDistribution(Graph.Colony, Graph.Start.Name, Graph.End.Name, Graph.AntsNumber)
	var i int
	ant := 1

	for distirbution[0] != 0 || ant <= Graph.AntsNumber {
		if i == len(distirbution) {
			i = 0
			continue
		}
		if distirbution[i] > 0 {
			Graph.Ants[ant-1] = graph.Ant{
				Name:    strconv.Itoa(ant),
				Path:    i,
				Current: Graph.Start.Name,
			}
			distirbution[i]--
			ant++
		}
		if distirbution[i] == 0 {
			i = 0
			continue
		}
		i++
	}
	// fmt.Println(paths)
	output := MoveAnts(Graph, paths)

	for _, l := range output {
		fmt.Println(l)
	}
	// ants := AntsWay(len(paths), Graph.AntsNumber, paths)
	// fmt.Println(ants)
	fmt.Println(Graph.Ants)
}
