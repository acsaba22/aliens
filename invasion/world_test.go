package invasion

import (
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
