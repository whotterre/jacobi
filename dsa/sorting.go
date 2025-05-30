package main

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
// Bubble sort 
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}

// Merge sort 
func MergeSort(arr []int) []int {
	// Already sorted if just zero or one elements 
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid + 1:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var sortedArr []int
	// Does the actual sorting via two pointer approach
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		// If the item on the left is smaller, append it to the result arr
		if left[i] <= right[j] {
			sortedArr = append(sortedArr, left[i])
			// Move the pointer forward by one
			i += 1
		} else {
			// Otherwise, append from the right 
			sortedArr = append(sortedArr, right[j])
			// Move the pointer forward by one
			j += 1
		}
	}
	// If there are any remaining items in the left or right, append them
	sortedArr = append(sortedArr, left[i:]...)
	sortedArr = append(sortedArr, right[j:]...)
	return sortedArr
}