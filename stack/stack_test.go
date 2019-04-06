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

		top := s.Top()
		if top.(int) != 30 {
			t.Errorf("invalid top value, should 30")
		}

		val, ok := s.Pop()
		if !ok {
			t.Errorf("error when pop stack")
		}

		if val.(int) != 30 {
			t.Errorf("invalid pop value, should 30")
		}
		top = s.Top()
		if top.(int) != 20 {
			t.Errorf("invalid top value, should 20")
		}

		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()

		if !s.IsEmpty() {
			t.Errorf("stack should be empty")
		}

		fmt.Println(s)
		if top = s.Top(); top != nil {
			t.Errorf("in empty stack, top should be nil")
		}
	})
}
