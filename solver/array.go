package solver

// Append to a slice without mutating the original slice.
// The native append() function in Go mutates the original slice.
func appendImmutable[T any](slice []T, elems ...T) []T {
	newSlice := make([]T, len(slice)+len(elems))
	copy(newSlice, slice)
	copy(newSlice[len(slice):], elems)
	return newSlice
}

// TODO
// func mapArray[T any, U any](array []T, f func(T) U) []U {
// 	result := make([]U, len(array))
// 	for i, elem := range array {
// 		result[i] = f(elem)
// 	}
// 	return result
// }

// TODO
// func removeDuplicates[T comparable](array []T) []T {
// 	seen := make(map[T]bool)
// 	var result []T
// 	for _, elem := range array {
// 		if !seen[elem] {
// 			result = append(result, elem)
// 			seen[elem] = true
// 		}
// 	}
// 	return result
// }
