package main

import "fmt"

type Graph struct {
	ADList map[int]map[int]bool
}

func newGraph() *Graph {
	return &Graph{ADList: make(map[int]map[int]bool)}
}
func (R *Graph) Addvertex(vertex int) {
	if R.ADList[vertex] == nil {
		R.ADList[vertex] = make(map[int]bool)
	}
}
func (R *Graph) AddEdges(vertex1, vertex2 int) {
	R.Addvertex(vertex1)
	R.Addvertex(vertex2)

	R.ADList[vertex1][vertex2] = true
	R.ADList[vertex2][vertex1] = true
}

func (R *Graph) Display() {
	for vertex, vertices := range R.ADList {
		fmt.Printf("%d--%v\n", vertex, vertices)
	}
}

func main() {
	graph := newGraph()

	graph.Addvertex(10)
	graph.Addvertex(20)
	graph.Addvertex(30)
	graph.Addvertex(40)

	graph.AddEdges(10, 40)
	graph.AddEdges(20, 30)
	graph.AddEdges(10, 20)

	graph.Display()
}
