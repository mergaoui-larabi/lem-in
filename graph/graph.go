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
	Ants       []Ant
	AntsNumber int
	Rooms      []*Room
	Start      *Room
	End        *Room
	Colony     map[string][]string
	RoomNumber int
	LinkNumber int
}
