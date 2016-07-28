package algorithms

// BubbleSort sorts the given array of integers in ascending order using an
// implementation of the Bubble Sort algorithm.
//
// It the runs in subquadtratic time O(n2) in the average and worst case.
//
// The implementation is optimized to exit early if no swaps are performed in a
// given pass resulting in O(n) in the best case of a pre-sorted input.
func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		c := false
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				x := a[j]
				a[j] = a[j-1]
				a[j-1] = x
				c = true
			}
		}

		if !c {
			break
		}
	}
}
