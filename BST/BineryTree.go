package main

import "fmt"

type Node struct {
	data  int   // Data stored in the node
	left  *Node // Pointer to the left child node
	right *Node // Pointer to the right child node
}

type BinaryTree struct {
	root *Node // Pointer to the root node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(data int) {
	if t.root == nil {
		// If the tree is empty, create a new node and make it the root
		newNode := &Node{data: data, left: nil, right: nil}
		t.root = newNode
	} else {
		// If the tree is not empty, call the insertNode function recursively to find the appropriate position for insertion
		t.insertNode(data, t.root)
	}
}

func (t *BinaryTree) insertNode(data int, node *Node) {
	if node.left == nil {
		// If there is no left child, create a new node and assign it as the left child
		newNode := &Node{data: data, left: nil, right: nil}
		node.left = newNode
	} else if node.right == nil {
		// If there is no right child, create a new node and assign it as the right child
		newNode := &Node{data: data, left: nil, right: nil}
		node.right = newNode
	} else {
		// If both left and right children exist, recursively try to insert the node into the left subtree
		t.insertNode(data, node.left)
	}
}

func (t *BinaryTree) TraverseInOrder(node *Node) {
	if node != nil {
		// Recursively traverse the left subtree
		t.TraverseInOrder(node.left)
		// Print the current node's data
		fmt.Printf("%d ", node.data)
		// Recursively traverse the right subtree
		t.TraverseInOrder(node.right)
	}
}

func main() {
	tree := NewBinaryTree() // Create a new binary tree

	// Insert nodes into the binary tree
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)

	fmt.Println("Binary Tree:")
	fmt.Println("----------------")
	fmt.Println("       1")
	fmt.Println("     /   \\")
	fmt.Println("    2     3")
	fmt.Println("   / \\   / \\")
	fmt.Println("  4   5 6   7")
	fmt.Println("----------------")

	fmt.Println("\nIn-order traversal of the binary tree:")
	tree.TraverseInOrder(tree.root) // Perform in-order traversal and print the nodes
}
