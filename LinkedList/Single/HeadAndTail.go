package main

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (ll *LinkedList) addNode(value int) {
	newNode := &node{data: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	fmt.Println("Enter values to add to the linked list (enter 'q' to quit):")
	for {
		var input string
		fmt.Scan(&input)

		if input == "q" {
			break
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer or 'q' to quit.")
			continue
		}

		list.addNode(value)
	}

	fmt.Println("Linked List:")
	list.display()
}
