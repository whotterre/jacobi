package main


// Two pointer approach to this
func isSubsequence(s string, t string) bool {
    // Initialize starting pointers for both strings
    i := 0
    j := 0
    // Go through the strings and check that the chars are in the right order 
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i += 1 // move the subseq ptr by 1 to check that the next char is in place
        }
        j += 1
    }
    // check that index at the end matches the length of the suspected subseq
    return i == len(s)
}
