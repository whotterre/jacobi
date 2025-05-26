package main

import "math"

func increasingTriplet(nums []int) bool {
	// Initialize two variables to hold the first and 
	// second smallest values to the maximum possible integer value.
	ersten := int(math.Inf(0))
	zweiten := int(math.Inf(0))

	// Iterate through the numbers in the array.
	// If the current number is less than or equal to the first smallest,
	// update the first smallest
	for _, zahl := range nums {
		if zahl <= ersten {
			ersten = zahl 
		} else if zahl <= zweiten {
			zweiten = zahl
		} else {
			return true
		}
	}
	return false
}