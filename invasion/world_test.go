package invasion

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var InputOriginal = []string{
	"Foo north=Bar west=Baz south=Qu-ux",
	"Bar south=Foo west=Bee",
}

func ToReader(lines ...string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}

func ExampleWorld_Simulate() {
	w, err := InitWorld(ToReader(
		"cityOne south=cityTwo",
		"cityTwo south=cityThree",
		"cityThree",
	))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	// Not a complete blackbox testing, we replace random aliens with predetermined ones the hard way.
	w.aliens = []int{0, 2}

	w.Simulate(os.Stdout)
	// Output:
	// cityTwo has been destroyed by 2 aliens!
	// Simulation stopped after 1 rounds, remaining cities:
	// =======
	// cityOne
	// cityThree
}

func ExampleInitWorld_onlyParsing() {
	w, err := parseWorldText(ToReader(InputOriginal...))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	w.Print(os.Stdout)
	// Output:
	// Foo north=Bar south=Qu-ux west=Baz
	// Bar south=Foo west=Bee
	// Baz
	// Bee
	// Qu-ux
}

func TestParseCitySpecifiedTwice(t *testing.T) {
	_, err := InitWorld(ToReader(
		"one south=two",
		"one",
	))
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "already seen")
}

func TestParseCheckGraph(t *testing.T) {
	w, err := InitWorld(ToReader(
		"one south=two",
		"two south=three",
	))
	require.Nil(t, err)
	require.Equal(t, [][]int{{1}, {0, 2}, {1}}, w.graph)
}

func TestGenerateAliensFull(t *testing.T) {
	aliens, err := generateAliens(5, 5)
	require.Nil(t, err)
	require.Equal(t, []int{0, 1, 2, 3, 4}, aliens)
}

func TestGenerateAliensHalf(t *testing.T) {
	aliens, err := generateAliens(5, 10)
	require.Nil(t, err)
	require.Equal(t, 5, len(aliens))
}

func TestDestroyOneCity(t *testing.T) {
	w, err := InitWorld(ToReader(
		"one south=two",
		"two south=three",
		"three",
	))
	require.Nil(t, err)
	w.aliens = []int{0, 2}
	w.move()
	require.Equal(t, []int{1, 1}, w.aliens)
	require.Equal(t, []int{0, 2}, w.graph[1])
	w.fight(&bytes.Buffer{})
	require.Equal(t, []int{}, w.aliens)
	require.Equal(t, []int{}, w.graph[1])
}

func ExampleWorld_Simulate_oneRoundFight() {
	w, err := InitWorld(ToReader(
		"cityOne south=cityTwo",
		"cityTwo south=cityThree",
		"cityThree",
	))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	w.aliens = []int{0, 2}
	w.move()
	w.fight(os.Stdout)
	// Output:
	// cityTwo has been destroyed by 2 aliens!
}
