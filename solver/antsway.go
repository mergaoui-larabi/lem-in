package solver

import "lem-in/graph"

func AntsWay(width, antsnumber int, paths [][]string) map[int][]string {
	ants := make(map[int][]string, antsnumber)
	var buffer []string
	for i := 1; i <= antsnumber; i++ {
		buffer = []string{}
		for j := 0; j < len(paths[(i+width-1)%width]); j++ {
			// fmt.Println("for ant", i, "-", paths[(i+width-1)%width][j])
			buffer = append(buffer, "L"+string(i+48)+"-"+paths[(i+width-1)%width][j])
		}
		ants[i] = buffer
	}

	return ants
}

func MoveAnts(Graph *graph.Graph, width int) [][]string {
	// var full = make(map[string]bool, Graph.RoomNumber)
	// var selected = 1
	for Graph.Ants[Graph.AntsNumber].Current != Graph.End.Name {
	}

	return [][]string{}
}
