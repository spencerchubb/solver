package solver

// Append a string to a slice without mutating the original slice.
// The native append() function in Go mutates the original slice.
func appendImmutable(slice []string, str string) []string {
	newSlice := make([]string, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = str
	return newSlice
}
