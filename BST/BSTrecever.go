package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) Insert(data int) {
	if t.Root == nil {
		t.Root = &Node{Data: data}
	} else {
		t.Root.Insert(data)
	}
}

func (n *Node) Insert(data int) {
	if data < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: data}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Data: data}
		} else {
			n.Right.Insert(data)
		}
	}
}

func (t *BinaryTree) Search(data int) bool {
	return t.Root.Search(data)
}

func (n *Node) Search(data int) bool {
	if n == nil {
		return false
	}

	if data == n.Data {
		return true
	} else if data < n.Data {
		return n.Left.Search(data)
	} else {
		return n.Right.Search(data)
	}
}

func (t *BinaryTree) InOrderTraverse() {
	t.Root.InOrderTraverse()
	fmt.Println()
}

func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}

	n.Left.InOrderTraverse()
	fmt.Printf("%d ", n.Data)
	n.Right.InOrderTraverse()
}

func (t *BinaryTree) PreOrderTraverse() {
	t.Root.PreOrderTraverse()
	fmt.Println()
}

func (n *Node) PreOrderTraverse() {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.Data)
	n.Left.PreOrderTraverse()
	n.Right.PreOrderTraverse()
}

func (t *BinaryTree) PostOrderTraverse() {
	t.Root.PostOrderTraverse()
	fmt.Println()
}

func (n *Node) PostOrderTraverse() {
	if n == nil {
		return
	}

	n.Left.PostOrderTraverse()
	n.Right.PostOrderTraverse()
	fmt.Printf("%d ", n.Data)
}

func main() {
	tree := BinaryTree{}

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	fmt.Println("In-order traversal:")
	tree.InOrderTraverse()

	fmt.Println("Pre-order traversal:")
	tree.PreOrderTraverse()

	fmt.Println("Post-order traversal:")
	tree.PostOrderTraverse()

	if tree.Search(10) {
		fmt.Println("10 is found in the tree")
	} else {
		fmt.Println("10 is not found in the tree")
	}

	if tree.Search(18) {
		fmt.Println("18 is found in the tree")
	} else {
		fmt.Println("18 is not found in the tree")
	}
}
