package sort

// CountingSort performs the counting sort algorithm with a fixed
// range of numbers.
func CountingSort(array []int) []int {
	return CountingSortRange(array, minInt(array), maxInt(array))
}

// CountingSortRange performs counting sort on an array containing
// integers in the closed interval [min, max].
func CountingSortRange(array []int, min, max int) []int {
	// Count the number of elements in the array
	counts := make([]int, max-min+1)
	for _, element := range array {
		counts[element-min]++
	}

	i := 0
	for element, count := range counts {
		for count > 0 {
			array[i] = element + min
			i++
			count--
		}
	}

	return array
}

func minInt(array []int) int {
	var min int
	for i, element := range array {
		if i == 0 || element < min {
			min = element
		}
	}
	return min
}

func maxInt(array []int) int {
	var max int
	for i, element := range array {
		if i == 0 || element > max {
			max = element
		}
	}
	return max
}
