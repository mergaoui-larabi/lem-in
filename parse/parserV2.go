package parse

import (
	"bufio"
	"fmt"
	"lem-in/graph"
	"os"
	"strconv"
	"strings"
)


func ParseV(file string, data *graph.Graph, coords *[]graph.Room) error {
	f, err := os.Open(file)
	if err != nil {
		return &errstr{fmt.Sprintf("ERROR: %v", err)}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	data.Colony = make(map[string][]string)
	coordinates := make(map[string]struct{})
	var antsParsed, startParsed, endParsed bool
	var lineNum int

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "##start"):
			if startParsed {
				return &errstr{"ERROR: multiple start rooms defined"}
			}
			startParsed = true
			if err := processSpecialRoom(scanner, data, coords, coordinates, &lineNum, &data.Start); err != nil {
				return err
			}

		case strings.HasPrefix(line, "##end"):
			if endParsed {
				return &errstr{"ERROR: multiple end rooms defined"}
			}
			endParsed = true
			if err := processSpecialRoom(scanner, data, coords, coordinates, &lineNum, &data.End); err != nil {
				return err
			}

		case strings.HasPrefix(line, "#"):
			continue

		case !antsParsed:
			if err := parseAnts(line, data); err != nil {
				return &errstr{fmt.Sprintf("ERROR: line %d: %v", lineNum, err)}
			}
			antsParsed = true

		case strings.Contains(line, "-"):
			if err := parseLink(line, data); err != nil {
				return &errstr{fmt.Sprintf("ERROR: line %d: %v", lineNum, err)}
			}

		default:
			if err := parseRoom(line, data, coords, coordinates,nil); err != nil {
				return &errstr{fmt.Sprintf("ERROR: line %d: %v", lineNum, err)}
			}
		}
	}

	if err := validateInput(antsParsed, startParsed, endParsed, data); err != nil {
		return err
	}

	return nil
}

func processSpecialRoom(scanner *bufio.Scanner, data *graph.Graph, coords *[]graph.Room, 
	coordinates map[string]struct{}, lineNum *int, target *string) error {
	if !scanner.Scan() {
		return &errstr{"ERROR: special room declaration not followed by room definition"}
	}
	*lineNum++
	return parseRoom(strings.TrimSpace(scanner.Text()), data, coords, coordinates, target)
}

func parseAnts(line string, data *graph.Graph) error {
	ants, err := strconv.Atoi(line)
	if err != nil || ants <= 0 {
		return fmt.Errorf("invalid number of ants")
	}
	data.Ants = ants
	return nil
}

func parseRoom(line string, data *graph.Graph, coords *[]graph.Room, 
	coordinates map[string]struct{}, target *string) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("invalid room format")
	}

	if strings.HasPrefix(parts[0], "L") {
		return fmt.Errorf("room name starts with 'L'")
	}

	if _, exists := data.Colony[parts[0]]; exists {
		return fmt.Errorf("duplicate room %q", parts[0])
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid X coordinate")
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid Y coordinate")
	}

	coordKey := fmt.Sprintf("%d,%d", x, y)
	if _, exists := coordinates[coordKey]; exists {
		return fmt.Errorf("duplicate coordinates %q", coordKey)
	}

	data.Colony[parts[0]] = []string{}
	coordinates[coordKey] = struct{}{}
	*coords = append(*coords, graph.Room{Name: parts[0], X: x, Y: y})

	if target != nil {
		*target = parts[0]
	}
	return nil
}

func parseLink(line string, data *graph.Graph) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid link format")
	}

	from, to := parts[0], parts[1]
	if from == to {
		return fmt.Errorf("self-linking room %q", from)
	}

	if _, exists := data.Colony[from]; !exists {
		return fmt.Errorf("undefined room %q", from)
	}

	if _, exists := data.Colony[to]; !exists {
		return fmt.Errorf("undefined room %q", to)
	}

	// Check for duplicate links
	for _, room := range data.Colony[from] {
		if room == to {
			return fmt.Errorf("duplicate link %q-%q", from, to)
		}
	}

	data.Colony[from] = append(data.Colony[from], to)
	data.Colony[to] = append(data.Colony[to], from)
	return nil
}

func validateInput(antsParsed, startParsed, endParsed bool, data *graph.Graph) error {
	if !antsParsed {
		return &errstr{"ERROR: missing ant count"}
	}
	if !startParsed {
		return &errstr{"ERROR: missing start room"}
	}
	if !endParsed {
		return &errstr{"ERROR: missing end room"}
	}
	if data.Start == data.End {
		return &errstr{"ERROR: start and end rooms are the same"}
	}
	return nil
}