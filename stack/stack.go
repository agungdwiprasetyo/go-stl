package stack

import (
	"fmt"
	"sync"
)

// Stack container
type Stack struct {
	size     int
	elements []interface{}
	mutex    sync.RWMutex
}

// New stack constructor
func New() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

// Push value to top stack
func (s *Stack) Push(value interface{}) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.size++
	s.elements = append(s.elements, value)
}

// Pop value in top stack
func (s *Stack) Pop() (interface{}, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.IsEmpty() {
		return nil, false
	}

	l := len(s.elements)
	val := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return val, true
}

// Top get peek value
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.elements[len(s.elements)-1]
}

// IsEmpty check value is empty
func (s *Stack) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.elements) == 0
}

// String print stack
func (s *Stack) String() string {
	if s.IsEmpty() {
		return "Stack is empty"
	}

	l := len(s.elements)
	str := fmt.Sprintf("┌──\t%+v\t──┐\n", s.elements[l-1])
	// fmt.Println([]byte(`┌`))

	for i := l - 2; i >= 0; i-- {
		// fmt.Println("────────────")
		str += fmt.Sprintf("│\t%+v\t│\n", s.elements[i])
	}
	return str
}
