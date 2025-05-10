package main

import (
	"fmt"
	"lem-in/parse"
	"log"
	"os"
)

func main() {

	filename := os.Args[1]
	if len(os.Args) > 2 {
		log.Fatalln("You Have Entered To Many Args")
	}
	graph, _ := parse.FileToGraph(filename)
	fmt.Println(graph)
}
