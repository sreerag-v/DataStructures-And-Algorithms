package main

import "fmt"

type node struct {
	data int
	next *node
}

type Stack struct {
	top *node
}

func (s *Stack) Push(value int) {
	newNode := &node{data: value}

	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return -1 // Return -1 or an appropriate value to indicate underflow
	}

	value := s.top.data
	s.top = s.top.next
	return value
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return -1 // Return -1 or an appropriate value to indicate underflow
	}

	return s.top.data
}
func (R *Stack) DeleteMiddle() {
	TempStack := Stack{}
	count := 0
	middle := 0

	current := R.top
	for current != nil {
		TempStack.Push(current.data)
		current = current.next
		count++
	}

	middle = count / 2

	R.top = nil
	for i := 0; i < count; i++ {
		if i != middle {
			R.Push(TempStack.Pop())
		} else {
			TempStack.Pop()
		}
	}
}

func (s *Stack) Display() {
	if s.top == nil {
		fmt.Println("Stack is empty.")
		return
	}

	fmt.Println("Stack elements:")
	currentNode := s.top
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	stack.Display()

	fmt.Println()

	stack.DeleteMiddle()
	stack.Display()

}
