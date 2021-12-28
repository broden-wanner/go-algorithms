package sort

// InsertionSort splits the array into sorted and unsorted parts and
// repeatedly iterates through the unsorted part and "inserts" it
// into the sorted part.
func InsertionSort(array []int) []int {
	n := len(array)
	for j := 1; j < n; j++ {
		key := array[j]
		// Insert array[j] into the sorted sequence array[0:j-1]
		i := j
		for ; i > 0 && array[i-1] >= key; i-- {
			// Move the greater elements to the left one space
			array[i] = array[i-1]
		}
		array[i] = key
	}
	return array
}
