package stack

import (
	"fmt"
	"sync"
)

// Stack serve data as first-in-last-out manner
type Stack struct {
	items []interface{}
	size  int
	mutex *sync.Mutex
}

// New return a stack
func New() *Stack {
	return &Stack{
		items: []interface{}{},
		size:  0,
		mutex: &sync.Mutex{},
	}
}

// Empty check if the stack is empty
func (s Stack) Empty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.size == 0
}

// Size return the number of items in the stack
func (s Stack) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.size
}

// Pop remove the first element from the stack and return it
func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.size == 0 {
		return nil
	}
	top := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	return top
}

// Top return the first element from the stack
func (s *Stack) Top() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.size == 0 {
		return nil
	}
	return s.items[s.size-1]
}

// Push put an element at the top of the stack
func (s *Stack) Push(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, v)
	s.size++
}

// String return the string representation of the stack
func (s Stack) String() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return fmt.Sprintf("%d %v", s.size, s.items)
}

// Clone return a deep copy of the stack
func (s *Stack) Clone() *Stack {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	copy := New()
	copy.size = s.size
	copy.items = make([]interface{}, copy.size)
	for i, v := range s.items {
		copy.items[i] = v
	}
	return copy
}

// Iterator return an iterator of the stack
func (s *Stack) Iterator() *Iterator {
	return &Iterator{s: s.Clone()}
}

// Iterator enable traveling elements in the stack
type Iterator struct {
	s *Stack
}

// HasNext check if there is another element
func (it Iterator) HasNext() bool {
	return !it.s.Empty()
}

// Next return the next element
func (it Iterator) Next() interface{} {
	if it.s.Empty() {
		return nil
	}
	v := it.s.Pop()
	return v
}
