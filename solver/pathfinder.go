package solver

import (
	"lem-in/graph"
	"lem-in/tools"
)

func FindPaths(graph *graph.Graph, start, end string) [][]string {
	var paths [][]string
	for _, link := range graph.Colony[start] {
		visited := make(map[string]bool)
		visited[start] = true
		path := tools.BFS(graph.Colony, link.Name, end, &visited)
		if len(path) != 0 {
			for _, room := range graph.Rooms {
				if existInPath(room.Name, path) {
					room.Explored = true
					room.Path = path
					room.Net = len(graph.Colony[room.Name])
				}
			}
			paths = append(paths, path)
		}
	}
	return paths
}

func existInPath(room string, path []string) bool {
	for _, v := range path {
		if v == room {
			return true
		}
	}
	return false
}
