package invasion

import (
	"fmt"
	"os"
	"strings"
)

const InputOriginal = `Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
`

func ExampleParseWorld() {
	w, err := ParseWorld(strings.NewReader(InputOriginal))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	w.Print(os.Stdout)
	// Output:
	// Foo north=Bar south=Qu-ux west=Baz
	// Bar south=Foo west=Bee
}
