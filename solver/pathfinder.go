package solver

import (
	"lem-in/graph"
)

func FindPaths(anthill *graph.Graph, start, end string) [][]string {
	var paths [][]string
	var hold_disjoint_rooms []*graph.Room
	visited := make(map[string]bool)
	var cost_after, cost_before float32
	shortestPath, _ := BFS(anthill, start, end, &visited)
	shortestPath = shortestPath[1:]
	for _, p := range shortestPath {
		anthill.Rooms[p].Explored = true
		anthill.Rooms[p].Path = shortestPath
		anthill.Rooms[p].Head = shortestPath[0]
		anthill.Rooms[p].Net = len(anthill.Colony[p])
		anthill.Rooms[p].Index = len(paths)

	}

	paths = append(paths, shortestPath)
	cost_before = CalculateCost(anthill, paths, nil)
	// to_skip := shortestPath[1]

	for _, link := range anthill.Colony[start] {
		visited = make(map[string]bool)
		visited[start] = true
		path, _ := BFS(anthill, link.Name, end, &visited)
		// fmt.Println("path before fight", path)
		if len(path) != 0 {
			// fmt.Println("path to add ", path)
			for _, p := range path {
				// fmt.Println("rooms", anthill.Rooms[p])
				if anthill.Rooms[p].Explored {
					hold_disjoint_rooms = append(hold_disjoint_rooms, anthill.Rooms[p])
				}
			}
			if hold_disjoint_rooms != nil {
				// fmt.Println("before filter", paths)
				FilterPath(anthill, hold_disjoint_rooms, path, &paths)
				hold_disjoint_rooms = []*graph.Room{}
				// fmt.Println("after filter", paths)
				continue
			}
			for _, p := range path {
				anthill.Rooms[p].Explored = true
				anthill.Rooms[p].Path = path
				anthill.Rooms[p].Head = path[0]
				anthill.Rooms[p].Net = len(anthill.Colony[p])
				anthill.Rooms[p].Index = len(paths)

			}
			cost_after = CalculateCost(anthill, paths, [][]string{path})
			// fmt.Println("cost and paths", paths, path, cost_before, cost_after)
			if cost_before > cost_after {
				paths = append(paths, path)
			}
		}
	}
	return paths
}
