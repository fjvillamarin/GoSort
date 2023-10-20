package quicksort

import (
	"fmt"
	"reflect"
	"testing"
	"testing/quick"
)

func ExampleQuickSorter_Sort() {
	ms := &QuickSorter{}
	arr := []int{5, 3, 9, 8, 1, 4, 7, 6, 2}
	fmt.Println(ms.Sort(arr))
	// Output: [1 2 3 4 5 6 7 8 9]
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Sort empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "Sort single element array",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "Sort array without duplicates",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Sort array with duplicates",
			arr:  []int{9, 9, 9, 8, 7, 6, 6, 5, 4, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 9, 9, 9},
		},
		{
			name: "Sort array with negative numbers",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3},
			want: []int{-3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Sort already sorted array",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Sort array with odd number of elements",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Sort array with even number of elements",
			arr:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	qs := &QuickSorter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qs.Sort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSorter.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortFuzzy(t *testing.T) {
	qs := &QuickSorter{}

	f := func(arr []int) bool {
		sorted := qs.Sort(arr)

		for i := 0; i < len(sorted)-1; i++ {
			if sorted[i] > sorted[i+1] {
				t.Errorf("Fuzzing found unsorted array: %v", sorted)
				return false
			}
		}

		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("Fuzzing failed:", err)
	}
}