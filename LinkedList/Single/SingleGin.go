package main

import "fmt"

type node struct {
	data int
	next *node
}

type linked struct {
	Head *node
	Tail *node
}

func (R *linked) Add(value int) {
	newnode := &node{data: value}

	if R.Tail == nil {
		R.Tail = newnode
		R.Head = newnode
	} else {
		R.Tail.next = newnode
		R.Tail = newnode
	}
}

// Insert After A node
func (R *linked) InsertAfternode(value int) {
	newnode := &node{data: value}
	i := 1
	var Pos int

	curr := R.Head
	fmt.Print("enter the position")
	fmt.Scan(&Pos)
	for i < Pos-1 {
		curr = curr.next
		i++
	}
	newnode.next = curr.next
	curr.next = newnode

}

func (R linked) display() {
	Var := R.Head

	for Var != nil {
		fmt.Print(Var.data)
		Var = Var.next
	}
}

// delete the fisrt position
func (R *linked) deletefirst() {
	if R.Head == nil {
		return
	} else {
		temp := R.Head
		R.Head = R.Head.next
		temp = temp.next
	}
}

// delete the last position
func (R linked) deleteEnd() {
	var prev *node
	temp := R.Head

	for temp.next != nil {
		prev = temp
		temp = temp.next
	}
	prev.next = nil
	R.Tail = prev
}

// delete specific position
func (R *linked) DeleteSpe() {
	var position int
	i := 1
	newTemp := R.Head

	fmt.Print("enter the position")
	fmt.Scan(&position)

	for i < position-1 {
		newTemp = newTemp.next
		i++
	}
	newnext := newTemp.next
	newTemp.next = newnext.next
}

// reverse the single linked list
func (R *linked) Revere() {
	var nextN *node = nil
	current := R.Head
	var prev *node = nil

	for current != nil {
		nextN = current.next
		current.next = prev
		prev = current
		current = nextN
	}
	R.Head = prev
}
func main() {
	new := linked{}
	new.Add(10)
	new.Add(20)
	new.Add(30)
	new.Add(40)
	new.InsertAfternode(25)
	new.display()
}
