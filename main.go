// Package main provides a command-line utility for sorting arrays using different algorithms.
package main

import (
	"flag"
	"fmt"
	"gosort/sort"
	"gosort/sort/mergesort"
	"gosort/sort/quicksort"
	"math/rand"
	"strconv"
)

const (
	QuickSort = "quicksort"
	MergeSort = "mergesort"
)

var size int
var max int
var algo string

// main is the entry point of the application. 
// It parses command-line flags and arguments to sort an array using the specified sorting algorithm.
func main() {
	flag.IntVar(&size, "size", 10, "Size of the array to be sorted")
	flag.IntVar(&max, "max", 100, "Max value of the array to be sorted")
	flag.StringVar(&algo, "algo", QuickSort, "Sorting algorithm to be used")

	flag.Parse()

	var arr []int

	// Check if there are non-flag arguments
	if len(flag.Args()) > 0 {
		for _, v := range flag.Args() {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error: %v is not a valid integer\n", v)
				return
			}
			arr = append(arr, num)
		}
	} else {
		// Use flags to generate a random array
		arr = generateArray(size, max)
	}

	var sorter sort.Sorter

	switch algo {
	case QuickSort:
		sorter = &quicksort.QuickSorter{}
	case MergeSort:
		sorter = &mergesort.MergeSorter{}
	default:
		fmt.Println("Unknown algorithm")
		return
	}

	fmt.Println("Unsorted array:", arr)
	fmt.Println("Sorted array:", sorter.Sort(arr))
}

// generateArray creates a slice of random integers of the given size with values ranging up to max.
func generateArray(size, max int) []int {
	var arr []int = make([]int, size)

	for i := range arr {
		arr[i] = rand.Intn(max)
	}

	return arr
}
