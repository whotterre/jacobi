package main

func maxVowels(s string, k int) int {
	// Define a set of vowels
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	currentVowelCount := 0
	sRunes := []rune(s)
	// Two stages:
	// Stage 1: Check the number of vowels in the string
	for i := 0; i < k; i++ {
		if _, ok := vowels[sRunes[i]]; ok {
			currentVowelCount++
		}
	}
	maxVowels := currentVowelCount
	// Stage 2: Iterate through the array
	// with a window of boundaries k and the length of the string
	for i := k; i < len(s); i++ {
		if _, ok := vowels[sRunes[i]]; ok {
			currentVowelCount++
		}
		// Considering that the upper limit is k >= 0
		// We can be sure we don't get a negative number
		if _, ok := vowels[sRunes[i-k]]; ok {
			currentVowelCount--
		}
		// Update maxVowels
		if currentVowelCount > maxVowels {
			maxVowels = currentVowelCount
		}
	}
	return maxVowels
}
