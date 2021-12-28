package sort

// BubbleSort repeatedly steps through the list, compares adjacent elements, and
// swaps them if they are in the wrong order.
func BubbleSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		// Only go until the (n-i)th element since the last i elements
		// are sorted.
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
