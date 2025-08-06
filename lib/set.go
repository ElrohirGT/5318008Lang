package lib

// Represents a bag of unique values.
// Iteration order is randomized.
type Set[T comparable] map[T]struct{}

// Creates a new Set for the given type.
func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Adds the given v to the set.
// If v is a new value to the Set True is returned, otherwise False is returned.
func (s *Set[T]) Add(v T) bool {
	_, alreadyAdded := (*s)[v]

	if !alreadyAdded {
		(*s)[v] = struct{}{}
	}

	return !alreadyAdded
}

// Checks if the given v exists.
func (s *Set[T]) Exists(v T) bool {
	_, exists := (*s)[v]
	return exists
}
