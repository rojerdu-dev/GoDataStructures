package Stacks

// Stack -> LIFO ordered list
// Add & Get elements at the top of the stack

// Methods:
// Push
// Pop
// Peek
// Size
// IsEmpty

type Stack struct {
	elements []any
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
