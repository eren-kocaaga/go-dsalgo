package sort

/*
SelectionSort is a sorting algorithm begins by finding the smallest element in an array
and interchanging it with data at, for instance, array index [0]. Starting
at index 0, we search for the smallest item in the list that exists between
index 1 and the index of the last element. When this element has been found,
it is exchanged with the data found at index 0. We simply repeat this process
until the list becomes sorted.

O(n^2)
*/
func SelectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
