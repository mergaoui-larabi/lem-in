package graph

type Room struct {
	Name string
	X    int
	Y    int
}

type Graph struct {
	Colony map[string][]string
}
