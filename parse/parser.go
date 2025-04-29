package parse

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/graph"
)

type errstr struct {
	data string
}

func (e *errstr) Error() string {
	return e.data
}

func Parse(file string, data *graph.Graph, coords *[]graph.Room) error {
	var found_start, found_end bool
	var no_space []string

	info, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	text := string(info)
	Splited := strings.FieldsFunc(text, func(r rune) bool {
		if r == 10 || r == 13 {
			return true
		}
		return false
	})

	for i := range Splited {

		if strings.HasPrefix(Splited[i], "L") {
			return &errstr{"Error: your room name starts with an L"}
		}
		if i == 0 {
			data.AntsNumber, err = strconv.Atoi(Splited[i])
			data.Ants = make([]graph.Ant, data.AntsNumber)
			if err != nil {
				return err
			}
			continue
		}
		if strings.HasPrefix(Splited[i], "##start") && i != len(Splited)-1 {
			found_start = true
			no_space = strings.Fields(Splited[i+1])
			if len(no_space) != 3 {
				return &errstr{"Error: room infos are invalid!"}
			}
			x, err := strconv.Atoi(no_space[1])
			if err != nil {
				return err
			}

			y, err := strconv.Atoi(no_space[2])
			if err != nil {
				return err
			}
			data.Start = &graph.Room{Name: no_space[0], X: x, Y: y}
			data.Rooms = append(data.Rooms, data.Start)
			continue
		}
		if strings.HasPrefix(Splited[i], "##end") && i != len(Splited)-1 {
			found_end = true
			no_space = strings.Fields(Splited[i+1])
			if len(no_space) != 3 {
				return &errstr{"Error: room infos are invalid!"}
			}
			x, err := strconv.Atoi(no_space[1])
			if err != nil {
				return err
			}

			y, err := strconv.Atoi(no_space[2])
			if err != nil {
				return err
			}
			data.End = &graph.Room{Name: no_space[0], X: x, Y: y}
			data.Rooms = append(data.Rooms, data.End)
			continue
		}
		if strings.HasPrefix(Splited[i], "#") {
			continue
		}
		if strings.Contains(Splited[i], "-") {
			copy := strings.Replace(Splited[i], "-", " ", 1)
			no_space = strings.Fields(copy)
			AddLink(no_space, data)
		}
		no_space = strings.Fields(Splited[i])

		AddRoom(no_space, data, coords)

	}
	if !found_start {
		fmt.Println("Error: the given file doesn't provide a start")
		return nil
	}
	if !found_end {
		fmt.Println("Error: the given file doesn't provide an end")
		return nil
	}
	// fmt.Printf("start:%v | end:%v | ants:%v\n", data.Start, data.End, data.Ants)
	return nil
}

func AddRoom(no_space []string, data *graph.Graph, coords *[]graph.Room) error {
	var x, y int
	var err error
	var r *graph.Room
	if len(no_space) != 3 {
		return &errstr{"Error: room infos are invalid!"}
	}

	x, err = strconv.Atoi(no_space[1])
	if err != nil {
		return err
	}

	y, err = strconv.Atoi(no_space[2])
	if err != nil {
		return err
	}
	r = &graph.Room{
		Name: no_space[0],
		X:    x,
		Y:    y,
	}
	data.Rooms = append(data.Rooms, r)
	data.Colony[r.Name] = []*graph.Room{}
	data.RoomNumber++
	*coords = append(*coords, *r)
	return nil
}

func AddLink(no_space []string, data *graph.Graph) error {
	if len(no_space) != 2 {
		return &errstr{"Error: some rooms have invalid links"}
	}

	data.Colony[no_space[0]] = append(data.Colony[no_space[0]], &graph.Room{Name: no_space[1]})
	data.Colony[no_space[1]] = append(data.Colony[no_space[1]], &graph.Room{Name: no_space[0]})
	data.LinkNumber++
	return nil
}
