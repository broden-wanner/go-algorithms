package sort

// SelectionSort "selects" the minimum element from the
// unsorted part of the array and places it at the end of
// the sorted part
func SelectionSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		// Iterate through the slice from i to the end and
		// select the minimum element, setting its index to
		// minIdx
		minIdx := i
		for j := i; j < n; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		// Swap the elements at indices i and minIdx to place
		// the minimum element at the ith index
		array[i], array[minIdx] = array[minIdx], array[i]
	}
	return array
}
