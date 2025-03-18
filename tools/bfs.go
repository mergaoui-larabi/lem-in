package tools

import (
	"fmt"
	"lem-in/queue"
)

func BFS(graph map[string][]string, visited map[string]bool, start, end string) (map[string]string, []string) {
	parent := make(map[string]string)
	parent[start] = ""
	list := queue.Queue{}
	list.Enqueue(start)

	visited[start] = true

	for !list.IsEmpty() {
		room := list.Dequeue()

		if room == end {
			fmt.Println("parent map:", parent)
			return parent, reconstructPath(parent, start, end)
		}

		for _, link := range graph[room] {
			if !visited[link] {
				visited[link] = true
				list.Enqueue(link)
				parent[link] = room
			}
		}
	}

	return parent, []string{}
}

func reconstructPath(parent map[string]string, start, end string) []string {
	path := []string{}
	for node := end; node != ""; node = parent[node] {
		path = append([]string{node}, path...) // Prepend to maintain order
	}
	return path
}
