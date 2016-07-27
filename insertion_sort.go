package algorithms

// InsertionSort sorts the given array of integers in ascending order using an
// implementation of the Insertion Sort algorithm.
//
// In the best case of a presorted array, Insertion sort runs in linear time
// O(n) while in the worst case of a reverse sorted array, it runs in quadratic
// time O(n2). The average case also runs in quadratic time making it
// inefficient for sorting large arrays.
func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		k := a[i]
		j := i - 1
		for ; j > -1 && a[j] > k; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = k
	}
}
