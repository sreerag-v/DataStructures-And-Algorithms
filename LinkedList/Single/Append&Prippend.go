package main

import "fmt"

type Node struct {
	Number interface{}
	next   *Node
}

type Linked struct {
	Head *Node
	Tail *Node
}

func (R *Linked) Append(Value interface{}) {
	Newnode := &Node{Number: Value}

	if R.Tail == nil {
		R.Tail = Newnode
		R.Head = Newnode
		return
	} else {
		R.Tail.next = Newnode
		R.Tail = Newnode
	}
}

func (R *Linked) Prepend(Value interface{}) {
	newNode := &Node{Number: Value}

	if R.Head == nil {
		R.Head = newNode
		R.Tail = newNode
		return
	} else {
		newNode.next = R.Head
		R.Head = newNode
	}
}

func (R *Linked) Display() {
	New := R.Head

	for New != nil {
		fmt.Println(New.Number)
		New = New.next
	}
	fmt.Println()
}

func main() {
	Print := Linked{}

	Print.Append(12)
	Print.Append(13)
	Print.Append(14)
	Print.Prepend(11)
	Print.Display()

}
