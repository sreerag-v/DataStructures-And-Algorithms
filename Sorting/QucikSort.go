package main

import "fmt"

func partition(arr []int, LB, UB int) int {
	Start := LB
	End := UB

	P := arr[LB]

	for Start < End {

		for arr[Start] <= P && Start < End {
			Start++
		}

		for arr[End] > P {
			End--
		}

		if Start < End {
			arr[Start], arr[End] = arr[End], arr[Start]
		} else {
			arr[LB], arr[End] = arr[End], arr[LB]
		}
	}
	return End
}

func Quick(arr []int, LB, UB int) {
	if LB < UB {
		lock := partition(arr, LB, UB)

		Quick(arr, LB, lock-1)
		Quick(arr, lock+1, UB)
	}

}

func main() {
	arr := []int{1, 34, 56, 78, 22}

	LB := 0
	UB := len(arr) - 1

	Quick(arr, LB, UB)

	len := len(arr)
	fmt.Println("Quick Sort:")
	for i := 0; i < len; i++ {
		fmt.Println(arr[i], " ")
	}
	smallet := arr[0]
	Larget := arr[len-1]

	fmt.Println("Larget:", Larget)
	fmt.Println("Smallet:", smallet)

}
