package main

func isSubsequence(s string, t string) bool {
    // Two pointer approach
    i := 0
    j := 0

    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i += 1
        }
        j += 1
    }
    return i == len(s)
}