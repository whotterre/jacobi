package main

// Helper to compute gcd aka hcf via the Euclidean algorithm
func gcd (x, y int) int {
    if (y == 0){
        return x
    } else {
        return gcd(y, x % y)
    }
}


func gcdOfStrings(str1 string, str2 string) string {
    first := str1 + str2
    second := str2 + str1
    // If the concatenation of the first and second string is the 
    // same as the concat of the reverse then it's possible
    if !(first == second) {
        return ""
    }
    // Compute the gcd of the string lengths,
    // The shorter one is the the longest substring
    gcdLength := gcd(len(str1), len(str2))

    return str1[:gcdLength]
}
