package solver

// Append to a slice without mutating the original slice.
// The native append() function in Go mutates the original slice.
func appendImmutable[T any](slice []T, elems ...T) []T {
	newSlice := make([]T, len(slice)+len(elems))
	copy(newSlice, slice)
	copy(newSlice[len(slice):], elems)
	return newSlice
}
