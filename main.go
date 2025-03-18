package main

import (
	"fmt"
	"lem-in/graph"
	extract "lem-in/parse"
	queue "lem-in/queue"
)

func main() {
	var coords []graph.Room

	newGraph := graph.Graph{Colony: make(map[string][]string)}
	err := extract.Parse("./tests/message.txt", &newGraph, &coords)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(newGraph.Colony)
	paths := FindPaths(newGraph.Colony, newGraph.Start, newGraph.End)
	fmt.Println("BFS paths :", paths)
}

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

func FindPaths(graph map[string][]string, start, end string) [][]string {
	var paths [][]string
	for _, link := range graph[start] {
		visited := make(map[string]bool)
		visited[start] = true
		_, path := BFS(graph, visited, link, end)
		paths = append(paths, path)
	}
	return paths
}
