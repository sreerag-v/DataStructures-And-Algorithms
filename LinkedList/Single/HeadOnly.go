package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

func (R *LinkedList) addnode(value int) {

	newNode := &node{data: value}

	if R.head == nil {
		R.head = newNode
	} else {
		common := R.head

		for common.next != nil {
			common = common.next
		}
		common.next = newNode
	}
}

func (R *LinkedList) display() {
	common := R.head
	for common != nil {
		fmt.Print(common.data, "")
		common = common.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.addnode(10)
	list.addnode(20)
	list.addnode(30)

	list.display()
}
