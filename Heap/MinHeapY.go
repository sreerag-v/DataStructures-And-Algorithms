package main

import (
	"fmt"
)

type MinHeap struct {
	Arr []int
}

// Insert adds elements to the heap
func (h *MinHeap) Insert(data int) {
	h.Arr = append(h.Arr, data)
	h.HeapifyUp(len(h.Arr) - 1)
}

// Extract returns the smallest key and removes it from the heap
func (h *MinHeap) Extract() int {
	if len(h.Arr) == 0 {
		fmt.Println("Cannot extract, array is empty")
		return -1
	}

	extracted := h.Arr[0]
	lastIndex := len(h.Arr) - 1
	h.Arr[0] = h.Arr[lastIndex]
	h.Arr = h.Arr[:lastIndex]

	h.HeapifyDown(0)

	return extracted
}

// HeapifyUp heapifies from bottom to top
func (h *MinHeap) HeapifyUp(index int) {
	for h.Arr[parent(index)] > h.Arr[index] {
		h.Swap(parent(index), index)
		index = parent(index)
	}
}

// HeapifyDown heapifies from top to bottom
func (h *MinHeap) HeapifyDown(index int) {
	lastIndex := len(h.Arr) - 1
	leftIdx, rightIdx := left(index), right(index)
	childIdx := 0

	for leftIdx <= lastIndex {
		if leftIdx == lastIndex { // Only left child exists
			childIdx = leftIdx
		} else if h.Arr[leftIdx] < h.Arr[rightIdx] { // Left child is smaller
			childIdx = leftIdx
		} else { // Right child is smaller
			childIdx = rightIdx
		}

		if h.Arr[index] > h.Arr[childIdx] {
			h.Swap(index, childIdx)
			index = childIdx
			leftIdx, rightIdx = left(index), right(index)
		} else {
			return
		}
	}
}

// Get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get left child index
func left(i int) int {
	return 2*i + 1
}

// Get right child index
func right(i int) int {
	return 2*i + 2
}

// Swap data in the array
func (h *MinHeap) Swap(i1, i2 int) {
	h.Arr[i1], h.Arr[i2] = h.Arr[i2], h.Arr[i1]
}

func main() {
	heap := MinHeap{}
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(30)
	heap.Insert(40)
	heap.Insert(50)
	heap.Insert(60)
	heap.Insert(70)
	fmt.Println(heap)
	// extracted := heap.Extract()
	// fmt.Println("Extracted:", extracted)
	// fmt.Println(heap)
}
