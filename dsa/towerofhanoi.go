package main

import "fmt"

// The Tower of Hanoi is a problem demonstrating recursion
// It has exponential TC of O(2^n)

// Helper function for TOH
func TOH(n int, from, to, temp string) {
	if n == 1 {
		fmt.Printf("Move disk, %d from peg %s to peg %s\n", n, from, to)
		return
	}
	// Move top n - 1 disks from A to B using C as auxillary
	TOH(n-1, from, temp, to)
	fmt.Printf("Move disk, %d from peg %s to peg %s\n", n, from, to)
	// Move n - 1 disks from B to C using A as auxillary
	TOH(n-1, temp, to, from)

}

func TowersOfHanoi(n int) {
	TOH(n, "A", "C", "B")
}

func main() {
	TowersOfHanoi(3)
}
