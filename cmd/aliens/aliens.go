package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/acsaba22/aliens/invasion"
)

const USAGE = `Usage:

cat input.txt | go run ./cmd/aliens N

N = integer, number of aliens, should be smaller or equal then number of cities
    in the input file
`

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatalf("Expect 1 argument %d given.\n%s", len(args), USAGE)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Failed to parse N (%s).\n%s", args[0], USAGE)
	}

	w, err := invasion.InitWorld(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	err = w.GenerateAliens(n)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	w.Simulate(os.Stdout)
}
