package main


type Graph struct {
	adjacencyList map[int][]int
}

// Breadth first seatch for adjacency list representation 
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