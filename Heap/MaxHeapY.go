package main

import (
	"fmt"
)

type Maxheap struct {
	Arr []int
}

// insert Add elements to the heap
func (R *Maxheap) Insert(Data int) {
	R.Arr = append(R.Arr, Data)
	R.HeapifyUp(len(R.Arr) - 1)
}

// HeapifyUp will heapify from bottom to to

func (R *Maxheap) HeapifyUp(index int) {
	for R.Arr[parent(index)] < R.Arr[index] {
		R.Swap(parent(index), index)
		// update the loop
		index = parent(index)
	}
}

// get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// to get the left child
func left(i int) int {
	return 2*i + 1
}

// to get the right child
func right(i int) int {
	return 2*i + 2
}

// swap Data in the array
func (R *Maxheap) Swap(i1, i2 int) {
	R.Arr[i1], R.Arr[i2] = R.Arr[i2], R.Arr[i1]
}

// extract returns the largest key,and remove it from the heap
func (R *Maxheap) Extract() int {

	extracted := R.Arr[0]
	L := len(R.Arr) - 1
	if R.Arr == nil {
		fmt.Print("Cannot Extraxct array is empty")
		return -1
	}
	R.Arr[0] = R.Arr[L]
	R.Arr = R.Arr[:L]
	R.HeapifyDown(0)
	return extracted
}

// Heapify will heapify the to top do bottom

func (R *Maxheap) HeapifyDown(index int) {
	lastindex := len(R.Arr) - 1
	LE, RE := left(index), right(index)
	ChildCom := 0
	// loop while index has atleast on child
	for LE < lastindex {
		if LE == lastindex { // when the left child is the only child
			ChildCom = LE
		} else if R.Arr[LE] > R.Arr[RE] { // when the left child is larger
			ChildCom = LE
		} else { //when the right child is larger
			ChildCom = RE
		}

		//compare the array value of current index to larger child and Swap if smaller

		if R.Arr[index] < R.Arr[ChildCom] {
			R.Swap(index, ChildCom)
			index = ChildCom
			LE, RE = left(index), right(index)
		} else {
			return
		}
	}
}

func main() {
	heap := Maxheap{}
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(30)
	heap.Insert(40)
	heap.Insert(50)
	fmt.Println(heap)
	extracted := heap.Extract()
	fmt.Println("Extracted:", extracted)
	fmt.Println(heap)
}
