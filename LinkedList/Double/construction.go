package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type Linked struct {
	head *node
	tail *node
}

func (R *Linked) adder(value int) {
	newnode := &node{data: value}

	if R.tail == nil && R.head == nil {
		R.tail = newnode
		R.head = newnode
	} else {
		newnode.prev = R.tail
		R.tail.next = newnode
		R.tail = newnode
	}
}

func (R *Linked) Display() {
	new := R.head

	for new != nil {
		fmt.Print(new.data)
		new = new.next
	}
	fmt.Print()
}
func (R *Linked) DisplayR() {
	new := R.tail

	for new != nil {
		fmt.Print(new.data)
		new = new.prev
	}
	fmt.Print()
}

func (R *Linked) deleteStart() {
	if R.head == R.tail {
		R.head = nil
		R.tail = nil
	} else {
		R.head = R.head.next
		R.head.prev = nil
	}
}

func (R *Linked) DeleteEnd() {
	if R.head == R.tail {
		R.head = nil
		R.tail = nil
	} else {
		R.tail = R.tail.prev
		R.tail.next = nil
	}
}

func main() {
	Var := Linked{}

	Var.adder(10)
	Var.adder(20)
	Var.adder(30)
	Var.DeleteEnd()
	Var.Display()

}
