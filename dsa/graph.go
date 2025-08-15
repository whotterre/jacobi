package main

type Graph struct {
	V             int
	adjacencyList map[int][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		V:             v,
		adjacencyList: make(map[int][]int, v),
	}
}

// Breadth first search for adjacency list representation
func bfs(graph Graph) {
	visited := make([]bool, len(graph.adjacencyList))
	queue := make([]int, 0)

	// Start from the first node
	queue = append(queue, 0)
	visited[0] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		println("Visiting node:", node)

		// Visit all adjacent nodes
		for _, neighbor := range graph.adjacencyList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
func UnweightedShortestPath(G *Graph, s int) ([]int, []int, error) {
	q := CreateQueue()
	q.Enqueue(s)
	distance := make([]int, G.V)
	path := make([]int, G.V)
	// Initialize all vertices in the distance list to -1
	for i := 0; i < G.V; i++ {
		distance[i] = -1
	}

	// Set the distance from the vertex at s to 0
	distance[s] = 0

	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, w := range G.adjacencyList[v] {
			if distance[w] == -1 {
				distance[w] = distance[v] + 1
				path[w] = v
				q.Enqueue(w)
			}
		}
	}
	return distance, path, nil
}
