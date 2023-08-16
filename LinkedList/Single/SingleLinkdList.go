package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) addNode(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}
func (list *LinkedList) printList() {
	curent := list.head
	for curent != nil {
		fmt.Print(curent.data, " ")
		curent = curent.next
	}
	fmt.Println()
}

func (list *LinkedList) deleteNode(data int) {
	if list.head == nil {
		fmt.Println("List is empty")
	} else if list.head.data == data {
		list.head = list.head.next
	} else {
		previous := list.head
		current := list.head.next

		for current != nil {
			if current.data == data {
				previous.next = current.next
			}
			previous = current
			current = current.next
		}

	}
}

// Method to delete a specific value after a value
func (list *LinkedList) delSpecified(data int) {
	if list.head == nil {
		fmt.Println("List is empty")
	} else if list.head.data == data {
		previous := list.head
		current := list.head.next
		if current != nil {
			previous.next = current.next
		} else {
			fmt.Println("List only contains two values!")
		}

	} else {
		current := list.head.next

		for current != nil {
			if current.data == data {
				if current.next != nil {
					current.next = current.next.next
				} else if current.data == data && current.next == nil {
					fmt.Println("There is no value after", data)
				}
			}
			current = current.next
		}
	}
}

// Method to Reverse linked list
func (list *LinkedList) reverse() {
	if list.head == nil || list.head.next == nil {
		return
	}
	var previousNode *Node
	currentNode := list.head

	for currentNode != nil {
		// fmt.Println("CurrentNode",currentNode)
		// fmt.Println()
		nextNode := currentNode.next
		// fmt.Println("NextNode",nextNode,"<- CurrentNode.next",currentNode.next)
		// fmt.Println()
		currentNode.next = previousNode
		// fmt.Println("CurrentNode.next",currentNode.next,"<- PreviousNode",previousNode)
		// fmt.Println()
		previousNode = currentNode
		// fmt.Println("PreviousNode",previousNode,"<- CurrentNode",currentNode)
		// fmt.Println()
		currentNode = nextNode
		// fmt.Println("CurrentNode",currentNode, "<- NextNode",nextNode)
		// fmt.Println()
	}
	list.head = previousNode
}
func main() {
	list := LinkedList{}
	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)
	list.addNode(60)

	list.printList()
	list.reverse()
	list.printList()
	list.reverse()
	list.printList()

}
