// Package mergesort provides a concurrent implementation of the merge sort algorithm.
package mergesort

import (
	"sync"
)

// MergeSorter is an empty struct that provides the Sort method
// to sort slices using the concurrent merge sort algorithm.
type MergeSorter struct{}

// Sort sorts the provided slice of integers using a concurrent version 
// of the merge sort algorithm and returns a new sorted slice. The function 
// leverages goroutines for sorting sub-slices concurrently.
//
// The function works as follows:
// 1. If the length of the slice is 0 or 1, it returns the slice as it's 
//    already sorted.
// 2. It then divides the slice into two roughly equal sub-slices.
// 3. Two goroutines are then launched:
//    a. The first one sorts the left sub-slice.
//    b. The second one sorts the right sub-slice.
// 4. The function waits for both goroutines to finish their computation 
//    using a WaitGroup.
// 5. Finally, the results are merged and returned as a single sorted slice.
func (ms *MergeSorter) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	var left, right []int

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		left = ms.Sort(arr[:mid])
	}()

	go func() {
		defer wg.Done()
		right = ms.Sort(arr[mid:])
	}()

	wg.Wait()

	return merge(left, right)
}

// merge takes two sorted slices and returns a single sorted slice by 
// merging them together. This function assumes that both input slices 
// are already sorted.
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}