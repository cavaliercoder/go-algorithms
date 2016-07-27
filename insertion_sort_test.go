package algorithms

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testSort(t, InsertionSort, 16384)
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	benchmakeSortRandom(b, InsertionSort, 4096)
}

func BenchmarkInsertionSortBestCase(b *testing.B) {
	benchmarkSortSorted(b, InsertionSort, 4096)
}

func BenchmarkInsertionSortWorstCase(b *testing.B) {
	benchmarkSortReverseSorted(b, InsertionSort, 4096)
}
