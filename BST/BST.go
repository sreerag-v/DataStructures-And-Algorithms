package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func Insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{Data: data}
	}

	if data < root.Data {
		root.Left = Insert(root.Left, data)
	} else {
		root.Right = Insert(root.Right, data)
	}

	return root
}

func Search(root *Node, data int) bool {
	if root == nil {
		return false
	}

	if data == root.Data {
		return true
	} else if data < root.Data {
		return Search(root.Left, data)
	} else {
		return Search(root.Right, data)
	}
}
func Delete(root *Node, data int) *Node {
	if root == nil {
		return nil
	}

	if data < root.Data {
		root.Left = Delete(root.Left, data)
	} else if data > root.Data {
		root.Right = Delete(root.Right, data)
	} else {
		// Node found, delete it
		if root.Left == nil {
			// Node has no left child, return right child
			return root.Right
		} else if root.Right == nil {
			// Node has no right child, return left child
			return root.Left
		} else {
			// Node has two children, find the inorder successor and swap data
			successor := root.Right
			for successor.Left != nil {
				successor = successor.Left
			}

			// Swap data with successor
			root.Data = successor.Data

			// Delete successor
			root.Right = Delete(root.Right, successor.Data)
		}
	}

	return root
}

func InOrderTraverse(root *Node, Value []int) []int {
	if root == nil {
		return Value
	}

	Value = InOrderTraverse(root.Left, Value)
	Value = append(Value, root.Data)
	Value = InOrderTraverse(root.Right, Value)

	return Value
}

func ISValid(Root *Node) bool {
	Value := InOrderTraverse(Root, []int{})

	for i := 1; i <= len(Value); i++ {
		if Value[i] <= Value[i-1] {
			return false
		}
	}
	return true
}

func PreOrderTraverse(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Data)
	PreOrderTraverse(root.Left)
	PreOrderTraverse(root.Right)
}

func PostOrderTraverse(root *Node) {
	if root == nil {
		return
	}

	PostOrderTraverse(root.Left)
	PostOrderTraverse(root.Right)
	fmt.Printf("%d ", root.Data)
}

func main() {
	var root *Node
	root = Insert(root, 10)
	root = Insert(root, 20)
	root = Insert(root, 30)
	root = Insert(root, 40)
	root = Insert(root, 50)
	root = Insert(root, 60)
	root = Insert(root, 70)

	// fmt.Println("In-order traversal:")
	// Values := InOrderTraverse(root, []int{})
	// fmt.Print(Values)

	fmt.Println("is the BST a Valid Tree?", ISValid(root))

	// fmt.Println("Pre-order traversal:")
	// PreOrderTraverse(root)
	// fmt.Println()

	// fmt.Println("Post-order traversal:")
	// PostOrderTraverse(root)
	// fmt.Println()

	// if Search(root, 10) {
	// 	fmt.Println("10 is found in the tree")
	// } else {
	// 	fmt.Println("10 is not found in the tree")
	// }

	// if Search(root, 18) {
	// 	fmt.Println("18 is found in the tree")
	// } else {
	// 	fmt.Println("18 is not found in the tree")
	// }
}
