package invasion

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
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
	cities       []cityDescription
	cityNameToId map[string]int
	graph        [][]int
}

func InitWorld(reader io.Reader) (*World, error) {
	w, err := parseWorldText(reader)
	if err != nil {
		return nil, err
	}

	w.calculateGraph()
	return w, nil
}

func parseWorldText(reader io.Reader) (*World, error) {
	ret := World{}

	ret.cityNameToId = map[string]int{}
	seenNeighbours := map[string]bool{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lineTokens := strings.Split(line, " ")
		cityFrom := lineTokens[0]

		_, alreadySpecifiedCity := ret.cityNameToId[cityFrom]
		if alreadySpecifiedCity {
			return nil, fmt.Errorf("Main city already seen [line: %s ]", line)
		}
		ret.cityNameToId[cityFrom] = len(ret.cities)

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
			seenNeighbours[cityTo] = true
		}
		ret.cities = append(ret.cities, cityDescription{
			name:       cityFrom,
			neighbours: neighbours})
	}
	unspecifiedNeighbours := []string{}
	for city := range seenNeighbours {
		_, specified := ret.cityNameToId[city]
		if !specified {
			unspecifiedNeighbours = append(unspecifiedNeighbours, city)
		}
	}
	sort.Strings(unspecifiedNeighbours)
	for _, city := range unspecifiedNeighbours {
		ret.cityNameToId[city] = len(ret.cities)
		ret.cities = append(ret.cities, cityDescription{name: city})
	}
	return &ret, nil
}

func (w *World) calculateGraph() {
	w.graph = make([][]int, len(w.cities))
	for fromId, city := range w.cities {
		for _, neigh := range city.neighbours {
			if neigh == "" {
				continue
			}

			toId, ok := w.cityNameToId[neigh]
			if !ok {
				// This should not happen, it's not user error but a programmer error.
				// We don't report error but exit with signal.
				// Depending on the project's task this might be not a good solution.
				log.Fatal("Logic error: map doesn't contain id")
			}
			w.graph[fromId] = append(w.graph[fromId], toId)
			w.graph[toId] = append(w.graph[toId], fromId)
		}
	}
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
