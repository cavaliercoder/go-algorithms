package algorithms

// BubbleSort sorts the given array of integers in ascending order using an
// implementation of the Bubble Sort algorithm.
//
// It the runs in subquadtratic time O(n2) for any input.
func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				x := a[j]
				a[j] = a[j-1]
				a[j-1] = x
			}
		}
	}
}
