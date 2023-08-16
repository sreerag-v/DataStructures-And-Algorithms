package main

import "fmt"

const (
	size = 10
)

type node struct {
	key   string
	Value int
	next  *node
}

type HashTable struct {
	Table [size]*node
}

// Make hash function

func HashFunc(key string) int {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash % size
}

// make insertion function

func (R *HashTable) Insert(key string, Value int) {
	index := HashFunc(key)

	node := &node{key: key, Value: Value}

	if R.Table[index] == nil {
		R.Table[index] = node
	} else {
		current := R.Table[index]

		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

// make search Function

func (R *HashTable) Search(key string) (int, bool) {

	index := HashFunc(key)
	curr := R.Table[index]

	for curr != nil {
		if curr.key == key {
			return curr.Value, true
		}
		curr = curr.next
	}
	return 0, false
}

// make Delete function

func (R *HashTable) Delete(key string) bool {
	index := HashFunc(key)
	Current := R.Table[index]

	var prev *node

	for Current != nil {
		if Current.key == key {
			if prev == nil {
				R.Table[index] = Current.next
			} else {
				prev.next = Current.next
			}
			return true
		}
		prev = Current
		Current = Current.next
	}
	return false
}

// Make Display Function

func (R *HashTable) Display() {
	for i := 0; i < size; i++ {
		Curr := R.Table[i]

		for Curr != nil {
			fmt.Printf("key:%s value :%d\n", Curr.key, Curr.Value)
			Curr = Curr.next
		}
	}
}

func main() {
	Var := HashTable{}

	Var.Insert("Sreerag", 55)
	Var.Insert("Vaishak", 550)
	Var.Insert("Faizu", 5500)

	Var.Display()

	value, Found := Var.Search("Sreerag")

	if Found {
		fmt.Println("value founded", value)
	} else {
		fmt.Println("Value not found")
	}

	delete := Var.Delete("Sreerag")

	if delete {
		fmt.Println("Value Deleted")
	} else {
		fmt.Println("Value not Funded")
	}
	Var.Display()

}
