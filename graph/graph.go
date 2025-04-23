package graph

type Room struct {
	Name string
	X    int
	Y    int
	FREE bool
}

type Ant struct {
	Name    string
	Path    int
	Movable bool
	Current string
}

type Graph struct {
	// TODO start , end *Room further implementation
	Ants       []Ant
	AntsNumber int
	Start      *Room
	End        *Room
	Colony     map[string][]*Room
	RoomNumber int
	LinkNumber int
}
