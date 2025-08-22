package main

import (
	"strings"
	"unicode"
)

func caeserCipher(s string, k int32) string {
	lowerCharSet := "abcdefghijklmnopqrstuvwxyz"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Rotate the alphabet by k in a map
	rotLowerAlphabet := make(map[rune]rune)
	rotUpperAlphabet := make(map[rune]rune)
	// Lower
	for i := 0; i < len(lowerCharSet); i++ {
		idx := (int32(i) + k) % int32(len(lowerCharSet))
		rotLowerAlphabet[rune(lowerCharSet[i])] = rune(lowerCharSet[idx])
	}
	// Upper
	for i := 0; i < len(upperCharSet); i++ {
		idx := (int32(i) + k) % int32(len(upperCharSet))
		println(idx)
		rotUpperAlphabet[rune(upperCharSet[i])] = rune(upperCharSet[idx])
	}
	for k, val := range rotUpperAlphabet {
		println(string(k), string(val))
	}
	// Go through the chars in the string
	// If it's an uppercase char, look it up in the rotatedAlphabet.
	// then change back to upperCase in the caeserCipher
	// println(len(rotatedAlphabet))
	var encryptedString strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) {
			if unicode.IsLower(char) {
				encryptedString.WriteRune(rune(rotLowerAlphabet[char]))
			}
			if unicode.IsUpper(char) {
				encryptedString.WriteRune(rune(rotUpperAlphabet[char]))
			}
		}
		// If it's a symbol rune, then just put it there
		if unicode.IsSymbol(char) || unicode.IsPunct(char) || unicode.IsDigit(char) {
			encryptedString.WriteRune(char)
		}
	}
	return encryptedString.String()
}

func main() {
	print(caeserCipher("D3q4", 0))
}
