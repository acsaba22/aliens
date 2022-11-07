package invasion

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var DirectionNames = []string{"north", "south", "east", "west"}
var DirectionIds map[string]Direction

func init() {
	DirectionIds = map[string]Direction{}
	for d := North; d < 4; d++ {
		DirectionIds[DirectionNames[d]] = d
	}
}

type cityDescription struct {
	name       string
	neighbours [4]string
}

type World struct {
	cities []cityDescription
}

func ParseWorld(reader io.Reader) (*World, error) {
	ret := World{}

	citiesSeen := map[string]bool{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lineTokens := strings.Split(line, " ")
		cityFrom := lineTokens[0]

		if citiesSeen[cityFrom] {
			return nil, fmt.Errorf("Main city already seen [line: %s ]", line)
		}
		citiesSeen[cityFrom] = true

		neighbours := [4]string{}
		for _, road := range lineTokens[1:] {
			roadTokens := strings.Split(road, "=")
			if len(roadTokens) != 2 {
				return nil, fmt.Errorf(
					"Road shoul have the format 'direction=city' [roadToken: %s, line: %s]", road, line)
			}
			direction, ok := DirectionIds[strings.ToLower(roadTokens[0])]
			if !ok {
				return nil, fmt.Errorf("Unreckognized direction [roadToken: %s, line: %s]", road, line)
			}
			if neighbours[direction] != "" {
				return nil, fmt.Errorf("Direction %s repeated in line [line: %s]", roadTokens[0], line)
			}
			cityTo := roadTokens[1]
			if cityTo == "" {
				return nil, fmt.Errorf("Empt city name [roadToken: %s, line: %s]", road, line)
			}

			neighbours[direction] = cityTo
		}
		ret.cities = append(ret.cities, cityDescription{
			name:       cityFrom,
			neighbours: neighbours})
	}
	return &ret, nil
}

func (w World) Print(writer io.Writer) {
	for _, city := range w.cities {
		fmt.Fprintf(writer, "%s", city.name)
		for d := North; d < 4; d++ {
			neigh := city.neighbours[d]
			if neigh != "" {
				fmt.Fprintf(writer, " %s=%s", DirectionNames[d], neigh)
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
