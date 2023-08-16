package main

import "fmt"

const alphabetSize = 26

type TrieNode struct {
	children [alphabetSize]*TrieNode
	isWord   bool
}

// trie reprtesent to point the root node
type Trie struct {
	root *TrieNode
}

// By initializing the root node with an empty TrieNode, we ensure that the trie starts with a clean state and has a valid root node to build upon.
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		index := char - 'a'                 // Calculate the index based on the character's ASCII value relative to 'a'
		if current.children[index] == nil { // If the child node does not exist, create a new one
			current.children[index] = &TrieNode{}
		}
		current = current.children[index] // Move to the child node
	}
	current.isWord = true // Mark the last node as the end of a word
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		index := char - 'a'                 // Calculate the index based on the character's ASCII value relative to 'a'
		if current.children[index] == nil { // If the child node does not exist, the word does not exist
			return false
		}
		current = current.children[index] // Move to the child node
	}
	return current.isWord // Return whether the last node represents the end of a word
}

func PrintTrie(node *TrieNode, Prefix string) {
	if node == nil {
		return
	}

	if node.isWord {
		fmt.Println(Prefix)
	}

	for i, child := range node.children {
		if child != nil {
			PrintTrie(child, Prefix+string('a'+i))
		}
	}
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("orange")

	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("banana"))  // Output: true
	fmt.Println(trie.Search("orange"))  // Output: true
	fmt.Println(trie.Search("grape"))   // Output: false
	fmt.Println(trie.Search("apricot")) // Output: false
}
