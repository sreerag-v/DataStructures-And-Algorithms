package main

import "fmt"

type node struct {
	Data int
	next *node
}

type Queue struct {
	Front *node
	Rear  *node
}

func (R *Queue) EnQueue(value int) {
	newnode := &node{Data: value}

	if R.Rear == nil {
		R.Rear = newnode
		R.Front = newnode
	} else {
		R.Rear.next = newnode
		R.Rear = newnode
	}
}

func (R *Queue) Display() {

	if R.Rear == nil {
		return
	}

	Temp := R.Front

	for Temp != nil {
		fmt.Println(Temp.Data)
		Temp = Temp.next
	}
	fmt.Println()
}

func (R Queue) Peak() int {

	if R.Rear == nil {
		return -1
	}

	return R.Front.Data
}

func (R *Queue) Dequeue() int {
	if R.Rear == nil {
		return -1
	}
	Pop := R.Front.Data
	R.Front = R.Front.next

	return Pop
}

func main() {
	Var := Queue{}

	Var.EnQueue(10)
	Var.EnQueue(20)
	Var.EnQueue(30)
	Var.EnQueue(40)

	Var.Display()

	Var.Dequeue()
	Var.Display()

}
