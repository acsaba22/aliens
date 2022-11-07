package invasion

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomName() string {
	const length = 10
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// c: number of citiees
// f: fill factor [0..100) how dense the graph should be.
//    number of roads is going to be 4*c*f/100
func GenerateWorld(c int, f int) World {
	w := World{}
	w.cities = make([]cityDescription, c)

	type cityRoadCount struct {
		cityId    int
		roadCount int
	}

	// cities with open outbound connections
	freeCities := make([]cityRoadCount, c)

	for i := 0; i < c; i++ {
		w.cities[i].name = randomName()
		freeCities[i] = cityRoadCount{cityId: i, roadCount: 0}
	}

	k := 4 * c * f / 100
	for i := 0; i < k; i++ {
		x := rand.Intn(len(freeCities))
		var fromCity *cityRoadCount = &freeCities[x]
		toId := rand.Intn(c)
		w.cities[fromCity.cityId].neighbours[fromCity.roadCount] = w.cities[toId].name

		fromCity.roadCount++
		if fromCity.roadCount == 4 {
			// city outbound connections filled, delete from free city list
			removeElement(&freeCities, x)
		}
	}
	return w
}
