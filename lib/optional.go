package lib

// Represents a value that may or may not exists when it's time to use it!
//
// By default the Optional instance is empty, you can only create valid Optional instances using the NewOpValue constructor!
type Optional[T any] struct {
	isValid bool
	value   T
}

func NewOpValue[T any](value T) Optional[T] {
	return Optional[T]{
		isValid: true,
		value:   value,
	}
}

func NewOpEmpty[T any]() Optional[T] {
	return Optional[T]{}
}

func (option *Optional[T]) HasValue() bool {
	return option.isValid
}

func (option *Optional[T]) GetValue() T {
	if !option.isValid {
		panic("Invalid access to optional!")
	}

	return option.value
}
