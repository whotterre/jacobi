package main

func isValid(s string) bool {
    var stack []rune

    // if we find a closing bracket, put it on the stack
    // pass 1
    for _, b := range s {
        if isOpenBracket(b) {
            stack = append(stack, b)
        } else {
            if len(stack) == 0 {
                return false
            }
            n := len(stack) - 1
            tos := rune(stack[n])
            stack = stack[:n]

        if !isMatchingPair(tos, b) {
            return false
        }
    }
    }
    return len(stack) == 0
}


func isMatchingPair(open, close rune) bool {
    return open == '(' && close == ')' ||
         open == '[' && close == ']' ||
         open == '{' && close == '}'
}


func isOpenBracket(r rune) bool {
    return r == '(' || r == '{' || r == '['
}