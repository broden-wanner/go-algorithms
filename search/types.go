package search

// Comparer is an interface that wraps the Compare method.
type Equaler interface {
	EqualTo(interface{}) bool
}
