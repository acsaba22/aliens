package main

import (
	"log"
	"os"

	"github.com/acsaba22/aliens/invasion"
)

func main() {
	w, err := invasion.ParseWorld(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	w.Print(os.Stdout)
}
