package solver

import (
	"fmt"
	"lem-in/graph"
)

func Solver(rooms *graph.Graph) {

	var _ [][]string
	distirbution, paths := AntDistribution(rooms.Colony, rooms.Start, rooms.End, rooms.Ants)
	var i int

	for distirbution[0] != 0 {
		if distirbution[i] > 0 {
			distirbution[i]--
		}
		i++
		if distirbution[i] == 0 {
			i = 0
			continue
		}
	}
	fmt.Println(distirbution)
	fmt.Println(paths)
	ants := AntsWay(3, rooms.Ants, paths)
	fmt.Println(ants)
}
