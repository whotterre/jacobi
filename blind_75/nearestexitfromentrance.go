package main

import (
	"container/list"
)

func nearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	if rows == 0 {
		return -1
	}
	cols := len(maze[0])

	// Initialize queue with starting position and steps
	queue := list.New()
	queue.PushBack([]int{entrance[0], entrance[1], 0})
	maze[entrance[0]][entrance[1]] = '+'

	// Directions: down, up, right, left
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		current := front.Value.([]int)
		r, c, steps := current[0], current[1], current[2]

		// Check all 4 directions
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]

			// Check boundaries and if cell is passable
			if nr >= 0 && nc >= 0 && nr < rows && nc < cols && maze[nr][nc] == '.' {
				// Check if we're at the boundary (exit)
				if nr == 0 || nc == 0 || nr == rows-1 || nc == cols-1 {
					return steps + 1
				}
				// Mark as visited and add to queue
				maze[nr][nc] = '+'
				queue.PushBack([]int{nr, nc, steps + 1})
			}
		}
	}

	return -1
}
