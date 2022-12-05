package search

/*
BinarySearch is a search strategy used to find elements within a list by
consistently reducing the amount of data to be searched and thereby increasing
the rate at which the search term is found. To use a binary search algorithm,
the list to be operated on must have already been sorted.

O(log n)
*/
func BinarySearch(key int, stack []int) bool {
	low, high := 0, len(stack)-1

	for low <= high {
		median := (low + high) / 2

		if stack[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(stack) || stack[low] != key {
		return false
	}

	return true
}
