package invasion

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
)

const MAXROUNDCOUNT = 10000

// Generates n aliens in different cities
func generateAliens(n, cityNum int) ([]int, error) {
	if cityNum < n {
		return nil, fmt.Errorf(
			"Number of aliens should be smaller or equal then the number of cities [%d vs %d]",
			n, cityNum)
	}

	cityIds := make([]int, cityNum)
	for i := 0; i < cityNum; i++ {
		cityIds[i] = i
	}

	ret := make([]int, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(len(cityIds))
		ret[i] = cityIds[x]
		removeElement(&cityIds, x)
	}
	sort.Ints(ret)
	return ret, nil
}

func (w *World) GenerateAliens(n int) error {
	var err error
	w.aliens, err = generateAliens(n, len(w.cities))
	return err
}

// Moves all aliens to a neighbouring city if possible.
func (w *World) move() {
	for i, cityId := range w.aliens {
		neighs := w.graph[cityId]
		if len(neighs) == 0 {
			continue
		}
		newLocation := neighs[rand.Intn(len(neighs))]
		w.aliens[i] = newLocation
	}
	sort.Ints(w.aliens)
}

// Finds aliens in the same city and removes the aliens and the cities
func (w *World) fight(writer io.Writer) {
	// find adjacent repeating values, remove them and destroy the cities
	for start := 0; start < len(w.aliens); {
		sameUntil := start + 1
		cityId := w.aliens[start]
		for ; sameUntil < len(w.aliens) && w.aliens[sameUntil] == cityId; sameUntil++ {
		}
		if 1 < sameUntil-start {
			// Destroy city and aliens!

			// TODO: check with client if alien ID's are necesarry.
			//   Requested error message by client was
			//   "CITY has been destroyed by alien NN and alien NN!"
			// But current design treats aliens anonymously.
			fmt.Fprintf(writer, "%s has been destroyed by %d aliens!\n",
				w.cities[cityId].name, sameUntil-start)
			w.destroyCity(cityId)
			removeElements(&w.aliens, start, sameUntil)
		} else {
			start++
		}
	}
}

func (w *World) Simulate(writer io.Writer) {
	round := 0
	for ; 0 < len(w.aliens); round++ {
		if MAXROUNDCOUNT <= round {
			fmt.Fprintf(writer, "Maximum iteration reached, stopping\n")
			break
		}
		w.move()
		w.fight(writer)
	}
	fmt.Fprintf(
		writer,
		"Simulation stopped after %d rounds, %d remaining cities:\n=======\n",
		round,
		len(w.cities)-len(w.isDestroyed))
	w.Print(writer)
}
