package main

import (
	"fmt"
	"math"
)

func attemptOne(chars []byte) int {
	// For strings with len = 1
	if len(chars) == 1 {
		return 1
	}

	freqMap := make(map[byte]int)

	// Iterate an count the frequency of items
	for _, char := range chars {
		freqMap[char]++
	}

	length := 0

	for char := range freqMap {
		if freqMap[char] != 0 {
			length++
			length += getNumDigits(freqMap[char])
		}
		length++
	}
	return length
}

func getNumDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)) + 1))
}

func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {
		char := chars[read]
		count := 0

		for read < len(chars) && chars[read] == char {
			read++
			count++
		}

		chars[write] = char
		write++

		if count > 1 {
			for _, digit := range []byte(fmt.Sprintf("%d", count)) {
				chars[write] = digit
				write++
			}
		}	
	}
	return write
}
