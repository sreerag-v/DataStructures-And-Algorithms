package main

import "fmt"

type maxheap struct {
	arr []int
}

func (R *maxheap) Insert(Data int) {
	R.arr = append(R.arr, Data)
	R.heapify(len(R.arr) - 1)
}

func (R *maxheap) heapify(index int) {
	for R.arr[parent(index)] < R.arr[index] {
		R.Swap(parent(index), index)
		index = parent(index)
	}
}

func (R *maxheap) Extract() int {
	extract := R.arr[0]
	L := len(R.arr) - 1

	if R.arr == nil {
		fmt.Println("The array is empty")
		return -1
	}

	R.arr[0] = R.arr[L]
	R.arr = R.arr[:L]

	R.heapifydown(0)
	return extract
}

func (R *maxheap) heapifydown(index int) {
	lastindex := len(R.arr) - 1
	LE, RE := Left(index), Right(index)
	child := 0

	if LE <= lastindex {
		if LE == lastindex {
			child = LE
		} else if R.arr[LE] > R.arr[RE] {
			child = LE
		} else {
			child = RE
		}

		if R.arr[index] < R.arr[child] {
			R.Swap(index, child)
			index = child

			LE, RE = Left(index), Right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*i + 2
}

func (R *maxheap) Swap(i1, i2 int) {
	R.arr[i1], R.arr[i2] = R.arr[i2], R.arr[i1]
}

func heapSort(arr []int) {
	// Create a max heap from the input array
	heap := maxheap{}
	// for _, val := range arr {
	// 	heap.Insert(val)
	// }

	// Extract elements from the heap to obtain the sorted array
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = heap.Extract()
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Array before sorting:", arr)

	heapSort(arr)

	fmt.Println("Array after sorting:", arr)
}
