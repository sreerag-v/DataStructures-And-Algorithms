package main

import (
	"fmt"
)

type Queue struct {
	Data  []int
	Front int
	Rear  int
}

func (R *Queue) EnQueue(Data int) {
	R.Data = append(R.Data, Data)

	if R.Front == -1 && R.Rear == -1 {
		R.Front = 0
	}
	R.Rear++
}

func (R *Queue) Dequeue() (int, error) {
	if R.Front == -1 && R.Rear == -1 {
		return 0, fmt.Errorf("its iempty")
	}

	picked := R.Data[R.Front]

	if R.Front == R.Rear {
		R.Front = -1
		R.Rear = -1
	} else {
		R.Front++
	}
	R.Data = R.Data[R.Front:]

	return picked, nil
}

func (R *Queue) front() (int, error) {
	if R.Front == -1 {
		return 0, fmt.Errorf("error")
	}

	frontE := R.Data[R.Front]

	return frontE, nil
}

func main() {
	Var := Queue{}

	Var.EnQueue(10)
	Var.EnQueue(20)
	Var.EnQueue(30)
	Var.EnQueue(40)

	fmt.Println(Var.Data)

	num, err := Var.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	fmt.Println(Var.Data)

	FrontE, err := Var.front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(FrontE)
	}

}
