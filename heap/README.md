# Binary Heap 

This package implements a [binary heap](https://en.wikipedia.org/wiki/Binary_heap) using Go generics. It can store items of any type with specified comparison function.

See [test file](heap_test.go) for examples of usage.

## Why this package exists?

The Go standard library has [implementation](https://pkg.go.dev/container/heap) of binary heap that does not support generics. It is cumbersome to use in practice. It requires to implement your own additional struct and functions, but also requires to use the functions from the standard 'heap' package when working with the heap implemented that way.

As of this writing, there is an [open issue](https://github.com/golang/go/issues/47632) to modernize that package in the standard library.


