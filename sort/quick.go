package sort

import "math/rand"

/*
QuickSort algorithm falls under the divide and conquer class
of algorithms, where we break (divide) a problem into smaller chunks
that are much simpler to solve (conquer). In this case, an unsorted
array is broken into sub-arrays that are partially sorted, until all
elements in the list are in the right position, by which time our
unsorted list will have become sorted.

O(log n)
*/
func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1

	pivot := rand.Int() % len(array)

	array[pivot], array[right] = array[right], array[pivot]

	for i, _ := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}
