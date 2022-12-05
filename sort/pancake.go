package sort

/*
PancakeSort is the colloquial term for the mathematical problem of sorting
a disordered stack of pancakes in order of size when a spatula can be inserted
at any point in the stack and used to flip all pancakes above it. A pancake
number is the maximum number of flips required for a given number of pancakes.
Here is source code of the Go Program to implement Pancake Sorting Algorithm.

O(n²)
*/
func PancakeSort(items []int) []int {
	for uns := len(items) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, items[0]
		for i := 1; i <= uns; i++ {
			if items[i] > lg {
				lx, lg = i, items[i]
			}
		}
		// move to final position in two flips
		items = flip(items, lx)
		items = flip(items, uns)
	}

	return items
}

func flip(items []int, r int) []int {
	for l := 0; l < r; l, r = l+1, r-1 {
		items[l], items[r] = items[r], items[l]
	}
	return items
}
