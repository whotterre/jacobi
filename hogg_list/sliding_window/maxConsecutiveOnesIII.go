package main

func longestOnes(nums []int, k int) int {
	i := 0
	zeroCount := 0
	maxLength := 0
	// Algo:
	// Initialize two pointers: i and j
	// The objective is to look for contingous subarrays of 1's
	// Keep track of the maximum subarray with a variable, maxOnes
	// Check for any k zeros adjacent to the maxSubarray,
	// return k + maxOnes

	// Move the right pointer j and count zeros we meet along the way
	for j := range nums {
		if nums[j] == 0 {
			zeroCount++
		}

		// If the zeroCount exceeds k, decrement zeroCount
		for zeroCount > k {
			if nums[i] == 0 {
				zeroCount--
			}
			// Slide to the right a bit
			i++
		}

		// Get the current window length
		currentWindowLength := j - i + 1
        // Return max length
		maxLength = max(currentWindowLength, maxLength)
	}

	return maxLength
}

// The n-th max function i've written atp 
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}