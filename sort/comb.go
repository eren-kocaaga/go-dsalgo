package sort

/*
CombSort is a variant of the Bubble Sort. Bubble sort always
compares adjacent values. So all inversions are removed one by one.
Comb Sort improves on Bubble Sort by using gap of size more than 1.
The gap starts with a large value and shrinks by a factor of 1.3
in every iteration until it reaches the value 1. Thus Comb Sort
removes more than one inversion counts with one swap and performs
better than Bubble Sort.

Although it works better than Bubble Sort on average, worst-case remains
O(n^2)
*/
func CombSort(items []int) {
	var (
		n       = len(items)
		gap     = len(items)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				swapped = true
			}
		}
	}
}
