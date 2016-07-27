package algorithms

// SelectionSort sorts the given array of integers in ascending order using an
// implementation of the Selection Sort algorithm.
//
// Selection sort run consistently in O(n2) time for the best, worst and average
// case.
func SelectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		pK := i
		k := a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] < k {
				k = a[j]
				pK = j
			}
		}

		a[pK] = a[i]
		a[i] = k
	}
}
