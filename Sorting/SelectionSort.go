package main

import "fmt"

func Selection(arr []int) {
	len := len(arr)

	Largest := arr[0]
	Smallest := arr[0]

	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	for i := 0; i < len; i++ {
		fmt.Println(arr[i], " ")

		if arr[i] > Largest {
			Largest = arr[i]
		} else {
			Smallest = arr[i]
		}
	}
	fmt.Println("Largest:", Largest)
	fmt.Println("Smallet:", Smallest)

}

func main() {
	arr := []int{4, 5, 7, 89, 23, 1, 3, 4, 5}

	fmt.Println("Selectin sort")

	Selection(arr)
}
