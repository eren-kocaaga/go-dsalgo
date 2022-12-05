package sort

/*
BubbleSort Given an unordered list, we compare adjacent elements in the list, each time,
putting in the right order of magnitude, only two elements. The algorithm
hinges on a swap procedure. Knowing how many times to swap is important when
implementing a bubble sort algorithm. To sort a list of numbers such as
[3, 2, 1], we need to swap the elements a maximum of twice. This is equal
to the length of the list minus 1

O(n^2)
*/
func BubbleSort(items []int) []int {
	size := len(items)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < size-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		size -= 1
	}

	return items
}
