package main

func longestPalindrome(s string) string {
   // Keep count of max 
   maxLength := 0
   n := len(s)
   longestPalindrome := ""

   // initialize longestPalindrome with the first char
   longestPalindrome += string(s[0])
   if n == 0 {
    return ""
   }
   
   if n == 1 {
    return s
   }

   // Stride across the string
   for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
        // Take the substring
        ss := s[i:j + 1]
        // Compare with isPalindrome
        if isPalindrome(ss) && len(ss) > len(longestPalindrome) {
            maxLength = max(maxLength, j - i)
            longestPalindrome = ss
        }
    }
   }
   return longestPalindrome
}

func isPalindrome(s string) bool {
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] != s[n - i - 1] {
            return false
        }
    }
    return true
}

