package main
func mergeAlternately(word1 string, word2 string) string {
    m := len(word1)
    n := len(word2)

    minLen := min(m, n)
    var result string 


    for i := 0; i < minLen; i++ {
        result += string(word1[i])
        result += string(word2[i])
    }

    if m > n {
        result += string(word1[minLen:])
    } else {
         result += string(word2[minLen:])
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a 
    }
    return b
}