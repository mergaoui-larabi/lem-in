package solver

import (
	"lem-in/graph"
	"lem-in/queue"
)

func BFS(graph *graph.Graph, start, end string, visited *map[string]bool) ([]string, bool) {
	parent := make(map[string]string)
	dupp := false
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
			path = path[:len(path)-1]
			return path, dupp
		}

		for _, link := range graph.Colony[room] {
			if !(*visited)[link.Name] {
				if link.Explored {
					dupp = true
				}
				(*visited)[link.Name] = true
				list.Enqueue(link.Name)
				parent[link.Name] = room
			}
		}
	}

	return []string{}, dupp
}

func reconstructPath(parent map[string]string, end string) []string {
	path := []string{}
	for node := end; node != ""; node = parent[node] {
		path = append([]string{node}, path...) // Prepend to maintain order
	}
	return path
}
