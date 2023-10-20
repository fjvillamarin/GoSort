// Package sort defines an interface for sorting algorithms
package sort

// Sorter is an interface that any sorting algorithm should implement.
// It provides a single method, Sort, that takes a slice of integers
// and returns a new sorted slice of integers.
type Sorter interface {
	// Sort takes a slice of integers and returns a new sorted slice.
	Sort([]int) []int
}
