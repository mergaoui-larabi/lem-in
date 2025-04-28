package solver

import (
	"lem-in/helpers"
)

func AntDistribution(graph map[string][]string, start, end string, antsnumber int) ([]int, [][]string) {

	paths := FindPaths(graph, start, end)
	helpers.SortPaths(&paths)

	ants_amount := make([]int, len(paths))

	for antsnumber > 0 {
		path_to_increment := 0
		path_hold := len(paths[0]) + ants_amount[0]
		for i := 1; i < len(paths); i++ {
			current_hold := len(paths[i]) + ants_amount[i]
			if path_hold > current_hold {
				path_hold = current_hold
				path_to_increment = i
			}
		}
		ants_amount[path_to_increment]++

		antsnumber--
	}
	return ants_amount, paths
}
