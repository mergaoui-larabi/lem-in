package solver

import (
	"lem-in/graph"
	"lem-in/tools"
)

func FindPaths(graph map[string][]*graph.Room, start, end string) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	for _, link := range graph[start] {
		visited[start] = true
		path := tools.BFS(graph, link.Name, end, &visited)
		if len(path) != 0 {
			paths = append(paths, path)
		}
	}
	return paths
}
