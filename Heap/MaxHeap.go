package main

import "fmt"

// MaxHeap struct represents a max heap data structure.
type MaxHeap struct {
	array []int
}

// NewMaxHeap creates and initializes a new MaxHeap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		array: []int{},
	}
}

// Insert adds a new element to the heap and maintains the heap property.
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value) // Add the value to the end of the array
	h.heapifyUp(len(h.array) - 1)    // Restore the heap property by moving the value up
}

// Delete removes the maximum element from the heap and maintains the heap property.
func (h *MaxHeap) Delete() {
	if len(h.array) == 0 {
		return
	}

	lastIndex := len(h.array) - 1
	h.swap(0, lastIndex)          // Swap the maximum element with the last element
	h.array = h.array[:lastIndex] // Remove the last element from the array
	h.heapifyDown(0)              // Restore the heap property by moving the swapped element down
}

// heapifyUp restores the heap property by moving an element up as needed.
func (h *MaxHeap) heapifyUp(index int) {
	parentIndex := (index - 1) / 2 // Calculate the parent index of the given index

	// Keep swapping the element with its parent until the heap property is restored
	for index > 0 && h.array[index] > h.array[parentIndex] {
		h.swap(index, parentIndex)
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

// heapifyDown restores the heap property by moving an element down as needed.
func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftChild := 2*index + 1  // Calculate the left child index
	rightChild := 2*index + 2 // Calculate the right child index
	largest := index

	// Find the largest child among the left and right children, if they exist
	if leftChild <= lastIndex && h.array[leftChild] > h.array[largest] {
		largest = leftChild
	}

	if rightChild <= lastIndex && h.array[rightChild] > h.array[largest] {
		largest = rightChild
	}

	// If the largest child is different from the current index, swap them and continue heapifying down
	if largest != index {
		h.swap(index, largest)
		h.heapifyDown(largest)
	}
}

// swap swaps two elements in the array.
func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	heap := NewMaxHeap() // Create a new MaxHeap
	elements := []int{4, 10, 3, 5, 1}

	// Insert elements into the heap
	for _, elem := range elements {
		heap.Insert(elem)
	}

	fmt.Println("Max Heap:")
	// Delete and print the maximum element until the heap is empty
	for len(heap.array) > 0 {
		fmt.Printf("%d ", heap.array[0])
		heap.Delete()
	}
}
