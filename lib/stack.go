package lib

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return []T{}
}

// Adds a value to the top of the stack
func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

// Checks if the given stack is empty
func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) <= 0
}

// Returns the value at the top of the stack (if there is one).
func (stack *Stack[T]) Peek() Optional[T] {
	if stack.IsEmpty() {
		return NewOpEmpty[T]()
	}

	val := (*stack)[len(*stack)-1]
	return NewOpValue(val)
}

// Returns the value at the top of the stack (if there is one).
// It also removes the value from the stack.
func (stack *Stack[T]) Pop() Optional[T] {
	if stack.IsEmpty() {
		return NewOpEmpty[T]()
	}

	val := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return NewOpValue(val)
}
