package sort

/*
ShellSort is mainly a variation of Insertion Sort. The idea of
shellSort is to allow exchange of far items. In shellSort, we
make the array N-sorted for a large value of N. We keep reducing the
value of N until it becomes 1. An array is said to be N-sorted
if all sub-lists of every N'th element is sorted.

O(n^2)
*/
func ShellSort(items []int) []int {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}

	return items
}

func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}
