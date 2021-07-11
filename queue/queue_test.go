package queue_test

import (
	"ds/misc"
	"ds/queue"
	"testing"
)

func TestString(t *testing.T) {
	q := queue.New()
	q.Enqueue(1)
	s := q.String()
	misc.Equals(t, "1 [1]", s)
}
