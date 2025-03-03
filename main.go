package main

import (
	"fmt"
	queue "lem-in/dsa"
	"lem-in/helpers"
)

func main() {

	graph := map[string][]string{
		"start": {"room1", "room2"},
		"room1": {"room2", "end"},
		"room2": {"room1", "end"},
		"end":   {"room1", "room2"},
	}
	visted := make(map[string]bool)

	// tst := queue.Queue{}
	// tst.Enqueue("start")
	// tst.Enqueue("dhhd")
	// fmt.Println("len 		:", tst.Length)
	// tst.Print()
	// fmt.Println("2	", tst.Dequeue())
	// fmt.Println("2	", tst.Dequeue())
	// fmt.Println("len 		:", tst.Length)
	// tst.Print()

	// fmt.Println("DFS :")
	// DFS(graph, visted, "start")

	fmt.Println("BFS :")
	out := BFS(graph, visted, "start")
	fmt.Println(out)
}

func DFS(graph map[string][]string, visted map[string]bool, start string) {
	// if start != "end" && start != "start" {
	// }
	visted[start] = true

	fmt.Println(start)
	for _, link := range graph[start] {
		if !visted[link] {
			DFS(graph, visted, link)
		}
	}
}

func BFS(graph map[string][]string, visted map[string]bool, start string) []string {
	var path []string
	list := queue.Queue{}
	list.Enqueue(start)
	visted[start] = false

	for !list.IsEmpty() {
		room := list.Dequeue()
		path = append(path, room)
		if !visted[room] {
			visted[room] = true
			if helpers.Contains("end", graph[room]) {
				return path
			}
			for _, link := range graph[room] {
				if !visted[link] {
					visted[link] = true
					list.Enqueue(link)
				}
			}
		}
	}
	return nil
}
