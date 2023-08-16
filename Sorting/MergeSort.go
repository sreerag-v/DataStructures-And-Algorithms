package main

import "fmt"

func MergeSort(arr []int, LB, UB int) {
	if LB < UB {
		Mid := (LB + UB) / 2

		MergeSort(arr, LB, Mid)
		MergeSort(arr, Mid+1, UB)

		Merge(arr, LB, Mid, UB)
	}
}

func Merge(arr []int, LB, Mid, UB int) {
	i := LB
	j := Mid + 1
	k := LB
	arr1 := make([]int, UB-LB+1)

	for i <= Mid && j <= UB {
		if arr[i] > arr[j] {
			arr1[k-LB] = arr[i]
			i++
		} else {
			arr1[k-LB] = arr[j]
			j++
		}
		k++
	}

	if i > Mid {
		for j <= UB {
			arr1[k-LB] = arr[j]
			j++
			k++
		}
	} else {
		for i <= Mid {
			arr1[k-LB] = arr[i]
			i++
			k++
		}
	}

	for k = LB; k <= UB; k++ {
		arr[k] = arr1[k-LB]
	}

}

func main() {
	arr := []int{23, 44, 55, 33, 42, 4, 5, 2, 44, 55}

	LB := 0
	UB := len(arr) - 1

	MergeSort(arr, LB, UB)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}

	// Largest := arr[0]
	// Smallest := arr[len(arr)-1]

	// fmt.Println("Largest:", Largest)
	// fmt.Println("Smallest:", Smallest)
}
