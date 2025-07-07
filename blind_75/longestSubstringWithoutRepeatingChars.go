func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	l, r, maxLength := 0, 0, 0
	hashSet := make(map[rune]bool)

	for r < n {
        currentChar := rune(s[r])
        if !hashSet[currentChar] {
            hashSet[currentChar] = true
            maxLength = max(maxLength, r - l + 1)
            r++
        } else {
            currentChar = rune(s[l])
            hashSet[currentChar] = false
            l++
        }
    }
    return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
