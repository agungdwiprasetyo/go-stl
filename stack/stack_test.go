package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Test Stack Data Structure", func(t *testing.T) {
		s := New()
		if s.elements == nil {
			t.Errorf("element is nil, should not nil")
		}

		s.Push(10)
		s.Push(20)
		s.Push(30)

		fmt.Println(s)

		peek := s.Peek()
		if peek.(int) != 30 {
			t.Errorf("invalid peek value, should 30")
		}

		val, ok := s.Pop()
		if !ok {
			t.Errorf("error when pop stack")
		}

		if val.(int) != 30 {
			t.Errorf("invalid pop value, should 30")
		}
		peek = s.Peek()
		if peek.(int) != 20 {
			t.Errorf("invalid peek value, should 20")
		}

		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()

		if !s.IsEmpty() {
			t.Errorf("stack should be empty")
		}

		fmt.Println(s)
		if peek = s.Peek(); peek != nil {
			t.Errorf("in empty stack, peek should be nil")
		}
	})
}
