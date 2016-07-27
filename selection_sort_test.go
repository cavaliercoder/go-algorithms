package algorithms

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testSort(t, SelectionSort, 16384)
}

func BenchmarkSelectionSortRandom(b *testing.B) {
	benchmakeSortRandom(b, SelectionSort, 4096)
}

func BenchmarkSelectionSortBestCase(b *testing.B) {
	benchmarkSortSorted(b, SelectionSort, 4096)
}

func BenchmarkSelectionSortWorstCase(b *testing.B) {
	benchmarkSortReverseSorted(b, SelectionSort, 4096)
}
