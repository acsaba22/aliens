package invasion

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Removes item index k.
// Modifies the slice original slice.
func removeElement[T any](s *[]T, k int) {
	copy((*s)[k:], (*s)[k+1:])
	(*s) = (*s)[:len(*s)-1]
}

// Removes items with index [from, until)
// Modifies the slice original slice.
func removeElements[T any](s *[]T, from, until int) {
	copy((*s)[from:], (*s)[until:])
	(*s) = (*s)[:len(*s)+from-until]
}

func removeValues[T comparable](s *[]T, v T) {
	writeDestination := 0
	for i := 0; i < len(*s); i++ {
		if (*s)[i] != v {
			(*s)[writeDestination] = (*s)[i]
			writeDestination++
		}
	}
	(*s) = (*s)[:writeDestination]
}
