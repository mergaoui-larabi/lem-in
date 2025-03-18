package generate

import (
	"fmt"
	"math/rand"
	"os"
)

func Generate() {
	numRooms := 1000
	linksPerRoom := 100
	antnum := "10"
	filename := "./tests/roomswithparams.txt"

	err := generateLemInFile(filename, antnum, numRooms, linksPerRoom)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File '%s' generated successfully!\n", filename)
}

func generateLemInFile(filename, antnum string, numRooms, linksPerRoom int) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(antnum + "\n")

	file.WriteString("##start\n")
	file.WriteString("S 0 1\n")

	file.WriteString("##end\n")
	file.WriteString("E 1000 1\n")

	for i := 1; i <= numRooms; i++ {
		file.WriteString(fmt.Sprintf("%d %d %d\n", i, i, i))
	}

	for i := 1; i <= numRooms; i++ {

		links := generateRandomLinks(i, numRooms, linksPerRoom)
		for _, link := range links {
			file.WriteString(fmt.Sprintf("%d-%d\n", i, link))
		}
	}

	file.WriteString("S-1\n")   
	file.WriteString("1000-E\n") 
	return nil
}

func generateRandomLinks(currentRoom, numRooms, linksPerRoom int) []int {
	
	allRooms := make([]int, 0, numRooms-1)
	for i := 1; i <= numRooms; i++ {
		if i != currentRoom {
			allRooms = append(allRooms, i)
		}
	}


	rand.Shuffle(len(allRooms), func(i, j int) {
		allRooms[i], allRooms[j] = allRooms[j], allRooms[i]
	})

	
	return allRooms[:linksPerRoom]
}
