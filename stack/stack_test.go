package stack_test

import (
	"ds/misc"
	"ds/stack"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	s := stack.New()
	misc.Equals(t, 0, s.Size())
	misc.Equals(t, true, s.Empty())
}

func TestString(t *testing.T) {
	s := stack.New()
	misc.Equals(t, "0 []", s.String())
	s.Push(1)
	misc.Equals(t, "1 [1]", s.String())
	s.Push(2)
	misc.Equals(t, "2 [1 2]", s.String())
}

func TestBasic(t *testing.T) {
	s := stack.New()
	for i := 1; i <= 9; i++ {
		s.Push(i)
		misc.Equals(t, i, s.Size())
		misc.Equals(t, i, s.Top())
	}
	misc.Equals(t, false, s.Empty())
	for i := 1; i <= 9; i++ {
		misc.Equals(t, 9-i+1, s.Pop())
		misc.Equals(t, 9-i, s.Size())
	}
	misc.Equals(t, true, s.Empty())
	misc.Equals(t, nil, s.Pop())
}

func TestIterator(t *testing.T) {
	s := stack.New()
	for i := 1; i <= 9; i++ {
		s.Push(i)
	}
	it := s.Iterator()
	i := 9
	for it.HasNext() {
		misc.Equals(t, i, it.Next())
		i--
	}
}

// TODO: Neet to figure out how to test the stack in cocurrnet
func TestCocurrent(t *testing.T) {
	s := stack.New()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10000; i++ {
			s.Push(i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			misc.Equals(t, 10000-i-1, s.Pop())
		}
		wg.Done()
	}()
	wg.Wait()
}
