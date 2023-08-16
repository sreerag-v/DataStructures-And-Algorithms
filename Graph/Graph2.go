package main

import (
	"fmt"
)

type Graph struct {
	adjacencyList map[int]map[int]bool
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int]map[int]bool),
	}
}

func (g *Graph) addVertex(vertex int) {
	if g.adjacencyList[vertex] == nil {
		g.adjacencyList[vertex] = make(map[int]bool)
	}
}

func (g *Graph) addEdge(vertex1, vertex2 int) {
	g.addVertex(vertex1)
	g.addVertex(vertex2)

	g.adjacencyList[vertex1][vertex2] = true
	g.adjacencyList[vertex2][vertex1] = true
}

func (g *Graph) hasEdge(vertex1, vertex2 int) bool {
	_, exists := g.adjacencyList[vertex1][vertex2]
	return exists
}

func (g *Graph) removeEdge(vertex1, vertex2 int) {
	delete(g.adjacencyList[vertex1], vertex2)
	delete(g.adjacencyList[vertex2], vertex1)
}

func (g *Graph) removeVertex(vertex int) {
	if g.adjacencyList[vertex] == nil {
		return
	}

	for adjacentVertex := range g.adjacencyList[vertex] {
		g.removeEdge(vertex, adjacentVertex)
	}

	delete(g.adjacencyList, vertex)
}

func (g *Graph) display() {
	for vertex, adjacentVertices := range g.adjacencyList {
		fmt.Printf("%d --> %v\n", vertex, getKeys(adjacentVertices))
	}
}

func getKeys(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	graph := NewGraph()

	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)

	graph.addEdge(1, 2)
	graph.addEdge(3, 2)

	fmt.Println("Before removing edge:")
	graph.display()

	// graph.removeEdge(1, 2)
	// graph.removeVertex(3)
	// // graph.removeVertex(1)
	// fmt.Println(graph.hasEdge(1, 2))

	// fmt.Println("After removing edge:")
	// graph.display()
}
