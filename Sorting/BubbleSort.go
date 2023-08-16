package main

import "fmt"

func Bubble(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		flag := 0
		for j := 0; j < size-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}

	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}

}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 4, 32, 1}

	fmt.Println("The Acending sorted array:")

	Bubble(arr)
}
