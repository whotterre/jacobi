package main

import (
	"unicode"
)

func camelcase(s string) int32 {
	var count int32

	if len(s) == 0 {
		count = 0
	}

	count = 1
	// Note the locations of uppercase characters
	// For each location, check that the following and preceding characters are both lowercase
	// If so, increase the count
	for i := 0; i < len(s); i++ {
		curr := rune(s[i])
		if unicode.IsUpper(curr) {
			count++
		}
	}
	return count
}
