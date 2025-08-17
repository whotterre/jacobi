package main

import (
	"errors"
	"fmt"
)

type GraphType string

const (
	DIRECTED   GraphType = "Directed"
	UNDIRECTED GraphType = "Undirected"
)

// Adjacency matrix graph representation
type AdjacencyMatrix struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjMatrix [][]int
}

// Initializes the graph
func (g *AdjacencyMatrix) Init() {
	g.AdjMatrix = make([][]int, g.Vertices)
	g.Edges = 0
	for i := 0; i < g.Vertices; i++ {
		g.AdjMatrix[i] = make([]int, g.Vertices)
	}
}

// Add edge
func (g *AdjacencyMatrix) AddEdge(vOne, vTwo int) error {
	// Handle edge cases
	if vOne >= g.Vertices || vTwo >= g.Vertices || vOne < 0 || vTwo < 0 {
		return errors.New("Index out of bounds")
	}
	// Two cases -> a directed and an undirected
	g.AdjMatrix[vOne][vTwo] = 1
	g.Edges++
	if g.GraphType == UNDIRECTED {
		// Add an opposite relationship
		g.AdjMatrix[vTwo][vOne] = 1
		g.Edges++
	}
	return nil
}

func (g *AdjacencyMatrix) AddWeightedEdge(vOne, vTwo, weight int) error {
	if vOne >= g.Vertices || vTwo >= g.Vertices || vOne < 0 || vTwo < 0 {
		return errors.New("Index out of bounds")
	}

	g.AdjMatrix[vOne][vTwo] = weight
	g.Edges++

	if g.GraphType == UNDIRECTED {
		g.AdjMatrix[vTwo][vOne] = weight
		g.Edges++
	}
	return nil
}

func (g *AdjacencyMatrix) RemoveEdge(vOne, vTwo int) error {
	// Add edge
	if vOne >= g.Vertices || vTwo >= g.Vertices || vOne < 0 || vTwo < 0 {
		return errors.New("Index out of bounds")
	}
	// Two cases -> a directed and an undirected
	// If directed, remove only one edge
	g.AdjMatrix[vOne][vTwo] = 0
	g.Edges--

	// If undirected, remove both edges
	if g.GraphType == UNDIRECTED {
		g.AdjMatrix[vTwo][vOne] = 0
		g.AdjMatrix[vOne][vTwo] = 0
		g.Edges--
	}
	return nil
}

func (g *AdjacencyMatrix) HasEdge(vOne, vTwo int) bool {
	if vOne >= g.Vertices || vTwo >= g.Vertices || vOne < 0 || vTwo < 0 {
		return false
	}

	return g.AdjMatrix[vOne][vTwo] != 0
}

func (g *AdjacencyMatrix) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	adjMatrixVertices := map[int]bool{}
	if vertex >= g.Vertices || vertex < 0 {
		return adjMatrixVertices
	}
	for i := 0; i < g.Vertices; i++ {
		if g.AdjMatrix[vertex][i] != 0 {
			adjMatrixVertices[i] = g.AdjMatrix[vertex][i] != 0
		}
	}
	return adjMatrixVertices
}

func main() {
	// Create a directed graph with 4 vertices
	testGraph := &AdjacencyMatrix{Vertices: 4, GraphType: DIRECTED}
	testGraph.Init()

	// Add edges
	testGraph.AddEdge(0, 1)
	testGraph.AddEdge(0, 2)
	testGraph.AddEdge(2, 3)

	// Get adjacent nodes for vertex 0
	adj := testGraph.GetAdjacentNodesForVertex(0)

	// Print the map directly
	fmt.Println("Adjacent nodes for vertex 0:", adj)

	// Or iterate for a cleaner display
	fmt.Print("Adjacent nodes for vertex 0: ")
	for v := range adj {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
