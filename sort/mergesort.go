package sort

func MergeSort(array []int) []int {
	if len(array) < 2 {
		// Base case
		return array
	} else {
		// Sort left and right slices, then merge them
		mid := len(array) / 2
		left := MergeSort(array[:mid])
		right := MergeSort(array[mid:])
		return merge(left, right)
	}
}

// Merge function merges two sorted slices into one
// sorted slice.
func merge(left, right []int) []int {
	merged := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) && k < len(merged) {
		if left[i] <= right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}
	// Put rest of left into array
	for i < len(left) {
		merged[k] = left[i]
		i++
		k++
	}
	// Put rest of right into array
	for j < len(right) {
		merged[k] = right[j]
		j++
		k++
	}
	return merged
}
