package main

import (
	"fmt"
	"lem-in/parse"
	"log"
	"os"
)

func main() {
	// file := []byte(strings.Repeat("a", 65*1024))
	// os.WriteFile("long_line.txt", file, 0755)
	filename := os.Args[1]
	if len(os.Args) > 2 {
		log.Fatalln("You Have Entered To Many Args")
	}
	graph, _ := parse.FileToGraph(filename)
	for _, v := range graph.Rooms {
		fmt.Println((**v))
	}
}
