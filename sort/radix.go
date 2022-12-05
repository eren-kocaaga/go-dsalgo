package sort

import (
	"bytes"
	"encoding/binary"
)

const (
	digit  = 4
	maxbit = -1 << 31
)

/*
RadixSort is a clever and intuitive little sorting algorithm. Radix Sort
puts the elements in order by comparing the digits of the numbers. Radix
sort is an integer sorting algorithm that sorts data with integer keys
by grouping the keys by individual digits that share the same significant
position and value (place value). Radix sort uses counting sort as a
subroutine to sort an array of numbers.

n: is the number of keys
w: is the key length
O(nw)
*/
func RadixSort(items []int) []int {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(items))
	for i, e := range items {
		binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		items[i] = int(w ^ maxbit)
	}

	return items
}
