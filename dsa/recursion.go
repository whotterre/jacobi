package main

import "fmt"

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		// Function to print the output
		fmt.Print(A[i])
	}
	fmt.Printf("\n")
}

// Function to generate all binary strings
func generateBinaryStrings(n int, A []int, i int) {
	if i == n {
		printResult(A, n)
		return
	}
	// assign "0" at ith position and try for all other permutations for remaining positions
	A[i] = 0
	generateBinaryStrings(n, A, i+1)
	// assign "1" at ith position and try for all other permutations for remaining positions
	A[i] = 1
	generateBinaryStrings(n, A, i+1)
}

func main() {
    n := 12 // The desired length of the binary strings
    // Initialize a slice of integers with length 'n'
    // All elements will be initialized to their zero value (0 for int)
    A := make([]int, n)
    generateBinaryStrings(n, A, 0) // Start generating from the first position (index 0)
}