package main

import (
	"gosort/sort/mergesort"
	"gosort/sort/quicksort"
	"testing"
)

func generateSortedArray(size int) []int {
	var arr []int = make([]int, size)

	for i := range arr {
		arr[i] = i
	}

	return arr
}

func benchmarkMergeSorter(arr []int, b *testing.B) {
	ms := &mergesort.MergeSorter{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ms.Sort(arr)
	}
}

func benchmarkQuickSorter(arr []int, b *testing.B) {
	qs := &quicksort.QuickSorter{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = qs.Sort(arr)
	}
}

func BenchmarkMergeSorterRandom10(b *testing.B) {
	benchmarkMergeSorter(generateArray(10, 10), b)
}

func BenchmarkQuickSorterRandom10(b *testing.B) {
	benchmarkQuickSorter(generateArray(10, 10), b)
}

func BenchmarkMergeSorterRandom100(b *testing.B) {
	benchmarkMergeSorter(generateArray(100, 100), b)
}

func BenchmarkQuickSorterRandom100(b *testing.B) {
	benchmarkQuickSorter(generateArray(100, 100), b)
}

func BenchmarkMergeSorterRandom1000(b *testing.B) {
	benchmarkMergeSorter(generateArray(1000, 1000), b)
}

func BenchmarkQuickSorterRandom1000(b *testing.B) {
	benchmarkQuickSorter(generateArray(1000, 1000), b)
}

func BenchmarkMergeSorterRandom10000(b *testing.B) {
	benchmarkMergeSorter(generateArray(10000, 10000), b)
}

func BenchmarkQuickSorterRandom10000(b *testing.B) {
	benchmarkQuickSorter(generateArray(10000, 10000), b)
}

func BenchmarkMergeSorterSorted100(b *testing.B) {
	benchmarkMergeSorter(generateSortedArray(100), b)
}

func BenchmarkQuickSorterSorted100(b *testing.B) {
	benchmarkQuickSorter(generateSortedArray(100), b)
}

func BenchmarkMergeSorterSorted1000(b *testing.B) {
	benchmarkMergeSorter(generateSortedArray(1000), b)
}

func BenchmarkQuickSorterSorted1000(b *testing.B) {
	benchmarkQuickSorter(generateSortedArray(1000), b)
}
