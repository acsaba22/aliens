package invasion

// Removes item index k.
// Modifies the slice original slice.
func removeElement[T any](s *[]T, k int) {
	copy((*s)[k:], (*s)[k+1:])
	(*s) = (*s)[:len(*s)-1]
}
