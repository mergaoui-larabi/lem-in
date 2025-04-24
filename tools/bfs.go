package tools

import (
	"lem-in/queue"
)

func BFS(graph map[string][]string, start, end string, visited *map[string]bool) []string {
	parent := make(map[string]string)
	parent[start] = ""
	list := queue.Queue{}
	list.Enqueue(start)

	(*visited)[start] = true

	for !list.IsEmpty() {
		room := list.Dequeue()

		if room == end {
			(*visited) = make(map[string]bool)
			path := reconstructPath(parent, end)
			for i := 0; i < len(path)-1; i++ {
				(*visited)[path[i]] = true
			}
			return path
		}

		for _, link := range graph[room] {
			if !(*visited)[link] {
				(*visited)[link] = true
				list.Enqueue(link)
				parent[link] = room
			}
		}
	}

	return []string{}
}

func reconstructPath(parent map[string]string, end string) []string {
	path := []string{}
	for node := end; node != ""; node = parent[node] {
		path = append([]string{node}, path...) // Prepend to maintain order
	}
	return path
}
