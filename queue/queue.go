package queue

import (
	"fmt"
	"sync"
)

// Queue serve data as first-in-first-out manner
type Queue struct {
	items []interface{}
	size  int
	mutex *sync.Mutex
}

// New return a queue
func New() *Queue {
	var q Queue
	q.items = []interface{}{}
	q.size = 0
	q.mutex = &sync.Mutex{}
	return &q
}

// Empty check if the queue is empty
func (q *Queue) Empty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.size == 0
}

// Size return the number of items in the queue
func (q *Queue) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.size
}

// Enqueue append an item at the end of the queue
func (q *Queue) Enqueue(v interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, v)
	q.size++
}

// Dequeue remove an itme at the begining of the queue
func (q *Queue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item
}

// Rare return an item at the end of the queue
func (q *Queue) Rare() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == 0 {
		return nil
	}
	return q.items[q.size-1]
}

// Rare return an item at the begining of the queue
func (q *Queue) Front() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.size == 0 {
		return nil
	}
	return q.items[0]
}

// Clone return a deep copy of the queue
func (q *Queue) Clone() *Queue {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	copy := New()
	copy.size = q.size
	copy.items = append(copy.items, q.items...)
	return copy
}

// String return string representation of the queue
func (q *Queue) String() string {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	str := fmt.Sprintf("%d %v", q.size, q.items)
	return str
}

// Iterator return an iterator of the queue
func (q *Queue) Iterator() *Iterator {
	return &Iterator{q: q.Clone()}
}

// Iterator enable traveling elements in the queue
type Iterator struct {
	q *Queue
}

// HasNext check if there is another element
func (it Iterator) HasNext() bool {
	return !it.q.Empty()
}

// Next return the next element
func (it Iterator) Next() interface{} {
	if it.q.Empty() {
		return nil
	}
	v := it.q.Dequeue()
	return v
}
