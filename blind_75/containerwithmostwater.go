package main

// Two pointer approach - Go implementation of NeetCode's solution
func maxArea(height []int) int {
	l := 0
	r := len(height)

	res := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return res
}

// Helper minimum function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Helper max function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
