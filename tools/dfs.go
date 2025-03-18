package tools

import "lem-in/graph"

func DFS(graph *graph.Graph, visited map[string]bool, current string, paths *[][]string, buffer *[]string) {

	visited[current] = true
	*buffer = append(*buffer, current)

	if current == graph.End {
		path := make([]string, len(*buffer))
		copy(path, *buffer)
		*paths = append(*paths, path)
	} else {

		for _, link := range graph.Colony[current] {
			if !visited[link] {
				DFS(graph, visited, link, paths, buffer)
			}
		}
	}

	// Backtrack:

	*buffer = (*buffer)[:len(*buffer)-1]
	visited[current] = false
}
