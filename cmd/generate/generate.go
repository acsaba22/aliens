package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/acsaba22/aliens/invasion"
)

const USAGE = `Usage:

go run ./cmd/generate N F > output.txt

C = integer, number of cities
F = integer [0..100], fill percentage. The tool will generate 4*C*F/100 roads randomly,
    which means that at 4 roads per city at 100 fill.
`

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatalf("Wrong number of arguments (%d). %s", len(args), USAGE)
	}

	c, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Failed to parse C (%s).\n%s", args[0], USAGE)
	}

	fmt.Fprintf(os.Stderr, "N = %d\n", c)

	f, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Failed to parse F (%s).\n%s", args[1], USAGE)
	}

	if f < 0 || 100 < f {
		log.Fatalf("F out of bounds (%d).\n%s", f, USAGE)
	}

	fmt.Fprintf(os.Stderr, "F = %d\n", f)

	w := invasion.GenerateWorld(c, f)

	w.Print(os.Stdout)
}
