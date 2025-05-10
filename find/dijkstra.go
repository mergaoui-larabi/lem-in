package find

import (
	"lem-in/dsa"
)

const (
	MaxRooms = 100000
)

func Dijkstra(graph *dsa.Antfarm) bool {

	return (*graph.Rooms[graph.End]).EdgeIn != "L"
}

func GraphReset(graph *dsa.Antfarm) {
	var node *dsa.Room
	for _, v := range graph.Rooms {
		node = *v
		node.CostIn = MaxRooms
		node.CostOut = MaxRooms
		node.EdgeIn = "L"
		node.EdgeOut = "L"
	}
	(*graph.Rooms[graph.Start]).CostIn = 0
	(*graph.Rooms[graph.Start]).CostOut = 0
}

func EdgeIn(graph *dsa.Antfarm, pq *dsa.Queue, in, out string) {
	nodeIn := (*graph.Rooms[in])
	nodeOut := (*graph.Rooms[out])
	if out == graph.Start || in == graph.End || nodeOut.Parent == in {
		return
	}
	if nodeIn.Parent == out && nodeIn.CostIn < MaxRooms && (1+nodeOut.CostOut > nodeIn.CostIn+nodeIn.PrinceIn-nodeOut.PrinceOut) {
		return
	} else if nodeIn.Parent != out && nodeIn.CostIn < MaxRooms && (-1+nodeOut.CostIn > nodeIn.CostOut+nodeIn.PrinceOut-nodeOut.PrinceIn) {
		return
	}
}

func EdgeOut(graph *dsa.Antfarm, pq *dsa.Queue, out string) {
	return
}
