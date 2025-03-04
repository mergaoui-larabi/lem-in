package graph

type Room struct {
	Name string
	X    int
	Y    int
}

type Graph struct {
	Ants   int
	Start  string
	End    string
	Colony map[string][]string
}
