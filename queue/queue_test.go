package queue_test

import (
	"ds/misc"
	"ds/queue"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	q := queue.New()
	misc.Equals(t, 0, q.Size())
	misc.Equals(t, true, q.Empty())
}

func TestString(t *testing.T) {
	q := queue.New()
	misc.Equals(t, "0 []", q.String())
	q.Enqueue(1)
	misc.Equals(t, "1 [1]", q.String())
	q.Enqueue(2)
	misc.Equals(t, "2 [1 2]", q.String())
}

func TestBasic(t *testing.T) {
	q := queue.New()
	for i := 1; i <= 9; i++ {
		q.Enqueue(i)
		misc.Equals(t, i, q.Size())
	}
	misc.Equals(t, false, q.Empty())
	for i := 1; i <= 9; i++ {
		misc.Equals(t, i, q.Dequeue())
		misc.Equals(t, 9-i, q.Size())
	}
	misc.Equals(t, true, q.Empty())
	misc.Equals(t, nil, q.Dequeue())
}

func TestIterator(t *testing.T) {
	q := queue.New()
	for i := 1; i <= 9; i++ {
		q.Enqueue(i)
	}
	it := q.Iterator()
	i := 1
	for it.HasNext() {
		misc.Equals(t, i, it.Next())
		i++
	}
}

func TestCocurrent(t *testing.T) {
	q := queue.New()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10000; i++ {
			q.Enqueue(i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			misc.Equals(t, i, q.Dequeue())
		}
		wg.Done()
	}()
	wg.Wait()
}
