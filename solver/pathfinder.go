package solver

import "lem-in/tools"

func FindPaths(graph map[string][]string, start, end string) [][]string {
	var paths [][]string
	for _, link := range graph[start] {
		visited := make(map[string]bool)
		visited[start] = true
		_, path := tools.BFS(graph, visited, link, end)
		paths = append(paths, path)
	}
	return paths
}
