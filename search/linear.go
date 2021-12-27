package search

// LinearSearchIntegers performs linear search on a slice of integers.
// Returns the index of the target if found and -1 otherwise.
func LinearSearchIntegers(collection []int, target int) (int, error) {
	for i := 0; i < len(collection); i++ {
		if collection[i] == target {
			return i, nil
		}
	}
	return -1, ErrorNotFound
}

// LinearSearchGeneral performs linear search on a slice of any comparable
// types. Returns the index of the target if found and -1 otherwise.
func LinearSearchGeneral(collection []Equaler, target Equaler) (int, error) {
	for i := 0; i < len(collection); i++ {
		if collection[i].EqualTo(target) {
			return i, nil
		}
	}
	return -1, ErrorNotFound
}
