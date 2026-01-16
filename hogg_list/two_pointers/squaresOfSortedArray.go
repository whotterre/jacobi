package main

func sortedSquares(nums []int) []int{
	// start with a pointer at the beginning and at the end
	n := len(nums) - 1
	left := 0
	right := n
	resIdx := n

	// initialize an empty slice, with the same size of the original slice
	res := make([]int, len(nums))

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			res[resIdx] = leftSquare
			left += 1
		} else {
			res[resIdx] = rightSquare
			right -= 1
		}

		resIdx -= 1
	}
	return res
}