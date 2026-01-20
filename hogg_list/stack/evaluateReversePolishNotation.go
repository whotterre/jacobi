package main

import "strconv"

func evalRPN(tokens []string) int {
    var stack []int

    i := 0

    for i < len(tokens) {
        curr := tokens[i]

        // if it's a digit, place it on top of the stack
        if isNumber(curr) {
            // cast to int
            num, _ := strconv.Atoi(curr)
            stack = append(stack, num)
        } else {
            a := stack[len(stack) - 2]
            stack = stack[:len(stack) - 2]
            res := evaluateOperation(a, b, curr)
            stack = append(stack, res)
        }
        i += 1
    }
    return stack[0]
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func evaluateOperation(a, b int, operator string) int {
    switch operator {
        case "+":
          return a + b
        case "-":
          return a - b
        case "*":
          return a * b
        case "/":
          return a / b
    }
    return 0
}