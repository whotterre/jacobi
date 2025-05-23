package main

/* LeetCode 1207 - Unique Number of Occurrences
 Given an array of integers arr, return true if the number of
occurrences of each value in the array is unique, or false otherwise.

*/

func uniqueOccurrences(arr []int) bool {
	// Count frequency of numbers in arr
	freqMap := make(map[int]int)

	for _, num := range arr {
		freqMap[num]++
	}
	// Use another hashmap to keep track of the occurences
	seenFreq := make(map[int]bool)
	for _, count := range freqMap {
		if seenFreq[count] {
			return false
		}
		seenFreq[count] = true
	}
	return true
}
