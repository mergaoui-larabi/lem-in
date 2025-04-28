package graph

type Room struct {
	Name string
	X    int
	Y    int
	Free bool
}

type Ant struct {
	Name    string
	Path    int
	Movable bool
	Current string
	UniquePath []string
}

type Graph struct {
	AntsNumber int
	RoomNumber int
	LinkNumber int
	Ants       []Ant
	Rooms      []*Room
	Start      *Room
	End        *Room
	Colony     map[string][]string
}
