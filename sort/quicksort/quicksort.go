// Package quicksort provides a concurrent implementation of the quicksort algorithm.
package quicksort

// QuickSorter is an empty struct which provides the Sort method
// to sort slices using the concurrent quicksort algorithm.
type QuickSorter struct{}

// Sort sorts the provided slice of integers using a concurrent version 
// of the quicksort algorithm and returns a new sorted slice. The function 
// leverages goroutines for sorting sub-slices concurrently.
//
// The function works as follows:
// 1. If the length of the slice is 0 or 1, it returns the slice as it's 
//    already sorted.
// 2. It then selects a pivot (the middle element of the slice).
// 3. Three goroutines are then launched:
//    a. The first one sorts all elements less than the pivot.
//    b. The second one collects all elements equal to the pivot.
//    c. The third one sorts all elements greater than the pivot.
// 4. The function then waits for all goroutines to finish their computation 
//    and sends their results through channels.
// 5. Finally, the results are combined and returned as a single sorted slice.
func (qs *QuickSorter) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	n := len(arr)

	pivot := arr[n / 2]

	lCh := make(chan []int)
	eCh := make(chan []int)
	rCh := make(chan []int)

	go func() {
		l := []int{}

		for _, v := range arr {
			if v < pivot {
				l = append(l, v)
			}
		}

		lCh <- qs.Sort(l)
	}()

	go func() {
		e := []int{}

		for _, v := range arr {
			if v == pivot {
				e = append(e, v)
			}
		}

		eCh <- e
	}()

	go func() {
		r := []int{}

		for _, v := range arr {
			if v > pivot {
				r = append(r, v)
			}
		}

		rCh <- qs.Sort(r)
	}()

	l := <-lCh
	e := <-eCh
	r := <-rCh

	return append(append(l, e...), r...)
}
