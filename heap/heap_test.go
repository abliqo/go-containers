package heap_test

import (
	"testing"

	"github.com/abliqo/go-containers/heap"
	"github.com/stretchr/testify/assert"
)

func TestEmptyHeap(t *testing.T) {
	compareFunc := func(a, b int) bool { return a < b }
	h := heap.NewHeap[int](compareFunc)

	assert.NotNil(t, h)
	assert.True(t, h.IsEmpty())
	assert.Equal(t, 0, h.Len())
	_, ok := h.Peek()
	assert.False(t, ok)
	_, ok = h.Pop()
	assert.False(t, ok)
}

func TestHeapBasicCases(t *testing.T) {
	compareFunc := func(a, b int) bool { return a < b }
	h := heap.NewHeap[int](compareFunc)

	assert.True(t, h.IsEmpty())
	h.Push(1)
	assert.False(t, h.IsEmpty())
	assert.Equal(t, 1, h.Len())

	v, ok := h.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	assert.False(t, h.IsEmpty())

	v, ok = h.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	assert.True(t, h.IsEmpty())

	h.Push(3)
	h.Push(7)
	h.Push(5)

	v, ok = h.Peek()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}

func TestMinHeapInt(t *testing.T) {
	compareFunc := func(a, b int) bool { return a < b }
	h := heap.NewHeap[int](compareFunc)
	assert.True(t, h.IsEmpty())

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	h.Push(7)
	h.Push(5)
	h.Push(1)
	h.Push(3)
	h.Push(2)
	h.Push(6)
	h.Push(4)
	h.Push(0)

	var actual []int
	for i := 0; i < len(expected); i++ {
		v, ok := h.Pop()
		assert.True(t, ok)
		actual = append(actual, v)
	}

	assert.Equal(t, expected, actual)
	assert.True(t, h.IsEmpty())
}

func TestMaxHeapInt(t *testing.T) {
	compareFunc := func(a, b int) bool { return a > b }
	h := heap.NewHeap[int](compareFunc)
	assert.True(t, h.IsEmpty())

	expected := []int{7, 6, 5, 4, 3, 3, 3, 2, 2, 1, 0}
	h.Push(3)
	h.Push(5)
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(0)
	h.Push(7)
	h.Push(2)
	h.Push(6)
	h.Push(4)
	h.Push(3)

	var actual []int
	for i := 0; i < len(expected); i++ {
		v, ok := h.Pop()
		assert.True(t, ok)
		actual = append(actual, v)
	}

	assert.Equal(t, expected, actual)
	assert.True(t, h.IsEmpty())
}

type Task struct {
	name     string
	priority int
}

func TestMaxHeapCustomType(t *testing.T) {
	compareFunc := func(a, b Task) bool { return a.priority > b.priority }
	h := heap.NewHeap[Task](compareFunc)
	assert.True(t, h.IsEmpty())

	expected := []Task{
		{name: "a", priority: 100},
		{name: "b", priority: 80},
		{name: "c", priority: 50},
		{name: "d", priority: 30},
		{name: "e", priority: 20},
		{name: "f", priority: 10},
	}

	h.Push(expected[4])
	h.Push(expected[2])
	h.Push(expected[5])
	h.Push(expected[0])
	h.Push(expected[3])
	h.Push(expected[1])

	var actual []Task
	for i := 0; i < len(expected); i++ {
		v, ok := h.Pop()
		assert.True(t, ok)
		actual = append(actual, v)
	}

	assert.Equal(t, expected, actual)
	assert.True(t, h.IsEmpty())
}
