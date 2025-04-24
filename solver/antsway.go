package solver

import (
	"fmt"
	"strconv"
	"time"

	"lem-in/graph"
)

func MoveAnts(Graph *graph.Graph, path [][]string) [][]string {
	var buffer []string
	var steps [][]string
	var availabe string

	selected := 0

	full := make(map[string]bool)

	for Graph.Ants[Graph.AntsNumber-1].Current != Graph.End.Name {
		fmt.Println("_________________________")
		time.Sleep(time.Millisecond * 1000)
		for _, l := range steps {
			fmt.Println(l)
		}

		if selected == Graph.AntsNumber {
			selected = 0
		}

		selected_path := path[Graph.Ants[selected].Path]
		current := Graph.Ants[selected].Current
		intersection := potentielRoom(selected_path, Graph.Colony[current])

		for _, r := range intersection {
			if !full[r] {
				availabe = r
			}
		}

		fmt.Println("full :", full)
		fmt.Println("selected :", selected)
		fmt.Println("room ", Graph.Ants[selected])
		fmt.Println("inter", intersection)
		fmt.Println("avail", availabe)
		fmt.Println("end", Graph.End.Name)

		if Graph.Ants[selected].Current == Graph.End.Name {
			fmt.Println("wasalna")
			selected++
			continue
		}

		if !full[availabe] {
			fmt.Println("if")
			full[Graph.Ants[selected].Current] = false
			Graph.Ants[selected].Movable = true
			Graph.Ants[selected].Current = availabe
			s := strconv.Itoa(selected + 1)
			move := "L" + s + "-" + availabe
			buffer = append(buffer, move)
			if availabe != Graph.End.Name {
				full[availabe] = true
			}
			selected++
			continue
		} else {
			fmt.Println("else")
			Graph.Ants[selected].Movable = false
			selected = 0
			steps = append(steps, buffer)
			buffer = []string{}
			continue
		}

	}

	return steps
}

func potentielRoom(antPath, availabeRooms []string) []string {
	var intersection []string

	set := make(map[string]bool)

	for _, room := range antPath {
		set[room] = true
	}

	for _, room := range availabeRooms {
		if set[room] {
			intersection = append(intersection, room)
		}
	}
	return intersection
}
