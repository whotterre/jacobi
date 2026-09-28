package leetcode

import "strings"

func evaluate(s string, knowledge [][]string) string {
	var res strings.Builder
	keyMap := make(map[string]string)

	for i := 0; i < len(knowledge); i++ {
		currentKey := knowledge[i][0]
		if _, exists := keyMap[currentKey]; !exists {
			keyMap[currentKey] = knowledge[i][1]
		}
	}

	lastIdxOpenBr := 0
	lastKeyIdx := 0
	currentKey := ""

	i := 0

	for i < len(s) {
		currentChar := s[i]
		if currentChar == '(' {
            res.WriteString(s[lastKeyIdx:i])
			lastIdxOpenBr = i
			i += 1
		}

		i += 1

		if currentChar == ')' {
			currentKey = s[lastIdxOpenBr+1 : i-1]
			if _, exists := keyMap[currentKey]; exists {
				lastKeyIdx = i
				res.WriteString(keyMap[currentKey])
				currentKey = ""
			} else {
				lastKeyIdx = i
				res.WriteString("?")
				currentKey = ""
			}
		}
	}
    res.WriteString(s[lastKeyIdx:])
	return res.String()
}