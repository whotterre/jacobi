package main

func sortColors(arr []int) []int {
	// Initialize 3 pointers 
	// low and mid at beginning
	low, mid := 0, 0
	// and high at end
	high := len(arr) - 1

	// Traverse from mid to high
	for mid <= high {
		// if current element is 0, swap with low and increment low and mid
		switch arr[mid] {
		case 0:
			arr[mid], arr[low] = arr[low], arr[mid]
			low++
			mid++
		case 1:
			// if current element is 1, move mid pointer forward
			mid++
		default:
			// if curr Element is 2, swap with high and decrement high
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	return arr
}