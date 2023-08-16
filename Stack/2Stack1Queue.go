package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type Stack struct {
	Top *node
}

func (s *Stack) Push(value int) {
	newNode := &node{value: value}

	if s.Top == nil {
		s.Top = newNode
	} else {
		newNode.next = s.Top
		s.Top = newNode
	}
}

func (s *Stack) Pop() (int, error) {
	if s.Top == nil {
		return 0, fmt.Errorf("stack is empty, cannot pop")
	}
	value := s.Top.value
	s.Top = s.Top.next
	return value, nil
}

func (s Stack) Display() {
	current := s.Top

	for current != nil {
		fmt.Println(current.value, " ")
		current = current.next
	}
}

type Queue struct {
	stack1 Stack
	stack2 Stack
}

func (q *Queue) Enqueue(value int) {
	q.stack1.Push(value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.stack2.Top == nil {
		for q.stack1.Top != nil {
			value, _ := q.stack1.Pop()
			q.stack2.Push(value)
		}
	}

	if q.stack2.Top == nil {
		return 0, fmt.Errorf("queue is empty, cannot dequeue")
	}

	value, _ := q.stack2.Pop()
	return value, nil
}

func main() {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)

	queue.stack1.Display()

	value, err := queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued value:", value)
	} else {
		fmt.Println(err)
	}

	queue.stack2.Display()
}
