package main

import "fmt"

type node struct {
	data int
	next *node
}

type linked struct {
	head *node
	tail *node
}

func (R *linked) Addon(value int) {
	newnode := &node{data: value}

	if R.head == nil {
		R.tail = newnode
		R.head = newnode
	} else {
		R.tail.next = newnode
		R.tail = newnode
	}
}

func (R *linked) display() {
	new := R.head

	for new != nil {
		fmt.Print(new.data, " ")
		new = new.next
	}
	fmt.Println()
}

func main() {
	Var := linked{}

	Var.Addon(10)
	Var.Addon(20)
	Var.Addon(30)
	Var.Addon(40)
	Var.Addon(50)
	Var.display()

}
