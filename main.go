package main

import (
	"fmt"
	"lem-in/graph"
	"lem-in/helpers"
	extract "lem-in/parse"
	queue "lem-in/queue"
)

func main() {

	// graph := map[string][]string{
	// 	"start": {"room1", "room2"},
	// 	"room1": {"room6", "room2"},
	// 	"room2": {"room1", "room3"},
	// 	"room3": {"end"},
	// 	"room4": {"room1", "room3"},
	// 	"room5": {"room2", "room5"},
	// 	"room6": {"room1", "room5"},
	// 	"end":   {"room1", "room2"},
	// }
	// visted := make(map[string]bool)

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

	// fmt.Println("BFS :")
	// out := BFS(graph, visted, "start")
	// fmt.Println(out)
	var coords []graph.Room
	newGraph := graph.Graph{Colony: make(map[string][]string)}
	err := extract.Parse("./tests/example05.txt", &newGraph, &coords)
	for i, v := range newGraph.Colony {
		fmt.Println(i, v)
	}
	fmt.Println(err)
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
