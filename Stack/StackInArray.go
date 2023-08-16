package main

import "fmt"

type Stack struct {
	Elements []int
}

// push
func (R *Stack) Push(Element int) {
	R.Elements = append(R.Elements, Element)
}

// pop
func (R *Stack) pop() (int, error) {
	if len(R.Elements) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}

	index := len(R.Elements) - 1
	popped := R.Elements[index]
	R.Elements = R.Elements[:index]

	return popped, nil
}

// peak
func (R *Stack) Peak() (int, error) {
	if len(R.Elements) == 0 {
		return 0, fmt.Errorf("error")
	}

	peak := len(R.Elements) - 1
	peak = R.Elements[peak]

	return peak, nil
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	fmt.Println(stack.Elements)

	popped, err := stack.pop()

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("popped:", popped)
	}

	fmt.Println(stack.Elements)

	// peak, err := stack.Peak()

	// if err != nil {
	// 	fmt.Println("error")
	// } else {
	// 	fmt.Print("Peak Element:", peak)
	// }

}
