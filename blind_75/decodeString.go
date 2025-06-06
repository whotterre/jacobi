package main

import (
	"strconv"
	"strings"
	"unicode"
)
// LeetCode 394 - Decode String
// Given an encoded string, return its decoded string.
func decodeString(s string) string {
    var stack []rune
    for _, char := range s {
        if char != ']' {
            stack = append(stack, char)
        } else {
            // Build the current string
            var curr []rune
            for len(stack) > 0 && stack[len(stack)-1] != '[' {
                curr = append([]rune{stack[len(stack)-1]}, curr...)
                stack = stack[:len(stack)-1]
            }
            
            // Remove the '[' by popping it from the stack
            stack = stack[:len(stack)-1]
            
            // Build the number
            var k []rune
            for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
                k = append([]rune{stack[len(stack)-1]}, k...)
                stack = stack[:len(stack)-1]
            }
            
            // Convert k to integer and repeat the current string
            num, _ := strconv.Atoi(string(k))
            repeated := strings.Repeat(string(curr), num)
            
            // Push the repeated string back to stack
            for _, r := range repeated {
                stack = append(stack, r)
            }
        }
    }
    return string(stack)
}