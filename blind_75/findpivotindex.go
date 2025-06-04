package main

// LeetCode 724 - Given an array of integers nums, calculate the pivot index of this array.
func pivotIndex(nums []int) int {
	totalSum := getSum(nums)
	leftSum := 0
	var rightSum int

	for i := 0; i < len(nums); i++ {
		// Compute right sum
		rightSum = totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
        // Update leftSum to include the current element 
        leftSum += nums[i] 
	}
    return -1
}

// Helper function to compute the sum of elements in a list
func getSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}