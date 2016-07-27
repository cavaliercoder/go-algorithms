package algorithms

import "math"

// MergeSort sorts the given array of integers in ascending order using an
// implementation of the top-down Merge Sort algorithm.
//
// Merge sort runs in O(n log n) in the worst and average case.
func MergeSort(a []int) {
	divide(a, 0, len(a)-1)
}

func divide(a []int, p, r int) {
	if p < r {
		q := int((p + r) / 2)
		divide(a, p, q)
		divide(a, q+1, r)
		conquer(a, p, q, r)
	}
}

func conquer(a []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = a[p+i]
	}
	L[len(L)-1] = math.MaxInt64

	for i := 0; i < n2; i++ {
		R[i] = a[q+i+1]
	}
	R[len(R)-1] = math.MaxInt64

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
	}
}
