package solver

import (
	"lem-in/graph"
)

func FilterPath(anthill *graph.Graph, rooms []*graph.Room, path []string, paths *[][]string) {
	var final [][]string
	var pathscost float32

	// fmt.Println("rooma", *rooms[0], *rooms[1])
	// fmt.Println("problematic paths:", rooms[0].Path, path)
	// fmt.Println("now", paths)

	if match(rooms[0].Path, path) {
		return
	}

	deletePath(paths, rooms[0].Path)
	pathscost = CalculateCost(anthill, *paths, nil)
	// fmt.Println("starter cost", pathscost)

	for len(rooms) > 0 {

		pop := rooms[0]
		rooms = rooms[1:]
		temp, cost, dupp := fight(anthill, paths, pop.Path, path, pop)
		if !dupp && cost < pathscost {
			final = temp
			pathscost = cost
		}
		// fmt.Println("f", temp, dupp)
		temp, cost, dupp = fight(anthill, paths, path, pop.Path, pop)
		if !dupp && cost < pathscost {
			final = temp
			pathscost = cost
		}
		// fmt.Println("s", temp, dupp)
	}
	// fmt.Println("final grou,p before append", final)

	(*paths) = append((*paths), final...)

}

func deletePath(paths *[][]string, path []string) {
	// fmt.Println("on delete ", (*paths))
	for i, p := range *paths {
		if len(p) == len(path) {
			if match(p, path) {
				(*paths) = append((*paths)[:i], (*paths)[i+1:]...)
			}
		}
	}
}

func fight(anthill *graph.Graph, paths *[][]string, first, second []string, join *graph.Room) ([][]string, float32, bool) {

	var set [][]string
	var cost float32
	visited := make(map[string]bool)
	visited[join.Name] = true
	visited[anthill.Start.Name] = true
	set = append(set, first)
	temp, dupp := BFS(anthill, second[0], anthill.End.Name, &visited)
	if len(temp) != 0 {
		set = append(set, temp)
	}
	cost = CalculateCost(anthill, *paths, set)
	return set, cost, dupp
}

func CalculateCost(anthill *graph.Graph, paths [][]string, set [][]string) float32 {
	var sum float32
	for i := range paths {
		sum += float32(len(paths[i]))
	}
	for i := range set {
		sum += float32(len(set[i]))
	}
	sum += float32(anthill.AntsNumber)
	denominator := float32(len(paths) + len(set))
	return sum / denominator

}

func match(p1, p2 []string) bool {
	set := make(map[string]bool)
	for _, s := range p1 {
		set[s] = true
	}
	for _, s := range p2 {
		if !set[s] {
			return false
		}
	}
	return true
}
