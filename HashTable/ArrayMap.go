package main

import "fmt"

const (
	size = 10
)

type node struct {
	key   string
	value int
	next  *node
}

type HashTable struct {
	Table [size]*node
}

func HashFunc(key string) int {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash % size
}

func (R *HashTable) Insert(key string, value int) {
	index := HashFunc(key)

	node := &node{key: key, value: value}

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

func (R *HashTable) Display() {
	for i := 0; i < size; i++ {
		curr := R.Table[i]

		for curr != nil {
			fmt.Printf("key:%s value:%d\n", curr.key, curr.value)
			curr = curr.next
		}
	}
}

func main() {
	node := HashTable{}

	node.Insert("Sreerag", 50)
	node.Insert("Hello", 55)

	node.Display()

}
