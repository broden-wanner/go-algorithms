package search

// BinarySearchIterativeIntegers performs binary search on a slice of integers
// using an iterative approach.
func BinarySearchIterativeIntegers(array []int, target int) (int, error) {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target < array[mid] {
			right = mid - 1
		} else if target > array[mid] {
			left = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrorNotFound
}

// BinarySearchRecursiveIntegers performs binary search on a slice of integers
// using a recursive approach.
func BinarySearchRecursiveIntegers(array []int, target int) (int, error) {
	var binaryHelper func(left, right int) (int, error)
	binaryHelper = func(left, right int) (int, error) {
		if right < left || len(array) == 0 {
			return -1, ErrorNotFound
		}
		mid := left + (right-left)/2
		if target < array[mid] {
			return binaryHelper(left, mid-1)
		} else if target > array[mid] {
			return binaryHelper(mid+1, right)
		} else {
			return mid, nil
		}
	}
	return binaryHelper(0, len(array)-1)
}
