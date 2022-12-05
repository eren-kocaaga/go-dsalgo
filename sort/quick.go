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
func QuickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	left, right := 0, len(items)-1

	pivot := rand.Int() % len(items)

	items[pivot], items[right] = items[right], items[pivot]

	for i, _ := range items {
		if items[i] < items[right] {
			items[left], items[i] = items[i], items[left]
			left++
		}
	}

	items[left], items[right] = items[right], items[left]

	QuickSort(items[:left])
	QuickSort(items[left+1:])

	return items
}
