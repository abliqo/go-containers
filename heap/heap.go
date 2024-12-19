package heap

// Heap is a generic implementation of a binary heap.
type Heap[T any] struct {
	compareFunc func(a, b T) bool
	items       []T
}

// CompareFunc determines how to order items of type 'T'.
// For a min-heap, the items must be ordered in ascending order 'a' < 'b'.
// For a max-heap (priority queue), the items must be ordered
// in descending order 'a' > 'b'.
type CompareFunc[T any] func(a, b T) bool

// NewHeap creates a new empty heap.
func NewHeap[T any](compareFunc CompareFunc[T]) *Heap[T] {
	return &Heap[T]{
		compareFunc: compareFunc,
		items:       make([]T, 0),
	}
}

// IsEmpty returns 'true' if the heap is empty, otherwise returns 'false'.
func (h *Heap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

// Len returns the number of itmes in the heap.
func (h *Heap[T]) Len() int {
	return len(h.items)
}

// Push inserts a new itme into the heap.
func (h *Heap[T]) Push(item T) {
	h.items = append(h.items, item)
	h.up(h.Len() - 1)
}

// Peek returns the top item from the heap without removing it.
func (h *Heap[T]) Peek() (_ T, _ bool) {
	if h.IsEmpty() {
		return
	}

	return h.items[0], true
}

// Removes the top item from the heap.
func (h *Heap[T]) Pop() (_ T, _ bool) {
	if h.IsEmpty() {
		return
	}

	n := h.Len() - 1
	h.swap(0, n)
	h.down(0, n)

	res := h.items[n]
	h.items = h.items[0:n]

	return res, true
}

// The implementations of below functions are copied from
// https://cs.opensource.google/go/go/+/refs/tags/go1.23.2:src/container/heap/heap.go
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

func (h *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.compareFunc(h.items[j], h.items[i]) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.compareFunc(h.items[j2], h.items[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.compareFunc(h.items[j], h.items[i]) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}
