package algorithms

import (
	"math/rand"
	"testing"
	"time"
)

const (
	BenchmarkSelectionSortInputSize = 4096
)

func TestSelectionSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 16384)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int()
	}

	SelectionSort(a)

	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			t.Fatalf("Selection sort failed with output: %v", a)
		}
	}
}

func BenchmarkSelectionSortRandom(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	v := make([]int, BenchmarkSelectionSortInputSize)
	for i := 0; i < len(v); i++ {
		v[i] = rand.Int()
	}

	for i := 0; i < b.N; i++ {
		a := make([]int, len(v))
		copy(a, v)

		SelectionSort(a)
	}
}

func BenchmarkSelectionSortBestCase(b *testing.B) {
	a := make([]int, BenchmarkSelectionSortInputSize)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	for i := 0; i < b.N; i++ {
		SelectionSort(a)
	}
}

func BenchmarkSelectionSortWorstCase(b *testing.B) {
	v := make([]int, BenchmarkSelectionSortInputSize)
	for i := 0; i < len(v); i++ {
		v[i] = len(v) - i
	}

	for i := 0; i < b.N; i++ {
		a := make([]int, len(v))
		copy(a, v)
		SelectionSort(a)
	}
}
