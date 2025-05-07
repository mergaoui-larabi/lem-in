package parse

import (
	"bufio"
	"container/list"
	"fmt"
	constant "lem-in/const"
	dsa "lem-in/dsa"
	"os"
	"strconv"
)

// implement the error interface{}
type ErrorMessage struct {
	Msg string
}

func (e *ErrorMessage) Error() string {
	return e.Msg
}

func FileExist(filename string) error {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return &ErrorMessage{Msg: "ERROR:The File " + filename + " Doesn't Exist in The Specifeid Path"}
	}
	if info.IsDir() {
		return &ErrorMessage{Msg: "ERROR:You Have Entered a Directory Path Istead Of a File Path"}
	}
	return nil
}

func FileToGraph(filename string) (*dsa.Antfarm, error) {
	var err error
	err = FileExist(filename)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, &ErrorMessage{Msg: constant.ErrFileIssue}
	}
	defer file.Close()
	graph := dsa.NewAntFarm()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := ParseLine(graph, scanner.Text())
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

	return graph, nil
}

func ParseLine(graph *dsa.Antfarm, line string) error {
	if IsComment(line) || line == "" {
		return nil
	}
	switch graph.Data.Phase {
	case constant.AntsField:
		return ParseAntNumber(graph, line)
	case constant.RoomsField:
		return ParseRooms(graph, line)
	case constant.LinksField:
		return ParseLinks(graph, line)
	}
	return nil
}

func ParseAntNumber(graph *dsa.Antfarm, line string) error {
	n, err := strconv.Atoi(line)
	if err != nil {
		return &ErrorMessage{Msg: constant.ErrAnts}
	}
	if n <= 0 || n > (1<<31-1) {
		return &ErrorMessage{Msg: constant.ErrAnts}
	}
	graph.Nants = n
	graph.Data.Phase = constant.RoomsField
	return nil
}

func ParseRooms(graph *dsa.Antfarm, line string) error {
	if IsStart(line) && !graph.Data.StartFound {
		graph.Data.StartFound = true
	} else if IsEnd(line) && !graph.Data.EndFound {
		graph.Data.EndFound = true
	} else {
		room, err := GetRoom(graph, line)
		if err != nil {
			return err
		}
		if graph.Rooms[room] != nil {
			return &ErrorMessage{Msg: "ERROR: the room " + room + " is dupplicated"}
		}
		if room != "" {
			node := &dsa.Room{Name: room, Parent: "L", Edges: make(map[string]byte)}
			if graph.Data.StartFound && graph.Start == "" {
				graph.Start = room
			} else if graph.Data.EndFound && graph.End == "" {
				graph.End = room
			}
			graph.Rooms[room] = &node
			graph.Exits = list.New()
		} else if graph.Start != "" && graph.End != "" {
			graph.Data.Phase = constant.LinksField
			graph.Data.Coords = nil // free up memory from rooms coords because they are unusable
			return ParseLine(graph, line)
		} else {
			return &ErrorMessage{Msg: constant.ErrNoStart + " or " + constant.ErrNoEnd}
		}
	}
	return nil
}

func ParseLinks(graph *dsa.Antfarm, line string) error {
	firstRoom, secondRoom := GetLink(line)
	if firstRoom == "" || secondRoom == "" {
		return &ErrorMessage{Msg: constant.ErrLink}
	}
	if firstRoom == secondRoom {
		return &ErrorMessage{Msg: constant.ErrLink}
	}
	if graph.Rooms[firstRoom] == nil || graph.Rooms[secondRoom] == nil {
		return &ErrorMessage{Msg: constant.ErrLink}
	}
	node1 := *graph.Rooms[firstRoom]
	node1.Edges[secondRoom] = 1

	node2 := *graph.Rooms[secondRoom]
	node2.Edges[firstRoom] = 1
	return nil
}
