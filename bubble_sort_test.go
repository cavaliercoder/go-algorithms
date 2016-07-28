package algorithms

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testSort(t, BubbleSort, 16384)
}

func BenchmarkBubbleSortRandom(b *testing.B) {
	benchmakeSortRandom(b, BubbleSort, 4096)
}

func BenchmarkBubbleSortBestCase(b *testing.B) {
	benchmarkSortSorted(b, BubbleSort, 4096)
}

func BenchmarkBubbleSortWorstCase(b *testing.B) {
	benchmarkSortReverseSorted(b, BubbleSort, 4096)
}
