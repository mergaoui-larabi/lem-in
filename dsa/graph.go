package dsa

import "container/list"

type Antfarm struct {
	Exits      *list.List
	Rooms      map[string]**Room
	Start, End string
	Nants      int
	Data       Info
}

type Room struct {
	Name                string
	Edges               map[string]byte
	Parent              string
	EdgeIn, EdgeOut     string
	PrinceIn, PrinceOut int
	CostIn, CostOut     int
	Split               bool
}

type Info struct {
	StartFound, EndFound bool
	Phase                byte
	Coords               map[[2]int]bool
}

func NewAntFarm() *Antfarm {
	return &Antfarm{Rooms: make(map[string]**Room), Data: Info{Coords: make(map[[2]int]bool)}}
}
