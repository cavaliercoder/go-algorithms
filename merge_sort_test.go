package algorithms

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testSort(t, MergeSort, 16384)
}

func BenchmarkMergeSortRandom(b *testing.B) {
	benchmakeSortRandom(b, MergeSort, 4096)
}

func BenchmarkMergeSortBestCase(b *testing.B) {
	benchmarkSortSorted(b, MergeSort, 4096)
}

func BenchmarkMergeSortWorstCase(b *testing.B) {
	benchmarkSortReverseSorted(b, MergeSort, 4096)
}
