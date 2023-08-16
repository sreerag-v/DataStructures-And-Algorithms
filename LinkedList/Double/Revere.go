package main

import "fmt"

type node struct {
	Value int
	next  *node
	prev  *node
}

type Double struct {
	Head *node
	Tail *node
}

func (R *Double) Append(value int) {
	NewNode := &node{Value: value}

	if R.Head == nil && R.Tail == nil {
		R.Head = NewNode
		R.Tail = NewNode
	} else {
		NewNode.prev = R.Tail
		R.Tail.next = NewNode
		R.Tail = NewNode
	}
}
func (R *Double) Preppend(value int) {
	NewNode := &node{Value: value}
	if R.Head == nil && R.Tail == nil {
		R.Head = NewNode
		R.Tail = NewNode
	} else {
		R.Head.prev = NewNode
		NewNode.next = R.Head
		R.Head = NewNode

	}

}

func (R *Double) Display() {
	Dis := R.Head

	for Dis != nil {
		fmt.Print(Dis.Value, " ")
		Dis = Dis.next
	}
	fmt.Println()
}

func (R Double) DisplayReverse() {
	Dis := R.Tail

	for Dis != nil {
		fmt.Print(Dis.Value, " ")
		Dis = Dis.prev
	}
}

func main() {
	Var := Double{}

	Var.Append(10)
	Var.Append(20)
	Var.Append(30)
	Var.Preppend(5)

	Var.Display()
}
