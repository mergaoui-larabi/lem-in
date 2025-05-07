package parse

import (
	constant "lem-in/const"
	"lem-in/dsa"
	"strconv"
	"strings"
)

func IsComment(s string) bool {
	return strings.HasPrefix(s, "#") && !IsStart(s) && !IsEnd(s)
}

func IsStart(s string) bool {
	return s == "##start"
}

func IsEnd(s string) bool {
	return s == "##end"
}

func IsRoom(s string) bool {
	return strings.HasPrefix(s, "L")
}

func IsCoords(x, y string, graph *dsa.Antfarm) bool {
	xn, err := strconv.Atoi(x)
	yn, err := strconv.Atoi(y)
	if err != nil {
		return false
	}
	if ok := graph.Data.Coords[[2]int{xn, yn}]; ok {
		return false
	}
	graph.Data.Coords[[2]int{xn, yn}] = true
	return true
}

func GetRoom(graph *dsa.Antfarm, s string) (string, error) {
	split := strings.Split(s, " ")
	if len(split) != 3 {
		return "", nil // to dodge links info if they comes first in order
	}
	if IsRoom(split[0]) {
		return "", &ErrorMessage{Msg: constant.ErrRoomName}
	}
	if !IsCoords(split[1], split[2], graph) {
		return "", &ErrorMessage{Msg: constant.ErrCoord + " : dupplicated " + split[0]}
	}
	return split[0], nil
}

func GetLink(s string) (string, string) {
	split := strings.Split(s, "-")
	if len(split) != 2 {
		return "", ""
	}
	return split[0], split[1]
}
