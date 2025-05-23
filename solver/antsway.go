package solver

import (
	"strconv"

	"lem-in/graph"
)

func MoveAnts(Graph *graph.Graph, path [][]string) [][]string {
	var buffer []string
	var steps [][]string
	var availabe string

	selected := 0

	full := make(map[string]bool)

	for Graph.Ants[Graph.AntsNumber-1].Current != Graph.End.Name {

		// time.Sleep(time.Millisecond * 1000)
		if selected == Graph.AntsNumber {
			selected = 0
		}

		selected_path := Graph.Ants[selected].UniquePath
		current := Graph.Ants[selected].Current
		intersection := potentielRoom(selected_path, Graph.Colony[current])

		for _, r := range intersection {
			if !full[r] {
				availabe = r
			}
		}

		// fmt.Println("selected :", selected)
		// fmt.Println(Graph.Ants)

		// fmt.Println("_________________________")
		// for _, l := range steps {
		// 	fmt.Println(l)
		// }

		// fmt.Println("selected :", selected)
		// fmt.Println("buffer", buffer)
		// fmt.Println("full :", full)
		// fmt.Println("room ", Graph.Ants[selected])
		// fmt.Println("inter", intersection)
		// fmt.Println("avail", availabe)
		// fmt.Println("end", Graph.End.Name)

		// fmt.Println("************************")

		if Graph.Ants[selected].Current == Graph.End.Name {
			selected++
			continue
		}

		if !full[availabe] {
			full[Graph.Ants[selected].Current] = false
			Graph.Ants[selected].Movable = true
			Graph.Ants[selected].Current = availabe
			Graph.Ants[selected].UniquePath = Graph.Ants[selected].UniquePath[1:]
			s := strconv.Itoa(selected + 1)
			move := "L" + s + "-" + availabe
			buffer = append(buffer, move)
			if availabe != Graph.End.Name {
				full[availabe] = true
			}
			if selected == Graph.AntsNumber-1 {
				selected = 0
				steps = append(steps, buffer)
				buffer = []string{}
				continue
			}
			selected++
			continue
		} else {
			Graph.Ants[selected].Movable = false
			selected = 0
			steps = append(steps, buffer)
			buffer = []string{}
			continue
		}

	}

	return steps
}

func potentielRoom(antPath []string, availabeRooms []*graph.Room) []string {
	var intersection []string

	set := make(map[string]bool)

	for _, room := range antPath {
		set[room] = true
	}

	for _, room := range availabeRooms {
		if set[room.Name] {
			intersection = append(intersection, room.Name)
		}
	}
	return intersection
}
