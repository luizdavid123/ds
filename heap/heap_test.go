package heap_test

import (
	"ds/heap"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Run Package Test: go test -v ./heap  --timeout=30s
// Run File Test: go test -v ./heap -run "^(TestNew|TestString)$" -timeout 30s
// Run Function Test: go test -v ./heap -run ^TestNew$ -timeout 30s

func TestNew(t *testing.T) {
	h := heap.New()
	equals(t, 0, h.Size())
	equals(t, 16, h.Capacity())
	equals(t, true, h.Empty())
}

func TestString(t *testing.T) {
	h := heap.New()
	equals(t, "0 16 []", h.String())
}

func TestHeapBasic(t *testing.T) {
	h := heap.New()
	for i := 8; i > 0; i-- {
		h.Push(i)
	}
	i := 1
	for i != 9 {
		equals(t, i, h.Pop())
		i++
	}
	equals(t, nil, h.Pop())
}

func TestIterator(t *testing.T) {
	h := heap.New()
	for i := 8; i > 0; i-- {
		h.Push(i)
	}
	it := h.Iterator()
	i := 1
	for it.HasNext() {
		equals(t, i, it.Next())
		i++
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\nexp: %#v\ngot: %#v\033[39m\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
