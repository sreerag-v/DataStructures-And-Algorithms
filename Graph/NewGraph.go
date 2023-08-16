package main

import "fmt"

type Graph struct {
	ADList map[int]map[int]bool
}

func NewGraph() *Graph {
	return &Graph{ADList: make(map[int]map[int]bool)}
}

func (R *Graph) AddVertex(vertex int) {
	if R.ADList[vertex] == nil {
		R.ADList[vertex] = make(map[int]bool)
	}
}

func (R *Graph) AddEdge(vertex1, vertex2 int) {
	R.AddVertex(vertex1)
	R.AddVertex(vertex2)

	R.ADList[vertex1][vertex2] = true
	R.ADList[vertex2][vertex1] = true
}

func (R *Graph) Display() {
	for vertex, ADvertices := range R.ADList {
		fmt.Printf("%d -- %v\n", vertex, ADvertices)
	}
}

func (R *Graph) BFS(startVertex int) []int {
	visited := make(map[int]bool)
	queue := []int{startVertex}
	result := []int{}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if visited[vertex] {
			continue
		}

		visited[vertex] = true
		result = append(result, vertex)

		for neighbor := range R.ADList[vertex] {
			queue = append(queue, neighbor)
		}
	}

	return result
}

func (R *Graph) DFS(startVertex int) []int {
	visited := make(map[int]bool)
	stack := []int{startVertex}
	result := []int{}

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[vertex] {
			continue
		}

		visited[vertex] = true
		result = append(result, vertex)

		for neighbor := range R.ADList[vertex] {
			stack = append(stack, neighbor)
		}
	}

	return result
}

func main() {
	graph := NewGraph()

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	fmt.Println("Graph:")
	graph.Display()

	fmt.Println("BFS Traversal:")
	bfsResult := graph.BFS(1)
	fmt.Println(bfsResult)

	fmt.Println("DFS Traversal:")
	dfsResult := graph.DFS(1)
	fmt.Println(dfsResult)
}
