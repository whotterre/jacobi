package main

// LeetCode 1004 - Given a binary array nums and an
//  integer k, return the maximum number of consecutive
//  1's in the array if you can flip at most k 0's.
func LongestOnes(nums []int, k int) int {
	i := 0
	zeroCount := 0
	maxLength := 0
	// Initial Algo:
	// Initialize two pointers: i and j
	// The objective is to look for contingous subarrays of 1's
	// Keep track of the maximum subarray with a variable, maxOnes
	// Check for any k zeros adjacent to the maxSubarray,
	// return maxOnes

	// Move the right pointer j and count zeros we meet along the way
	for j := 0; j < len(nums); j++ {
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
