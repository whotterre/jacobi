package main

/*
	 LeetCode 1732 - Given an integer array gain,
	 where gain[i] is the net gain in altitude between points i and i + 1.
		Return the highest altitude that can be reached.
	   The highest altitude is the maximum altitude reached at any point.
*/
func largestAltitude(gain []int) int {
	// Start with 0
	altitude := 0
	maxAlt := 0
	// Traverse the gain arr
	for _, change := range gain {
		// Adding the change to the altitude at each step
		altitude += change
		// Get the max and update the maxAlt
		maxAlt = maximum(maxAlt, altitude)
	}
	return maxAlt
}

// Helper max function
func maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
