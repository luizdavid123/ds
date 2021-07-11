package heap

import (
	"fmt"
	"strings"
)

// CmpFunc is a func that compare two values and return
// 1 if left > right
// 0 if left == right
// -1 if left < right
type CmpFunc func(left interface{}, right interface{}) int

// GetCmpFunc return comparator for specific type
func GetCmpFunc(v interface{}) CmpFunc {
	switch v.(type) {
	case int:
		return CmpInt
	case string:
		return CmpStr
	default:
		return nil
	}
}

// CmpInt is the comparator for integer
func CmpInt(left interface{}, right interface{}) int {
	return left.(int) - right.(int)
}

// CmpStr is the comparator for string
func CmpStr(left interface{}, right interface{}) int {
	return strings.Compare(left.(string), right.(string))
}

// Heap is the data structure that ensures that
// a element being popped is maximun
type Heap struct {
	size     int
	capacity int
	elements []interface{}
}

// New return a heap
func New() *Heap {
	return &Heap{
		size:     0,
		capacity: 16,
		elements: make([]interface{}, 0, 16),
	}
}

// Peek return the root element
func (h *Heap) Peek() interface{} {
	if h.size == 0 {
		return nil
	}
	return h.elements[0]
}

// Pop return and remove the root element
func (h *Heap) Pop() interface{} {
	if h.size == 0 {
		return nil
	}
	v := h.elements[0]
	l := h.elements[h.size-1]
	h.elements[0] = l
	h.elements[h.size-1] = nil
	h.elements = h.elements[:h.size-1]
	h.size--
	h.heapifyDown(0)
	return v
}

// Push add a element into heap
func (h *Heap) Push(v interface{}) {
	if h.size == h.capacity {
		h.resize()
	}
	h.elements = append(h.elements, v)
	h.size++
	h.heapifyUp(h.size - 1)
}

// Size return the number of elements in heap
func (h Heap) Size() int {
	return h.size
}

// Empty check if a heap has no elements
func (h Heap) Empty() bool {
	return h.size == 0
}

// Capacity return the capacity of a heap
func (h Heap) Capacity() int {
	return h.capacity
}

// String return string representation of a heap
func (h Heap) String() string {
	return fmt.Sprintf("%d %d %v", h.size, h.capacity, h.elements)
}

func (h *Heap) heapifyUp(index int) {
	if !h.hasParent(index) {
		return
	}
	parentIndex := h.getParentIndex(index)
	parent := h.elements[parentIndex]
	curr := h.elements[index]
	cmp := GetCmpFunc(curr)
	if cmp(curr, parent) > 0 {
		return
	}
	h.swap(index, parentIndex)
	h.heapifyUp(parentIndex)
}

func (h *Heap) heapifyDown(index int) {
	smallerIndex := h.getSmallerChildIndex(index)
	if smallerIndex == -1 {
		return
	}
	curr := h.elements[index]
	smaller := h.elements[smallerIndex]
	cmp := GetCmpFunc(curr)
	if cmp(curr, smaller) < 0 {
		return
	}
	h.swap(index, smallerIndex)
	h.heapifyDown(smallerIndex)
}

func (h *Heap) resize() {
	h.capacity = h.capacity * 3 / 2
	tmp := make([]interface{}, h.size, h.capacity)
	copy(tmp, h.elements)
	h.elements = tmp
}

func (h *Heap) swap(left int, right int) {
	tmp := h.elements[left]
	h.elements[left] = h.elements[right]
	h.elements[right] = tmp
}

func (h Heap) getParentIndex(index int) int {
	return (index - 1) / 2
}
func (h Heap) getLeftChildIndex(index int) int {
	return 2*index + 1
}
func (h Heap) getRightChildIndex(index int) int {
	return 2*index + 2
}

func (h Heap) hasParent(index int) bool {
	if index == 0 {
		return false
	}
	return h.getParentIndex(index) >= 0
}
func (h Heap) hasChild(index int) bool {
	return h.hasLeftChild(index) || h.hasRightChild(index)
}
func (h Heap) hasLeftChild(index int) bool {
	if index == h.size-1 {
		return false
	}
	return h.getLeftChildIndex(index) < h.size
}
func (h Heap) hasRightChild(index int) bool {
	if index == h.size-1 {
		return false
	}
	return h.getRightChildIndex(index) < h.size
}

func (h Heap) parent(index int) interface{} {
	if index == 0 {
		return nil
	}
	return h.elements[h.getParentIndex(index)]
}
func (h Heap) leftChild(index int) interface{} {
	if index == h.size-1 {
		return nil
	}
	return h.elements[h.getLeftChildIndex(index)]
}
func (h Heap) rightChild(index int) interface{} {
	if index == h.size-1 {
		return nil
	}
	return h.elements[h.getRightChildIndex(index)]
}

func (h Heap) getSmallerChildIndex(index int) int {
	if !h.hasChild(index) {
		return -1
	}
	var left, right interface{}
	if h.hasLeftChild(index) {
		left = h.leftChild(index)
		if !h.hasRightChild(index) {
			return h.getLeftChildIndex(index)
		}
	}
	if h.hasRightChild(index) {
		right = h.rightChild(index)
		if !h.hasLeftChild(index) {
			return h.getRightChildIndex(index)
		}
	}
	cmp := GetCmpFunc(left)
	if cmp(left, right) >= 0 {
		return h.getRightChildIndex(index)
	}
	return h.getLeftChildIndex(index)
}

func (h Heap) smallerChild(index int) interface{} {
	if !h.hasChild(index) {
		return nil
	}
	var left, right interface{}
	if h.hasLeftChild(index) {
		left = h.leftChild(index)
		if !h.hasRightChild(index) {
			return left
		}
	}
	if h.hasRightChild(index) {
		right = h.rightChild(index)
		if !h.hasLeftChild(index) {
			return right
		}
	}
	cmp := GetCmpFunc(left)
	if cmp(left, right) >= 0 {
		return left
	}
	return right
}

// Iterator return iterator of a heap
func (h Heap) Iterator() *Iterator {
	return &Iterator{h: h.Snapshot()}
}

// Snapshot return latest copy of a heap
func (h Heap) Snapshot() *Heap {
	var snap Heap
	snap.size = h.size
	snap.capacity = h.capacity
	tmp := make([]interface{}, h.size, h.capacity)
	copy(tmp, h.elements)
	snap.elements = tmp
	return &snap
}

// Iterator enabe traveling elements in the heap
type Iterator struct {
	h *Heap
}

// HasNext check if there is another element
func (it Iterator) HasNext() bool {
	return it.h.size != 0
}

// Next return the next element
func (it Iterator) Next() interface{} {
	if it.h.size == 0 {
		return nil
	}
	v := it.h.Pop()
	return v
}
