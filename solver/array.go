package solver

// Append a string to a slice without mutating the original slice.
// The native append() function in Go mutates the original slice.
func appendImmutable[T any](slice []T, val T) []T {
	newSlice := make([]T, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = val
	return newSlice
}
