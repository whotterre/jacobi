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

