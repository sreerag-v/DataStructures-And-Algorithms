package main

import "fmt"

func insertion(arr []int) {

	// var Smallest int = arr[0]
	// var Largest int = arr[0]
	len := len(arr)
	var temp int
	for i := 1; i < len; i++ {
		temp = arr[i]
		j := i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}

	for i := 0; i < len; i++ {
		fmt.Println(arr[i], " ")

		// if arr[i] > Largest {
		// 	Largest = arr[i]
		// } else {
		// 	Smallest = arr[i]
		// }
	}
	// fmt.Println("Largest:", Largest)
	// fmt.Println("Smallest:", Smallest)

}

func main() {

	arr := []int{1, 2, 4, 6, 5, 8, 6, 9, 10}

	fmt.Print("After the the sort:")

	insertion(arr)

}
