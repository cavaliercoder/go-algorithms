package algorithms

import (
	"math/rand"
	"testing"
	"time"
)

const (
	BenchmarkInsertionSortInputSize = 4096
)

func TestInsertionSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 16384)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int()
	}

	InsertionSort(a)

	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			t.Fatalf("Insertion sort failed with output: %v", a)
		}
	}
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	v := make([]int, BenchmarkInsertionSortInputSize)
	for i := 0; i < len(v); i++ {
		v[i] = rand.Int()
	}

	for i := 0; i < b.N; i++ {
		a := make([]int, len(v))
		copy(a, v)

		InsertionSort(a)
	}
}

func BenchmarkInsertionSortBestCase(b *testing.B) {
	a := make([]int, BenchmarkInsertionSortInputSize)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	for i := 0; i < b.N; i++ {
		InsertionSort(a)
	}
}

func BenchmarkInsertionSortWorstCase(b *testing.B) {
	v := make([]int, BenchmarkInsertionSortInputSize)
	for i := 0; i < len(v); i++ {
		v[i] = len(v) - i
	}

	for i := 0; i < b.N; i++ {
		a := make([]int, len(v))
		copy(a, v)
		InsertionSort(a)
	}
}
