package lib

import "testing"

func Test_BasicStack(t *testing.T) {
	stack := NewStack[string]()

	stack.Push("Flavio")
	stack.Push("Jose")

	if op := stack.Peek(); op.GetValue() != "Jose" {
		t.Error("The top of the stack should be Jose")
	}

	if op := stack.Pop(); op.GetValue() != "Jose" {
		t.Error("The top of the stack should be Jose")
	}

	if op := stack.Peek(); op.GetValue() != "Flavio" {
		t.Error("The top of the stack should be Flavio")
	}

	if op := stack.Pop(); op.GetValue() != "Flavio" {
		t.Error("The top of the stack should be Flavio")
	}

	if op := stack.Peek(); op.HasValue() {
		t.Error("The stack should be empty!")
	}

	if op := stack.Pop(); op.HasValue() {
		t.Error("The stack should be empty!")
	}
}
