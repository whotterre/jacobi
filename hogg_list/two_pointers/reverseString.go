package main

func reverseString(s []byte) {
    n := len(s) - 1

    for i := 0; i <= n/2; i++ {
        compIdx := n - i

        temp := s[i]
        s[i] = s[compIdx]
        s[compIdx] = temp
    }
}
