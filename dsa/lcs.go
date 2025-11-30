package main

func LCS(stringOne, stringTwo string) string {
	// Start at the beginning of both the strings
	return lcsHelper(stringOne, stringTwo, 0, 0)
}

func lcsHelper(stringOne, stringTwo string, i int, j int) string {
	// Check if the strings are of the same length
	if i == len(stringOne) || j == len(stringTwo) {
		return ""
	} else if stringOne[i] == stringTwo[j] {
	// If the subsequence at indices i and j match, move the pointers forward by 1
      return lcsHelper(stringOne, stringTwo, i + 1, j + 1)  
	}
	// For the last case, recurse by skipping one character in either string
	subseqOne := lcsHelper(stringOne, stringTwo, i+1, j)
	subseqTwo := lcsHelper(stringOne, stringTwo, i, j+1)

	// and taking the max substring
	if subseqOne > subseqTwo {
		return subseqOne
	}
	return subseqTwo
}