package invasion

import (
	"bytes"
	_ "embed"
	"log"
	"strings"
	"testing"
)

//go:embed benchmark_input_1000_70.txt
var bigFile string

// Run benchmark tests with
// $ go test  -v -run=NONE -bench=. ./...
//
// For the 1000 cities example the parsing is 1.4 ms and the whole
// simulation is 4.3 ms.
//
// BenchmarkParse-8             835           1452513 ns/op
// BenchmarkAll-8               274           4314164 ns/op

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(bigFile)
		_, err := InitWorld(reader)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(bigFile)
		w, err := InitWorld(reader)
		if err != nil {
			log.Fatal(err)
		}

		err = w.GenerateAliens(800)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		w.Simulate(&bytes.Buffer{})
	}
}
