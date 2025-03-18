package graph

type Room struct {
	Name string
	//TODO: Nieghbours []string  further implementation
	X    int
	Y    int
}

type Graph struct {
	//TODO start , end *Room further implementation
	Ants   int
	Start  string
	End    string
	Colony map[string][]string
}
