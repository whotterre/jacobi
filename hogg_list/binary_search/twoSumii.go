package main
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low := i + 1
		high := len(numbers) - 1
		// use bin search to search
		for low <= high {
			mid := low + (high-low)/2
			sum := numbers[i] + numbers[mid]

			if sum == target {
				return []int{i + 1, mid + 1}
			} else if sum > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return nil
}