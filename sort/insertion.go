package sort

// InsertionSort The idea of swapping adjacent elements to sort a list of items
// can also be used to implement the insertion sort. In the insertion sort
// algorithm, we assume that a certain portion of the list has already been
// sorted, while the other portion remains unsorted. With this assumption,
// we move through the unsorted portion of the list, picking one element at
// a time. With this element, we go through the sorted portion of the list
// and insert it in the right order so that the sorted portion of the list
// remains sorted./*
//
// O(n²)
func InsertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
