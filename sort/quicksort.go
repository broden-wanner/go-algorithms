package sort

import (
	"math/rand"
)

// QuickSort performs the quick sort algorithm on the given
// slice using the recursive algorithm.
func QuickSort(array []int) []int {
	quickSortHelper(array, 0, len(array)-1)
	return array
}

// RandomizedQuickSort performs the quick sort algorithm on the given
// slice but uses a random partition to improve runtime performance.
func RandomizedQuickSort(array []int) []int {
	randomizedQuickSortHelper(array, 0, len(array)-1)
	return array
}

// Performs the work of quicksort by first partitioning
// the array[p:r] and recursively calling the helper function
// on the partitions.
func quickSortHelper(array []int, p, r int) {
	if len(array) < 2 {
		return
	} else if p < r {
		q := partition(array, p, r)
		quickSortHelper(array, p, q-1)
		quickSortHelper(array, q+1, r)
		return
	}
}

// Performs the work of quicksort by first randomly partitioning
// the array[p:r] and recursively calling the helper function
// on the partitions.
func randomizedQuickSortHelper(array []int, p, r int) {
	if len(array) < 2 {
		return
	} else if p < r {
		q := randomizedPartition(array, p, r)
		randomizedQuickSortHelper(array, p, q-1)
		randomizedQuickSortHelper(array, q+1, r)
		return
	}
}

// Partitions the array using partition but first
// performs a random swap between the elements in
// indexes p through r
func randomizedPartition(array []int, p, r int) int {
	i := rand.Intn(r-p+1) + p
	array[r], array[i] = array[i], array[r]
	return partition(array, p, r)
}

// Partitions the array[p:r] into two subarrays array[p:q-1]
// and array[q+1:r] such that each element of array[p:q-1] is
// less than or equal to array[q], which is then less than or equal
// to array[q+1:r]
func partition(array []int, p, r int) int {
	pivot := array[r]
	i := p - 1
	for j := p; j < r; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[r] = array[r], array[i+1]
	return i + 1
}
