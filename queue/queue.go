package queue

// Queue serve data as first-in-first-out manner
type Queue struct {
	// the empty channel conveys metadata about the items channel:
	// empty indicates that no goroutine is sending to items
	items chan []interface{}
	// holds true if the queue is empty
	empty chan bool
}

// New return a queue
func New() *Queue {
	items := make(chan []interface{}, 1)
	empty := make(chan bool, 1)
	empty <- true
	return &Queue{items, empty}
}

// Empty check if the queue is empty
func (q *Queue) Empty() bool {
	empty := <-q.empty
	return empty
}

// Size return the number of items in the queue
func (q *Queue) Size() int {
	items := <-q.items
	size := len(items)
	q.items <- items
	return size
}

// Enqueue append an item at the end of the queue
func (q *Queue) Enqueue(v interface{}) {
	var items []interface{}
	select {
	case items = <-q.items:
	case <-q.empty:
	}
	items = append(items, v)
	q.items <- items
}

// Dequeue remove an itme at the begining of the queue
func (q *Queue) Dequeue() interface{} {
	items := <-q.items
	item := items[0]
	items = items[1:]
	if len(items) == 0 {
		q.empty <- true
	} else {
		q.items <- items
	}
	return item
}

// Rare return an item at the end of the queue
func (q *Queue) Rare() interface{} {
	items := <-q.items
	item := items[len(items)-1]
	q.items <- items
	return item
}

// Rare return an item at the begining of the queue
func (q *Queue) Front() interface{} {
	items := <-q.items
	item := items[0]
	q.items <- items
	return item
}

// Clone return a deep copy of the queue
func (q *Queue) Clone() *Queue {
	copy := New()
	empty := <-q.empty
	copy.empty <- empty
	q.empty <- empty
	items := <-q.items
	copy.items <- items
	q.items <- items
	return copy
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
