package heap_test

import (
	"ds/heap"
	"ds/misc"
	"testing"
)

// Run Package Test: go test -v ./heap  --timeout=30s
// Run File Test: go test -v ./heap -run "^(TestNew|TestString)$" -timeout 30s
// Run Function Test: go test -v ./heap -run ^TestNew$ -timeout 30s

func TestNew(t *testing.T) {
	h := heap.New()
	misc.Equals(t, 0, h.Size())
	misc.Equals(t, 16, h.Capacity())
	misc.Equals(t, true, h.Empty())
}

func TestString(t *testing.T) {
	h := heap.New()
	misc.Equals(t, "0 16 []", h.String())
}

func TestHeapBasic(t *testing.T) {
	h := heap.New()
	for i := 8; i > 0; i-- {
		h.Push(i)
	}
	i := 1
	for i != 9 {
		misc.Equals(t, i, h.Pop())
		i++
	}
	misc.Equals(t, nil, h.Pop())
}

func TestIterator(t *testing.T) {
	h := heap.New()
	for i := 8; i > 0; i-- {
		h.Push(i)
	}
	it := h.Iterator()
	i := 1
	for it.HasNext() {
		misc.Equals(t, i, it.Next())
		i++
	}
}
