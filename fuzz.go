// +build gofuzz
package algorithms

import (
	"fmt"
)

func Fuzz(data []byte) int {
	// insertion sort
	a := getIntSlice(data)
	InsertionSort(a)
	panicOnUnsortedArray(a)

	// selection sort
	a = getIntSlice(data)
	SelectionSort(a)
	panicOnUnsortedArray(a)

	return 1
}

func getIntSlice(data []byte) []int {
	a := make([]int, len(data)/4)
	for i := 0; i < len(a); i++ {
		a[i] = int(data[i*4]) << 24
		a[i] += int(data[(i*4)+1]) << 16
		a[i] += int(data[(i*4)+2]) << 8
		a[i] += int(data[(i*4)+3])
	}

	return a
}

func panicOnUnsortedArray(a []int) {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			panic(fmt.Errorf("Value at offset %v is less than the value at offset %v", i, i-1))
		}
	}
}
