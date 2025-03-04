package main

import (
	"fmt"
	"lem-in/graph"
	"lem-in/helpers"
	extract "lem-in/parse"
	queue "lem-in/queue"
)

func main() {
	var coords []graph.Room
	var paths [][]string
	var buffer []string
	newGraph := graph.Graph{Colony: make(map[string][]string)}
	err := extract.Parse("./tests/example01.txt", &newGraph, &coords)
	v := make(map[string]bool)
	// for k, v := range newGraph.Colony {
	// 	fmt.Println(k, v)
	// }
	fmt.Println("start: ", newGraph.Colony[newGraph.Start])
	DFS(&newGraph, v, newGraph.Start, &paths, &buffer)
	for _, v := range paths {
		fmt.Println(v)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DFS(graph *graph.Graph, visited map[string]bool, current string, paths *[][]string, buffer *[]string) {
	// Mark the current room as visited
	visited[current] = true

	// Add the current room to the buffer
	*buffer = append(*buffer, current)

	// If the current room is the end room, save the path
	if current == graph.End {
		// Save a copy of the buffer into paths
		path := make([]string, len(*buffer))
		copy(path, *buffer)
		*paths = append(*paths, path)
	} else {
		// Explore all unvisited neighbors
		for _, link := range graph.Colony[current] {
			if !visited[link] {
				DFS(graph, visited, link, paths, buffer)
			}
		}
	}

	// Backtrack:
	// Remove the current room from the buffer and mark it as unvisited
	*buffer = (*buffer)[:len(*buffer)-1]
	visited[current] = false
}

func BFS(graph map[string][]string, visted map[string]bool, start string) []string {
	var path []string
	list := queue.Queue{}
	list.Enqueue(start)
	visted[start] = false

	for !list.IsEmpty() {
		list.Print()
		room := list.Dequeue()
		path = append(path, room)

		visted[room] = true
		if helpers.Contains("end", graph[room]) {
			path = append(path, "end")
			return path
		}
		for _, link := range graph[room] {
			if !visted[link] {
				visted[link] = true
				list.Enqueue(link)
			}
		}

	}
	return nil
}
